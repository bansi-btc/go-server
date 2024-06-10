package main

import (
	"chi/routes"
	"fmt"
	"net/http"
)

func main() {

	r := routes.GetRouter()

	fmt.Println("Server started at port 8080")

	http.ListenAndServe(":8080", r)
}
