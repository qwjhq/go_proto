package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go_proto/client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//单向证书
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt","adams.com")
	//双向证书方式
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //客户端证书
		ServerName: "adams.com",
		RootCAs: certPool,
	})
	//1、无证书时请求
	//conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := services.NewProdServiceClient(conn)
	res, err := client.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId:10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.ProdStock)
}
