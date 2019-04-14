package main

import (
	"google.golang.org/grpc"
	"log"
	"context"
	pb "github.com/OdaDaisuke/grpc_shopping/pb"
	"fmt"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection err: ", err)
	}
	defer conn.Close()

	client := pb.NewItemsClient(conn)
	message := &pb.GetAllMessage{
		Page: "0",
	}

	res, err := client.GetAll(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request response: %v", res)

}