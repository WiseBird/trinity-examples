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

func (controller *HomeController) Index() mvc.ActionResultInterface {
	return mvc.ShowView("", "", &Model { "John", &AddressModel { "City Center Plaza 516 Main St.", "Elgin", "OR", "USA" } })
}
	