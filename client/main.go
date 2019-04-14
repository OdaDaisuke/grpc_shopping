package main

import (
	"github.com/OdaDaisuke/grpc_shopping/client/handlers"
	"net/http"
	"fmt"
	"github.com/OdaDaisuke/grpc_shopping/client/gateways"
)

func main() {
	grpcConn := gateways.InitGrpcConn()
	defer grpcConn.Close()

	handlerFactory := handlers.NewHandlerFactory(grpcConn)

	http.HandleFunc("/items", handlerFactory.ItemsHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}