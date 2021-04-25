package entrest

import (
	"entdemo/entrest"
	"time"

	"entdemo/ent"
	"entdemo/ent/user"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

// User .
type User struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy string     `json:"updatedBy,omitempty"`
	Remark    string     `json:"remark,omitempty"`
	Name      string     `json:"name,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	Pets      []*ent.Pet `json:"pets,omitempty"`
}

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
	ID string `json:"id,omitempty" v:"required"`
}

// CreateUserRequest .
type CreateUserRequest struct {
	Remark string  `json:"remark,omitempty" v:"required#请输入备注"`
	Name   string  `json:"name,omitempty" v:"required#请输入用户的名字"`
	Phone  *string `json:"phone,omitempty"`
	Pets   []*Pet  `json:"pets,omitempty"`
}

// UpdateUserRequest .
type UpdateUserRequest struct {
	ID     string  `json:"id,omitempty" v:"required"`
	Remark string  `json:"remark,omitempty" v:"required#请输入备注"`
	Phone  *string `json:"phone,omitempty"`
	Pets   []*Pet  `json:"pets,omitempty"`
}

// DeleteUserRequest .
type DeleteUserRequest struct {
	ID string `json:"id,omitempty" v:"required"`
}

func entUser2restUser(user *ent.User) *User {
	return &User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		Remark:    user.Remark,
		Name:      user.Name,
		Phone:     user.Phone,
		Pets:      user.Edges.Pets,
	}
}

func entUsers2restUsers(users []*ent.User) []*User {
	restUsers := make([]*User, len(users))
	for _, user := range users {
		restUsers = append(restUsers, entUser2restUser(user))
	}
	return restUsers
}

func restUserIDs(items []*User) []string {
	ids := make([]string, len(items))
	for _, item := range items {
		ids = append(ids, item.ID)
	}
	return ids
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

			users, err := client.User.
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

			res := entUsers2restUsers(users)
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

			user, err := client.User.
				Query().
				Where(
					user.IDEQ(req.ID),
				).
				WithPets().
				First(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorGet,
					Error:     err,
				})
			}

			res := entUser2restUser(user)
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
			duplicateCount, err := client.User.
				Query().
				Where(
					user.NameEQ(req.Name),
				).
				Count(r.Context())
			if err != nil {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorCheckDuplicate,
					Error:     err,
				})
			} else if duplicateCount > 0 {
				respHandler(r, &entrest.Result{
					ErrorType: entrest.ErrorDuplicate,
					Error:     err,
				})
			}

			res, err := client.User.
				Create().
				SetRemark(req.Remark).
				SetName(req.Name).
				SetNillablePhone(req.Phone).
				AddPetIDs(restPetIDs(req.Pets)...).
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

			res, err := client.User.
				UpdateOneID(req.ID).
				SetRemark(req.Remark).
				SetNillablePhone(req.Phone).
				AddPetIDs(restPetIDs(req.Pets)...).
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

			ids := gstr.SplitAndTrimSpace(req.ID, ",")
			_, err := client.User.
				Delete().
				Where(
					user.IDIn(ids...),
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
