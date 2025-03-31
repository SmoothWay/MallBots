package monolith

import (
	"context"
	"database/sql"

	"github.com/SmoothWay/MallBots/internal/config"
	"github.com/SmoothWay/MallBots/internal/waiter"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Mux() *chi.Mux
	Waiter() waiter.Waiter
	Logger() zerolog.Logger
	RPC() *grpc.Server
}

type Module interface {
	Startup(context.Context, Monolith) error
}
