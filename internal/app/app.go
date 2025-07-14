package app

import (
	"fmt"
	"log/slog"
	"os"
	"test/internal/config"
	v1 "test/internal/controller/rest/v1"
	psg "test/internal/infra/postgres"
	services "test/internal/usecase"
	"test/pkg/httpserver"
	"test/pkg/postgres"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	s   *httpserver.Server
	db  *postgres.Postgres
	log *slog.Logger
}

func Run(
	log *slog.Logger,
	cfg *config.Config,
) *HttpServer {
	// Database
	pg, err := postgres.New(cfg.Database.URL, postgres.MaxPoolSize(cfg.Database.PoolMax))
	if err != nil {
		log.Error(fmt.Errorf("app - Run - postgres.New: %w", err).Error())
		os.Exit(1)
	}

	// Repository
	subRepo := psg.NewSubscriptionRepository(pg)

	// Services
	subService := services.NewSubscriptionService(log, subRepo)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(log, handler, subService)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	return &HttpServer{s: httpServer, db: pg, log: log}
}

func (s *HttpServer) Shutdown() {
	defer s.db.Close()
	err := s.s.Shutdown()
	if err != nil {
		s.log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}
}
