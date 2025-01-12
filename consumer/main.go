package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/callback", func(c *gin.Context) {

		data, _ := io.ReadAll(c.Request.Body)

		fmt.Printf("Callback received: %s", string(data))

		if rand.Float64() < 0.5 { // Simulate 70% failure rate
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Something went wrong..."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Ack"})
	})

	if err := r.Run(":8090"); err != nil {
		panic(err)
	}

}
