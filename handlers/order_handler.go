package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/order-service/models"
)

func OrderHandler(c *gin.Context) {
	orderID := c.Param("id")
	order, err := models.GetOrder(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := order.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
