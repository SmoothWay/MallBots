package grpc

import (
	"context"

	"github.com/SmoothWay/MallBots/notifications/internal/application"
	"github.com/SmoothWay/MallBots/notifications/notificationspb"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	notificationspb.UnimplementedNotificationServiceServer
}

var _ notificationspb.NotificationServiceServer = (*server)(nil)

func RegisterServer(_ context.Context, app application.App, registrar grpc.ServiceRegistrar) error {
	notificationspb.RegisterNotificationServiceServer(registrar, &server{app: app})
	return nil
}

func (s *server) NotifyOrderCreated(ctx context.Context, req *notificationspb.NotifyOrderCreatedRequest) (*notificationspb.NotifyOrderCreatedResponse, error) {
	err := s.app.NotifyOrderCreated(ctx, application.OrderCreated{
		OrderID:    req.GetOrderId(),
		CustomerID: req.GetCustomerId()})
	return &notificationspb.NotifyOrderCreatedResponse{}, err
}

func (s *server) NotifyOrderCanceled(ctx context.Context, req *notificationspb.NotifyOrderCanceledRequest) (*notificationspb.NotifyOrderCanceledResponse, error) {
	err := s.app.NotifyOrderCanceled(ctx, application.OrderCanceled{
		OrderID:    req.GetOrderId(),
		CustomerID: req.GetCustomerId()})
	return &notificationspb.NotifyOrderCanceledResponse{}, err
}

func (s *server) NotifyOrderReady(ctx context.Context, req *notificationspb.NotifyOrderReadyRequest) (*notificationspb.NotifyOrderReadyResponse, error) {
	err := s.app.NotifyOrderReady(ctx, application.OrderReady{
		OrderID:    req.GetOrderId(),
		CustomerID: req.GetCustomerId()})
	return &notificationspb.NotifyOrderReadyResponse{}, err
}
