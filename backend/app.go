package backend

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Source struct {
	Path      string  `json:"path"`
	Tables    []Table `json:"tables"`
	ProfileID int     `json:"profile"`
	Agents    bool    `json:"agents"`
}

type Table struct {
	Name              string       `json:"name"`
	RecordsCount      int          `json:"rec_count,omitempty"`
	RecordsCountTotal int          `json:"rec_count_total,omitempty"`
	Dependencies      []Dependency `json:"dependencies,omitempty"`
	Success           bool         `json:"success,omitempty"`
	Seconds           float64      `json:"seconds,omitempty"`
}

type Dependency struct {
	Name string `json:"name"`
}

type Destination struct {
	Path         string  `json:"path"`
	Tables       []Table `json:"tables"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Pather interface {
	path() string
	setPath(string)
}

type App struct {
	ctx context.Context
	src Source
	dst Destination
}

func NewApp() *App {
	return &App{}
}

func (a *Source) path() string {
	return a.Path
}

func (a *Source) setPath(s string) {
	a.Path = s
}

func (a *App) Src() Source {
	return a.src
}

func (a *App) SetSrc(v Source) {
	a.src = v
}

func (a *App) srcChanged() {
	runtime.EventsEmit(a.ctx, "src.changed", a.src)
}

func (a *Destination) path() string {
	return a.Path
}

func (a *Destination) setPath(s string) {
	a.Path = s
}

func (a *App) Dst() Destination {
	return a.dst
}

func (a *App) SetDst(v Destination) {
	a.dst = v
}

func (a *App) dstChanged() {
	runtime.EventsEmit(a.ctx, "dst.changed", a.dst)
}

func (a *App) selectFile(v Pather, f func()) {
	s, e := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select a file",
	})
	if e == nil && v.path() != s {
		v.setPath(s)
		if f != nil {
			f()
		}
	}
}

func (a *App) SelectSrc() {
	a.selectFile(&a.src, a.srcChanged)
}

func (a *App) SelectDst() {
	a.selectFile(&a.dst, a.dstChanged)
}

// func (a *App) selectSrcDragnDrop() {
// 	//drag'n'drop file
// 	var (
// 		e error
// 		s string
// 	)
// 	if e == nil && a.src.Path != s {
// 		a.src.Path = s
// 		a.srcChanged()
// 	}
// }

// -------------- defaults -----------------

func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) OnShutdown(ctx context.Context) {
}
