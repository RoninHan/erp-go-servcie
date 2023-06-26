package models

import "time"

type Employee struct {
	ID            string       `json:"id" gorm:"primaryKey"`
	Username      string       `json:"user_name" gorm:"not null;unique"`
	Password      string       `json:"password" gorm:"not null"`
	FirstName     string       `json:"first_name" gorm:"not null"`
	LastName      string       `json:"last_name" gorm:"not null"`
	DepartmentID  string       `json:"department_id" gorm:"not null"`
	Position      string       `json:"position" gorm:"not null"`
	HireDate      time.Time    `json:"hire_date" gorm:"not null"`
	Salary        float64      `json:"salary" gorm:"not null"`
	Email         string       `json:"email" gorm:"not null"`
	ContactNumber string       `json:"contact_number" gorm:"not null"`
	Address       string       `json:"address" gorm:"not null"`
	Termination   *Termination // 员工离职关联
}

func (Employee) TableName() string {
	return "employee"
}
