package handler

import (
	"gorm-echo/model"

	"gorm.io/gorm"
)

func CreateUser() {
	//
}

func CreateUsers() {
	//
}

func DeleteUser() {
	//
}

func Login() {
	//
}

func GetUsers(db *gorm.DB) (res interface{}, httpStatus int, err error) {
	//
	var users []model.User
	db.Find(&users)

	//still don't know how to get error code when user not found
	return res, httpStatus, err
}

func GetUser(db *gorm.DB) (res interface{}, httpStatus int, err error) {
	var users []model.User
	db.Find(&users)

	//still don't know how to get error code when user not found
	return res, httpStatus, err
}
