package home

import (
	"logmontire/app/routing"
	"logmontire/common/errors"
	"logmontire/network/http"
)

func registerHomeController() {
	http.RegisterController(register)
}

func register() {
	http.AddGetRoute(
		"Home",
		"/",
		apiHome,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
	http.AddGetRoute(
		"Home",
		"/home",
		apiHome,
		routing.DefaultRouteGroup(),
		http.CreateValidationRules())
}

func apiHome(_ map[string]interface{}) (interface{}, *errors.Error) {
	return "home", nil
}
