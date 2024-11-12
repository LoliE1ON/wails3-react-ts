package config

import (
	service2 "github.com/LoliE1ON/wails3-react-ts/domain/application/service"
)

type ServicesConfig struct {
	Application *service2.ApplicationService
}

var services = ServicesConfig{
	Application: service2.CreateService(),
}

func GetServices() ServicesConfig {
	return services
}
