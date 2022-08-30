package models

import (
	"UserManagementSystem/pkg/config"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (r *Role) CreateRole() *Role {
	db.NewRecord(r)
	db.Create(&r)
	return r
}

func GetAllRoles() []Role {
	var roles []Role
	db.Find(&roles)
	return roles
}

func GetRoleById(Id int64) (*Role, *gorm.DB) {
	var role Role

	db := db.Where("ID=?", Id).Find(&role)
	return &role, db
}

func DeleteRole(Id int64) Role {
	var role Role
	db.Where("ID=?", Id).Delete(role)
	return role
}
