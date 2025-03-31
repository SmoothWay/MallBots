package notifications

import (
	"context"

	"github.com/SmoothWay/MallBots/internal/monolith"
	"github.com/SmoothWay/MallBots/notifications/internal/application"
	"github.com/SmoothWay/MallBots/notifications/internal/grpc"
	"github.com/SmoothWay/MallBots/notifications/internal/logging"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	defer conn.Close()

	customers := grpc.NewCustomerRepository(conn)

	var app application.App
	app = application.NewApplication(customers)
	app = logging.LogApplicationAccess(app, mono.Logger())

	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	return nil
}
