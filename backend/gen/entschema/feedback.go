// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"fmt"
	"strings"
	"tabelf/backend/gen/entschema/feedback"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Feedback is the model entity for the Feedback schema.
type Feedback struct {
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
	// UserUID holds the value of the "user_uid" field.
	// 用户uid
	UserUID string `json:"user_uid,omitempty"`
	// Category holds the value of the "category" field.
	// 问题分类
	Category string `json:"category,omitempty"`
	// Description holds the value of the "description" field.
	// 问题描述
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feedback) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case feedback.FieldID:
			values[i] = new(sql.NullInt64)
		case feedback.FieldUID, feedback.FieldUserUID, feedback.FieldCategory, feedback.FieldDescription:
			values[i] = new(sql.NullString)
		case feedback.FieldCreatedAt, feedback.FieldUpdatedAt, feedback.FieldDeactivatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Feedback", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feedback fields.
func (f *Feedback) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feedback.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = uint64(value.Int64)
		case feedback.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				f.UID = value.String
			}
		case feedback.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case feedback.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case feedback.FieldDeactivatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated_at", values[i])
			} else if value.Valid {
				f.DeactivatedAt = new(time.Time)
				*f.DeactivatedAt = value.Time
			}
		case feedback.FieldUserUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_uid", values[i])
			} else if value.Valid {
				f.UserUID = value.String
			}
		case feedback.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				f.Category = value.String
			}
		case feedback.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Feedback.
// Note that you need to call Feedback.Unwrap() before calling this method if this Feedback
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feedback) Update() *FeedbackUpdateOne {
	return (&FeedbackClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Feedback entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Feedback) Unwrap() *Feedback {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("entschema: Feedback is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feedback) String() string {
	var builder strings.Builder
	builder.WriteString("Feedback(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", uid=")
	builder.WriteString(f.UID)
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	if v := f.DeactivatedAt; v != nil {
		builder.WriteString(", deactivated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", user_uid=")
	builder.WriteString(f.UserUID)
	builder.WriteString(", category=")
	builder.WriteString(f.Category)
	builder.WriteString(", description=")
	builder.WriteString(f.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Feedbacks is a parsable slice of Feedback.
type Feedbacks []*Feedback

func (f Feedbacks) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
