package logging

import (
	"context"

	"github.com/rs/zerolog"
)

type Application struct {
	app    application.App
	logger *zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(app application.App, logger *zerolog.Logger) application.App {
	return &Application{
		app:    app,
		logger: logger,
	}
}

func (a Application) RegisterCustomer(ctx context.Context, register application.RegisterCustomer) error {
	a.logger.Info().Msgf("--> Customers.RegisterCustomer: %s", register.CustomerID)
	err := a.app.RegisterCustomer(ctx, register)
	if err != nil {
		a.logger.Info().Err(err).Msgf("%s <-- Customers.RegisterCustomer", register.CustomerID)
	}
	return err
}

func (a Application) AuthorizeCustomer(ctx context.Context, authorize application.AuthorizeCustomer) error {
	a.logger.Info().Msgf("--> Customers.AuthorizeCustomer: %s", authorize.CustomerID)
	err := a.app.AuthorizeCustomer(ctx, authorize)
	if err != nil {
		a.logger.Info().Err(err).Msgf("%s <-- Customers.AuthorizeCustomer", authorize.CustomerID)
	}
	return err
}

func (a Application) GetCustomer(ctx context.Context, get application.GetCustomer) (*application.Customer, error) {
	a.logger.Info().Msgf("--> Customers.GetCustomer: %s", get.CustomerID)
	customer, err := a.app.GetCustomer(ctx, get)
	if err != nil {
		a.logger.Info().Err(err).Msgf("%s <-- Customers.GetCustomer", get.CustomerID)
	}
	return customer, err
}

func (a Application) EnableCustomer(ctx context.Context, enable application.EnableCustomer) error {
	a.logger.Info().Msgf("--> Customers.EnableCustomer: %s", enable.CustomerID)
	err := a.app.EnableCustomer(ctx, enable)
	if err != nil {
		a.logger.Info().Err(err).Msgf("%s <-- Customers.EnableCustomer", enable.CustomerID)
	}
	return err
}

func (a Application) DisableCustomer(ctx context.Context, disable application.DisableCustomer) error {
	a.logger.Info().Msgf("--> Customers.DisableCustomer: %s", disable.CustomerID)
	err := a.app.DisableCustomer(ctx, disable)
	if err != nil {
		a.logger.Info().Err(err).Msgf("%s <-- Customers.DisableCustomer", disable.CustomerID)
	}
	return err
}
