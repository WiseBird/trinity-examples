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
	Account mvc.Controller
}

func newControllerNames() *ControllerNames {
	return &ControllerNames{"home", "account"}
}

type ActionNames struct {
	Home     *homeActions
	Account *accountActions
}

func newActionNames() *ActionNames {
	return &ActionNames{
		&homeActions{"index", "private"},
		&accountActions{"login", "logout"},
	}
}

type homeActions struct {
	Index mvc.Action
	Private mvc.Action
}

type accountActions struct {
	Login mvc.Action
	Logout mvc.Action
}

type ViewNames struct {
	Home     *homeViews
}

func newViewNames() *ViewNames {
	return &ViewNames{
		&homeViews{"index", "private"},
	}
}

type homeViews struct {
	Index mvc.Action
	Private mvc.Action
}