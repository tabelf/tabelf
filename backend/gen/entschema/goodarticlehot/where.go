// Code generated by entc, DO NOT EDIT.

package goodarticlehot

import (
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func IDNotIn(ids ...uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func IDGT(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAt applies equality check predicate on the "deactivated_at" field. It's identical to DeactivatedAtEQ.
func DeactivatedAt(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// ArticleUID applies equality check predicate on the "article_uid" field. It's identical to ArticleUIDEQ.
func ArticleUID(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleUID), v))
	})
}

// HasExpired applies equality check predicate on the "has_expired" field. It's identical to HasExpiredEQ.
func HasExpired(v bool) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasExpired), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...string) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func UIDNotIn(vs ...string) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func UIDGT(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// UIDContains applies the Contains predicate on the "uid" field.
func UIDContains(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUID), v))
	})
}

// UIDHasPrefix applies the HasPrefix predicate on the "uid" field.
func UIDHasPrefix(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUID), v))
	})
}

// UIDHasSuffix applies the HasSuffix predicate on the "uid" field.
func UIDHasSuffix(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUID), v))
	})
}

// UIDEqualFold applies the EqualFold predicate on the "uid" field.
func UIDEqualFold(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUID), v))
	})
}

// UIDContainsFold applies the ContainsFold predicate on the "uid" field.
func UIDContainsFold(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeactivatedAtEQ applies the EQ predicate on the "deactivated_at" field.
func DeactivatedAtEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtNEQ applies the NEQ predicate on the "deactivated_at" field.
func DeactivatedAtNEQ(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIn applies the In predicate on the "deactivated_at" field.
func DeactivatedAtIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func DeactivatedAtNotIn(vs ...time.Time) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func DeactivatedAtGT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtGTE applies the GTE predicate on the "deactivated_at" field.
func DeactivatedAtGTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLT applies the LT predicate on the "deactivated_at" field.
func DeactivatedAtLT(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtLTE applies the LTE predicate on the "deactivated_at" field.
func DeactivatedAtLTE(v time.Time) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeactivatedAt), v))
	})
}

// DeactivatedAtIsNil applies the IsNil predicate on the "deactivated_at" field.
func DeactivatedAtIsNil() predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeactivatedAt)))
	})
}

// DeactivatedAtNotNil applies the NotNil predicate on the "deactivated_at" field.
func DeactivatedAtNotNil() predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeactivatedAt)))
	})
}

// ArticleUIDEQ applies the EQ predicate on the "article_uid" field.
func ArticleUIDEQ(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDNEQ applies the NEQ predicate on the "article_uid" field.
func ArticleUIDNEQ(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDIn applies the In predicate on the "article_uid" field.
func ArticleUIDIn(vs ...string) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldArticleUID), v...))
	})
}

// ArticleUIDNotIn applies the NotIn predicate on the "article_uid" field.
func ArticleUIDNotIn(vs ...string) predicate.GoodArticleHot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldArticleUID), v...))
	})
}

// ArticleUIDGT applies the GT predicate on the "article_uid" field.
func ArticleUIDGT(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDGTE applies the GTE predicate on the "article_uid" field.
func ArticleUIDGTE(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDLT applies the LT predicate on the "article_uid" field.
func ArticleUIDLT(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDLTE applies the LTE predicate on the "article_uid" field.
func ArticleUIDLTE(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDContains applies the Contains predicate on the "article_uid" field.
func ArticleUIDContains(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDHasPrefix applies the HasPrefix predicate on the "article_uid" field.
func ArticleUIDHasPrefix(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDHasSuffix applies the HasSuffix predicate on the "article_uid" field.
func ArticleUIDHasSuffix(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDEqualFold applies the EqualFold predicate on the "article_uid" field.
func ArticleUIDEqualFold(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldArticleUID), v))
	})
}

// ArticleUIDContainsFold applies the ContainsFold predicate on the "article_uid" field.
func ArticleUIDContainsFold(v string) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldArticleUID), v))
	})
}

// HasExpiredEQ applies the EQ predicate on the "has_expired" field.
func HasExpiredEQ(v bool) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHasExpired), v))
	})
}

// HasExpiredNEQ applies the NEQ predicate on the "has_expired" field.
func HasExpiredNEQ(v bool) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHasExpired), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodArticleHot) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodArticleHot) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
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
func Not(p predicate.GoodArticleHot) predicate.GoodArticleHot {
	return predicate.GoodArticleHot(func(s *sql.Selector) {
		p(s.Not())
	})
}
