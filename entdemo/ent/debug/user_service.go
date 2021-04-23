package entrest

import (
	"entdemo/ent"
	"entdemo/entrest"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// test

// ListUserRequest .
type ListUserRequest struct {
	PageSize  int     `json:"pageSize,omitempty"`
	PageToken int     `json:"pageToken,omitempty"`
	Sort      string  `json:"sort,omitempty"`
	Name      string  `json:"name,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

// GetUserRequest .
type GetUserRequest struct {
	ID string `json:"id,omitempty"`
}

// CreateUserRequest .
type CreateUserRequest struct {
	Remark string  `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string  `json:"name,omitempty" v:"required#请输入用户的名字"`
	Phone  *string `json:"phone,omitempty"`
}

// UpdateUserRequest .
type UpdateUserRequest struct {
	Remark string  `json:"remark,omitempty" v:"required#请输入备注"`
	Phone  *string `json:"phone,omitempty"`
}

// DeleteUserRequest .
type DeleteUserRequest struct {
	ID string `json:"id,omitempty"`
}

func NewUserServiceHandler(client *ent.Client, respHandler func(r *ghttp.Request, result *entrest.Result)) {
	s := g.Server()

	s.BindMiddleware(
		"GET:/api/v1/user",
		func(r *ghttp.Request) {
			var req ListUserRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			res, err := client.User.
				Query().
				Limit(req.PageSize).
				Offset((req.PageToken - 1) * req.PageSize).
				All(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorList,
					Error:     err,
				})
			}
			respHandler(r, &entrest.Result{
				Data:   res,
				IsList: true,
			})
		},
	)

	s.BindMiddleware(
		"GET:/api/v1/user/{id}",
		func(r *ghttp.Request) {
			var req GetUserRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			res, err := client.User.
				Get(r.Context(), id)
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorGet,
					Error:     err,
				})
			}
			respHandler(r, &entrest.Result{
				Data: res,
			})
		},
	)

	s.BindMiddleware(
		"POST:/api/v1/user",
		func(r *ghttp.Request) {
			var req CreateUserRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			res, err := client.User.
				Create().
				SetRemark(req.Remark).
				SetName(req.Name).
				SetNillablePhone(req.Phone).
				Save(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorCreate,
					Error:     err,
				})
			}
			respHandler(r, &entrest.Result{
				Data: res,
			})
		},
	)

	s.BindMiddleware(
		"PUT:/api/v1/user/{id}",
		func(r *ghttp.Request) {
			var req UpdateUserRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			res, err := client.User.
				UpdateOneID(id).
				SetRemark(req.Remark).
				SetNillablePhone(req.Phone).
				Save(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorUpdate,
					Error:     err,
				})
			}
			respHandler(r, &entrest.Result{
				Data: res,
			})
		},
	)

	s.BindMiddleware(
		"DELETE:/api/v1/user/{id}",
		func(r *ghttp.Request) {
			var req DeleteUserRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			err := client.User.
				DeleteOneID(id).
				Exec(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorDelete,
					Error:     err,
				})
			}
			respHandler(r, &entrest.Result{})
		},
	)
}
