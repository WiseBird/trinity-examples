package main

import (
	mvc "github.com/cihub/trinity"
	"net/http"
	names "github.com/cihub/trinity-examples/8-authentication-authorization/mvcnames"
)

var (
	authProvider = new(AuthProvider)
)

func main() {
	mvcI := mvc.NewMvcInfrastructure()
	mvcI.SetAccessChecker(authProvider)
	err := mvcI.ParseViewsFolder("templates")
	if err != nil {	panic(err) }
	
	mvcI.BindController(NewHomeController)
	mvcI.BindController(NewAccountController)
	mvcI.BindUrl(names.C.Home, names.A.Home.Index, "/")
	
	http.Handle("/", mvcI.Router)
	err = http.ListenAndServe(":8000", nil)
	if err != nil {	panic(err) }
}