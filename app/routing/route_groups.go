package routing

import (
	"logmontire/app/routing/middleware"
	"logmontire/network/http"
)

func DefaultRouteGroup() *http.RouteGroup {
	return &http.RouteGroup{
		Chain: http.CreateMiddlewareChain(middleware.EnableCORSHeader, middleware.SetResponseType, middleware.ValidateRequest),
	}
}
