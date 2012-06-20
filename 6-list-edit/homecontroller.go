package main

import (
	mvc "github.com/cihub/trinity"
	"sync"
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
	controllerInfo.ActionInfos["EditPost"].Method("POST").Action("edit")
	return controllerInfo
}

func (controller *HomeController) List() mvc.ActionResultInterface {
	return mvc.ShowView("", "", products)
}

func (controller *HomeController) Create() mvc.ActionResultInterface {
	return mvc.ShowView("", "edit", ProductVMFromProduct(new(Product)))
}

type HomeEditInput struct {
	ProductId int
}
func (controller *HomeController) Edit(input *HomeEditInput) mvc.ActionResultInterface {
	product := controller.findProduct(input.ProductId)
	if product == nil {
		product = new(Product)
	}
	
	return mvc.ShowView("", "", ProductVMFromProduct(product))
}

func (controller *HomeController) EditPost(productVM *ProductVM) mvc.ActionResultInterface {
	if len(productVM.Name) == 0 {
		productVM.Error = "Name can't be empty"
		return mvc.ShowView("", "", productVM)
	}
	
	productToUpdate := controller.findProduct(productVM.ProductId)
	isNewProduct := false
	if productToUpdate == nil {
		isNewProduct = true
		productAddSync.Lock()
		defer productAddSync.Unlock()
		
		productToUpdate = &Product { len(products) + 1, "", "" }
		products = append(products, productToUpdate)
	}
	
	productToUpdate.Name = productVM.Name
	productToUpdate.Description = productVM.Description
	
	if isNewProduct {
		return mvc.RedirectToAction("", "list", nil)
	}
	
	return mvc.ShowView("", "", ProductVMFromProduct(productToUpdate))
}

func (controller *HomeController) findProduct(productId int) *Product {
	for _, product := range products {
		if product.ProductId == productId {
			return product
		}
	}
	
	return nil
}