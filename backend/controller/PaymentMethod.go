package controller

import (
	"github.com/N0hara/PaymentSystem/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /PaymentMethod
// List all PaymentMethod
func ListPaymentMethod(c *gin.Context) {
	var paymentmethod []entity.PaymentMethod
	if err := entity.DB().Raw("SELECT * FROM payment_methods").Scan(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// GET /PaymentMethod/:id
// Get PaymentMethod by id
func GetPaymentMethod(c *gin.Context) {
	var paymentmethod entity.PaymentMethod
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM payment_methods WHERE id = ?", id).Scan(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// POST /PaymentMethod
func CreatePaymentMethod(c *gin.Context) {
	var paymentmethod entity.PaymentMethod
	if err := c.ShouldBindJSON(&paymentmethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// PATCH /PaymentMethod
func UpdatePaymentMethod(c *gin.Context) {
	var paymentmethod entity.PaymentMethod
	if err := c.ShouldBindJSON(&paymentmethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", paymentmethod.ID).First(&paymentmethod); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentMethod not found"})
		return
	}

	if err := entity.DB().Save(&paymentmethod).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentmethod})
}

// DELETE /PaymentMethod/:id
func DeletePaymentMethod(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payment_methods WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentMethod not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
