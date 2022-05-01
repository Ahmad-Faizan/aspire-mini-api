package handlers

import (
	"log"
	"net/http"

	"github.com/Ahmad-Faizan/aspire-mini-api/models"
	"github.com/gin-gonic/gin"
)

// fetch all users from db
func GetUsers(c *gin.Context) {

	allUsers := models.GetAllUsers()

	c.JSON(http.StatusOK, allUsers)
}

// add a user to db
func AddUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		log.Print(err)
	}

	u = models.AddUser(u)
	c.JSON(http.StatusOK, u)
}
