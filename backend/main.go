package main

import (
	"github.com/N0hara/PaymentSystem/controller"

	"github.com/N0hara/PaymentSystem/entity"

	"github.com/N0hara/PaymentSystem/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	/*gin.SetMode(gin.ReleaseMode)*/
	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		p := api.Use(middlewares.Authorizes())
		{
			// FinancialOfficer Routes
			p.GET("/financialofficers", controller.ListFinancialOfficer)
			p.GET("/financialofficer/:id", controller.GetFinancialOfficer)
			p.POST("/financialofficers", controller.CreateFinancialOfficer)
			//p.PATCH("/financialofficers", controller.UpdateFinancialOfficer)
			//p.DELETE("/financialofficer/:id", controller.DeleteFinancialOfficer)

			// RightTreatment Routes
			p.GET("/righttreatments", controller.ListRightTreatment)
			p.GET("/righttreatment/:id", controller.GetRightTreatment)
			p.POST("/righttreatments", controller.CreateRightTreatment)
			//p.PATCH("/righttreatments", controller.UpdateRightTreatment)
			//p.DELETE("/righttreatment/:id", controller.DeleteRightTreatment)

			// PaymentMethod Routes
			p.GET("/paymentmethods", controller.ListPaymentMethod)
			p.GET("/paymentmethod/:id", controller.GetPaymentMethod)
			p.POST("/paymentmethods", controller.CreatePaymentMethod)
			//p.PATCH("/paymentmethods", controller.UpdatePaymentMethod)
			//p.DELETE("/paymentmethod/:id", controller.DeletePaymentMethod)

			// patient Routes
			p.GET("/patients", controller.ListPatient)
			p.GET("/patient/:id", controller.GetPatient)
			p.POST("/patients", controller.CreatePatient)
			//p.PATCH("/patients", controller.UpdatePatient)
			//p.DELETE("/patient/:id", controller.DeletePatient)

			// admission Routes
			p.GET("/admissions", controller.ListAdmission)
			p.GET("/admission/:id", controller.GetAdmission)
			p.POST("/admissions", controller.CreateAdmission)
			//p.PATCH("/admissions", controller.UpdateAdmission)
			//p.DELETE("/admission/:id", controller.DeleteAdmission)
			p.GET("/patient/admissions", controller.PatientByAdmission)
			//PatientByAdmission

			// TreatmentRecord Routes
			p.GET("/treatmentrecords", controller.ListTreatmentRecord)
			p.GET("/treatmentrecord/:id", controller.GetTreatmentRecord)
			p.POST("/treatmentrecords", controller.CreateTreatmentRecord)
			//p.PATCH("/treatmentrecords", controller.UpdateTreatmentRecord)
			//p.DELETE("/treatmentrecord/:id", controller.DeleteTreatmentRecord)
			p.GET("/admission/treatmentrecords", controller.AdmissionByTreatmentRecord)
			//AdmissionByTreatmentRecord

			// bill Routes
			p.GET("/bills", controller.ListBill)
			p.GET("/bill/:id", controller.GetBill)
			p.POST("/bills", controller.CreateBill)
			//p.PATCH("/bills", controller.UpdateBill)
			//p.DELETE("/bill/:id", controller.DeleteBill)

		}
	}

	// Run the server

	// User Routes
	//r.POST("/FinancialOfficers/Creat", controller.CreateFinancialOfficer)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}

}
