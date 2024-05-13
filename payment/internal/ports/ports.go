package ports

import (
	"context"

	"github.com/udovichenk0/microservices/payment/internal/application/domain"
)

type PaymentPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
