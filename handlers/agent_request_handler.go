package handlers

import (
	"github.com/fernandonogueira/golly/models"
	"github.com/dghubble/sling"
	"net/http"
	"io/ioutil"
	"time"
	"strings"
	"log"
	"os"
)

type AgentRequestHandler struct {
}

func NewRequestHandler() *AgentRequestHandler {
	return &AgentRequestHandler{
	}
}

func doRequest(builder *sling.Sling) models.AgentResponse {
	httpClient := http.Client{}
	httpRequest, rErr := builder.Request()
	if (rErr != nil) {
		log.Println(rErr.Error())
	}

	response := models.AgentResponse{}

	resp, err := httpClient.Do(httpRequest)
	if (err != nil) {
		response.Result = "ERROR"
		log.Println(err)
	} else {
		defer resp.Body.Close()
		response.Result = "OK"
		body, err := ioutil.ReadAll(resp.Body)
		if (err != nil) {
			log.Println(err)
		}
		strBody := string(body)
		response.Body = &strBody
		response.Status = resp.StatusCode
	}

	return response;
}

func (r *AgentRequestHandler) Execute(request models.AgentRequest) models.AgentResponse {
	method := request.HttpMethod;

	builder := sling.New()

	switch method {
	case "GET":
		builder = builder.Get(request.Url)
		break
	case "POST":
		builder = builder.Post(request.Url)
		break
	case "HEAD":
		builder = builder.Head(request.Url)
	default:
		return models.AgentResponse{}
	}

	start := time.Now().UnixNano()
	response := doRequest(builder)
	end := time.Now().UnixNano()

	took := end - start

	response.DurationMs = took / 1000000

	if (strings.EqualFold("OK", response.Result) && !request.AlwaysReturnBody) {
		response.Body = nil
	}

	assignRegionInfo(&response)

	return response
}

func assignRegionInfo(response *models.AgentResponse) {
	regionInfo := models.RegionInfo{}
	regionInfo.Region = os.Getenv("REGION")

	response.RegionInfo = regionInfo
}