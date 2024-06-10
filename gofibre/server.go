package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	Task string `json:"task"`
	Id   int    `json:"id"`
}

var todoArr = []todo{}

func handleAddTodo(c *fiber.Ctx) error {

	var reqData todo

	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Success": true,
			"Message": err.Error(),
		})
	}

	reqData.Id = len(todoArr)

	todoArr = append(todoArr, reqData)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Success": true,
		"Message": "Todo added successfully",
		"Todos":   todoArr,
	})
}

func handleGetTodo(c *fiber.Ctx) error {

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Success": true,
		"Message": "Fetched successfully",
		"Todos":   todoArr,
	})
}

func handleDeleteTodo(c *fiber.Ctx) error {

	var reqData struct {
		Id int `json:"id"`
	}

	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"Success": false,
			"Message": err.Error(),
		})
	}

	id := reqData.Id

	// fmt.Printf("%T", id)
	fmt.Println(id)

	var ind int = -1

	for index, value := range todoArr {

		if value.Id == id {
			ind = index
		}
	}

	if ind == -1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Success": false,
			"Message": "Invalid delete",
		})
	}

	todoArr = append(todoArr[:ind], todoArr[ind+1:]...)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Success": true,
		"Message": "Todo deleted successfully",
		"Todos":   todoArr,
	})
}

func handleUpdateTodo(c *fiber.Ctx) error {
	var reqData struct {
		Id   int    `json:"id"`
		Task string `json:"task"`
	}

	if err := c.BodyParser(&reqData); err != nil {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
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

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Success": true,
		"Message": "Updated successfully",
		"Todos":   todoArr,
	})

}

func main() {

	app := fiber.New()

	app.Post("/addTodo", handleAddTodo)
	app.Get("/getTodo", handleGetTodo)
	app.Delete("/deleteTodo", handleDeleteTodo)
	app.Put("/updateTodo", handleUpdateTodo)

	app.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	})

	app.Listen(":8080")
}
