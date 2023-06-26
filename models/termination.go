package models

import "time"

type Termination struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	EmployeeID      string    `json:"employee_id" gorm:"not null"`
	TerminationDate time.Time `json:"termination_date" gorm:"type:datetime"`
	Reason          string    `json:"reason" gorm:"not null"`
}

func (Termination) TableName() string {
	return "termination"
}
