package main

import (
	"log"

	"github.com/nutreet/common"
)

func main() {
	log.Println("Starting API Gateway server on port " + common.Constants.API_GATEWAY_PORT + "...")
	RunUserGatewayServer()
}
