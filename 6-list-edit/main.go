package main

import (
	mvc "github.com/cihub/trinity"
	"net/http"
)

func main() {
	mvcI := mvc.NewMvcInfrastructure()
	err := mvcI.ParseViewsFolder("templates")
	if err != nil {	panic(err) }
	
	mvcI.BindController(NewHomeController)
	mvcI.BindUrl("home", "list", "/")
	
	http.Handle("/", mvcI.Router)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {	panic(err) }
}