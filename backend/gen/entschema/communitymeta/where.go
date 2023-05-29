// Code generated by entc, DO NOT EDIT.

package communitymeta

import (
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func IDGT(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAt applies equality check predicate on the "deactivated_at" field. It's identical to DeactivatedAtEQ.
func DeactivatedAt(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// CommunityUID applies equality check predicate on the "community_uid" field. It's identical to CommunityUIDEQ.
func CommunityUID(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommunityUID), v))
	})
}

// UserUID applies equality check predicate on the "user_uid" field. It's identical to UserUIDEQ.
func UserUID(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserUID), v))
	})
}

// HasPraise applies equality check predicate on the "has_praise" field. It's identical to HasPraiseEQ.
func HasPraise(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasPraise), v))
	})
}

// HasStar applies equality check predicate on the "has_star" field. It's identical to HasStarEQ.
func HasStar(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasStar), v))
	})
}

// HasView applies equality check predicate on the "has_view" field. It's identical to HasViewEQ.
func HasView(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasView), v))
	})
}

// HasUsed applies equality check predicate on the "has_used" field. It's identical to HasUsedEQ.
func HasUsed(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasUsed), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func UIDNotIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func UIDGT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUID), v))
	})
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUID), v))
	})
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUID), v))
	})
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUID), v))
	})
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAtEQ applies the EQ predicate on the "deactivated_at" field.
func DeactivatedAtEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtNEQ applies the NEQ predicate on the "deactivated_at" field.
func DeactivatedAtNEQ(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIn applies the In predicate on the "deactivated_at" field.
func DeactivatedAtIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func DeactivatedAtNotIn(vs ...time.Time) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func DeactivatedAtGT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtGTE applies the GTE predicate on the "deactivated_at" field.
func DeactivatedAtGTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLT applies the LT predicate on the "deactivated_at" field.
func DeactivatedAtLT(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLTE applies the LTE predicate on the "deactivated_at" field.
func DeactivatedAtLTE(v time.Time) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIsNil applies the IsNil predicate on the "deactivated_at" field.
func DeactivatedAtIsNil() predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeactivatedAt)))
	})
}

// DeactivatedAtNotNil applies the NotNil predicate on the "deactivated_at" field.
func DeactivatedAtNotNil() predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeactivatedAt)))
	})
}

// CommunityUIDEQ applies the EQ predicate on the "community_uid" field.
func CommunityUIDEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDNEQ applies the NEQ predicate on the "community_uid" field.
func CommunityUIDNEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDIn applies the In predicate on the "community_uid" field.
func CommunityUIDIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommunityUID), v...))
	})
}

// CommunityUIDNotIn applies the NotIn predicate on the "community_uid" field.
func CommunityUIDNotIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommunityUID), v...))
	})
}

// CommunityUIDGT applies the GT predicate on the "community_uid" field.
func CommunityUIDGT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDGTE applies the GTE predicate on the "community_uid" field.
func CommunityUIDGTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDLT applies the LT predicate on the "community_uid" field.
func CommunityUIDLT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDLTE applies the LTE predicate on the "community_uid" field.
func CommunityUIDLTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDContains applies the Contains predicate on the "community_uid" field.
func CommunityUIDContains(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDHasPrefix applies the HasPrefix predicate on the "community_uid" field.
func CommunityUIDHasPrefix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDHasSuffix applies the HasSuffix predicate on the "community_uid" field.
func CommunityUIDHasSuffix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDEqualFold applies the EqualFold predicate on the "community_uid" field.
func CommunityUIDEqualFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCommunityUID), v))
	})
}

// CommunityUIDContainsFold applies the ContainsFold predicate on the "community_uid" field.
func CommunityUIDContainsFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCommunityUID), v))
	})
}

// UserUIDEQ applies the EQ predicate on the "user_uid" field.
func UserUIDEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserUID), v))
	})
}

// UserUIDNEQ applies the NEQ predicate on the "user_uid" field.
func UserUIDNEQ(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserUID), v))
	})
}

// UserUIDIn applies the In predicate on the "user_uid" field.
func UserUIDIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserUID), v...))
	})
}

// UserUIDNotIn applies the NotIn predicate on the "user_uid" field.
func UserUIDNotIn(vs ...string) predicate.CommunityMeta {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CommunityMeta(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserUID), v...))
	})
}

// UserUIDGT applies the GT predicate on the "user_uid" field.
func UserUIDGT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserUID), v))
	})
}

// UserUIDGTE applies the GTE predicate on the "user_uid" field.
func UserUIDGTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserUID), v))
	})
}

// UserUIDLT applies the LT predicate on the "user_uid" field.
func UserUIDLT(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserUID), v))
	})
}

// UserUIDLTE applies the LTE predicate on the "user_uid" field.
func UserUIDLTE(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserUID), v))
	})
}

// UserUIDContains applies the Contains predicate on the "user_uid" field.
func UserUIDContains(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserUID), v))
	})
}

// UserUIDHasPrefix applies the HasPrefix predicate on the "user_uid" field.
func UserUIDHasPrefix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserUID), v))
	})
}

// UserUIDHasSuffix applies the HasSuffix predicate on the "user_uid" field.
func UserUIDHasSuffix(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserUID), v))
	})
}

// UserUIDEqualFold applies the EqualFold predicate on the "user_uid" field.
func UserUIDEqualFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserUID), v))
	})
}

// UserUIDContainsFold applies the ContainsFold predicate on the "user_uid" field.
func UserUIDContainsFold(v string) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserUID), v))
	})
}

// HasPraiseEQ applies the EQ predicate on the "has_praise" field.
func HasPraiseEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasPraise), v))
	})
}

// HasPraiseNEQ applies the NEQ predicate on the "has_praise" field.
func HasPraiseNEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasPraise), v))
	})
}

// HasStarEQ applies the EQ predicate on the "has_star" field.
func HasStarEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasStar), v))
	})
}

// HasStarNEQ applies the NEQ predicate on the "has_star" field.
func HasStarNEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasStar), v))
	})
}

// HasViewEQ applies the EQ predicate on the "has_view" field.
func HasViewEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasView), v))
	})
}

// HasViewNEQ applies the NEQ predicate on the "has_view" field.
func HasViewNEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasView), v))
	})
}

// HasUsedEQ applies the EQ predicate on the "has_used" field.
func HasUsedEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasUsed), v))
	})
}

// HasUsedNEQ applies the NEQ predicate on the "has_used" field.
func HasUsedNEQ(v bool) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasUsed), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CommunityMeta) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CommunityMeta) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
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
func Not(p predicate.CommunityMeta) predicate.CommunityMeta {
	return predicate.CommunityMeta(func(s *sql.Selector) {
		p(s.Not())
	})
}
