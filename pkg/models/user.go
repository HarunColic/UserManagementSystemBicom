package models

import (
	"UserManagementSystem/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string `gorm:"" json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    int    `json:"roleID"`
	Role      Role
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var usr User

	db := db.Where("ID=?", Id).Find(&usr)
	return &usr, db
}

func DeleteUser(Id int64) User {
	var usr User
	db.Where("ID=?", Id).Delete(usr)
	return usr
}
