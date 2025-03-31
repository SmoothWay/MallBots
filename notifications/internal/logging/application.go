package logging

import (
	"context"

	"github.com/SmoothWay/MallBots/notifications/internal/application"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(app application.App, logger zerolog.Logger) application.App {
	return Application{
		App:    app,
		logger: logger,
	}
}

func (a Application) NotifyOrderCreated(ctx context.Context, notify application.OrderCreated) error {
	var err error
	a.logger.Info().Msgf("--> Notifications.NotifyOrderCreated: %s", notify.OrderID)
	defer func() {
		a.logger.Info().Err(err).Msgf("%s <-- Notifications.NotifyOrderCreated", notify.OrderID)
	}()
	err = a.App.NotifyOrderCreated(ctx, notify)
	return err
}

func (a Application) NotifyOrderCanceled(ctx context.Context, notify application.OrderCanceled) error {
	var err error
	a.logger.Info().Msgf("--> Notifications.NotifyOrderCanceled: %s", notify.OrderID)
	defer func() {
		a.logger.Info().Msgf("%s <-- Notifications.NotifyOrderCanceled", notify.OrderID)
	}()
	err = a.App.NotifyOrderCanceled(ctx, notify)
	return err
}

func (a Application) NotifyOrderReady(ctx context.Context, notify application.OrderReady) error {
	var err error
	a.logger.Info().Msg("Notifications.NotifyOrderReady")
	defer func() {
		a.logger.Info().Msgf("%s <-- NotifyOrderReady", notify.OrderID)
	}()
	err = a.App.NotifyOrderReady(ctx, notify)
	return err
}
