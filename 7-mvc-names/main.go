package main

import (
	mvc "github.com/cihub/trinity"
	"net/http"
	names "github.com/cihub/trinity-examples/7-mvc-names/mvcnames"
)

func main() {
	mvcI := mvc.NewMvcInfrastructure()
	err := mvcI.ParseViewsFolder("templates")
	if err != nil {	panic(err) }
	
	mvcI.BindController(NewHomeController)
	mvcI.BindUrl(names.C.Home, names.A.Home.Index, "/")
	
	http.Handle("/", mvcI.Router)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {	panic(err) }
}