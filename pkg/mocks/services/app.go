package services

import (
	"github.com/mjosc/rp-cache/pkg/app"
	"github.com/mjosc/rp-cache/pkg/mocks/services/api"
	"github.com/mjosc/rp-cache/pkg/mocks/services/ui"
	"go.uber.org/fx"
)

func Register(config *app.Config) fx.Option {
	return fx.Options(
		api.Register(config),
		ui.Register(config),
	)
}
