package controllers

import (
	"UserManagementSystem/pkg/models"
	"UserManagementSystem/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var newUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["userId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	user, _ := models.GetUserById(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.CreateUser()
	res, _ := json.Marshal(u)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	UserId := vars["userId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	UserId := vars["userId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	user, db := models.GetUserById(ID)

	user = updateUser
	user.ID = uint(ID)
	db.Save(user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func AssignRole(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	UserId := vars["userId"]
	ID, err := strconv.ParseInt(UserId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	user, db := models.GetUserById(ID)

	if updateUser.RoleID != 0 {
		user.RoleID = updateUser.RoleID
	}

	db.Save(user)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}
