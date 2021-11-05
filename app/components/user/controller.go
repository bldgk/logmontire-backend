package user

import (
	"logmontire/app/routing"
	"logmontire/common/errors"
	"logmontire/network/http"
)

func registerController() {
	http.RegisterController(registerRoutes)
}

func registerRoutes() {
	http.AddPostRoute(
		"RegisterUser",
		"/user/register",
		apiRegistration,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules().AddValidation("username", []string{"!empty"}, "string", true))
	http.AddPostRoute(
		"FindUser",
		"/user/find",
		apiFind,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddGetRoute(
		"AllUsers",
		"/user/all",
		apiAll,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
}

func apiRegistration(request map[string]interface{}) (interface{}, *errors.Error) {
	return registerUser(request)
}

func apiFind(request map[string]interface{}) (interface{}, *errors.Error) {
	return findUser(request)
}

func apiAll(request map[string]interface{}) (interface{}, *errors.Error) {
	return allUsers(request)
}
