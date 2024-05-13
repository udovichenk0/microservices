package adapters

import (
	"context"
	"fmt"
	"log"

	order "github.com/udovichenk0/microservices/order/golang"
	"github.com/udovichenk0/microservices/payment/internal/application/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, r *order.CreatePaymentRequest) (*order.CreatePaymentResponse, error) {
	log.Println("HELLO")
	newPayment := domain.NewPayment(r.UserId, r.OrderId, r.TotalPrice)
	res, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &order.CreatePaymentResponse{PaymentId: int64(res.TotalPrice)}, nil
}
