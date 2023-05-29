// Code generated by entc, DO NOT EDIT.

package teamfolder

import (
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAt applies equality check predicate on the "deactivated_at" field. It's identical to DeactivatedAtEQ.
func DeactivatedAt(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// TeamUID applies equality check predicate on the "team_uid" field. It's identical to TeamUIDEQ.
func TeamUID(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamUID), v))
	})
}

// FileName applies equality check predicate on the "file_name" field. It's identical to FileNameEQ.
func FileName(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileName), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUID), v))
	})
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUID), v))
	})
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUID), v))
	})
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUID), v))
	})
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAtEQ applies the EQ predicate on the "deactivated_at" field.
func DeactivatedAtEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtNEQ applies the NEQ predicate on the "deactivated_at" field.
func DeactivatedAtNEQ(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIn applies the In predicate on the "deactivated_at" field.
func DeactivatedAtIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeactivatedAt), v...))
	})
}

// DeactivatedAtNotIn applies the NotIn predicate on the "deactivated_at" field.
func DeactivatedAtNotIn(vs ...time.Time) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeactivatedAt), v...))
	})
}

// DeactivatedAtGT applies the GT predicate on the "deactivated_at" field.
func DeactivatedAtGT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtGTE applies the GTE predicate on the "deactivated_at" field.
func DeactivatedAtGTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLT applies the LT predicate on the "deactivated_at" field.
func DeactivatedAtLT(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLTE applies the LTE predicate on the "deactivated_at" field.
func DeactivatedAtLTE(v time.Time) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIsNil applies the IsNil predicate on the "deactivated_at" field.
func DeactivatedAtIsNil() predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeactivatedAt)))
	})
}

// DeactivatedAtNotNil applies the NotNil predicate on the "deactivated_at" field.
func DeactivatedAtNotNil() predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeactivatedAt)))
	})
}

// TeamUIDEQ applies the EQ predicate on the "team_uid" field.
func TeamUIDEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamUID), v))
	})
}

// TeamUIDNEQ applies the NEQ predicate on the "team_uid" field.
func TeamUIDNEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeamUID), v))
	})
}

// TeamUIDIn applies the In predicate on the "team_uid" field.
func TeamUIDIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeamUID), v...))
	})
}

// TeamUIDNotIn applies the NotIn predicate on the "team_uid" field.
func TeamUIDNotIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeamUID), v...))
	})
}

// TeamUIDGT applies the GT predicate on the "team_uid" field.
func TeamUIDGT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeamUID), v))
	})
}

// TeamUIDGTE applies the GTE predicate on the "team_uid" field.
func TeamUIDGTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeamUID), v))
	})
}

// TeamUIDLT applies the LT predicate on the "team_uid" field.
func TeamUIDLT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeamUID), v))
	})
}

// TeamUIDLTE applies the LTE predicate on the "team_uid" field.
func TeamUIDLTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeamUID), v))
	})
}

// TeamUIDContains applies the Contains predicate on the "team_uid" field.
func TeamUIDContains(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeamUID), v))
	})
}

// TeamUIDHasPrefix applies the HasPrefix predicate on the "team_uid" field.
func TeamUIDHasPrefix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeamUID), v))
	})
}

// TeamUIDHasSuffix applies the HasSuffix predicate on the "team_uid" field.
func TeamUIDHasSuffix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeamUID), v))
	})
}

// TeamUIDEqualFold applies the EqualFold predicate on the "team_uid" field.
func TeamUIDEqualFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeamUID), v))
	})
}

// TeamUIDContainsFold applies the ContainsFold predicate on the "team_uid" field.
func TeamUIDContainsFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeamUID), v))
	})
}

// FileNameEQ applies the EQ predicate on the "file_name" field.
func FileNameEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileName), v))
	})
}

// FileNameNEQ applies the NEQ predicate on the "file_name" field.
func FileNameNEQ(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFileName), v))
	})
}

// FileNameIn applies the In predicate on the "file_name" field.
func FileNameIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFileName), v...))
	})
}

// FileNameNotIn applies the NotIn predicate on the "file_name" field.
func FileNameNotIn(vs ...string) predicate.TeamFolder {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TeamFolder(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFileName), v...))
	})
}

// FileNameGT applies the GT predicate on the "file_name" field.
func FileNameGT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFileName), v))
	})
}

// FileNameGTE applies the GTE predicate on the "file_name" field.
func FileNameGTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFileName), v))
	})
}

// FileNameLT applies the LT predicate on the "file_name" field.
func FileNameLT(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFileName), v))
	})
}

// FileNameLTE applies the LTE predicate on the "file_name" field.
func FileNameLTE(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFileName), v))
	})
}

// FileNameContains applies the Contains predicate on the "file_name" field.
func FileNameContains(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFileName), v))
	})
}

// FileNameHasPrefix applies the HasPrefix predicate on the "file_name" field.
func FileNameHasPrefix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFileName), v))
	})
}

// FileNameHasSuffix applies the HasSuffix predicate on the "file_name" field.
func FileNameHasSuffix(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFileName), v))
	})
}

// FileNameIsNil applies the IsNil predicate on the "file_name" field.
func FileNameIsNil() predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFileName)))
	})
}

// FileNameNotNil applies the NotNil predicate on the "file_name" field.
func FileNameNotNil() predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFileName)))
	})
}

// FileNameEqualFold applies the EqualFold predicate on the "file_name" field.
func FileNameEqualFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFileName), v))
	})
}

// FileNameContainsFold applies the ContainsFold predicate on the "file_name" field.
func FileNameContainsFold(v string) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFileName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TeamFolder) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TeamFolder) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TeamFolder) predicate.TeamFolder {
	return predicate.TeamFolder(func(s *sql.Selector) {
		p(s.Not())
	})
}
