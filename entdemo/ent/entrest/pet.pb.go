package entrest

import (
	"entdemo/ent"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Pet struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

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
	PageSize int    `json:"pageSize,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	Sort     string `json:"sort,omitempty"`
}

// toProtoPet transforms the ent type to the pb type
func toProtoPet(e *ent.Pet) *Pet {
	return &Pet{
		ID:   e.ID,
		Name: e.Name,
	}
}

// type PetServiceHandler interface {
// 	CreatePet(context.Context, *CreatePetRequest) (*Pet, error)
// 	UpdatePet(context.Context, *UpdatePetRequest) (*Pet, error)
// 	DeletePet(context.Context, *DeletePetRequest) error
// 	GetPet(context.Context, *GetPetRequest) (*Pet, error)
// 	ListPet(context.Context, *ListPetRequest) ([]*Pet, error)
// }

func NewPetServiceHandler(client *ent.Client) {
	s := g.Server()

	s.BindHandler("POST:/api/pet", func(r *ghttp.Request) {
		var req CreatePetRequest
		if err := r.Parse(&req); err != nil {
			r.Response.WriteExit(err.Error())
		}

		res, err := client.Pet.
			Create().
			SetName(req.Name).
			Save(r.Context())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit(res)
	})

	s.BindHandler("PUT:/api/pet/{id}", func(r *ghttp.Request) {
		var req UpdatePetRequest
		if err := r.Parse(&req); err != nil {
			r.Response.WriteExit(err.Error())
		}

		id := r.GetInt("id")
		res, err := client.Pet.
			UpdateOneID(id).
			SetName(req.Name).
			Save(r.Context())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit(res)
	})

	s.BindHandler("DELETE:/api/pet/{id}", func(r *ghttp.Request) {
		var req CreatePetRequest
		if err := r.Parse(&req); err != nil {
			r.Response.WriteExit(err.Error())
		}

		id := r.GetInt("id")
		err := client.Pet.
			DeleteOneID(id).
			Exec(r.Context())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit()
	})

	s.BindHandler("GET:/api/pet/{id}", func(r *ghttp.Request) {
		var req CreatePetRequest
		if err := r.Parse(&req); err != nil {
			r.Response.WriteExit(err.Error())
		}

		id := r.GetInt("id")
		res, err := client.Pet.
			Get(r.Context(), id)
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit(res)
	})

	s.BindHandler("GET:/api/pet", func(r *ghttp.Request) {
		var req ListPetRequest
		if err := r.Parse(&req); err != nil {
			r.Response.WriteExit(err.Error())
		}

		res, err := client.Pet.
			Query().
			Limit(req.PageSize).
			Offset((req.PageNum - 1) * req.PageSize).
			All(r.Context())
		if err != nil {
			r.Response.WriteExit(err.Error())
		}
		r.Response.WriteExit(res)
	})
}
