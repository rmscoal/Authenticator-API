package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/rmscoal/Authenticator-API/config"
	v1 "github.com/rmscoal/Authenticator-API/internal/controller/http/v1"
	"github.com/rmscoal/Authenticator-API/internal/usecase"
	"github.com/rmscoal/Authenticator-API/internal/usecase/repo"
	"github.com/rmscoal/Authenticator-API/pkg/httpserver"
	"github.com/rmscoal/Authenticator-API/pkg/logger"
	"github.com/rmscoal/Authenticator-API/pkg/postgres"
	"github.com/rmscoal/Authenticator-API/pkg/tokenizer"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Use case
	userUseCase := usecase.New(
		repo.New(pg),
	)

	// Tokenizer
	t := tokenizer.New(tokenizer.MinCost(cfg.Token.MinCost))

	// HTTP Server
	handler := echo.New()
	v1.NewRouter(handler, l, userUseCase, *t)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
