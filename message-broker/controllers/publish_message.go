package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublishMessageRequest struct {
	Topic   string `json:"topic" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}

func PublishMessage(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request PublishMessageRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
			return
		}

		msg := broker.Message{
			Topic:   request.Topic,
			Payload: request.Payload,
		}

		b.Publish(msg)
		c.JSON(http.StatusOK, gin.H{"status": "Message published"})
	}
}
