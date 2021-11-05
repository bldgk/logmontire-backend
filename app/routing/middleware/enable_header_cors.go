package middleware

import (
	"github.com/kirillbeldyaga/fasthttp"
	"logmontire/network/http"
)

func EnableCORSHeader(_ *http.Route, h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		h(ctx)
	}
}
