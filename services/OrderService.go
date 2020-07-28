package services

import (
	"context"
	"fmt"
)

type OrdersService struct {

}

//func (this *OrdersService) 	NewOrder(ctx context.Context, orderMain *OrderMain) (*OrderResponse, error) {
func (this *OrdersService) 	NewOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	fmt.Println(request.OrderMain)
	return &OrderResponse{Status:"OK",Message:"success"}, nil
}
