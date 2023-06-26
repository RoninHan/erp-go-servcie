package models

import "time"

type SalaryHistory struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	EmployeeID    string    `json:"employee_id" gorm:"not null"`
	EffectiveDate time.Time `json:"effective_date" gorm:"not null"`
	Salary        float64   `json:"salary" gorm:"not null"`
	Employee      Employee  `gorm:"foreignKey:EmployeeID"`
}
