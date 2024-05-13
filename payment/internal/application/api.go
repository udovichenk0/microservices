package application

import (
	"context"
	"log"

	"github.com/udovichenk0/microservices/payment/internal/application/domain"
)

type Application struct {
}

func NewApplication() Application {
	return Application{}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	log.Println("Charge user from api")
	return payment, nil
}
