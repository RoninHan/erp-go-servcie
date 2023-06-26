package forms

import "time"

type EmployeeForm struct {
	Username      string
	Password      string
	FirstName     string
	LastName      string
	DepartmentID  string
	Position      string
	HireDate      time.Time
	Salary        float64
	Email         string
	ContactNumber string
	Address       string
}
