package main

import (
	"log"

	"github.com/udovichenk0/microservices/order/config"
	"github.com/udovichenk0/microservices/order/internal/adapters/db"
	"github.com/udovichenk0/microservices/order/internal/adapters/grpc"
	"github.com/udovichenk0/microservices/order/internal/adapters/payment"
	"github.com/udovichenk0/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter := db.NewDBAdapter(config.GetDataSourceURL())
	paymentAdapter, err := payment.NewPaymentAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}
	app := api.NewApplication(dbAdapter, paymentAdapter)
	// port := config.GetEnv()
	grpcAdapter := grpc.NewAdapter(app, 3000)

	grpcAdapter.Run()
}
