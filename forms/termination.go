package forms

import "time"

type TerminationForm struct {
	EmployeeID      string
	TerminationDate time.Time
	Reason          string
}
