package backend

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	rtc "github.com/jesc7/zync/backend/rtc"
	signal "github.com/jesc7/zync/backend/signal"
	"github.com/jesc7/zync/backend/util"
	"github.com/pion/webrtc/v4"
)

type DataPart struct {
	val      webrtc.SessionDescription
	Key      string `json:"key"`
	Password string `json:"password"`
	e        error
}

type Data struct {
	Offer  DataPart `json:"offer"`
	Answer DataPart `json:"answer"`
}

func (o *Data) IsOfferReady() bool {
	return len(o.Offer.val.SDP) != 0 &&
		len(o.Offer.Key) != 0 &&
		len(o.Offer.Password) != 0
}

func (o *Data) IsAnswerReady() bool {
	return len(o.Answer.val.SDP) != 0
}

func (p *DataPart) IsError() (bool, string) {
	if p.e != nil {
		return true, p.e.Error()
	}
	return false, ""
}

func (o *Data) IsOfferError() (bool, string) {
	return o.Offer.IsError()
}

func (o *Data) IsAnswerError() (bool, string) {
	return o.Answer.IsError()
}

func (o *Data) Get() DataPart {
	return o.Offer
}

func (o *Data) Set(part DataPart) {
	o.Answer.Key = part.Key
	o.Answer.Password = part.Password
}

var (
	MyData Data
	Conn   *webrtc.PeerConnection
)

type Config struct {
	Signal struct {
		Addr string `json:"addr"`
	} `json:"signal"`
	Stuns []string `json:"stuns"`
}

type App struct {
	ctx context.Context
	cfg Config
	sig *signal.Client
}

func NewApp() *App {
	return &App{}
}

func (a *App) OnStartup(ctx context.Context) {
	var e error

	a.ctx = ctx
	a.cfg = Config{
		Stuns: []string{
			"stun:stun.l.google.com:19302",
		},
	}

	func() {
		if util.IsFileExists(filepath.Join(filepath.Dir(os.Args[0]), "cfg.json")) {
			f, e := os.ReadFile(filepath.Join(filepath.Dir(os.Args[0]), "cfg.json"))
			if e != nil {
				return
			}
			if e = json.Unmarshal(f, &a.cfg); e != nil {
				return
			}
		}
	}()

	a.sig, e = signal.NewClient(a.cfg.Signal.Addr)
	_ = e

	go func() {
		Conn, MyData.Offer.val, MyData.Offer.e = rtc.CreateOffer(a.cfg.Stuns)
	}()
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
