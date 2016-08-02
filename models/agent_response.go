package models

type AgentResponse struct {
	Token      *string `json:"token"`
	DurationMs int64 `json:"durationMs"`
	Body       *string `json:"body"`
	Status     int `json:"status"`
	Result     string `json:"result"`
	RegionInfo RegionInfo `json:"regionInfo"`
}