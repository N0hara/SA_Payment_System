package controller

import (
	"github.com/N0hara/PaymentSystem/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /FinancialOfficer
// List all FinancialOfficer
func ListFinancialOfficer(c *gin.Context) {
	var financialofficer []entity.FinancialOfficer
	if err := entity.DB().Raw("SELECT * FROM financial_officers").Scan(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// GET /FinancialOfficer/:id
// Get FinancialOfficer by id
func GetFinancialOfficer(c *gin.Context) {
	var financialofficer entity.FinancialOfficer
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM financial_officers WHERE id = ?", id).Scan(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// POST /FinancialOfficer
func CreateFinancialOfficer(c *gin.Context) {
	var financialofficer entity.FinancialOfficer
	if err := c.ShouldBindJSON(&financialofficer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// PATCH /FinancialOfficer
func UpdateFinancialOfficer(c *gin.Context) {
	var financialofficer entity.FinancialOfficer
	if err := c.ShouldBindJSON(&financialofficer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", financialofficer.ID).First(&financialofficer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "financialofficer not found"})
		return
	}

	if err := entity.DB().Save(&financialofficer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": financialofficer})
}

// DELETE /FinancialOfficer/:id
func DeleteFinancialOfficer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM financial_officers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FinancialOfficer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
