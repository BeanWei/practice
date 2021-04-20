package app

import (
	"github.com/BeanWei/gfent/internal/app/smc"
	"github.com/gogf/gf/frame/g"
)

// Run 应用启动
func Run() {
	s := g.Server()

	// 业务系统初始化
	smc.Init()

	// 启动 http 服务
	s.Run()
}
