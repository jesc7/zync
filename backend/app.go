package backend

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DataPart struct {
	Value    string `json:"value"`
	Key      string `json:"key"`
	Password string `json:"password"`
}

type Data struct {
	Offer  DataPart `json:"offer"`
	Answer DataPart `json:"answer"`
}

func (o *Data) IsAnswerer() bool {
	return len(o.Answer.Key) != 0 && len(o.Answer.Password) == 0
}

func (o *Data) Get() DataPart {
	return o.Offer
}

func (o *Data) Set(val DataPart) {
	o.Answer.Key = val.Key
	o.Answer.Password = val.Password
}

var MyData Data

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
