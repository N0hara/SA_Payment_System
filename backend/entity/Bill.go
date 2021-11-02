package entity

import (
	"time"

	"gorm.io/gorm"
)

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
