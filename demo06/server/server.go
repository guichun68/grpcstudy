package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcstudy/demo06/server/services"
	"io/ioutil"
	"net"
)

func main() {
	//tls 是 ssl 的升级版
	cert, _ := tls.LoadX509KeyPair("cert/server.pem","cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},//服务端证书
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp",":8081")
	rpcServer.Serve(lis)
}
