package main

import (
	"github.com/mjosc/rp-cache/pkg/app"
	"github.com/mjosc/rp-cache/pkg/handlers"
	mocks "github.com/mjosc/rp-cache/pkg/mocks/services"
	"github.com/mjosc/rp-cache/pkg/server"
)

func main() {
	config := &app.Config{
		ProxyServicePort: 8000,
		APIServicePort:   8100,
		UIServicePort:    8200,
	}

	app := app.New(
		config,
		handlers.Register,
		server.Register,
		mocks.Register,
	)

	app.Run()
}
