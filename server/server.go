package main

import (
	"net"

	"github.com/go_proto/helper"
	"github.com/go_proto/services"
	"google.golang.org/grpc"
)

func main(){

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCerds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)

}

