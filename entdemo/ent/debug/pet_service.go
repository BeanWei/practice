package entrest

import (
	"entdemo/ent"
	"entdemo/entrest"

	"entdemo/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// ListPetRequest .
type ListPetRequest struct {
	PageSize  int    `json:"pageSize,omitempty"`
	PageToken int    `json:"pageToken,omitempty"`
	Sort      string `json:"sort,omitempty"`
	Name      string `json:"name,omitempty"`
}

// GetPetRequest .
type GetPetRequest struct {
	ID string `json:"id,omitempty"`
}

// CreatePetRequest .
type CreatePetRequest struct {
	Remark string `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// UpdatePetRequest .
type UpdatePetRequest struct {
	Remark string `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// DeletePetRequest .
type DeletePetRequest struct {
	ID string `json:"id,omitempty"`
}

func NewPetServiceHandler(client *ent.Client, respHandler func(r *ghttp.Request, result *entrest.Result)) {
	s := g.Server()

	s.BindMiddleware(
		"GET:/api/v1/pet",
		middleware.AllMiddleware,
		middleware.ListMiddleware,
		func(r *ghttp.Request) {
			var req ListPetRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			res, err := client.Pet.
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
		"GET:/api/v1/pet/{id}",
		middleware.AllMiddleware,
		middleware.GetMiddleware,
		func(r *ghttp.Request) {
			var req GetPetRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			res, err := client.Pet.
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
		"POST:/api/v1/pet",
		middleware.AllMiddleware,
		middleware.CreateMiddleware,
		func(r *ghttp.Request) {
			var req CreatePetRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			res, err := client.Pet.
				Create().
				SetRemark(req.Remark).
				SetName(req.Name).
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
		"PUT:/api/v1/pet/{id}",
		middleware.AllMiddleware,
		middleware.UpdateMiddleware,
		func(r *ghttp.Request) {
			var req UpdatePetRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			res, err := client.Pet.
				UpdateOneID(id).
				SetRemark(req.Remark).
				SetName(req.Name).
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
		"DELETE:/api/v1/pet/{id}",
		middleware.AllMiddleware,
		middleware.DeleteMiddleware,
		func(r *ghttp.Request) {
			var req DeletePetRequest
			if err := r.Parse(&req); err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorParameterBind,
					Error:     err,
				})
			}

			id := r.GetString("id")
			err := client.Pet.
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
