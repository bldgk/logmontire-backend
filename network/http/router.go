package http

import (
	"github.com/kirillbeldyaga/fasthttprouter"
)

var (
	API = map[string]*Route{}
)

func RegisterHandlers(router *fasthttprouter.Router) {
	for _, route := range API {
		if route.Method == GET {
			router.GET(route.Pattern, route.Handler)
		} else {
			router.POST(route.Pattern, route.Handler)
		}
	}
}
