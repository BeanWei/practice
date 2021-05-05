package resp

import (
	"entdemo/apikit/errors"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Resp 统一的 RESTful 接口响应处理

// Error 错误处理
func Error(r *ghttp.Request, err *errors.ErrorInfo) {
	r.Response.WriteStatusExit(err.Code(), err)
}

// OK 常规数据的成功处理
func OK(r *ghttp.Request, data interface{}) {
	r.Response.WriteJsonExit(data)
}

// PageOK 列表数据的成功处理
func PageOK(r *ghttp.Request, data interface{}, total int) {
	r.Response.WriteJsonExit(g.Map{
		"total": total,
		"data":  data,
	})
}
