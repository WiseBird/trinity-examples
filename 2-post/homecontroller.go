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
	controllerInfo := mvc.NewToLowerControllerInfoExtracter(controller)
	controllerInfo.ActionInfos["IndexPost"].Method("POST").Action("index")
	return controllerInfo
}

func (controller *HomeController) Index() mvc.ActionResultInterface {
	return mvc.ShowView("", "", "")
}

func (controller *HomeController) IndexPost(model *Model) mvc.ActionResultInterface {
	return mvc.ShowView("", "", model.Name)
}