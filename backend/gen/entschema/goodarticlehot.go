// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"fmt"
	"strings"
	"tabelf/backend/gen/entschema/goodarticlehot"
	"time"

	"entgo.io/ent/dialect/sql"
)

// GoodArticleHot is the model entity for the GoodArticleHot schema.
type GoodArticleHot struct {
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
	// ArticleUID holds the value of the "article_uid" field.
	// 好文uid
	ArticleUID string `json:"article_uid,omitempty"`
	// HasExpired holds the value of the "has_expired" field.
	// 是否过期
	HasExpired bool `json:"has_expired,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodArticleHot) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodarticlehot.FieldHasExpired:
			values[i] = new(sql.NullBool)
		case goodarticlehot.FieldID:
			values[i] = new(sql.NullInt64)
		case goodarticlehot.FieldUID, goodarticlehot.FieldArticleUID:
			values[i] = new(sql.NullString)
		case goodarticlehot.FieldCreatedAt, goodarticlehot.FieldUpdatedAt, goodarticlehot.FieldDeactivatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodArticleHot", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodArticleHot fields.
func (gah *GoodArticleHot) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodarticlehot.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gah.ID = uint64(value.Int64)
		case goodarticlehot.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				gah.UID = value.String
			}
		case goodarticlehot.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gah.CreatedAt = value.Time
			}
		case goodarticlehot.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gah.UpdatedAt = value.Time
			}
		case goodarticlehot.FieldDeactivatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated_at", values[i])
			} else if value.Valid {
				gah.DeactivatedAt = new(time.Time)
				*gah.DeactivatedAt = value.Time
			}
		case goodarticlehot.FieldArticleUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field article_uid", values[i])
			} else if value.Valid {
				gah.ArticleUID = value.String
			}
		case goodarticlehot.FieldHasExpired:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_expired", values[i])
			} else if value.Valid {
				gah.HasExpired = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodArticleHot.
// Note that you need to call GoodArticleHot.Unwrap() before calling this method if this GoodArticleHot
// was returned from a transaction, and the transaction was committed or rolled back.
func (gah *GoodArticleHot) Update() *GoodArticleHotUpdateOne {
	return (&GoodArticleHotClient{config: gah.config}).UpdateOne(gah)
}

// Unwrap unwraps the GoodArticleHot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gah *GoodArticleHot) Unwrap() *GoodArticleHot {
	tx, ok := gah.config.driver.(*txDriver)
	if !ok {
		panic("entschema: GoodArticleHot is not a transactional entity")
	}
	gah.config.driver = tx.drv
	return gah
}

// String implements the fmt.Stringer.
func (gah *GoodArticleHot) String() string {
	var builder strings.Builder
	builder.WriteString("GoodArticleHot(")
	builder.WriteString(fmt.Sprintf("id=%v", gah.ID))
	builder.WriteString(", uid=")
	builder.WriteString(gah.UID)
	builder.WriteString(", created_at=")
	builder.WriteString(gah.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gah.UpdatedAt.Format(time.ANSIC))
	if v := gah.DeactivatedAt; v != nil {
		builder.WriteString(", deactivated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", article_uid=")
	builder.WriteString(gah.ArticleUID)
	builder.WriteString(", has_expired=")
	builder.WriteString(fmt.Sprintf("%v", gah.HasExpired))
	builder.WriteByte(')')
	return builder.String()
}

// GoodArticleHots is a parsable slice of GoodArticleHot.
type GoodArticleHots []*GoodArticleHot

func (gah GoodArticleHots) config(cfg config) {
	for _i := range gah {
		gah[_i].config = cfg
	}
}
