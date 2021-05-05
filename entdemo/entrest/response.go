package entrest

import (
	"net/http"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Resp 统一的 RESTful 接口响应处理
var Resp = response{}

type response struct{}

// Error 错误处理
func (resp *response) Error(r *ghttp.Request, code int, reason string, err error) {
	g.Log().Error(err)
	r.Response.WriteStatusExit(code, g.Map{
		"code":    code,
		"reason":  reason,
		"message": gerror.Current(err).Error(),
	})
}

// OK 常规数据的成功处理
func (resp *response) OK(r *ghttp.Request, data interface{}) {
	r.Response.WriteJsonExit(data)
}

// PageOK 列表数据的成功处理
func (resp *response) PageOK(r *ghttp.Request, data interface{}, total int) {
	r.Response.WriteJsonExit(g.Map{
		"total": total,
		"data":  data,
	})
}

// UnknownError new UnknownError error that is mapped to a 500 response.
func (resp *response) UnknownError(r *ghttp.Request, reason string, err error) {
	resp.Error(r, http.StatusInternalServerError, reason, err)
}

// BadRequest new BadRequest error that is mapped to a 400 response.
func (resp *response) BadRequest(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusBadRequest, "bad_request", err)
}

// InvalidArgument new InvalidArgument error that is mapped to a 400 response.
func (resp *response) InvalidArgument(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusBadRequest, "invalid_argument", err)
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func (resp *response) Unauthorized(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusUnauthorized, "unauthorized", err)
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func (resp *response) Forbidden(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusForbidden, "forbidden", err)
}

// NotFound new NotFound error that is mapped to a 404 response.
func (resp *response) NotFound(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusNotFound, "not_found", err)
}

// Conflict new Conflict error that is mapped to a 409 response.
func (resp *response) Conflict(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusConflict, "conflict", err)
}

// AlreadyExists new AlreadyExists error that is mapped to a 409 response.
func (resp *response) AlreadyExists(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusConflict, "already_exists", err)
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func (resp *response) InternalServer(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusInternalServerError, "internal_server_error", err)
}

// DatabaseError new DatabaseError error that is mapped to a 500 response.
func (resp *response) DatabaseError(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusInternalServerError, "database_error", err)
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to a HTTP 503 response.
func (resp *response) ServiceUnavailable(r *ghttp.Request, err error) {
	resp.Error(r, http.StatusServiceUnavailable, "service_unavailable", err)
}
