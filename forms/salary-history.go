package forms

import "time"

type SalaryHistoryForm struct {
	EmployeeID    string    `json:"employee_id" gorm:"not null"`
	EffectiveDate time.Time `json:"effective_date" gorm:"not null"`
	Salary        float64   `json:"salary" gorm:"not null"`
}
