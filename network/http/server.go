package http

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kirillbeldyaga/fasthttp"
	"github.com/kirillbeldyaga/fasthttprouter"
	"log"
	"time"
)

type ServerConfiguration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (cfg *ServerConfiguration) Get() {

}

func (cfg *ServerConfiguration) Set() {

}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (httpServer *Server) Address() string {
	return fmt.Sprintf("%s:%d", httpServer.Host, httpServer.Port)
}

var (
	server = &Server{}
)

func Init(host string, port int) {
	server.Host = host
	server.Port = port
	lg("Network module initialized")
}

func NotFound(ctx *fasthttp.RequestCtx) {
	err := RouteNotFindError(string(ctx.Path()))
	ctx.Error(err.Error(), err.Code)
	ctx.Response.Header.Set("Content-Type", "application/json")
}

func MethodNotAllowed(ctx *fasthttp.RequestCtx) {
	err := MethodNotAllowedError(string(ctx.Method()))
	ctx.Error(err.Error(), err.Code)
	ctx.Response.Header.Set("Content-Type", "application/json")
}

func Start() {
	router := fasthttprouter.New()
	router.NotFound = NotFound
	router.MethodNotAllowed = MethodNotAllowed

	RegisterHandlers(router)

	lg("Listening at " + server.Address())

	log.Fatal(fasthttp.ListenAndServe(server.Address(), router.Handler))
}

func lg(text string) {
	color.New(color.FgHiCyan).Printf("%s\t%s\n", time.Now().Format("2006-01-02 15:04:05.000000"), "HttpServer: "+text)
}
