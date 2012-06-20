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
}

func newControllerNames() *ControllerNames {
	return &ControllerNames{"home"}
}

type ActionNames struct {
	Home     *homeActions
}

func newActionNames() *ActionNames {
	return &ActionNames{
		&homeActions{"index"},
	}
}

type homeActions struct {
	Index mvc.Action
}

type ViewNames struct {
	Home     *homeViews
}

func newViewNames() *ViewNames {
	return &ViewNames{
		&homeViews{"index"},
	}
}

type homeViews struct {
	Index mvc.Action
}
