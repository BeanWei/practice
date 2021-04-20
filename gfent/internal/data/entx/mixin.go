package entx

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/BeanWei/gfent/internal/shared"
	"github.com/gogf/gf/util/guid"
)

// BaseMixin implements the ent.Mixin for sharing
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return guid.S()
			}).
			Comment("编号"),
		field.Time("createdAt").
			Immutable().
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
		field.Time("deletedAt").
			Optional().
			Nillable().
			Comment("删除时间"),
		field.String("createdBy").
			Optional().
			Comment("创建者"),
		field.String("updatedBy").
			Optional().
			Comment("修改者"),
	}
}

// Indexes .
func (BaseMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("createdAt"),
		index.Fields("createdBy"),
	}
}

// Hooks of the BaseMixin.
func (BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			type BaseMutation interface {
				SetID(string)
				ID() (id string, exists bool)
				SetCreatedAt(time.Time)
				CreatedAt() (value time.Time, exists bool)
				SetUpdatedAt(time.Time)
				UpdatedAt() (value time.Time, exists bool)
				SetDeletedAt(time.Time)
				DeletedAt() (value time.Time, exists bool)
				SetCreatedBy(string)
				CreatedBy() (id string, exists bool)
				SetUpdatedBy(string)
				UpdatedBy() (id string, exists bool)
			}
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				bm, ok := m.(BaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected base-mutate call from mutation type %T", m)
				}

				var username string
				if customCtx := shared.Context.Get(ctx); customCtx != nil && customCtx.User != nil {
					username = customCtx.User.Username
				}

				switch op := m.Op(); {
				case op.Is(ent.OpCreate):
					if _, exists := bm.CreatedBy(); !exists {
						bm.SetCreatedBy(username)
					}
				case op.Is(ent.OpUpdateOne | ent.OpUpdate):
					if _, exists := bm.UpdatedBy(); !exists {
						bm.SetUpdatedBy(username)
					}
				case op.Is(ent.OpDeleteOne | ent.OpDelete):
					bm.SetDeletedAt(time.Now())
					if _, exists := bm.UpdatedBy(); !exists {
						bm.SetUpdatedBy(username)
					}
				}

				return next.Mutate(ctx, m)
			})
		},
	}
}
