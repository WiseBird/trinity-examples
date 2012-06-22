package mvcnames

import (
	mvc "github.com/cihub/trinity"
)

var (
	C = newControllerNames()
	A = newActionNames()
	V = newViewNames()
)

type ControllerNames struct {
	Home     mvc.Controller
	Shared mvc.Controller
}

func newControllerNames() *ControllerNames {
	return &ControllerNames{"home", "shared"}
}

type ActionNames struct {
	Home     *homeActions
}

func newActionNames() *ActionNames {
	return &ActionNames{
		&homeActions{"index", "error"},
	}
}

type homeActions struct {
	Index mvc.Action
	Error mvc.Action
}

type ViewNames struct {
	Home     *homeViews
	Shared *sharedViews
}

func newViewNames() *ViewNames {
	return &ViewNames{
		&homeViews{"index"},
		&sharedViews{ "notfound404", "error"},
	}
}

type homeViews struct {
	Index mvc.Action
}

type sharedViews struct {
	NotFound404 mvc.Action
	Error mvc.Action
}