package main

import (
	"net/http"

	"github.com/vibinnair/user-management-in-mem/internal/handler"
	"github.com/vibinnair/user-management-in-mem/internal/repository"
	"github.com/vibinnair/user-management-in-mem/internal/service"
)

func main() {
	mux := http.NewServeMux()
	repository := repository.NewUserRepository()
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	mux.HandleFunc("GET /ping", handler.GetPing)
	mux.HandleFunc("GET /users/{id}", handler.GetUser)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users", handler.CreateUser)

	http.ListenAndServe(":8282", mux)
}
