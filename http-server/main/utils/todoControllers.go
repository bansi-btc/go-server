package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type todo struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

var todoArr = []todo{}

func HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	timeStampSt := time.Now().UnixNano()

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

	// res := responseType1{
	// 	Success: true,
	// 	Message: "Todo added successfully",
	// 	Todo:    todoArr,
	// }

	jsondata, _ := json.Marshal(jsonRes)
	timeStampFin := time.Now().UnixNano()

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsondata)

	processingTime := timeStampFin - timeStampSt

	fmt.Println("processing time is", processingTime)
}

func HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	timeStampSt := time.Now().UnixNano()

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

	// fmt.Println(todoArr)
	timeStampFin := time.Now().UnixNano()

	processingTime := timeStampFin - timeStampSt

	fmt.Println("processing time is", processingTime)

}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	timeStampSt := time.Now().UnixNano()

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

	fmt.Println(idx)

	todoArr = append(todoArr[0:idx], todoArr[idx+1:]...)

	fmt.Println(todoArr)

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
	timeStampFin := time.Now().UnixNano()

	processingTime := timeStampFin - timeStampSt

	fmt.Println("processing time is", processingTime)

}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	timeStampSt := time.Now().UnixNano()

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		jsonErrRes := map[string]interface{}{
			"Success": false,
			"Message": err,
		}

		jsonData, _ := json.Marshal(jsonErrRes)

		w.Write(jsonData)
	}

	reqData := make(map[string]interface{})

	json.Unmarshal(data, &reqData)

	id, ok := reqData["id"]

	if !ok {
		fmt.Println("Error present")
	}

	floatId := id.(float64)

	intId := int(floatId)

	task := reqData["task"]

	fmt.Printf("%T", task)

	for index, value := range todoArr {
		if index == intId {
			value.Task = "him"
		}
	}

	jsonRes := map[string]interface{}{
		"Success": true,
		"Message": "Updated todo successfully",
		"Todo":    todoArr,
	}

	jsonData, _ := json.Marshal(jsonRes)

	w.Write(jsonData)

	timeStampFin := time.Now().UnixNano()

	processingTime := timeStampFin - timeStampSt

	fmt.Println("processing time is", processingTime)
}

func SayHello() {
	fmt.Println("Hello")
}
