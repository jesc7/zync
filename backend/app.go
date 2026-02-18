package backend

import (
	"context"
	"encoding/json"
	"log"
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

type Config struct {
	Signal struct {
		Addr string `json:"addr"`
	} `json:"signal"`
	Stuns []string `json:"stuns"`
}

type SignalStatus int
type OfferStatus int
type AnswerStatus int

const (
	SIGNAL_ERROR SignalStatus = iota - 1
	SIGNAL_OK
)

const (
	OFFER_ERROR OfferStatus = iota - 1
	OFFER_OK
	OFFER_CONNECTED
)

const (
	ANSWER_ERROR AnswerStatus = iota - 1
	ANSWER_OK
	ANSWER_CONNECTED
)

type App struct {
	ctx          context.Context
	cfg          Config
	sig          *signal.Client
	conn         *webrtc.PeerConnection
	MyData       Data
	statusSignal SignalStatus
	statusOffer  OfferStatus
	statusAnswer AnswerStatus
}

func NewApp() *App {
	return &App{}
}

func (a *App) onSignalOk() {
	a.statusSignal = SIGNAL_OK
}

func (a *App) onSignalError(e error) {
	a.statusSignal = SIGNAL_ERROR
	log.Println(e)
}

func (a *App) onOfferOk() {
	a.statusOffer = OFFER_OK

	if a.statusSignal == SIGNAL_OK {
		offer, _ := rtc.Encode(a.MyData.Offer.val)
		if a.MyData.Offer.Key, a.MyData.Offer.Password, a.MyData.Offer.e = a.sig.SendOffer(offer); a.MyData.Offer.e != nil {
			a.onOfferError(a.MyData.Offer.e)
		}
		log.Println(a.MyData.Offer.Key, a.MyData.Offer.Password)
	}
}

func (a *App) onOfferError(e error) {
	a.statusOffer = OFFER_ERROR
	log.Println(e)
}

func (a *App) onOfferConnected() {
	a.statusOffer = OFFER_CONNECTED
}

func (a *App) onAnswerOk() {
	a.statusAnswer = ANSWER_OK
}

func (a *App) onAnswerError(e error) {
	a.statusAnswer = ANSWER_ERROR
	log.Println(e)
}

func (a *App) onAnswerConnected() {
	a.statusAnswer = ANSWER_CONNECTED
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
		pwd, _ := os.Getwd()
		log.Println(pwd)
		if util.IsFileExists(filepath.Join(pwd, "cfg.json")) {
			f, e := os.ReadFile(filepath.Join(pwd, "cfg.json"))
			if e != nil {
				log.Println(e)
				return
			}
			if e = json.Unmarshal(f, &a.cfg); e != nil {
				log.Println(e)
				return
			}
		}
	}()

	if a.sig, e = signal.NewClient(a.ctx, a.cfg.Signal.Addr); e != nil {
		a.onSignalError(e)
	}

	go func() {
		if a.conn, a.MyData.Offer.val, a.MyData.Offer.e = rtc.CreateOffer(a.cfg.Stuns); a.MyData.Offer.e != nil {
			a.onOfferError(a.MyData.Offer.e)
			return
		}
		a.onOfferOk()
	}()
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
