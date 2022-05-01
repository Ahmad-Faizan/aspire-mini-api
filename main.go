package main

import (
	"fmt"

	"github.com/Ahmad-Faizan/aspire-mini-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello World")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", handlers.AddUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:userId/loans", handlers.GetLoans)
	r.POST("/users/:userId/loans", handlers.AddLoan)
	r.GET("/users/:userId/loans/:loanId", handlers.GetLoan)
	r.PUT("/users/:userId/loans/:loanId", handlers.UpdateLoan)
	r.POST("/loans/:loanId/payments", handlers.RepayLoan)
	// r.GET("/users/:userId/loans/:loanId/payments", handlers.GetAllPayments)

	r.Run()

}
