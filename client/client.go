package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go_proto/services"
	"github.com/golang/protobuf/ptypes/timestamp"
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
	var i int32
	userClient := services.NewUserServiceClient(conn)
	fmt.Println("get User message by server/client bothway stream")
	bothStream, err := userClient.GetUserScoreByTWS(ctx)
	var uid int32 = 1
	for j := 1; j <= 3; j++ {
		clientReq := services.UserScoreRequest{}
		clientReq.Users =  make([]*services.UserInfo, 0)
		for i = 1; i < 20; i ++ { //假设这是一个耗时的过程
			clientReq.Users = append(clientReq.Users, &services.UserInfo{UserId:uid})
			uid++
		}
		err := bothStream.Send(&clientReq)
		if err != nil {
			log.Println(err)
		}
		res, err := bothStream.Recv()
		if err == io.EOF{break}
		if err != nil {
			log.Println(err)
		}
		fmt.Println("result", res.Users)
	}

	fmt.Println("get User message by client stream")
	clientStream, err := userClient.GetUserScoreByClientStream(ctx)
	for j := 1; j <= 3; j++ {
		clientReq := services.UserScoreRequest{}
		clientReq.Users =  make([]*services.UserInfo, 0)
		for i = 1; i < 20; i ++ { //假设这是一个耗时的过程
			clientReq.Users = append(clientReq.Users, &services.UserInfo{UserId:i})
		}
		err := clientStream.Send(&clientReq)
		if err != nil {
			log.Println(err)
		}
	}
	clientStreamRes,_ := clientStream.CloseAndRecv()
	fmt.Println("result ",clientStreamRes.Users)


	fmt.Println("get User message")
	userReq := services.UserScoreRequest{}
	userReq.Users =  make([]*services.UserInfo, 0)
	for i = 1; i < 20; i ++ {
		userReq.Users = append(userReq.Users, &services.UserInfo{UserId:i})
	}
	// 一次性全部接收完成
	userRes, _ := userClient.GetUserScore(ctx, &userReq)
	fmt.Println("result ", userRes.Users)

	fmt.Println("get User message by server stream")
	userStreamRes, err := userClient.GetUserScoreByServerStream(ctx, &userReq)
	if err != nil {log.Fatal(err)}
	for {
		streamRes, err := userStreamRes.Recv()
		if err == io.EOF {break} //是否结束
		if err != nil {log.Fatal(err)} //是否有错
		fmt.Println("result ",streamRes.Users)
	}

	fmt.Println("get Order message")
	orderClient := services.NewOrderServiceClient(conn)
	t := timestamp.Timestamp{Seconds:time.Now().Unix()}
	/*orderInfo, err := orderClient.NewOrder(ctx, &services.OrderMain{
		OrderId:     1001,
		OrderNo:     "20190809 ",
		OrderMoneny: 100,
		OrderTime:   &t,
	})*/
	orderInfo, err := orderClient.NewOrder(ctx,&services.OrderRequest{OrderMain: &services.OrderMain{
		OrderId:     1001,
		OrderNo:     "20190809",
		OrderMoney: 100,
		OrderTime:   &t,
	},
	} )
	fmt.Println("result", orderInfo)

	client := services.NewProdServiceClient(conn)
	fmt.Println("get ProdStock info")
	info, err := client.GetProdInfo(ctx,
		&services.ProdRequest{ProdId:10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
	res, err := client.GetProdStock(ctx,
		&services.ProdRequest{ProdId:10, ProdArea:services.ProdAreas_A},)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result ",res.ProdStock)

	fmt.Println("get ProdStock list")
	response, err := client.GetProdStocks(ctx, &services.QuerySize{Size:10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result ",response.ProdRes)
	fmt.Println("result ",response.ProdRes[2].ProdStock)
}
