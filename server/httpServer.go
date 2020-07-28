package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/go_proto/helper"
	"github.com/go_proto/services"
)

func main() {
	gwmux := runtime.NewServeMux()
	grpcEndPoint := "localhost:8081"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCerds())}
	if err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, grpcEndPoint, opts ); err != nil {
		log.Fatal(err)
	}

	if err := services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, grpcEndPoint, opts ); err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr: ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
