package main

import (
	mvc "github.com/cihub/trinity"
)

type HomeController struct {
	*mvc.BaseController
}

func NewHomeController() *HomeController {
	controller := new(HomeController)
	controller.BaseController = mvc.NewBaseController()
	return controller
}

func (controller *HomeController) GetInfo() mvc.ControllerInfoInterface {
	return mvc.NewToLowerControllerInfoExtracter(controller)
}

type HomeIndexInput struct {
	Restricted bool
}

func (controller *HomeController) Index(input *HomeIndexInput) mvc.ActionResultInterface {
	return mvc.ShowView("", "", 
		NewModel(authProvider.IsAuthenticated(controller.Request), input.Restricted))
}

func (controller *HomeController) Private() mvc.ActionResultInterface {
	return mvc.ShowView("", "", 
		NewModel(authProvider.IsAuthenticated(controller.Request), false))
}