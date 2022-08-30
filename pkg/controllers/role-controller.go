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

var newRole models.Role

func GetRoles(w http.ResponseWriter, r *http.Request) {
	roles := models.GetAllRoles()
	res, _ := json.Marshal(roles)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoleId := vars["roleId"]
	ID, err := strconv.ParseInt(RoleId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	role, _ := models.GetRoleById(ID)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	newRole := &models.Role{}
	utils.ParseBody(r, newRole)
	u := newRole.CreateRole()
	res, _ := json.Marshal(u)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	RoleId := vars["roleId"]
	ID, err := strconv.ParseInt(RoleId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	role := models.DeleteRole(ID)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var updateRole = &models.Role{}
	utils.ParseBody(r, updateRole)
	vars := mux.Vars(r)
	RoleId := vars["roleId"]
	ID, err := strconv.ParseInt(RoleId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	role, db := models.GetRoleById(ID)

	role = updateRole
	role.ID = uint(ID)
	db.Save(role)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "pkglication/json")
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("Error while writing header", err.Error())

	}
	w.WriteHeader(http.StatusOK)
}
