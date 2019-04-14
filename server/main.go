package main

import (
	"fmt"
	pb "github.com/OdaDaisuke/grpc_shopping/pb"
	"github.com/OdaDaisuke/grpc_shopping/server/service"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/OdaDaisuke/grpc_shopping/server/middlewares"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"database/sql"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("mysql", "daisuke:password@/shopping")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middlewares.Auth),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)
	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
	))
	itemService := &service.ItemService{}

	pb.RegisterItemsServer(server, itemService)
	fmt.Println("Start server on port http://localhost:19003")
	server.Serve(listenPort)
}