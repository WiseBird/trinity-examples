package main

import (
	"code.google.com/p/gorilla/securecookie"
	mvc "github.com/cihub/trinity"
	"net/http"
	names "github.com/cihub/trinity-examples/8-authentication-authorization/mvcnames"
)

var (
	hashKey = securecookie.GenerateRandomKey(32)
	blockKey = securecookie.GenerateRandomKey(32)
	
	authCookieName = "auth"
	authCookieValue = "valid"
)

type AuthProvider struct {
}

func (provider *AuthProvider) IsAccessAllowed(
	c mvc.Controller, 
	a mvc.Action, 
	response http.ResponseWriter, 
	request *http.Request) mvc.ActionResultInterface {
		
	if (c == names.C.Home && a == names.A.Home.Index) ||
		(c == names.C.Account && a == names.A.Account.Login) {
		return nil
	}

	if !provider.IsAuthenticated(request) {
		return mvc.RedirectToAction(names.C.Home, names.A.Home.Index, 
			map[string]string { "Restricted" : "true" })
	}

	return nil
}

func (provider *AuthProvider) IsAuthenticated(request *http.Request) bool {
	cookie, err := request.Cookie(authCookieName)
	if err != nil {
		return false
	}
	
	var s = securecookie.New(hashKey, blockKey)
	
	value := ""
	err = s.Decode(authCookieName, cookie.Value, &value);
	if err != nil {
		return false
	}

	return value == authCookieValue
}

func (provider *AuthProvider) Login(response http.ResponseWriter, request *http.Request) (error) {
	var s = securecookie.New(hashKey, blockKey)
	encoded, err := s.Encode(authCookieName, authCookieValue)
	if err != nil {
		return err
	}
	
	cookie := &http.Cookie{
		Name:  authCookieName,
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(response, cookie)

	return nil
}

func (provider *AuthProvider) Logout(response http.ResponseWriter, request *http.Request) {
	cookie := &http.Cookie{
		Name:  authCookieName,
		Value: "",
		Path:  "/",
	}
	
	http.SetCookie(response, cookie)
}