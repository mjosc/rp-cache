package handlers

import (
	"github.com/mjosc/rp-cache/pkg/mocks/services/ui/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	Hello shared.Hello
}
