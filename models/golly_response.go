package models

type GollyResponse struct {
	Token             *string `json:"token"`
	DurationMs        int64 `json:"durationMs"`
	RequestStartEpoch int64 `json:"requestStartEpoch"`
	RequestEndEpoch   int64 `json:"requestEndEpoch"`
	Body              *string `json:"body"`
	StatusCode        int `json:"statusCode"`
	Result            string `json:"result"`
	RegionInfo        RegionInfo `json:"regionInfo"`
}