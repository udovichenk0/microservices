package ports

type Payment struct {
	price float64
}

func NewPayment(price float64) Payment {
	return Payment{price}
}
