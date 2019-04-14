package service

import (
	pb "github.com/OdaDaisuke/grpc_shopping/pb"
	"context"
	"fmt"
)

type ItemService struct {
}

func (s *ItemService) GetAll(ctx context.Context, in *pb.GetAllMessage) (*pb.GetAllResponse, error) {
	var items []*pb.Item
	items = append(items, &pb.Item{
		Name: "itemA",
	})

	fmt.Println("Page is ", in.GetPage())
	return &pb.GetAllResponse{
		Items: items,
	}, nil
}

/*------------ mock ---------------*/
