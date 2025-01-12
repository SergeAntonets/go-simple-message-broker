package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishMessage(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg broker.Message
		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
			return
		}
		b.Publish(msg)
		c.JSON(http.StatusOK, gin.H{"status": "Message published"})
	}
}
