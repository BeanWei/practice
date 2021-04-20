package smc

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Init 系统管理中心初始化
func Init() {
	s := g.Server()

	s.BindHandler("/smc/ping", func(r *ghttp.Request) {
		r.Response.Write("pong")
	})
}
