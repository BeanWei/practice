package entrest

import (
	"entdemo/ent"
	"entdemo/entrest"

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
	ID int `json:"id,omitempty"`
}

// CreatePetRequest .
type CreatePetRequest struct {
	Name string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// UpdatePetRequest .
type UpdatePetRequest struct {
	Name string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// DeletePetRequest .
type DeletePetRequest struct {
	ID int `json:"id,omitempty"`
}

func NewPetServiceHandler(client *ent.Client, respHandler func(r *ghttp.Request, result *entrest.Result)) {
	s := g.Server()

	s.BindHandler("GET:/api/v1/pet", func(r *ghttp.Request) {
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
	})

	s.BindHandler("GET:/api/v1/pet/{id}", func(r *ghttp.Request) {
		var req GetPetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		id := r.GetInt("id")
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
	})

	s.BindHandler("POST:/api/v1/pet", func(r *ghttp.Request) {
		var req CreatePetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		res, err := client.Pet.
			Create().
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
	})

	s.BindHandler("PUT:/api/v1/pet/{id}", func(r *ghttp.Request) {
		var req UpdatePetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		id := r.GetInt("id")
		res, err := client.Pet.
			UpdateOneID(id).
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
	})

	s.BindHandler("DELETE:/api/v1/pet/{id}", func(r *ghttp.Request) {
		var req DeletePetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		id := r.GetInt("id")
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
	})
}
