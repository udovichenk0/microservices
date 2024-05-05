package grpc

import (
	"context"
	"fmt"

	order "github.com/udovichenk0/microservices/order/golang"
	"github.com/udovichenk0/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   float32(orderItem.UnitPrice),
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	err = status.Errorf(codes.InvalidArgument, fmt.Sprintf("failed to charge user: %d", request.UserId))
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

// should be in separate repo
func (a *PaymentAdapter) Charge(o *domain.Order) error {
	//payment stub
	_, err := a.payment.Create(context.Background(),
		//this order is payment. order because it lays in golang folder with order generated files
		&order.CreatePaymentRequest{
			UserId:     o.CustomerID,
			OrderId:    o.ID,
			TotalPrice: o.TotalPrice(),
		})
	return err
}
