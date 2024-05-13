package payment

import (
	"context"
	"log"

	"github.com/sony/gobreaker"
	order "github.com/udovichenk0/microservices/order/golang"
	"github.com/udovichenk0/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment order.PaymentClient
}

func CircuitBreakerClientInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		_, cbErr := cb.Execute(func() (interface{}, error) {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil {
				return nil, err
			}
			return nil, nil
		})
		return cbErr
	}
}

func NewPaymentAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption

	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "payment",
		MaxRequests: 3,
		Timeout:     4,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio > 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("CircuitBreaker %s changed from %s to %s", name, from, to)
		},
	})

	opts = append(opts, grpc.WithUnaryInterceptor(CircuitBreakerClientInterceptor(cb)))

	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := order.NewPaymentClient(conn)

	return &Adapter{client}, nil
}

func (a Adapter) Charge(o *domain.Order) error {
	log.Println(o.CustomerID, o.ID, o.TotalPrice())
	_, err := a.payment.Create(context.Background(), &order.CreatePaymentRequest{
		UserId:     o.CustomerID,
		OrderId:    o.ID,
		TotalPrice: o.TotalPrice(),
	})
	return err
}
