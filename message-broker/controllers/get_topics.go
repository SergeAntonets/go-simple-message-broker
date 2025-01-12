package controllers

import (
	"message-broker/broker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTopics(b *broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Response struct {
			Name string `json:"name"`
		}

		topics := make([]Response, 0)
		// for _, topicList := range b.topics {
		// 	for _, topic := range topicList {
		// 		topics = append(topics, Response{Name: topic.name})
		// 	}
		// }

		c.JSON(http.StatusOK, topics)
	}
}
