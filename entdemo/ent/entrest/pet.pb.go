package entrest

import (
	"entdemo/apikit/errors"
	"entdemo/apikit/resp"
	"entdemo/ent"
	"entdemo/ent/pet"
	"entdemo/ent/predicate"
	"entdemo/entrest"

	"github.com/gogf/gf/errors/gerror"
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
	PageSize  int     `json:"pageSize,omitempty" v:"required|min:1"`
	PageToken int     `json:"pageToken,omitempty" v:"required|min:1"`
	Sort      string  `json:"sortField,omitempty"`
	Name      *string `json:"name,omitempty"`
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

func getPetColumnForListRequest(field string) string {
	switch field {
	case "name":
		return pet.FieldName
	default:
		return ""
	}
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
			resp.Error(r, errors.InvalidArgument("pet", err.Error()))
		}

		existsCount, err := client.Pet.
			Query().
			Where(
				pet.NameEQ(req.Name),
			).
			Count(r.Context())
		if err != nil {
			resp.Error(r, errors.DatabaseError("user", "check duplicate error"))
		} else if existsCount > 0 {
			resp.Error(r, errors.AlreadyExists("user", "It already exists. Do not create it again."))
		}

		res, err := client.Pet.
			Create().
			SetName(req.Name).
			Save(r.Context())
		if err != nil {
			resp.Error(r, errors.DatabaseError("user", "insert error"))
		}

		resp.OK(r, res)
	})

	s.BindHandler("PUT:/api/pet/{id}", func(r *ghttp.Request) {
		var req UpdatePetRequest
		if err := r.Parse(&req); err != nil {
			entrest.Resp.InvalidArgument(r, err)
		}

		id := r.GetString("id")

		existsCount, err := client.Pet.
			Query().
			Where(
				pet.IDNEQ(id),
				pet.NameEQ(req.Name),
			).
			Count(r.Context())
		if err != nil {
			entrest.Resp.DatabaseError(r, gerror.Wrap(err, "check duplicate error"))
		} else if existsCount > 0 {
			entrest.Resp.AlreadyExists(r, gerror.New("It already exists."))
		}

		res, err := client.Pet.
			UpdateOneID(id).
			SetName(req.Name).
			Save(r.Context())
		if err != nil {
			entrest.Resp.DatabaseError(r, gerror.Wrap(err, "update error"))
		}

		respHandler(r, &entrest.Result{
			Data: res,
		})
	})

	s.BindHandler("DELETE:/api/pet/{id}", func(r *ghttp.Request) {
		var req DeletePetRequest
		if err := r.Parse(&req); err != nil {
			entrest.Resp.InvalidArgument(r, err)
		}

		ids := gstr.SplitAndTrimSpace(r.GetString("id"), ",")
		_, err := client.Pet.
			Delete().
			Where(
				pet.IDIn(ids...),
			).
			Exec(r.Context())
		if err != nil {
			entrest.Resp.DatabaseError(r, gerror.Wrap(err, "delete error"))
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
			if ent.IsNotFound(err) {
				entrest.Resp.NotFound(r, gerror.Wrapf(err, "not found: %s", id))
			} else {
				entrest.Resp.DatabaseError(r, gerror.Wrap(err, "get error"))
			}
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

		petQuery := client.Pet.Query()

		if req.Sort != "" {
			sortName, sortOp := entrest.GetSortArg(req.Sort)
			if fieldName := getPetColumnForListRequest(sortName); fieldName != "" {
				if sortOp == entrest.ASC {
					petQuery.Order(
						ent.Asc(fieldName),
					)
				} else if sortOp == entrest.DESC {
					petQuery.Order(
						ent.Desc(fieldName),
					)
				}
			}

		}

		wherePlaceholder := make([]predicate.Pet, 0)
		if req.Name != nil {
			wherePlaceholder = append(wherePlaceholder, pet.NameEQ(*req.Name))
		}
		if len(wherePlaceholder) > 0 {
			petQuery.Where(wherePlaceholder...)
		}

		qty, err := petQuery.Count(r.Context())
		if err != nil {
			respHandler(r, &entrest.Result{
				ErrorType: entrest.ErrorList,
				Error:     err,
			})
		}

		res, err := petQuery.
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
			Total:  qty,
			IsList: true,
		})
	})
}
