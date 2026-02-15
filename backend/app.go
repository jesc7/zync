package backend

import (
	"context"
	"time"
)

type Data struct {
	Offer    string
	Key      string
	Password string
	Time     time.Time
}

type App struct {
	ctx  context.Context
	data Data
}

func (a *App) SetTime(value time.Time) {
	a.data.Time = value
}

func NewApp() *App {
	return &App{}
}

// -------------- defaults -----------------

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for v := range t.C {
			a.SetTime(v)
		}
	}()
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
