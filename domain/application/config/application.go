package config

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var Application = application.Options{
	Name:        "Application",
	Description: "Application",
	Services:    []application.Service{},
	Assets:      application.AssetOptions{},
}
