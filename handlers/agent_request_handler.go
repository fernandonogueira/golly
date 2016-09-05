package handlers

import (
	"github.com/fernandonogueira/golly/models"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"log"
	"os"
	"bytes"
)

type AgentRequestHandler struct {
}

func NewRequestHandler() *AgentRequestHandler {
	return &AgentRequestHandler{
	}
}

func doRequest(request http.Request) models.AgentResponse {
	httpClient := http.Client{}

	response := models.AgentResponse{}

	resp, err := httpClient.Do(&request)
	defer resp.Body.Close()
	if (err != nil) {
		response.Result = "ERROR"
		log.Println(err)
	} else {
		response.Result = "OK"
		body, err := ioutil.ReadAll(resp.Body)
		if (err != nil) {
			log.Println(err)
		}
		strBody := string(body)
		response.Body = &strBody
		response.StatusCode = resp.StatusCode
	}

	return response;
}

func (r *AgentRequestHandler) Execute(request models.AgentRequest) models.AgentResponse {
	method := request.HttpMethod;

	var prepRequest *http.Request
	var err error

	if request.Body != "" {
		prepRequest, err = http.NewRequest(method, request.Url, bytes.NewBuffer([]byte(request.Body)))
	} else {
		prepRequest, err = http.NewRequest(method, request.Url, nil)
	}

	if len(request.Headers) > 0 {
		for k, v := range request.Headers {
			prepRequest.Header.Add(k, v)
		}
	}

	if err != nil {
		// handle error
	}

	start := time.Now()
	response := doRequest(*prepRequest)
	end := time.Now()

	took := end.UnixNano() - start.UnixNano()

	response.RequestStartEpoch = start.UnixNano() / 1000000
	response.RequestEndEpoch = end.UnixNano() / 1000000
	response.DurationMs = took / 1000000

	if (strings.EqualFold("OK", response.Result) && !request.AlwaysReturnBody) {
		response.Body = nil
	}

	assignRegionInfo(&response)
	assignToken(&request, &response)

	log.Println("Request finished.")
	return response
}
func assignToken(request *models.AgentRequest, response *models.AgentResponse) {
	if (request.Token != nil) {
		response.Token = request.Token
	}
}

func assignRegionInfo(response *models.AgentResponse) {
	regionInfo := models.RegionInfo{}
	regionInfo.Region = os.Getenv("REGION")

	response.RegionInfo = regionInfo
}