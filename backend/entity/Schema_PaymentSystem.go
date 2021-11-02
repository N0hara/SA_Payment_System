package entity

import (
	"time"

	"gorm.io/gorm"
)

type FinancialOfficer struct {
	gorm.Model
	FinancialName string
	Pid           string `gorm:"uniqueIndex"`
	Password      string
}

type RightTreatment struct {
	gorm.Model

	RightTreatmentName   string `gorm:"uniqueIndex"`
	RightTreatmentDetail string
	Price                int

	Admissions []Admission `gorm:"foreignKey:RightTreatmentID"`
	Bills      []Bill      `gorm:"foreignKey:RightTreatmentID"`
}

type PaymentMethod struct {
	gorm.Model
	PaymentMethodName string `gorm:"uniqueIndex"`

	Bills []Bill `gorm:"foreignKey:PaymentMethodID"`
}

type Patient struct {
	gorm.Model
	IdentificationID string `gorm:"uniqueIndex"`
	PatientName      string

	Admissions []Admission `gorm:"foreignKey:PatientID"`
}

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

type Bill struct {
	gorm.Model
	BillDateTime      time.Time
	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord
	PatientPID        string
	RightTreatmentID  *uint
	RightTreatment    RightTreatment
	PaymentMethodID   *uint
	PaymentMethod     PaymentMethod
	AmountPaid        int
}
