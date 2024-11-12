package service

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

type ApplicationService struct {
}

var Application *ApplicationService = &ApplicationService{}

func CreateService() *ApplicationService {
	return &ApplicationService{}
}

func (service *ApplicationService) Quit() {
	application.Get().Quit()
}

type Model struct {

}

func (service *ApplicationService) GetModel() Model {
	return Model{}
}
