package main

import (
	mvc "github.com/cihub/trinity"
	"sync"
	names "github.com/cihub/trinity-examples/10-list-edit-ajax/mvcnames"
)

var (
	products = make([]*Product, 0)
	productAddSync sync.Mutex
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
	controllerInfo.ActionInfos["EditPost"].Method("POST").Action(names.A.Home.Edit)
	return controllerInfo
}

func (controller *HomeController) Index() mvc.ActionResultInterface {
	return mvc.ShowView("", "", NewModel())
}

func (controller *HomeController) List() mvc.ActionResultInterface {
	return mvc.JsonResult(products)
}

type HomeEditInput struct {
	ProductId int
}
func (controller *HomeController) Edit(input *HomeEditInput) mvc.ActionResultInterface {
	product := controller.findProduct(input.ProductId)
	if product == nil {
		product = new(Product)
	}
	
	return mvc.JsonResult(ProductVMFromProduct(product))
}

func (controller *HomeController) EditPost(productVM *ProductVM) mvc.ActionResultInterface {
	if len(productVM.Name) == 0 {
		productVM.Error = "Name can't be empty"
		return mvc.JsonResult(productVM)
	}
	
	productToUpdate := controller.findProduct(productVM.ProductId)
	if productToUpdate == nil {
		productAddSync.Lock()
		defer productAddSync.Unlock()
		
		productToUpdate = &Product { len(products) + 1, "", "" }
		products = append(products, productToUpdate)
	}
	
	productToUpdate.Name = productVM.Name
	productToUpdate.Description = productVM.Description
	
	return mvc.JsonResult(ProductVMFromProduct(productToUpdate))
}

func (controller *HomeController) findProduct(productId int) *Product {
	for _, product := range products {
		if product.ProductId == productId {
			return product
		}
	}
	
	return nil
}