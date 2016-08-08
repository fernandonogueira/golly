package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/fernandonogueira/golly/models"
	"github.com/fernandonogueira/golly/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	region := os.Getenv("REGION")

	if region == "" {
		log.Fatal("$REGION must be set")
	}

	requestHandler := handlers.NewRequestHandler()
	webhookHandler := handlers.NewWebhookHandler()
	requestValidator := handlers.NewRequestValidator()

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/syncAnalysis", func(c *gin.Context) {
		agentRequest := models.AgentRequest{}
		c.Bind(&agentRequest)
		response := requestHandler.Execute(agentRequest);
		c.JSON(http.StatusOK, response)
	})

	router.POST("/", func(c *gin.Context) {
		log.Println("Hello!")
	})

	router.POST("/analyze", func(c *gin.Context) {
		agentRequest := models.AgentRequest{}
		c.Bind(&agentRequest)

		validationErrors := requestValidator.Validate(agentRequest, true)

		if validationErrors.Error != "" {
			log.Println("Validation errors found")
			c.JSON(http.StatusBadRequest, validationErrors)
			return
		}

		go func() {
			response := requestHandler.Execute(agentRequest)
			webhookHandler.NotifyEndpoint(&agentRequest, &response)
		}()
		c.Done()
	})

	router.Run(":" + port)
}
