package ui

import (
	"github.com/mjosc/rp-cache/pkg/app"
	"github.com/mjosc/rp-cache/pkg/mocks/services/ui/handlers"
	"github.com/mjosc/rp-cache/pkg/mocks/services/ui/server"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		handlers.Register(config),
		server.Register(config),
	)
}
