package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Level int

type LogConfig struct {
	Environment string
	LogLevel    Level
}

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	PANIC
)

const (
	TimeFormat = "03:04:05.000PM"
)

func New(cfg LogConfig) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	switch cfg.Environment {
	case "production":
		return zerolog.New(os.Stdout).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Logger()
	default:
		return zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = TimeFormat
		})).
			Level(logLevelToZero(cfg.LogLevel)).
			With().
			Timestamp().
			Logger()
	}
}

func logLevelToZero(level Level) zerolog.Level {
	switch level {
	case TRACE:
		return zerolog.TraceLevel
	case DEBUG:
		return zerolog.DebugLevel
	case INFO:
		return zerolog.InfoLevel
	case WARN:
		return zerolog.WarnLevel
	case ERROR:
		return zerolog.ErrorLevel
	case PANIC:
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}
