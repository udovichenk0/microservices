package main

import (
	"github.com/udovichenk0/microservices/payment/internal/adapters"
	"github.com/udovichenk0/microservices/payment/internal/application"
)

func main() {
	app := application.NewApplication()
	adapters := adapters.NewAdapter(app, 3001)

	adapters.Run()
}
