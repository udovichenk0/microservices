package ports

import "github.com/udovichenk0/microservices/order/internal/application/core/domain"

type OrderPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}

//in other github repo should be smth like this
//but i use monolith for simplicity
// import github.com/udovichenko/grpc/order/internal/application/core/domain

type PaymentPort interface {
	Charge(*domain.Order) error
}
