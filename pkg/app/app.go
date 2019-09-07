package app

import (
	"go.uber.org/fx"
)

type Config struct {
	ProxyServicePort int
	APIServicePort   int
	UIServicePort    int
}

type FXRegistrationFunc func(*Config) fx.Option

func New(config *Config, registrations ...FXRegistrationFunc) *App {
	options := make([]fx.Option, 0, len(registrations))
	for _, f := range registrations {
		options = append(options, f(config))
	}
	return &App{
		inner: fx.New(options...),
	}
}

type App struct {
	inner *fx.App
}

func (app *App) Run() {
	app.inner.Run()
}
