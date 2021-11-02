package entity

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	IdentificationID string `gorm:"uniqueIndex"`
	PatientName      string

	Admissions []Admission `gorm:"foreignKey:PatientID"`
}
