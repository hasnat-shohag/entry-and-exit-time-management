package models

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	Name        string
	Designation string
	Date        time.Time
	EntryTime   time.Time
	ExitTime    time.Time
}
