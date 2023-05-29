// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"fmt"
	"strings"
	"tabelf/backend/gen/entschema/goodarticlecategory"
	"time"

	"entgo.io/ent/dialect/sql"
)

// GoodArticleCategory is the model entity for the GoodArticleCategory schema.
type GoodArticleCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID string `json:"uid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeactivatedAt holds the value of the "deactivated_at" field.
	DeactivatedAt *time.Time `json:"deactivated_at,omitempty"`
	// Name holds the value of the "name" field.
	// 分类名称
	Name string `json:"name,omitempty"`
	// Status holds the value of the "status" field.
	// 状态
	Status bool `json:"status,omitempty"`
	// Sequence holds the value of the "sequence" field.
	// 序号
	Sequence int `json:"sequence,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodArticleCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodarticlecategory.FieldStatus:
			values[i] = new(sql.NullBool)
		case goodarticlecategory.FieldID, goodarticlecategory.FieldSequence:
			values[i] = new(sql.NullInt64)
		case goodarticlecategory.FieldUID, goodarticlecategory.FieldName:
			values[i] = new(sql.NullString)
		case goodarticlecategory.FieldCreatedAt, goodarticlecategory.FieldUpdatedAt, goodarticlecategory.FieldDeactivatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodArticleCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodArticleCategory fields.
func (gac *GoodArticleCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodarticlecategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gac.ID = uint64(value.Int64)
		case goodarticlecategory.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				gac.UID = value.String
			}
		case goodarticlecategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gac.CreatedAt = value.Time
			}
		case goodarticlecategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gac.UpdatedAt = value.Time
			}
		case goodarticlecategory.FieldDeactivatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated_at", values[i])
			} else if value.Valid {
				gac.DeactivatedAt = new(time.Time)
				*gac.DeactivatedAt = value.Time
			}
		case goodarticlecategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gac.Name = value.String
			}
		case goodarticlecategory.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				gac.Status = value.Bool
			}
		case goodarticlecategory.FieldSequence:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence", values[i])
			} else if value.Valid {
				gac.Sequence = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodArticleCategory.
// Note that you need to call GoodArticleCategory.Unwrap() before calling this method if this GoodArticleCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (gac *GoodArticleCategory) Update() *GoodArticleCategoryUpdateOne {
	return (&GoodArticleCategoryClient{config: gac.config}).UpdateOne(gac)
}

// Unwrap unwraps the GoodArticleCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gac *GoodArticleCategory) Unwrap() *GoodArticleCategory {
	tx, ok := gac.config.driver.(*txDriver)
	if !ok {
		panic("entschema: GoodArticleCategory is not a transactional entity")
	}
	gac.config.driver = tx.drv
	return gac
}

// String implements the fmt.Stringer.
func (gac *GoodArticleCategory) String() string {
	var builder strings.Builder
	builder.WriteString("GoodArticleCategory(")
	builder.WriteString(fmt.Sprintf("id=%v", gac.ID))
	builder.WriteString(", uid=")
	builder.WriteString(gac.UID)
	builder.WriteString(", created_at=")
	builder.WriteString(gac.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gac.UpdatedAt.Format(time.ANSIC))
	if v := gac.DeactivatedAt; v != nil {
		builder.WriteString(", deactivated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(gac.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", gac.Status))
	builder.WriteString(", sequence=")
	builder.WriteString(fmt.Sprintf("%v", gac.Sequence))
	builder.WriteByte(')')
	return builder.String()
}

// GoodArticleCategories is a parsable slice of GoodArticleCategory.
type GoodArticleCategories []*GoodArticleCategory

func (gac GoodArticleCategories) config(cfg config) {
	for _i := range gac {
		gac[_i].config = cfg
	}
}
