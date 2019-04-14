package main

import (
	"fmt"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"database/sql"
	"github.com/OdaDaisuke/grpc_shopping/server/configs"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/OdaDaisuke/grpc_shopping/server/middlewares"
	"github.com/OdaDaisuke/grpc_shopping/server/service"
	"github.com/OdaDaisuke/grpc_shopping/pb/user"
	"github.com/OdaDaisuke/grpc_shopping/pb/item"
)

func main() {
	listenPort, err := net.Listen("tcp", configs.PORT)
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
			grpc_recovery.StreamServerInterceptor(),
		)),
	)
	serverWithAdmin := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middlewares.Auth),
		)),
	)

	itemService := &service.ItemService{}
	userService := &service.UserService{}

	item.RegisterItemsServer(server, itemService)
	user.RegisterUsersServer(serverWithAdmin, userService)

	fmt.Printf("Start server on port http://localhost%s\n", configs.PORT)
	server.Serve(listenPort)
}