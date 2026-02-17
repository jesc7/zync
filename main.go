package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/jesc7/zync/backend"
)

//go:embed all:frontend/dist/spa
var assets embed.FS

////go:embed build/appicon.png
//var icon []byte

func main() {
	flog, e := os.OpenFile(strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0]))+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln(e)
	}
	defer flog.Close()
	log.SetOutput(flog)

	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Zync",
		Width:            800,
		Height:           600,
		MinWidth:         800,
		MinHeight:        600,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Logger:        nil,
		LogLevel:      logger.DEBUG,
		OnStartup:     app.OnStartup,
		OnBeforeClose: app.OnBeforeClose,
		OnShutdown:    app.OnShutdown,
		Bind: []any{
			app,
			//&backend.MyData,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			//WebviewIsTransparent: false,
			//WindowIsTranslucent:  false,
			//DisableWindowIcon:    false,
			//WebviewUserDataPath:  "",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
