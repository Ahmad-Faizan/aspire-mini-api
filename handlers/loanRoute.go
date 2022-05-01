package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ahmad-Faizan/aspire-mini-api/models"
	"github.com/gin-gonic/gin"
)

func AddLoan(c *gin.Context) {
	var l *models.Loan
	userRaw := c.Param("userId")
	user, err := strconv.ParseInt(userRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	if err := c.ShouldBindJSON(&l); err != nil {
		log.Print(err)
	}

	l.UserID = user
	l, err = models.AddLoan(l)
	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, l)
}

func GetLoans(c *gin.Context) {
	userRaw := c.Param("userId")
	userID, err := strconv.ParseInt(userRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	allLoans := models.GetAllLoans(userID)

	c.JSON(http.StatusOK, allLoans)
}

func GetLoan(c *gin.Context) {
	userRaw := c.Param("userId")
	userID, err := strconv.ParseInt(userRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	loanRaw := c.Param("loanId")
	loanID, err := strconv.ParseInt(loanRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	l, err := models.GetLoan(userID, loanID)
	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, l)
}

func UpdateLoan(c *gin.Context) {
	userRaw := c.Param("userId")
	userID, err := strconv.ParseInt(userRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	loanRaw := c.Param("loanId")
	loanID, err := strconv.ParseInt(loanRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	var l models.Loan
	if err := c.ShouldBindJSON(&l); err != nil {
		log.Print(err)
	}

	err = models.ApproveLoan(userID, loanID)
	if err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "approved"})
}

func RepayLoan(c *gin.Context) {

	loanRaw := c.Param("loanId")
	loanID, err := strconv.ParseInt(loanRaw, 0, 64)
	if err != nil {
		log.Print(err)
	}

	l, err := models.GetSpecificLoan(loanID)
	if err != nil {
		log.Print(err)
	}

	var loan models.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		log.Print(err)
	}

	amount := loan.Amount
	l.Repay(float32(amount))
}
