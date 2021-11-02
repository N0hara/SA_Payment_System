package controller

import (
	"github.com/N0hara/PaymentSystem/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /RightTreatment
// List all RightTreatment
func ListRightTreatment(c *gin.Context) {
	var righttreatment []entity.RightTreatment
	if err := entity.DB().Raw("SELECT * FROM right_treatments").Scan(&righttreatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": righttreatment})
}

// GET /RightTreatment/:id
// Get RightTreatment by id
func GetRightTreatment(c *gin.Context) {
	var righttreatment entity.RightTreatment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM right_treatments WHERE id = ?", id).Scan(&righttreatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": righttreatment})
}

// POST /RightTreatment
func CreateRightTreatment(c *gin.Context) {
	var righttreatment entity.RightTreatment
	if err := c.ShouldBindJSON(&righttreatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&righttreatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": righttreatment})
}

// PATCH /RightTreatment
func UpdateRightTreatment(c *gin.Context) {
	var righttreatment entity.RightTreatment
	if err := c.ShouldBindJSON(&righttreatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", righttreatment.ID).First(&righttreatment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RightTreatment not found"})
		return
	}

	if err := entity.DB().Save(&righttreatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": righttreatment})
}

// DELETE /RightTreatment/:id
func DeleteRightTreatment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM right_treatments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RightTreatment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
