package handlers

import (
	"github.com/mjosc/rp-cache/pkg/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	APIProxy shared.APIProxy
	UIProxy  shared.UIProxy
}
