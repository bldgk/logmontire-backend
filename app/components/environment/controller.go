package environment

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
		"CreateEnvironment",
		"/environment/create",
		apiCreate,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules().AddValidation("name", []string{"!empty"}, "string", true))
	http.AddPostRoute(
		"FindEnvironment",
		"/environment/find",
		apiFind,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddGetRoute(
		"AllEnvironments",
		"/environment/all",
		apiAll,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
}

func apiCreate(request map[string]interface{}) (interface{}, *errors.Error) {
	return createEnvironment(request)
}

func apiFind(request map[string]interface{}) (interface{}, *errors.Error) {
	return findEnvironment(request)
}

func apiAll(request map[string]interface{}) (interface{}, *errors.Error) {
	return allEnvironments(request)
}
