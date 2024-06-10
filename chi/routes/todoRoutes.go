package routes

import (
	utils "chi/controllers"

	"github.com/go-chi/chi/v5"
)

func GetRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Post("/addTodo", utils.HandleAddTodo)
	r.Get("/getTodo", utils.HandleGetTodo)
	r.Delete("/deleteTodo", utils.HandleDeleteTodo)
	r.Put("/updateTodo", utils.HandleUpdateTodo)

	return r
}
