package ports

import "github.com/udovichenk0/microservices/order/internal/application/core/domain"

type OrderPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
