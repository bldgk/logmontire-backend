package channel

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
		"CreateChannel",
		"/channel/create",
		apiCreate,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddPostRoute(
		"FindChannel",
		"/channel/find",
		apiFind,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddGetRoute(
		"AllChannels",
		"/channel/all",
		apiAll,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
}

func apiCreate(request map[string]interface{}) (interface{}, *errors.Error) {
	return CreateChannel(request)
}

func apiFind(request map[string]interface{}) (interface{}, *errors.Error) {
	return FindChannel(request)
}

func apiAll(request map[string]interface{}) (interface{}, *errors.Error) {
	return allChannels(request)
}
