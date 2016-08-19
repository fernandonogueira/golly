package handlers

import (
	"github.com/aeden/traceroute"
	"log"
	"strconv"
)

type TraceRouteHandler struct {

}

func NewTraceRouteHandler() *TraceRouteHandler {
	return &TraceRouteHandler{
	}
}

func (t *TraceRouteHandler) TraceRoute(url string) {
	opts := traceroute.TracerouteOptions{
	}
	opts.SetTimeoutMs(30000)
	opts.SetMaxHops(10)
	opts.SetRetries(2)
	result, err := traceroute.Traceroute(url, &opts)

	if err != nil {
		log.Println("Error found during traceroute. Err: " + err.Error())
	}

	hops := result.Hops
	for v := 0; v < len(hops); v++ {
		log.Println("ElapsedTimeMs: " + strconv.FormatInt(hops[v].ElapsedTime.Nanoseconds() / 1000000, 10))
		log.Println(hops[v].HostOrAddressString() + " (" + hops[v].AddressString() + ")")
	}

}
