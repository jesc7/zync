package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/jesc7/zync/backend"
)

//go:embed all:frontend/dist/spa
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "ZFix UI",
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
