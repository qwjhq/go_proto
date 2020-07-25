package helper

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"google.golang.org/grpc/credentials"
)

//获取服务端证书
func GetServerCerds() credentials.TransportCredentials{
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ClientAuth: tls.RequestClientCert,
		ClientCAs: certPool,
	})
	return creds

}

//获取客户端证书
func GetClientCerds() credentials.TransportCredentials {
	//双向证书方式
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")

	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //客户端证书
		ServerName: "adams.com",
		RootCAs: certPool,
	})
	return creds
}