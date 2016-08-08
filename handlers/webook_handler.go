package handlers

import (
	"github.com/fernandonogueira/golly/models"
	"log"
	"net/http"
	"bytes"
	"encoding/json"
)

type WebhookHandler struct {
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
	}
}

func (w *WebhookHandler) NotifyEndpoint(request *models.AgentRequest, response *models.AgentResponse) {
	if (request.WebhookAddress != nil) {
		log.Println("Sending webhook..." + *request.WebhookAddress)

		jsonResponse, err := json.Marshal(response)

		if err != nil {
			// handle error
		}

		request, err := http.NewRequest("POST", *request.WebhookAddress, bytes.NewBuffer(jsonResponse))
		request.Header.Add("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(request)

		if err != nil {
			// handle error
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}
	}
}