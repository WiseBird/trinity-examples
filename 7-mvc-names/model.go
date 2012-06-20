package main

import (
	names "github.com/cihub/trinity-examples/7-mvc-names/mvcnames"
)

type Model struct {
	C *names.ControllerNames
	A *names.ActionNames
	Name string
}

func NewModel(name string) *Model {
	model := new(Model)
	
	model.Name = name
	model.C = names.C
	model.A = names.A
	
	return model
}