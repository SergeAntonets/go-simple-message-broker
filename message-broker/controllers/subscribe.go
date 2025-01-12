package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Subscribe(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Topic       string `json:"topic" binding:"required"`
			CallbackUrl string `json:"callback_url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		sub := &broker.Subscriber{
			Topic:       broker.Topic{Name: request.Topic},
			CallbackUrl: request.CallbackUrl,
		}
		b.Subscribe(sub)

		c.JSON(http.StatusOK, gin.H{"status": "Subscribed to topic"})
	}
}
