package main

import (
	"fmt"
	"log"
	"native/utils"
	"net/http"
)

func main() {

	utils.SayHello()
	fmt.Println("server starting at port 8080")

	http.HandleFunc("/addTodo", utils.HandleAddTodo)
	http.HandleFunc("/getTodo", utils.HandleGetTodo)
	http.HandleFunc("/deleteTodo", utils.HandleDeleteTodo)
	http.HandleFunc("/updateTodo", utils.HandleUpdateTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
