package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vibinnair/user-management-in-mem/internal/service"
	"github.com/vibinnair/user-management-in-mem/model"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetPing(writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) CreateUser(writer http.ResponseWriter, reader *http.Request) {
	var user model.User

	error := json.NewDecoder(reader.Body).Decode(&user)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusBadRequest)
		return
	}

	error = handler.service.CreateUser(&user)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) UpdateUser(writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) GetUser(writer http.ResponseWriter, reader *http.Request) {
	userID, error := strconv.Atoi(reader.PathValue("id"))
	if error != nil {
		http.Error(writer, "Empty User ID", http.StatusNotFound)
		fmt.Println("Empty User ID")
		return
	}

	fmt.Printf("User ID: %d", userID)
	user, err := handler.service.GetUser(userID)
	if err != nil {
		http.Error(writer, "User Not Found", http.StatusNotFound)
		fmt.Println("User Not Found")
		return
	}

	fmt.Println("User Found")
	json.NewEncoder(writer).Encode(user)
}

func (handler *UserHandler) DeleteUser(writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
