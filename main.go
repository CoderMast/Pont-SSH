package main

import (
	"PontSsh/backend/constant"
	"PontSsh/backend/service"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed all:backend/database/pontssh.db
var database embed.FS

func main() {

	connection := service.NewConnection()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Pont SSH 连接工具",
		Width:     1024,
		Height:    720,
		MinWidth:  constant.MIN_WINDOW_WIDTH,
		MinHeight: constant.MIN_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  connection.Startup,
		OnShutdown: connection.Shutdown,
		Bind: []interface{}{
			connection,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
