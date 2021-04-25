package entrest

import (
	"entdemo/ent"
	"entdemo/ent/pet"
	"entdemo/entrest"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

type CreatePetRequest struct {
	Name string `json:"name,omitempty" v:"required#请输入宠物名字"`
}

type UpdatePetRequest struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty" v:"required#请输入宠物名字"`
}

type DeletePetRequest struct {
	ID string `json:"id,omitempty"`
}

type GetPetRequest struct {
	ID string `json:"id,omitempty"`
}

type ListPetRequest struct {
	PageSize  int    `json:"pageSize,omitempty"`
	PageToken int    `json:"pageToken,omitempty"`
	Sort      string `json:"sort,omitempty"`
}

type Pet struct {
	ID   string
	Name string
}

func entPet2restPet(e *ent.Pet) *Pet {
	return &Pet{
		ID:   e.ID,
		Name: e.Name,
	}
}

func restPet2entPet(r *Pet) *ent.Pet {
	return &ent.Pet{
		ID:   r.ID,
		Name: r.Name,
	}
}

func Middleware(r *ghttp.Request) {
	// 中间件处理逻辑
	r.Middleware.Next()
}

func NewPetServiceHandler(client *ent.Client, respHandler func(r *ghttp.Request, result *entrest.Result)) {
	s := g.Server()

	s.BindMiddleware(
		"POST:/api/pet",
		Middleware,
	)

	s.BindHandler("POST:/api/pet", func(r *ghttp.Request) {
		var req CreatePetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		duplicateID, err := client.Pet.
			Query().
			Where(
				pet.NameEQ(req.Name),
			).
			FirstID(r.Context())
		if err == nil || !ent.IsNotFound(err) {
			respHandler(r, &entrest.Result{
				ErrorType:   entrest.ErrorDuplicate,
				Error:       err,
				DuplicateID: duplicateID,
			})
		} else {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorCheckDuplicate,
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

	s.BindHandler("PUT:/api/pet/{id}", func(r *ghttp.Request) {
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

	s.BindHandler("DELETE:/api/pet/{id}", func(r *ghttp.Request) {
		var req DeletePetRequest
		if err := r.Parse(&req); err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorParameterBind,
				Error:     err,
			})
		}

		ids := gstr.SplitAndTrimSpace(r.GetString("id"), ",")
		_, err := client.Pet.
			Delete().
			Where(
				pet.IDIn(ids...),
			).
			Exec(r.Context())
		if err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorDelete,
				Error:     err,
			})
		}
		respHandler(r, &entrest.Result{})
	})

	s.BindHandler("GET:/api/pet/{id}", func(r *ghttp.Request) {
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
	})

	s.BindHandler("GET:/api/pet", func(r *ghttp.Request) {
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
}
