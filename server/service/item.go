package service

import (
	"github.com/OdaDaisuke/grpc_shopping/pb/item"
	"context"
	"fmt"
)

type ItemService struct {
}

func (s *ItemService) GetAll(ctx context.Context, in *item.GetAllMessage) (*item.GetAllResponse, error) {
	var items []*item.Item
	items = append(items, &item.Item{
		Name: "itemA",
	})

	fmt.Println("Page is ", in.GetPage())
	return &item.GetAllResponse{
		Items: items,
	}, nil
}

/*------------ mock ---------------*/
