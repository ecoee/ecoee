package main

import (
	"context"
	"ecoee/pkg/config"
	"ecoee/pkg/ecoee/domain/service"
	"ecoee/pkg/ecoee/infrastructure/cloudstorage"
	db2 "ecoee/pkg/ecoee/infrastructure/db"
	"ecoee/pkg/ecoee/infrastructure/db/mongo"
	"ecoee/pkg/ecoee/infrastructure/gemini"
	"ecoee/pkg/ecoee/presentation/rest/assessment"
	"ecoee/pkg/ecoee/presentation/rest/campaign"
	"ecoee/pkg/ecoee/presentation/rest/health"
	"ecoee/pkg/ecoee/presentation/rest/organization"
	"ecoee/pkg/ecoee/presentation/rest/point"
	"ecoee/pkg/ecoee/presentation/rest/user"
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	ctx := context.Background()
	config := config.NewConfig(viper.New())
	// init infrastructure layer
	db, err := mongo.NewDB(ctx, config)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to connect to database: %v", errors.WithStack(err)))
		return
	}

	cloudStroageRepository, err := cloudstorage.NewRepository(ctx, config)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create cloud storage repository %v", errors.WithStack(err)))
		return
	}
	pointRepository := db2.NewPointRepository(db)
	userRepository := db2.NewUserRepository(db)
	organizationRepository := db2.NewOrganizationRepository(db)
	assessRepository, err := gemini.NewRepository(ctx, config)
	campaignRepository := db2.NewCampaignRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create assess repository %v", errors.WithStack(err)))
		return
	}

	// init domain layer
	pointService := service.NewPointService(pointRepository, userRepository)

	// init presentation layer
	healthRegistry := health.NewRegistry()
	userRegistry := user.NewRegistry(userRepository, organizationRepository, pointService)
	organizationRegistry := organization.NewRegistry(organizationRepository, pointService)
	assessmentRegistry := assessment.NewRegistry(assessRepository)
	pointRegistry := point.NewRegistry(userRepository, organizationRepository, pointRepository)
	campaignRegistry := campaign.NewRegistry(campaignRepository, userRepository, cloudStroageRepository)

	//init gin
	r := gin.New()
	r.Use(gin.Recovery())

	healthRegistry.Register(r)
	userRegistry.Register(r)
	organizationRegistry.Register(r)
	assessmentRegistry.Register(r)
	pointRegistry.Register(r)
	campaignRegistry.Registry(r)
	serverPort := ":8080"

	srv := &http.Server{
		Addr:         serverPort,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      h2c.NewHandler(r, &http2.Server{}), // h2c enables HTTP/2 without TLS
	}

	// Handle SIGINT(CTRL+C), SIGTERM gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srvErr := make(chan error, 1)
	go func() {
		slog.Info(fmt.Sprintf("starting server on port %s", serverPort))
		srvErr <- srv.ListenAndServe()
	}()
	// Wait for interruption.
	select {
	case err := <-srvErr:
		// Error when starting HTTP server.
		slog.Error(fmt.Sprintf("failed to start server: %v", err))
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		slog.Info("shutting down server")
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error(fmt.Sprintf("failed to shutdown server: %v", err))
		return
	}

	slog.Info("server shutdown successfully")
}
