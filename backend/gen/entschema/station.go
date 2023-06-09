// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"encoding/json"
	"fmt"
	"strings"
	"tabelf/backend/gen/entschema/station"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Station is the model entity for the Station schema.
type Station struct {
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
	// Title holds the value of the "title" field.
	// 标题
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	// 描述
	Description string `json:"description,omitempty"`
	// Image holds the value of the "image" field.
	// 图片
	Image string `json:"image,omitempty"`
	// Tags holds the value of the "tags" field.
	// 标签列表
	Tags []string `json:"tags,omitempty"`
	// Icon holds the value of the "icon" field.
	// 图标
	Icon string `json:"icon,omitempty"`
	// Source holds the value of the "source" field.
	// 来源
	Source string `json:"source,omitempty"`
	// Link holds the value of the "link" field.
	// 链接
	Link string `json:"link,omitempty"`
	// Praise holds the value of the "praise" field.
	// 点赞量
	Praise int `json:"praise,omitempty"`
	// Star holds the value of the "star" field.
	// 收藏量
	Star int `json:"star,omitempty"`
	// View holds the value of the "view" field.
	// 查看量
	View int `json:"view,omitempty"`
	// UserUID holds the value of the "user_uid" field.
	// 用户uid
	UserUID string `json:"user_uid,omitempty"`
	// Status holds the value of the "status" field.
	// 状态
	Status bool `json:"status,omitempty"`
	// CategoryUID holds the value of the "category_uid" field.
	// 分类uid
	CategoryUID string `json:"category_uid,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Station) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case station.FieldTags:
			values[i] = new([]byte)
		case station.FieldStatus:
			values[i] = new(sql.NullBool)
		case station.FieldID, station.FieldPraise, station.FieldStar, station.FieldView:
			values[i] = new(sql.NullInt64)
		case station.FieldUID, station.FieldTitle, station.FieldDescription, station.FieldImage, station.FieldIcon, station.FieldSource, station.FieldLink, station.FieldUserUID, station.FieldCategoryUID:
			values[i] = new(sql.NullString)
		case station.FieldCreatedAt, station.FieldUpdatedAt, station.FieldDeactivatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Station", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Station fields.
func (s *Station) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case station.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = uint64(value.Int64)
		case station.FieldUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				s.UID = value.String
			}
		case station.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case station.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case station.FieldDeactivatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deactivated_at", values[i])
			} else if value.Valid {
				s.DeactivatedAt = new(time.Time)
				*s.DeactivatedAt = value.Time
			}
		case station.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case station.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case station.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				s.Image = value.String
			}
		case station.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case station.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				s.Icon = value.String
			}
		case station.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				s.Source = value.String
			}
		case station.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				s.Link = value.String
			}
		case station.FieldPraise:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field praise", values[i])
			} else if value.Valid {
				s.Praise = int(value.Int64)
			}
		case station.FieldStar:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field star", values[i])
			} else if value.Valid {
				s.Star = int(value.Int64)
			}
		case station.FieldView:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field view", values[i])
			} else if value.Valid {
				s.View = int(value.Int64)
			}
		case station.FieldUserUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_uid", values[i])
			} else if value.Valid {
				s.UserUID = value.String
			}
		case station.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.Bool
			}
		case station.FieldCategoryUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category_uid", values[i])
			} else if value.Valid {
				s.CategoryUID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Station.
// Note that you need to call Station.Unwrap() before calling this method if this Station
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Station) Update() *StationUpdateOne {
	return (&StationClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Station entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Station) Unwrap() *Station {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("entschema: Station is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Station) String() string {
	var builder strings.Builder
	builder.WriteString("Station(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", uid=")
	builder.WriteString(s.UID)
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	if v := s.DeactivatedAt; v != nil {
		builder.WriteString(", deactivated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteString(", image=")
	builder.WriteString(s.Image)
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", s.Tags))
	builder.WriteString(", icon=")
	builder.WriteString(s.Icon)
	builder.WriteString(", source=")
	builder.WriteString(s.Source)
	builder.WriteString(", link=")
	builder.WriteString(s.Link)
	builder.WriteString(", praise=")
	builder.WriteString(fmt.Sprintf("%v", s.Praise))
	builder.WriteString(", star=")
	builder.WriteString(fmt.Sprintf("%v", s.Star))
	builder.WriteString(", view=")
	builder.WriteString(fmt.Sprintf("%v", s.View))
	builder.WriteString(", user_uid=")
	builder.WriteString(s.UserUID)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", category_uid=")
	builder.WriteString(s.CategoryUID)
	builder.WriteByte(')')
	return builder.String()
}

// Stations is a parsable slice of Station.
type Stations []*Station

func (s Stations) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
