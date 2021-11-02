package entity

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	PaymentMethodName string `gorm:"uniqueIndex"`

	Bills []Bill `gorm:"foreignKey:PaymentMethodID"`
}
