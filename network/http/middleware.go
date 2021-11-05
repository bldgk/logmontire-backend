package http

import (
	"github.com/kirillbeldyaga/fasthttp"
)

type Middleware func(*Route, fasthttp.RequestHandler) fasthttp.RequestHandler

type MiddlewareChain struct {
	middlewares []Middleware
}

func CreateMiddlewareChain(middlewares ...Middleware) *MiddlewareChain {
	return &MiddlewareChain{append(([]Middleware)(nil), middlewares...)}
}

func (c *MiddlewareChain) Then(r *Route, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	if h == nil {
		h = func(ctx *fasthttp.RequestCtx) {}
	}

	for i := range c.middlewares {

		h = c.middlewares[len(c.middlewares)-1-i](r, h)
	}

	return h
}

func (c *MiddlewareChain) Append(middlewares ...Middleware) *MiddlewareChain {
	newCons := make([]Middleware, 0, len(c.middlewares)+len(middlewares))
	newCons = append(newCons, c.middlewares...)
	newCons = append(newCons, middlewares...)

	return &MiddlewareChain{newCons}
}

func (c *MiddlewareChain) Extend(chain MiddlewareChain) *MiddlewareChain {
	return c.Append(chain.middlewares...)
}
