package grpc

import (
	"fmt"
	"log"
	"net"

	order "github.com/udovichenk0/microservices/order/golang"
	"github.com/udovichenk0/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.OrderPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.OrderPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	// if config.GetEnv() == "development" {
	reflection.Register(grpcServer)
	// }
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

// in other repo
type PaymentAdapter struct {
	payment order.PaymentClient
}

func NewPaymentAdapter(paymentServiceUrl string) (*PaymentAdapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := order.NewPaymentClient(conn)
	return &PaymentAdapter{payment: client}, nil
}
