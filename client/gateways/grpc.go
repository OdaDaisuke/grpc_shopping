package gateways

import (
	"google.golang.org/grpc"
	"log"
)

func InitGrpcConn() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection err: ", err)
	}
	return conn
}
