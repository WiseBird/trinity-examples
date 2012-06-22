package main

import (
	names "github.com/cihub/trinity-examples/10-list-edit-ajax/mvcnames"
)

type Model struct {
	C *names.ControllerNames
	A *names.ActionNames
}

func NewModel() *Model {
	model := new(Model)
	
	model.C = names.C
	model.A = names.A
	
	return model
}