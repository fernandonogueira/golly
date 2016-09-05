package models

type GollyRequest struct {
	/**
	URL to be analysed
	 */
	Url                string `json:"url"`
	/**
	Headers to be sent
	 */
	Headers            map[string]string `json:"headers"`
	/**
	Protocol (for now, only HTTP is supported)
	 */
	Protocol           string `json:"protocol"`
	/**
	Http method to be used
	 */
	HttpMethod         string `json:"httpMethod"`
	/**
	Expected status code for response
	 */
	ExpectedStatusCode int `json:"expectedStatusCode"`
	/**
	Body must have this content
	 */
	BodyContains       string `json:"bodyContains"`
	/**
	Request body
	 */
	Body               string `json:"body"`
	/**
	Request token. Needed for async analysis
	 */
	Token              *string `json:"token"`
	/**
	Always returns body even if status code == 200
	 */
	AlwaysReturnBody   bool `json:"alwaysReturnBody"`
	/**
	Enables analysis if responde != 200
	 */
	AnalysisEnabled    bool `json:"analysisEnabled"`
	/**
	WebhookAddress to send result to
	 */
	WebhookAddress     *string `json:"webhookAddress"`
}