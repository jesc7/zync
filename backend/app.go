package backend

import (
	"context"
)

type Data struct {
	Offer  string
	Answer string
}

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// -------------- defaults -----------------

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
