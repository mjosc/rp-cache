package handlers

import (
	"github.com/mjosc/rp-cache/pkg/mocks/services/api/shared"
	"go.uber.org/fx"
)

type Handlers struct {
	fx.In
	Hello shared.Hello
}
