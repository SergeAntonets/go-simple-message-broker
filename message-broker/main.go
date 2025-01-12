package main

import (
	"fmt"
	"message-broker/broker"
	"message-broker/database"
	"message-broker/routers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client := &http.Client{}

	db := database.Create()
	defer db.Close()

	database.CreateTopics(db)

	bufferSize := 10
	b := broker.NewBroker(bufferSize)
	b.StartWorkers(5, client)

	engine := gin.Default()
	engine.Use(cors.Default())

	routers.SetApiRouter(engine, b)

	b.AddTopic("topic1")

	sub := &broker.Subscriber{
		Topic:       broker.Topic{Name: "topic1"},
		CallbackUrl: "http://localhost:8090/callback",
	}
	b.Subscribe(sub)

	for i := 0; i < 100; i++ {
		msg := broker.Message{
			Topic:   "topic1",
			Payload: fmt.Sprintf("Payload-%d", i),
		}
		b.Publish(msg)
	}

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}

}
