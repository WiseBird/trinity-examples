package main

import (
	mvc "github.com/cihub/trinity"
	"net/http"
	names "github.com/cihub/trinity-examples/9-views-for-errors/mvcnames"
)

func main() {
	mvcI := mvc.NewMvcInfrastructure()
	err := mvcI.ParseViewsFolder("templates")
	if err != nil {	panic(err) }
	
	mvcI.SetNotFoundView(&mvc.ControllerAction{names.C.Shared, names.V.Shared.NotFound404})
	mvcI.SetInternalErrorView(&mvc.ControllerAction{names.C.Shared, names.V.Shared.Error})
	
	mvcI.BindController(NewHomeController)
	mvcI.BindUrl(names.C.Home, names.A.Home.Index, "/")
	
	http.Handle("/", mvcI.Router)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {	panic(err) }
}