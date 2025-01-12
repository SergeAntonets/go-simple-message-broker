package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTopic(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request struct {
			Name string `json:"name" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		b.AddTopic(request.Name)

		c.JSON(http.StatusOK, gin.H{"status": "Topic created"})
	}
}
