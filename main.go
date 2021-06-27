package main

import (
	"github.com/gogf/gf/frame/g"
	_ "myapp/boot"
	_ "myapp/router"
)

// @title       `gf-demo`示例服务API
// @version     1.0
// @description `GoFrame`基础开发框架示例服务API接口文档。
// @schemes     http
func main() {
	//go ghttp.StartPProfServer(8299)
	//s := g.Server()
	//s.EnablePProf()
	//s.Run()
	g.Server().Run()
}
