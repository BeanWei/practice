package middleware

import "github.com/gogf/gf/net/ghttp"

func AllMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

func ListMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

func GetMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

func CreateMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

func UpdateMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}

func DeleteMiddleware(r *ghttp.Request) {
	r.Middleware.Next()
}
