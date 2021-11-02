package entity

import (
	"gorm.io/gorm"
)

type FinancialOfficer struct {
	gorm.Model
	FinancialName string
	Pid           string `gorm:"uniqueIndex"`
	Password      string
}
