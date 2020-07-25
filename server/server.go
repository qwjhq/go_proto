package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/go_proto/server/services"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

func main(){
	//单向证书方式
	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	//双向证书方式
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ClientAuth: tls.RequestClientCert,
		ClientCAs: certPool,
	})

	//1、无证书
	// rpcServer := grpc.NewServer()
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	//2、tcp的服务方式
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
	/*mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request.Proto",request.Proto)
		fmt.Println("request",request)
		rpcServer.ServeHTTP(writer, request)
	})
	 httpServer := &http.Server{
		Addr: ":8081",
		Handler: mux,
	}
	 */
	// http的方式
	//err = httpServer.ListenAndServe()
	/// https的方式
	///err = httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_passwd.key")
	///if err != nil {
	///	log.Fatal(err)
	///}
}

