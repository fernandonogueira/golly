package handlers

import (
	"github.com/fernandonogueira/golly/models"
)

type RequestValidator struct {

}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{}
}

func (v *RequestValidator) Validate(request models.AgentRequest, async bool) models.ErrorResponse {

	errorResponse := models.ErrorResponse{}

	if async {
		if request.WebhookAddress == nil {
			errorResponse.Error = "Webhook address is mandatory"
		}
	}

	return errorResponse
}