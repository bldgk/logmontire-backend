package http

import (
	"fmt"
	"github.com/kirillbeldyaga/fasthttp"
	"logmontire/common/errors"
	"logmontire/crypto/encoding/json"
	"logmontire/log"
)

func RegisterController(registration func()) {
	registration()
}

func AddRoute(method Method, name string, pattern string, handler func(map[string]interface{}) (interface{}, *errors.Error), group *RouteGroup, rules ValidationRules) {
	route := &Route{
		Name:        name,
		Method:      method,
		Pattern:     pattern,
		RouteGroup:  group,
		Validations: rules,
	}
	route.CreateHandler(createHandler(route, handler))
	route.Register()
}

func AddGetRoute(name string, pattern string, handler func(map[string]interface{}) (interface{}, *errors.Error), group *RouteGroup, rules ValidationRules) {
	AddRoute(GET, name, pattern, handler, group, rules)
}

func AddPostRoute(name string, pattern string, handler func(map[string]interface{}) (interface{}, *errors.Error), group *RouteGroup, rules ValidationRules) {
	AddRoute(POST, name, pattern, handler, group, rules)
}

func createHandler(route *Route, handler func(map[string]interface{}) (interface{}, *errors.Error)) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		//log.Log(log.LStr("API: %s method handled", route.Name), log.ConsoleLog)

		fmt.Println(ctx.Request.PostArgs())
		request, _ := json.Ti(ctx.PostBody())
		//
		//var p fastjson.Parser
		//v, _ := p.Parse(`{
		//        "str": "bar",
		//        "int": 123,
		//        "float": 1.23,
		//        "bool": true,
		//        "arr": [1, "foo", {}]
		//}`)
		//
		//fmt.Println(v.Get("str"))
		//fmt.Println(v.GetBool("int"))

		response, err := handler(request)

		if err != nil {
			ctx.Error(err.Error(), err.Code)
		} else {
			ctx.SetBody(json.Fi(response))
			ctx.SetStatusCode(200)
		}
		//ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		//ctx.Response.Header.Set("Content-Type", "application/json")

		//ctx.Response.SetStatusCode(code)
		//start := time.Now()
		//if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		//	elapsed := time.Since(start)
		//	logrus.Errorfp("", elapsed, err.Error(), obj)
		//	ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		//}

		log.Log(log.LStr("API: %s method handled", route.Name), log.ConsoleLog)
	}
}
