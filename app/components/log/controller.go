package log

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
		"StoreLog",
		"/log/store",
		apiStore,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddPostRoute(
		"FindLog",
		"/log/find",
		apiFind,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddGetRoute(
		"AllLog",
		"/log/all",
		apiAll,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
}

func apiStore(request map[string]interface{}) (interface{}, *errors.Error) {
	return storeLog(request)
}

func apiFind(request map[string]interface{}) (interface{}, *errors.Error) {
	return FindLog(request)
}

func apiAll(request map[string]interface{}) (interface{}, *errors.Error) {
	//return allLogs(request)
	return nil, nil
}
