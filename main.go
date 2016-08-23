package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/fernandonogueira/golly/models"
	"github.com/fernandonogueira/golly/handlers"
	"strconv"
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

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status":"UP"})
	})

	router.POST("/syncAnalysis", func(c *gin.Context) {
		agentRequest := models.AgentRequest{}
		c.Bind(&agentRequest)
		response := requestHandler.Execute(agentRequest);
		c.JSON(http.StatusOK, response)
	})

	router.POST("/printResponse", func(c *gin.Context) {
		log.Println("Hello!")
		agentResponse := models.AgentResponse{}
		c.Bind(&agentResponse)
		log.Println("status: " + strconv.Itoa(agentResponse.Status))
		log.Println("duration: " + strconv.FormatInt(agentResponse.DurationMs, 10))
	})

	router.POST("/analyze", func(c *gin.Context) {
		agentRequest := models.AgentRequest{}
		err := c.Bind(&agentRequest)

		if err != nil {
			log.Println("Body content invalid. Err: " + err.Error())
			validationErr := models.ErrorResponse{Error:"Args.Invalid"}
			c.JSON(http.StatusBadRequest, validationErr)
			return
		}

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
