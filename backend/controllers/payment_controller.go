package controllers

import (
	"net/http"

	"bus-booking/utils"
	"github.com/gin-gonic/gin"
)

// InitiatePayment mocks payment initiation and returns a payment ID
func InitiatePayment(c *gin.Context) {
	// In real app, integrate with payment gateway here
	paymentID := "PAY123456789"
	c.JSON(http.StatusOK, gin.H{
		"payment_id": paymentID,
		"status":     "initiated",
	})
}

// GetPaymentStatus mocks payment status retrieval
func GetPaymentStatus(c *gin.Context) {
	paymentID := c.Param("payment_id")
	// In real app, query payment gateway for status
	c.JSON(http.StatusOK, gin.H{
		"payment_id": paymentID,
		"status":     "success",
	})
}
