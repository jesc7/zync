package backend

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type OfferData struct {
	Value    string `json:"value"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

func (o *OfferData) Get() OfferData {
	return *o
}

func (o *OfferData) Set(val OfferData) {
	*o = val
}

var Offer OfferData

type App struct {
	ctx context.Context
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
			//a.data.Time = v.Unix()
			_ = v
			runtime.EventsEmit(a.ctx, "changeTime")
		}
	}()
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
