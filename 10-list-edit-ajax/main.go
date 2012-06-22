package main

import (
	mvc "github.com/cihub/trinity"
	"net/http"
	names "github.com/cihub/trinity-examples/10-list-edit-ajax/mvcnames"
)

func main() {
	mvcI := mvc.NewMvcInfrastructure()
	err := mvcI.ParseViewsFolder("templates")
	if err != nil {	panic(err) }
	
	mvcI.BindController(NewHomeController)
	mvcI.BindUrl(names.C.Home, names.A.Home.Index, "/")
	
	mvcI.ServeStatic("/static/", "./static/")
	
	http.Handle("/", mvcI.Router)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {	panic(err) }
}