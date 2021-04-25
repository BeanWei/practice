package entrest

import (
	"entdemo/entrest"
	"time"

	"entdemo/ent"
	"entdemo/ent/pet"

	"entdemo/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// Pet .
type Pet struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy string     `json:"updatedBy,omitempty"`
	Remark    string     `json:"remark,omitempty"`
	Name      string     `json:"name,omitempty"`
	Owner     *ent.User  `json:"owner,omitempty"`
}

// ListPetRequest .
type ListPetRequest struct {
	PageSize  int    `json:"pageSize,omitempty"`
	PageToken int    `json:"pageToken,omitempty"`
	Sort      string `json:"sort,omitempty"`
	Name      string `json:"name,omitempty"`
}

// GetPetRequest .
type GetPetRequest struct {
	ID string `json:"id,omitempty" v:"required"`
}

// CreatePetRequest .
type CreatePetRequest struct {
	Remark string `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string `json:"name,omitempty" v:"required#请输入宠物的名字"`
	Owner  *User  `json:"owner,omitempty"`
}

// UpdatePetRequest .
type UpdatePetRequest struct {
	ID     string `json:"id,omitempty" v:"required"`
	Remark string `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string `json:"name,omitempty" v:"required#请输入宠物的名字"`
	Owner  *User  `json:"owner,omitempty"`
}

// DeletePetRequest .
type DeletePetRequest struct {
	ID string `json:"id,omitempty" v:"required"`
}

func entPet2restPet(pet *ent.Pet) *Pet {
	return &Pet{
		ID:        pet.ID,
		CreatedAt: pet.CreatedAt,
		UpdatedAt: pet.UpdatedAt,
		DeletedAt: pet.DeletedAt,
		CreatedBy: pet.CreatedBy,
		UpdatedBy: pet.UpdatedBy,
		Remark:    pet.Remark,
		Name:      pet.Name,
		Owner:     pet.Edges.Owner,
	}
}

func entPets2restPets(pets []*ent.Pet) []*Pet {
	restPets := make([]*Pet, len(pets))
	for _, pet := range pets {
		restPets = append(restPets, entPet2restPet(pet))
	}
	return restPets
}

func restPetIDs(items []*Pet) []string {
	ids := make([]string, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
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

			pets, err := client.Pet.
				Query().
				WithOwner().
				Limit(req.PageSize).
				Offset((req.PageToken - 1) * req.PageSize).
				All(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorList,
					Error:     err,
				})
			}

			res := entPets2restPets(pets)
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

			pet, err := client.Pet.
				Query().
				Where(
					pet.IDEQ(req.ID),
				).
				WithOwner().
				First(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorGet,
					Error:     err,
				})
			}

			res := entPet2restPet(pet)
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
				SetOwnerID(req.Owner.ID).
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

			res, err := client.Pet.
				UpdateOneID(req.ID).
				SetRemark(req.Remark).
				SetName(req.Name).
				SetOwnerID(req.Owner.ID).
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

			ids := gstr.SplitAndTrimSpace(req.ID, ",")
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
		},
	)
}
