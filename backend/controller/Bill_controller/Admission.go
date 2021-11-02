package controller

import (
	"github.com/N0hara/PaymentSystem/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /Admission
// List all Admission

func ListAdmission(c *gin.Context) {
	var admission []entity.Admission
	if err := entity.DB().Preload("Patient").Raw("SELECT * FROM admissions").Find(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// GET /Admission/:id
// Get Admission by id

func GetAdmission(c *gin.Context) {
	var admission entity.Admission
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Raw("SELECT * FROM admissions WHERE id = ?", id).Find(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// POST /Admission
func CreateAdmission(c *gin.Context) {
	var admission entity.Admission
	if err := c.ShouldBindJSON(&admission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&admission).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admission})
}

// GET /Patient/Admission
func PatientByAdmission(context *gin.Context) {
	var admission []entity.Admission

	if err := entity.DB().Joins("Patient").
		Find(&admission).Where("").Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": admission})
}
