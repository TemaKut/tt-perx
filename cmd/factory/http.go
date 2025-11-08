package factory

import (
	"context"
	"errors"
	"fmt"
	"github.com/TemaKut/tt-perx/internal/app/config"
	httphandler "github.com/TemaKut/tt-perx/internal/app/handlers/http/math"
	"github.com/TemaKut/tt-perx/internal/app/logger"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

var HttpSet = wire.NewSet(
	ProvideHttpProvider,
	ProvideHttpServer,
	httphandler.NewHandler,
)

type HttpProvider struct{}

func ProvideHttpProvider(_ *HttpServer) *HttpProvider {
	return &HttpProvider{}
}

type HttpServer struct{}

func ProvideHttpServer(cfg *config.Config, handler *httphandler.Handler, log *logger.Logger) (*HttpServer, func(), error) {
	server := echo.New()

	server.Use(
		middleware.Recover(),
		middleware.Logger(),
	)

	addHttpHandlers(server, handler)

	errCh := make(chan error, 1)

	go func() {
		log.Infof("starting http server at %s", cfg.Http.Addr)

		if err := server.Start(cfg.Http.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error start http server: %s", err)

			errCh <- fmt.Errorf("error starting http server: %w", err)
		}
	}()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	select {
	case err := <-errCh:
		return nil, nil, fmt.Errorf("error from errCh: %w", err)
	case <-ticker.C:
	}

	return &HttpServer{}, func() {
		log.Infof("stopping http server at %s", cfg.Http.Addr)

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Errorf("error shutdown http server: %s", err)
		}
	}, nil
}

func addHttpHandlers(server *echo.Echo, handler *httphandler.Handler) {
	server.POST("/math/arithmetic-progression/tasks/add", handler.HandleArithmeticProgressionTasksAdd)
	server.GET("/math/arithmetic-progression/tasks", handler.HandleArithmeticProgressionTasks)
}
