package backend

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Data struct {
	Offer    string
	Key      string
	Password string
	Time     int64
}

type App struct {
	ctx  context.Context
	data Data
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
			a.data.Time = v.Unix()
			runtime.EventsEmit(a.ctx, "changeTime")
		}
	}()
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
