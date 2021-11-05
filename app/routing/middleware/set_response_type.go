package middleware

import (
	"github.com/kirillbeldyaga/fasthttp"
	"logmontire/network/http"
)

func SetResponseType(_ *http.Route, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Content-Type", "application/json")
		h(ctx)
	}
}
