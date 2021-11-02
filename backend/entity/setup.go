package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("PaymentSystem.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		&FinancialOfficer{}, &RightTreatment{}, &PaymentMethod{}, &Patient{}, &Admission{}, &TreatmentRecord{}, &Bill{},
	)
	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	//FinancialOfficer
	db.Model(&FinancialOfficer{}).Create(&FinancialOfficer{
		FinancialName: "Somchai",
		Pid:           "1234567890123",
		Password:      string(password),
	})
	db.Model(&FinancialOfficer{}).Create(&FinancialOfficer{
		FinancialName: "Samhon",
		Pid:           "3210987654321",
		Password:      string(password),
	})

	//RightTreatment
	rt1 := RightTreatment{
		RightTreatmentName:   "ไม่ใช้สิทธิค่ารักษา",
		RightTreatmentDetail: "ไม่มีส่วนลด",
		Price:                0,
	}
	db.Model(&RightTreatment{}).Create(&rt1)
	rt2 := RightTreatment{
		RightTreatmentName:   "บัตร30",
		RightTreatmentDetail: "ลดเหลือ 30 บาท",
		Price:                30,
	}
	db.Model(&RightTreatment{}).Create(&rt2)
	rt3 := RightTreatment{
		RightTreatmentName:   "ข้าราชการ",
		RightTreatmentDetail: "ลดสูงสุด 20000 บาท",
		Price:                -20000,
	}
	db.Model(&RightTreatment{}).Create(&rt3)
	rt4 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น1",
		RightTreatmentDetail: "ลดสูงสุด 30000 บาท",
		Price:                -30000,
	}
	db.Model(&RightTreatment{}).Create(&rt4)
	rt5 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น2",
		RightTreatmentDetail: "ลดสูงสุด 15000 บาท",
		Price:                -15000,
	}
	db.Model(&RightTreatment{}).Create(&rt5)
	rt6 := RightTreatment{
		RightTreatmentName:   "ประกันชั้น3",
		RightTreatmentDetail: "ลดสูงสุด 7000 บาท",
		Price:                -7000,
	}
	db.Model(&RightTreatment{}).Create(&rt6)

	//PaymentMethod
	pm1 := PaymentMethod{
		PaymentMethodName: "เงินสด",
	}
	db.Model(&PaymentMethod{}).Create(&pm1)
	pm2 := PaymentMethod{
		PaymentMethodName: "บัตรเครดิต",
	}
	db.Model(&PaymentMethod{}).Create(&pm2)
	pm3 := PaymentMethod{
		PaymentMethodName: "ออนไลน์",
	}
	db.Model(&PaymentMethod{}).Create(&pm3)

	//Patient
	pt1 := Patient{
		IdentificationID: "1400000000001",
		PatientName:      "Somchai Saichom",
	}
	db.Model(&Patient{}).Create(&pt1)
	pt2 := Patient{
		IdentificationID: "1400000000002",
		PatientName:      "Somsri Sridang",
	}
	db.Model(&Patient{}).Create(&pt2)
	pt3 := Patient{
		IdentificationID: "1400000000003",
		PatientName:      "Sombut Sasbom",
	}
	db.Model(&Patient{}).Create(&pt3)

	//Admission
	ad1 := Admission{
		Patient:        pt1,
		PatientName:    "Somchai Saichom",
		Roomid:         "RM2002",
		RightTreatment: rt1,
	}
	db.Model(&Admission{}).Create(&ad1)
	ad2 := Admission{
		Patient:        pt2,
		PatientName:    "Somsri Sridang",
		Roomid:         "RM2005",
		RightTreatment: rt5,
	}
	db.Model(&Admission{}).Create(&ad2)
	ad3 := Admission{
		Patient:        pt3,
		PatientName:    "Sombut Sasbom",
		Roomid:         "RM2004",
		RightTreatment: rt6,
	}
	db.Model(&Admission{}).Create(&ad3)

	//TreatmentRecord

	tr1 := TreatmentRecord{
		Admission:   ad1,
		DoctorID:    "DR6001",
		Treatment:   "Heart Transplant",
		FoodType:    3000,
		MedID:       "MED6001",
		MedAmount:   3,
		EquipmentID: 002,
		Cost:        50000,
	}
	db.Model(&TreatmentRecord{}).Create(&tr1)

	tr2 := TreatmentRecord{
		Admission:   ad3,
		DoctorID:    "DR6001",
		Treatment:   "Gastric lavage",
		FoodType:    3001,
		MedID:       "MED6001",
		MedAmount:   3,
		EquipmentID: 001,
		Cost:        20000,
	}
	db.Model(&TreatmentRecord{}).Create(&tr2)

}
