package routers

import (
	"message-broker/broker"
	"message-broker/controllers"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine, b *broker.Broker) {
	api := router.Group("/api/v1")
	{
		api.POST("/topics", controllers.CreateTopic(b))
		api.POST("/publish", controllers.PublishMessage(b))
		api.POST("/subscribe", controllers.Subscribe(b))

		api.GET("/subscribers", controllers.GetSubscribers(b))
		api.GET("/topics", controllers.GetTopics(b))
	}
}
