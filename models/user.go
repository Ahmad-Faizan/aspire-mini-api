package models

import (
	"fmt"
)

type User struct {
	Id      int64  `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	IsAdmin bool   `form:"isAdmin" json:"isAdmin"`
}

type Users []User

func GetAllUsers() Users {
	return userList
}

func AddUser(u User) User {
	userList = append(userList, u)
	return u
}

func GetUser(userID int64) (User, error) {
	for _, user := range userList {
		if user.Id == userID {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("no user found")
}

func isAdminUser(userID int64) bool {
	u, err := GetUser(userID)

	if err != nil {
		return false
	}

	return u.IsAdmin
}

var userList = []User{
	{
		Id:      1001,
		Name:    "John",
		IsAdmin: false,
	},
	{
		Id:      1002,
		Name:    "Sam",
		IsAdmin: false,
	}, {
		Id:      1003,
		Name:    "Kate",
		IsAdmin: true,
	},
}
