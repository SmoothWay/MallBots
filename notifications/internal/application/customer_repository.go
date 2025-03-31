package application

import (
	"context"

	"github.com/SmoothWay/MallBots/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, id string) (*models.Customer, error)
}
