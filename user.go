package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nutreet/common"
	gen "github.com/nutreet/common/gen/user"
	"google.golang.org/grpc"
)

func RunUserGatewayServer() {
	grpcMux := runtime.NewServeMux()
	//nolint:staticcheck // ignore SA1019
	err := gen.RegisterUserServiceHandlerFromEndpoint(context.Background(), grpcMux, fmt.Sprintf("localhost:%s", common.Constants.USER_GRPC_PORT), []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", common.Constants.API_GATEWAY_PORT), // The port to listen on
		Handler: grpcMux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
