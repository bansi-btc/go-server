package main

import (
	"github.com/kataras/iris/v12"
)

type todo struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

var todoArr = []todo{}

func handleAddTodo(ctx iris.Context) {

	var reqData todo

	if err := ctx.ReadJSON(&reqData); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"Success": false,
			"Message": err.Error(),
		})
	}

	reqData.Id = len(todoArr)

	todoArr = append(todoArr, reqData)

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(iris.Map{
		"Success": true,
		"Message": "Todo added successfully",
		"Todos":   todoArr,
	})
}

func handleGetTodo(ctx iris.Context) {

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(iris.Map{
		"Success": true,
		"Message": "Fetched successfully",
		"Todos":   todoArr,
	})
}

func handleDeleteTodo(ctx iris.Context) {

	var reqData struct {
		Id int `json:"id"`
	}

	if err := ctx.ReadJSON(&reqData); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"Success": false,
			"Message": err.Error(),
		})
	}

	id := reqData.Id

	var ind int = -1

	for index, value := range todoArr {

		if value.Id == id {
			ind = index
		}
	}

	if ind == -1 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"Success": false,
			"Message": "Invalid request",
		})

		return
	}
	todoArr = append(todoArr[:ind], todoArr[ind+1:]...)

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(iris.Map{
		"Success": true,
		"Message": "Todo deleted successfully",
		"Todos":   todoArr,
	})
}

func handleUpdateTodo(ctx iris.Context) {

	var reqData struct {
		Id   int    `json:"id"`
		Task string `json:"task"`
	}

	if err := ctx.ReadJSON(&reqData); err != nil {
		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(iris.Map{
			"Success": false,
			"Message": err.Error(),
		})
	}

	id := reqData.Id

	for index, value := range todoArr {

		if value.Id == id {
			todoArr[index].Task = reqData.Task
			break
		}
	}

	ctx.StatusCode(iris.StatusAccepted)
	ctx.JSON(iris.Map{
		"Success": true,
		"Message": "Todo updated successfully",
		"Todos":   todoArr,
	})
}

func main() {
	app := iris.New()

	api := app.Party("/api/v1")
	{
		api.Post("/addTodo", handleAddTodo)
		api.Get("/getTodo", handleGetTodo)
		api.Delete("/deleteTodo", handleDeleteTodo)
		api.Put("/updateTodo", handleUpdateTodo)

	}

	app.Listen(":8080")
}
