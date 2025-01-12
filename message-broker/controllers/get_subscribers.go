package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubscribers(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Response struct{}

		subscribers := make([]Response, 0)

		c.JSON(http.StatusOK, subscribers)
	}
}
