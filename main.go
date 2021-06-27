package main

import (
	"github.com/gogf/gf/net/ghttp"
	_ "myapp/boot"
	_ "myapp/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	go ghttp.StartPProfServer(8299)

	s := g.Server()
	s.EnablePProf()
	s.Run()
}
