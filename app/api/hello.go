package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Hello = helloApi{}

type helloApi struct{}

// Index is a demonstration route handler for output "Hello World!".
// @summary hello接口
// @tags    hello
// @router  /hello [GET]
// @success 200  "hello"
func (*helloApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
