package main

type ProductVM struct {
	Error string
	
	ProductId int
	Name string
	Description string
}

func ProductVMFromProduct(product *Product) *ProductVM {
	vm := new(ProductVM)
	
	vm.ProductId = product.ProductId
	vm.Name = product.Name
	vm.Description = product.Description
	
	return vm
}