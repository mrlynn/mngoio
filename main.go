package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrlynn/mngoio/config"
	"github.com/mrlynn/mngoio/storage"
	"github.com/mrlynn/mngoio/storage/mongodb"
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

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	router.GET("/:code", Redirect)
	router.Run()
}

// Index handles GETs sent to "/"
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

// Redirect handles POSTs sent to "/"
func Redirect(c *gin.Context) {
	code := c.Param("code")
	url, err := storage.GetURL(code)

	if err != nil {
		log.Println("ERROR", err)
		// c.Writer.WriteHeader(http.StatusNotFound)
		// c.Writer.Write([]byte("URL Not Found"))
		errMsg := "URL not Found"
		c.HTML(http.StatusOK, "index.tmpl.html", errMsg)
		return
	}

	http.Redirect(c.Writer, c.Request, url, http.StatusMovedPermanently)
}
