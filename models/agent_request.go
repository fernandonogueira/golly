package models

type AgentRequest struct {
	Url string `json:"url"`
	CheckId int64 `json:"checkId"`
	Headers map[string]string `json:"headers"`
	Protocol string `json:"protocol"`
	HttpMethod string `json:"httpMethod"`
	ExpectedStatusCode int `json:"expectedStatusCode"`
	BodyContains string `json:"bodyContains"`
	Body string `json:"body"`
	Token string `json:"token"`
}