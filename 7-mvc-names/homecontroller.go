package main

import (
	mvc "github.com/cihub/trinity"
	names "github.com/cihub/trinity-examples/7-mvc-names/mvcnames"
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
	controllerInfo.ActionInfos["IndexPost"].Method("POST").Action(names.A.Home.Index)
	return controllerInfo
}

func (controller *HomeController) Index() mvc.ActionResultInterface {
	return mvc.ShowView("", "", NewModel(""))
}

func (controller *HomeController) IndexPost(model *Model) mvc.ActionResultInterface {
	return mvc.ShowView("", "", NewModel(model.Name))
}