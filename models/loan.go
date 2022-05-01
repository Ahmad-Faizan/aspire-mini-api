package models

import (
	"fmt"
	"log"
	"time"
)

type Loan struct {
	Id         int64        `json:"id" form:"id"`
	UserID     int64        `json:"userId"`
	ApprovedBy int64        `json:"approvedBy" form:"approvedBy"`
	Amount     int64        `json:"amount" form:"amount"`
	Status     LoanStatus   `json:"loanStatus"`
	StartDate  string       `json:"startDate" form:"startDate"`
	Term       int          `json:"term" form:"term"`
	Repayments []*Repayment `json:"repayments"`
}

type LoanStatus string

const (
	Pending  LoanStatus = "Pending"
	Approved LoanStatus = "Approved"
	Paid     LoanStatus = "Paid"
)

type Loans []*Loan

var loanList Loans

func AddLoan(l *Loan) (*Loan, error) {
	var r []*Repayment

	sd, err := time.Parse("2006-01-02", l.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error in parsing start date %s", err)
	}

	for i := 1; i <= l.Term; i++ {
		var rp Repayment

		rp.Amount = float32(l.Amount) / float32(l.Term)
		rp.PendingAmount = rp.Amount
		rp.ScheduledDate = sd.Add(time.Duration(i*24*7) * time.Hour).String()
		rp.Status = RepaymentStatus(Awaited)

		r = append(r, &rp)
	}

	newLoan := Loan{
		Id:         l.Id,
		UserID:     l.UserID,
		ApprovedBy: 0,
		Amount:     l.Amount,
		Status:     LoanStatus(Pending),
		StartDate:  l.StartDate,
		Term:       l.Term,
		Repayments: r,
	}
	loanList = append(loanList, &newLoan)
	return &newLoan, nil
}

func GetAllLoans(userID int64) Loans {
	if isAdminUser(userID) {
		return loanList
	}

	var userLoans Loans
	for _, l := range loanList {
		if l.UserID != userID {
			continue
		}
		userLoans = append(userLoans, l)
	}
	return userLoans
}

func GetLoan(userID int64, loanID int64) (Loan, error) {
	var loan Loan
	for _, l := range loanList {
		if l.Id == loanID {
			loan = *l
			break
		}
	}

	if loan.UserID != userID && !isAdminUser(userID) {
		return Loan{}, fmt.Errorf("access denied")
	}

	return loan, nil
}

func GetSpecificLoan(loanID int64) (Loan, error) {
	var loan Loan
	for _, l := range loanList {
		if l.Id == loanID {
			loan = *l
			break
		}
	}
	return loan, nil
}

func ApproveLoan(userID int64, loanID int64) error {
	if !isAdminUser(userID) {
		return fmt.Errorf("only admins can approve loan")
	}

	l, err := GetLoan(userID, loanID)
	if err != nil {
		log.Print(err)
	}

	l.Approve(userID)
	for i, v := range loanList {
		if v.Id == loanID {
			loanList[i] = &l
			break
		}
	}

	return nil
}

func (l *Loan) Approve(userID int64) {
	l.ApprovedBy = userID
	l.Status = LoanStatus(Approved)
}

func (l *Loan) Repay(amount float32) {
	var extra float32
	var paidAmount float32
	var paidTerms int

	for i, rp := range l.Repayments {
		if rp.Status == RepaymentStatus(Completed) {
			paidAmount = paidAmount + rp.Amount
			paidTerms++
			continue
		}

		if rp.PendingAmount < amount {
			rp.PendingAmount = 0
			rp.Status = RepaymentStatus(Completed)
		}
		if rp.PendingAmount == amount {
			rp.PendingAmount = 0
			rp.Status = RepaymentStatus(Completed)
		}
		if rp.PendingAmount > amount {
			rp.PendingAmount = rp.PendingAmount - amount
		}

		extra = amount - rp.PendingAmount
		l.Repayments[i] = rp
		break
	}

	log.Printf("paidAmount %f", paidAmount)
	log.Printf("paidTerms %d", paidTerms)
	log.Printf("extra %f", extra)
	if extra > 0 {
		finalPendingAmount := (float32(l.Amount) - paidAmount - extra) / (float32(l.Term) - float32(paidTerms))
		log.Printf("finalPendingAmount %f", finalPendingAmount)

		for i, rp := range l.Repayments {
			if rp.Status == RepaymentStatus(Completed) {
				continue
			}
			rp.PendingAmount = finalPendingAmount
			if rp.PendingAmount == 0 {
				rp.Status = RepaymentStatus(Completed)
			}
			l.Repayments[i] = rp
			break
		}
		if finalPendingAmount < 0 {
			l.Status = LoanStatus(Paid)
		}
	}
}
