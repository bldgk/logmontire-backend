package http

import (
	"fmt"
	"github.com/kirillbeldyaga/fasthttp"
)

type Method int

const (
	GET  = 0
	POST = 1
)

type RouteGroup struct {
	Chain *MiddlewareChain
}

type Route struct {
	Name        string
	Method      Method
	Pattern     string
	Handler     func(*fasthttp.RequestCtx)
	RouteGroup  *RouteGroup
	Validations ValidationRules
}

func (route *Route) Register() {
	if route.RouteGroup == nil {
		route.RouteGroup = &RouteGroup{
			Chain: CreateMiddlewareChain(),
		}
	}

	API[route.Pattern] = route
}

func (route *Route) CreateHandler(handler func(ctx *fasthttp.RequestCtx)) {
	route.Handler = route.RouteGroup.Chain.Then(route, handler)
}

func (route *Route) String() string {
	return fmt.Sprintf("%s [%s] %s", route.Name, route.Method, route.Pattern)
}
