package controller

import (
	"github.com/N0hara/PaymentSystem/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /TreatmentRecord
// List all TreatmentRecord
func ListTreatmentRecord(c *gin.Context) {
	var treatmentrecord []entity.TreatmentRecord
	if err := entity.DB().Preload("Admission").Raw("SELECT * FROM treatment_records").Find(&treatmentrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}

// GET /TreatmentRecord/:id
// Get TreatmentRecord by id
func GetTreatmentRecord(c *gin.Context) {
	var treatmentrecord entity.TreatmentRecord
	id := c.Param("id")
	if err := entity.DB().Preload("Admission").Raw("SELECT * FROM treatment_records WHERE id = ?", id).Find(&treatmentrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}

// POST /TreatmentRecord
func CreateTreatmentRecord(c *gin.Context) {
	var treatmentrecord entity.TreatmentRecord
	if err := c.ShouldBindJSON(&treatmentrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&treatmentrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}

// GET /Admission/TreatmentRecord
func AdmissionByTreatmentRecord(context *gin.Context) {
	var treatmentrecord []entity.TreatmentRecord

	if err := entity.DB().Joins("Admission").
		Find(&treatmentrecord).Where("").Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "admission not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": treatmentrecord})
}
