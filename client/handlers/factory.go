package handlers

import (
	"google.golang.org/grpc"
	"github.com/OdaDaisuke/grpc_shopping/pb/item"
	"net/http"
	"encoding/json"
	"context"
	"log"
	"errors"
)

type HandlerFactory struct {
	itemsClient item.ItemsClient
}

func NewHandlerFactory(conn *grpc.ClientConn) *HandlerFactory {
	return &HandlerFactory{
		itemsClient: item.NewItemsClient(conn),
	}
}

func (h *HandlerFactory) ItemsHandler(w http.ResponseWriter, r *http.Request) {
	page := getQueryParam(w, r, "page")
	res, err := h.itemsClient.GetAll(context.TODO(), &item.GetAllMessage{
		Page: page,
	})
	if err != nil {
		Err500(w, err)
	}

	js, err := json.Marshal(res)
	if err != nil {
		Err500(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

/*------------- utils ---------------*/

func getQueryParam(w http.ResponseWriter, r *http.Request, key string) string {
	keys, ok := r.URL.Query()[key]
	if !ok || len(keys[0]) < 1 {
		Err500(w, errors.New("param page is not valid"))
	}
	return string(keys[0])
}

func Err500(w http.ResponseWriter, err error) {
	log.Fatal(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}