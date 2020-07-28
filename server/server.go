package main

import (
	"net"

	"github.com/go_proto/helper"
	"github.com/go_proto/services"
	"google.golang.org/grpc"
)

func main(){

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCerds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService)) //商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrdersService)) //订单服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService)) //用户服务
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)

}

