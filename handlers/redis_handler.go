package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/order-service/models"
)

func RedisHandler(c *gin.Context) {
	client, err := models.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	data, err := client.Get(context.Background(), "key").Bytes()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, string(data))
}
