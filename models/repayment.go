package models

type Repayment struct {
	Amount        float32
	PendingAmount float32
	ScheduledDate string
	Status        RepaymentStatus
}

type RepaymentStatus string

const (
	Awaited   RepaymentStatus = "Awaited"
	Completed RepaymentStatus = "Completed"
)
