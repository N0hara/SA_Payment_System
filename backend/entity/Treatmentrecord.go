package entity

import (
	"time"

	"gorm.io/gorm"
)

type TreatmentRecord struct {
	gorm.Model
	AdmissionID *uint
	Admission   Admission
	DoctorID    string
	RecordDate  time.Time
	Treatment   string
	FoodType    uint
	MedID       string
	MedAmount   uint
	EquipmentID uint
	Cost        int

	Bills []Bill `gorm:"foreignKey:TreatmentRecordID"`
}
