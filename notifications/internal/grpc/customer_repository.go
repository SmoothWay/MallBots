package grpc

import (
	"context"

	"github.com/SmoothWay/MallBots/notifications/internal/application"
	"github.com/SmoothWay/MallBots/notifications/internal/models"
	"google.golang.org/grpc"
)

type CustomerRepository struct {
	client customerspb.CustomerServiceClient
}

var _ application.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(conn *grpc.ClientConn) *CustomerRepository {
	return &CustomerRepository{
		client: customerspb.NewCustomerServiceClient(conn),
	}
}

func (r CustomerRepository) Find(ctx context.Context, customerId string) (*models.Customer, error) {
	resp, err := r.client.GetCustomer(ctx, &customerspb.GetCustomerRequest{Id: customerId})
	if err != nil {
		return nil, err
	}
	return &models.Customer{
		ID:        resp.GetCustomer().GetId(),
		Name:      resp.GetCustomer().GetFirstName(),
		SmsNumber: resp.GetCustomer().GetSmsNumber(),
	}, nil
}
