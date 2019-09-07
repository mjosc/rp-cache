package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mjosc/rp-cache/pkg/app"
	"github.com/mjosc/rp-cache/pkg/mocks/services/ui/handlers"
	"go.uber.org/fx"
)

var configuration *app.Config

func Register(config *app.Config) fx.Option {
	configuration = config
	return fx.Options(
		fx.Invoke(
			setup,
		),
	)
}

func setup(lc fx.Lifecycle, handlers handlers.Handlers) {
	mux := http.NewServeMux()
	mux.Handle("/", handlers.Hello)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", configuration.UIServicePort),
		Handler: mux,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					if err := server.ListenAndServe(); err != nil {
						if err == http.ErrServerClosed {
							log.Println("Server stopped")
						} else {
							log.Println(err, "Error shutting down server")
						}
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				if err := server.Shutdown(ctx); err != nil {
					log.Println(err, "Error shutting down server")
				}
				return nil
			},
		},
	)
}
