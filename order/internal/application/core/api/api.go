package api

import (
	"log"

	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"github.com/ruandg/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payment  ports.PaymentPort
	shipping ports.ShippingPort 
}


func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:       db,
		payment:  payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	
	var totalItems int32
	for _, item := range order.OrderItems {
		totalItems += item.Quantity
	}
	if totalItems > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Order cannot have more than 50 items in total.")
	}

	
	for _, item := range order.OrderItems {
		
		exists, err := a.db.CheckStock(item.ProductCode)
		if err != nil {
			return domain.Order{}, status.Errorf(codes.Internal, "Failed to verify stock for product %s.", item.ProductCode)
		}
		if !exists {
			
			return domain.Order{}, status.Errorf(codes.NotFound, "Product %s does not exist in stock.", item.ProductCode)
		}
	}

	
	order.Status = "Pending"
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		if status.Code(paymentErr) == codes.DeadlineExceeded {
			log.Println("[LOG] A chamada ao serviço de pagamento falhou por estouro de tempo (Timeout).")
		}

		order.Status = "Canceled"
		_ = a.db.Save(&order) 
		return domain.Order{}, paymentErr
	}

	
	shippingErr := a.shipping.CreateShipment(&order)
	if shippingErr != nil {
		
		log.Printf("[LOG] Erro ao agendar entrega via gRPC para o pedido %d: %v", order.ID, shippingErr)
	}

	
	order.Status = "Paid"
	err = a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}