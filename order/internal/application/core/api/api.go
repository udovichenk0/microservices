package api

import (
	"log"
	"strings"

	"github.com/udovichenk0/microservices/order/internal/application/core/domain"
	"github.com/udovichenk0/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db, payment}
}

func (app *Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// err := app.db.Save(&order)
	// if err != nil {
	// 	return domain.Order{}, err
	// }
	// calling grpc stub(payment.Charge)
	paymentErr := app.payment.Charge(&order)
	log.Println(paymentErr)
	if paymentErr != nil {
		// st, _ := status.FromError(paymentErr)
		st := status.Convert(paymentErr)
		var allErrors []string
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}

		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
	}
	return order, nil
}
