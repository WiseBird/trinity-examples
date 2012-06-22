package main

import (
	names "github.com/cihub/trinity-examples/8-authentication-authorization/mvcnames"
)

type Model struct {
	C *names.ControllerNames
	A *names.ActionNames
	
	Authenticated bool
	Restricted bool
}

func NewModel(authenticated bool, restricted bool) *Model {
	model := new(Model)
	
	model.Authenticated = authenticated
	model.Restricted = restricted
	model.C = names.C
	model.A = names.A
	
	return model
}