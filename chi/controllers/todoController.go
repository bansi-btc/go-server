package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type todo struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

var todoArr = []todo{}

func HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("invalid request")
	}

	data, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var reqData todo

	json.Unmarshal(data, &reqData)

	reqData.Id = len(todoArr)

	todoArr = append(todoArr, reqData)

	jsonRes := map[string]interface{}{
		"Success": true,
		"Message": "todo added successfully",
		"Todo":    todoArr,
	}

	jsondata, err := json.Marshal(jsonRes)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsondata)

}

func HandleGetTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		fmt.Println("Invalid request")
		return
	}

	jsonRes := map[string]interface{}{
		"Success": true,
		"Message": "Fetched successfully",
		"Todos":   todoArr,
	}

	jsonData, err := json.Marshal(jsonRes)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)

}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	reqData := make(map[string]interface{})

	json.Unmarshal(data, &reqData)

	id := reqData["id"]
	idFloat, ok := id.(float64)

	if !ok {
		fmt.Println(ok)
	}

	idVal := int(idFloat)

	var idx int

	for index, val := range todoArr {
		if val.Id == idVal {

			idx = index
		}
	}

	todoArr = append(todoArr[0:idx], todoArr[idx+1:]...)

	jsonRes := map[string]interface{}{
		"Success": true,
		"Message": "Todo deleted successfully",
		"Todos":   todoArr,
	}

	jsonData, err := json.Marshal(jsonRes)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {

	var reqData struct {
		Id   int    `json:"id"`
		Task string `json:"task"`
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		jsonRes := map[string]interface{}{
			"Success": false,
			"Message": err.Error(),
		}

		jsonData, _ := json.Marshal(jsonRes)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

	json.Unmarshal(data, &reqData)

	id := reqData.Id

	for index, value := range todoArr {

		if value.Id == id {
			todoArr[index].Task = reqData.Task
			break
		}
	}

	jsonRes := map[string]interface{}{
		"Success": true,
		"Message": "Todo updated successfully",
		"Todos":   todoArr,
	}

	jsonData, _ := json.Marshal(jsonRes)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
