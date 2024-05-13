package adapters

import (
	"fmt"
	"log"
	"net"

	order "github.com/udovichenk0/microservices/order/golang"
	"github.com/udovichenk0/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api  ports.PaymentPort
	port int
	order.UnimplementedPaymentServer
}

func NewAdapter(api ports.PaymentPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	order.RegisterPaymentServer(grpcServer, a)
	// if config.GetEnv() == "development" {
	reflection.Register(grpcServer)
	// }
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}
