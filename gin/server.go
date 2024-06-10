package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

var todoArr = []todo{}

func handleAddTodo(c *gin.Context) {

	var reqData todo

	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{
			"Success": false,
			"Message": err.Error(),
		})
	}

	reqData.Id = len(todoArr)

	todoArr = append(todoArr, reqData)

	c.JSON(200, gin.H{
		"Success": true,
		"Message": "Todo added successfully",
		"Todo":    todoArr,
	})

}

func handleGetTodo(c *gin.Context) {

	c.JSON(200, gin.H{
		"Success": true,
		"Message": "Todos fetched successfully",
		"Todos":   todoArr,
	})
}

func handleDeleteTodo(c *gin.Context) {

	var reqData struct {
		Id int `json:"id"`
	}

	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{
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
		c.JSON(400, gin.H{
			"Success": false,
			"Message": "Invalid delete",
		})
		return

	}

	todoArr = append(todoArr[:ind], todoArr[ind+1:]...)

	c.JSON(200, gin.H{
		"Success": true,
		"Message": "Deleted succesfully",
		"Todos":   todoArr,
	})
}

func handleUpdateTodo(c *gin.Context) {
	var reqData struct {
		Id   int    `json:"id"`
		Task string `json:"task"`
	}

	if err := c.BindJSON(&reqData); err != nil {
		c.JSON(400, gin.H{
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

	c.JSON(200, gin.H{
		"Success": true,
		"Message": "Todo updated successfully",
		"Todos":   todoArr,
	})
}

func main() {
	// fmt.Println("Gin")

	r := gin.Default()

	r.POST("/addTodo", handleAddTodo)
	r.GET("/getTodo", handleGetTodo)
	r.DELETE("/deleteTodo", handleDeleteTodo)
	r.PUT("/updateTodo", handleUpdateTodo)

	r.Run(":8080")
}
