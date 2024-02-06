package types

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type EmployeeRequest struct {
	Name        string    `json:"name"`
	Designation string    `json:"designation"`
	Date        time.Time `json:"date"`
	EntryTime   time.Time `json:"entry_time,omitempty"`
	ExitTime    time.Time `json:"exit_time,omitempty"`
}

func (emp EmployeeRequest) Validate() error {
	return validation.ValidateStruct(&emp,
		validation.Field(&emp.Name,
			validation.Required.Error("Employee name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&emp.Designation,
			validation.Required.Error("Designation ID is required")),
	)
}
