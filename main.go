package main

import (
	"./config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.GetConfigFromJSON("config.json")

	if err != nil {
		log.Fatal(err)
	}

	client, err := mongodb.NewMongoClient(cfg.Mongo.URI)

	if err != nil {
		log.Fatal(err)
	}

	repository := mongodb.NewMongoRepository(cfg.Mongo.DB, cfg.Mongo.Collection, client)

	storage.SetStorage(repository)

	router := gin.Default()
	router.GET("/", Index)
	router.Run()
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
