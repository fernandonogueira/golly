package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/fernandonogueira/golly/models"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/syncAnalysis", func(c *gin.Context) {
		agentRequest := models.AgentRequest{}
		c.Bind(&agentRequest)
		c.JSON(http.StatusOK, agentRequest)
	})

	router.Run(":" + port)
}
