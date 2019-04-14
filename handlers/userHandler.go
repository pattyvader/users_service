package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pattyvader/users_service/dao"
	"github.com/pattyvader/users_service/models"
)

//GetAllUsersHandler retrives all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := dao.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	FormatResponseToJSON(w, http.StatusOK, users)
}

//GetUserByIDHandler retrieves only one user
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	paramID := param["id"]

	userID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Println(err)
	}

	user, err := dao.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	FormatResponseToJSON(w, http.StatusOK, user)
}

//CreateUserHandler creates a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	newUser, err := dao.CreateUser(user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	FormatResponseToJSON(w, http.StatusCreated, newUser)
}

//UpdateUserHandler updates an user
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	paramID := param["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	userID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Println(err)
	}

	updateUser, code, err := dao.UpdateUser(user, userID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if code == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	FormatResponseToJSON(w, http.StatusOK, updateUser)
}

// DeleteUserHandler removes an user
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	paramID := param["id"]

	userID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Println(err)
	}

	result, err := dao.RemoveUser(userID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if result == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	FormatResponseToJSON(w, http.StatusAccepted, "Accepted")
}

//FormatResponseToJSON format the response to json
func FormatResponseToJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
