package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// The existing single-page application is embedded into the desktop binary so
// the Wails build does not require a separate frontend runtime or web server.
//
//go:embed index.html
var assets embed.FS

func main() {
	err := wails.Run(&options.App{
		Title:                            "Voz P2P",
		Width:                            960,
		Height:                           800,
		MinWidth:                         420,
		MinHeight:                        620,
		BackgroundColour:                 options.NewRGBA(11, 16, 32, 255),
		AssetServer:                      &assetserver.Options{Assets: assets},
		EnableDefaultContextMenu:         false,
		EnableFraudulentWebsiteDetection: false,
	})
	if err != nil {
		log.Fatal(err)
	}
}
