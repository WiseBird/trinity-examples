package main

import (
	mvc "github.com/cihub/trinity"
	names "github.com/cihub/trinity-examples/8-authentication-authorization/mvcnames"
)

type AccountController struct {
	*mvc.BaseController
}

func NewAccountController() *AccountController {
	controller := new(AccountController)
	controller.BaseController = mvc.NewBaseController()
	return controller
}

func (controller *AccountController) GetInfo() mvc.ControllerInfoInterface {
	return mvc.NewToLowerControllerInfoExtracter(controller)
}

func (controller *AccountController) Login() mvc.ActionResultInterface {
	authProvider.Login(controller.Response, controller.Request)
	
	return mvc.RedirectToAction(names.C.Home, names.A.Home.Index, nil)
}

func (controller *AccountController) Logout() mvc.ActionResultInterface {
	authProvider.Logout(controller.Response, controller.Request)
	
	return mvc.RedirectToAction(names.C.Home, names.A.Home.Index, nil)
}