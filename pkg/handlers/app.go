package handlers

import (
	"fmt"

	"github.com/mjosc/rp-cache/pkg/app"
	"github.com/mjosc/rp-cache/pkg/shared"
	"go.uber.org/fx"
)

var configuration *app.Config

func Register(config *app.Config) fx.Option {
	configuration = config
	return fx.Option(
		fx.Provide(
			newAPIProxy,
			newUIProxy,
		),
	)
}

func newAPIProxy() (shared.APIProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.APIServicePort)
	return NewAPIProxy(dst)
}

func newUIProxy() (shared.UIProxy, error) {
	dst := fmt.Sprintf("http://localhost:%v", configuration.UIServicePort)
	return NewUIProxy(dst)
}
