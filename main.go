package main

import (
	"embed"
	_ "embed"
	config2 "github.com/LoliE1ON/wails3-react-ts/domain/application/config"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"log"
)

//go:embed renderer/dist
var assets embed.FS

func main() {
	var services = config2.GetServices()

	config2.Application.Assets.Handler = application.AssetFileServerFS(assets)
	config2.Application.Services = []application.Service{
		application.NewService(services.Application),
	}

	var instance = application.New(config2.Application)

	instance.On(events.Common.ApplicationStarted, func(event *application.Event) {
		//
	})

	instance.NewWebviewWindowWithOptions(config2.Window)

	var err = instance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
