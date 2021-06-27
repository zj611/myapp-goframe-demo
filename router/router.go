package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"myapp/app/api"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.Middleware(
		//	service.Middleware.Ctx,
		//	service.Middleware.CORS,
		//)
		group.GET("/hello", api.Hello)
		//group.ALL("/chat", api.Chat)
		group.POST("/user", api.User)

		//group.Group("/", func(group *ghttp.RouterGroup) {
		//	group.Middleware(service.Middleware.Auth)
		//	group.ALL("/user/profile", api.User.Profile)
		//})
	})
}
