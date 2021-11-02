package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admission struct {
	gorm.Model
	PatientID          *uint
	Patient            Patient
	PatientName        string
	Roomid             string
	Admissionentrytime time.Time
	RightTreatmentID   *uint
	RightTreatment     RightTreatment

	TreatmentRecords []TreatmentRecord `gorm:"foreignKey:AdmissionID"`
}
