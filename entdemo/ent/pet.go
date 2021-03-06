// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entdemo/ent/pet"
	"entdemo/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	// 编号
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	// 创建时间
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// DeletedAt holds the value of the "deletedAt" field.
	// 删除时间
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// CreatedBy holds the value of the "createdBy" field.
	// 创建者
	CreatedBy string `json:"createdBy,omitempty"`
	// UpdatedBy holds the value of the "updatedBy" field.
	// 修改者
	UpdatedBy string `json:"updatedBy,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
	// Name holds the value of the "name" field.
	// 宠物名字
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PetQuery when eager-loading is set.
	Edges     PetEdges `json:"edges"`
	user_pets *string
}

// PetEdges holds the relations/edges for other nodes in the graph.
type PetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PetEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pet.FieldID, pet.FieldCreatedBy, pet.FieldUpdatedBy, pet.FieldRemark, pet.FieldName:
			values[i] = new(sql.NullString)
		case pet.FieldCreatedAt, pet.FieldUpdatedAt, pet.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case pet.ForeignKeys[0]: // user_pets
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pe.ID = value.String
			}
		case pet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case pet.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case pet.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deletedAt", values[i])
			} else if value.Valid {
				pe.DeletedAt = new(time.Time)
				*pe.DeletedAt = value.Time
			}
		case pet.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field createdBy", values[i])
			} else if value.Valid {
				pe.CreatedBy = value.String
			}
		case pet.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updatedBy", values[i])
			} else if value.Valid {
				pe.UpdatedBy = value.String
			}
		case pet.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				pe.Remark = value.String
			}
		case pet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case pet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_pets", values[i])
			} else if value.Valid {
				pe.user_pets = new(string)
				*pe.user_pets = value.String
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Pet entity.
func (pe *Pet) QueryOwner() *UserQuery {
	return (&PetClient{config: pe.config}).QueryOwner(pe)
}

// Update returns a builder for updating this Pet.
// Note that you need to call Pet.Unwrap() before calling this method if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Pet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	if v := pe.DeletedAt; v != nil {
		builder.WriteString(", deletedAt=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", createdBy=")
	builder.WriteString(pe.CreatedBy)
	builder.WriteString(", updatedBy=")
	builder.WriteString(pe.UpdatedBy)
	builder.WriteString(", remark=")
	builder.WriteString(pe.Remark)
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
