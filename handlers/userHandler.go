package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pattyvader/users_service/dao"
	"github.com/pattyvader/users_service/models"
)

//GetAllUsersHandler retorna em formato json todos os usuários cadastrados no banco de dados
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	response := users
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

//GetUserByIDHandler retorna em formato json um usuário
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID := param["id"]

	user, err := dao.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	response := user
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

//CreateUserHandler cria um novo usuário
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	userCreate, err := dao.CreateUser(user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	response := userCreate
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

//UpdateUserHandler atualiza um usuário
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID := param["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	userUpdate, err := dao.UpdateUser(user, userID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	response := userUpdate
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// DeleteUserHandler remove um usuário
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID := param["id"]

	user, err := dao.RemoveUser(userID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	response := user
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
