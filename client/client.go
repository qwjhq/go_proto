package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go_proto/services"
	"google.golang.org/grpc"
	"github.com/go_proto/helper"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCerds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx := context.Background()
	client := services.NewProdServiceClient(conn)
	res, err := client.GetProdStock(ctx,
		&services.ProdRequest{ProdId:10, ProdArea:services.ProdAreas_A},)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.ProdStock)
	fmt.Println("get ProdStock list")
	response, err := client.GetProdStocks(ctx, &services.QuerySize{Size:10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.ProdRes)
	fmt.Println(response.ProdRes[2].ProdStock)
}
