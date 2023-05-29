// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/urlcrawl"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UrlCrawlCreate is the builder for creating a UrlCrawl entity.
type UrlCrawlCreate struct {
	config
	mutation *UrlCrawlMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUID sets the "uid" field.
func (ucc *UrlCrawlCreate) SetUID(s string) *UrlCrawlCreate {
	ucc.mutation.SetUID(s)
	return ucc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableUID(s *string) *UrlCrawlCreate {
	if s != nil {
		ucc.SetUID(*s)
	}
	return ucc
}

// SetCreatedAt sets the "created_at" field.
func (ucc *UrlCrawlCreate) SetCreatedAt(t time.Time) *UrlCrawlCreate {
	ucc.mutation.SetCreatedAt(t)
	return ucc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableCreatedAt(t *time.Time) *UrlCrawlCreate {
	if t != nil {
		ucc.SetCreatedAt(*t)
	}
	return ucc
}

// SetUpdatedAt sets the "updated_at" field.
func (ucc *UrlCrawlCreate) SetUpdatedAt(t time.Time) *UrlCrawlCreate {
	ucc.mutation.SetUpdatedAt(t)
	return ucc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableUpdatedAt(t *time.Time) *UrlCrawlCreate {
	if t != nil {
		ucc.SetUpdatedAt(*t)
	}
	return ucc
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (ucc *UrlCrawlCreate) SetDeactivatedAt(t time.Time) *UrlCrawlCreate {
	ucc.mutation.SetDeactivatedAt(t)
	return ucc
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableDeactivatedAt(t *time.Time) *UrlCrawlCreate {
	if t != nil {
		ucc.SetDeactivatedAt(*t)
	}
	return ucc
}

// SetURL sets the "url" field.
func (ucc *UrlCrawlCreate) SetURL(s string) *UrlCrawlCreate {
	ucc.mutation.SetURL(s)
	return ucc
}

// SetCommunityUID sets the "community_uid" field.
func (ucc *UrlCrawlCreate) SetCommunityUID(s string) *UrlCrawlCreate {
	ucc.mutation.SetCommunityUID(s)
	return ucc
}

// SetNillableCommunityUID sets the "community_uid" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableCommunityUID(s *string) *UrlCrawlCreate {
	if s != nil {
		ucc.SetCommunityUID(*s)
	}
	return ucc
}

// SetCount sets the "count" field.
func (ucc *UrlCrawlCreate) SetCount(i int) *UrlCrawlCreate {
	ucc.mutation.SetCount(i)
	return ucc
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableCount(i *int) *UrlCrawlCreate {
	if i != nil {
		ucc.SetCount(*i)
	}
	return ucc
}

// SetHasComplete sets the "has_complete" field.
func (ucc *UrlCrawlCreate) SetHasComplete(b bool) *UrlCrawlCreate {
	ucc.mutation.SetHasComplete(b)
	return ucc
}

// SetNillableHasComplete sets the "has_complete" field if the given value is not nil.
func (ucc *UrlCrawlCreate) SetNillableHasComplete(b *bool) *UrlCrawlCreate {
	if b != nil {
		ucc.SetHasComplete(*b)
	}
	return ucc
}

// SetID sets the "id" field.
func (ucc *UrlCrawlCreate) SetID(u uint64) *UrlCrawlCreate {
	ucc.mutation.SetID(u)
	return ucc
}

// Mutation returns the UrlCrawlMutation object of the builder.
func (ucc *UrlCrawlCreate) Mutation() *UrlCrawlMutation {
	return ucc.mutation
}

// Save creates the UrlCrawl in the database.
func (ucc *UrlCrawlCreate) Save(ctx context.Context) (*UrlCrawl, error) {
	var (
		err  error
		node *UrlCrawl
	)
	ucc.defaults()
	if len(ucc.hooks) == 0 {
		if err = ucc.check(); err != nil {
			return nil, err
		}
		node, err = ucc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UrlCrawlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucc.check(); err != nil {
				return nil, err
			}
			ucc.mutation = mutation
			if node, err = ucc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucc.hooks) - 1; i >= 0; i-- {
			if ucc.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = ucc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UrlCrawlCreate) SaveX(ctx context.Context) *UrlCrawl {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UrlCrawlCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UrlCrawlCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucc *UrlCrawlCreate) defaults() {
	if _, ok := ucc.mutation.UID(); !ok {
		v := urlcrawl.DefaultUID()
		ucc.mutation.SetUID(v)
	}
	if _, ok := ucc.mutation.CreatedAt(); !ok {
		v := urlcrawl.DefaultCreatedAt()
		ucc.mutation.SetCreatedAt(v)
	}
	if _, ok := ucc.mutation.UpdatedAt(); !ok {
		v := urlcrawl.DefaultUpdatedAt()
		ucc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ucc.mutation.Count(); !ok {
		v := urlcrawl.DefaultCount
		ucc.mutation.SetCount(v)
	}
	if _, ok := ucc.mutation.HasComplete(); !ok {
		v := urlcrawl.DefaultHasComplete
		ucc.mutation.SetHasComplete(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UrlCrawlCreate) check() error {
	if _, ok := ucc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`entschema: missing required field "UrlCrawl.uid"`)}
	}
	if v, ok := ucc.mutation.UID(); ok {
		if err := urlcrawl.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf(`entschema: validator failed for field "UrlCrawl.uid": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entschema: missing required field "UrlCrawl.created_at"`)}
	}
	if _, ok := ucc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`entschema: missing required field "UrlCrawl.updated_at"`)}
	}
	if _, ok := ucc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`entschema: missing required field "UrlCrawl.url"`)}
	}
	if _, ok := ucc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`entschema: missing required field "UrlCrawl.count"`)}
	}
	if _, ok := ucc.mutation.HasComplete(); !ok {
		return &ValidationError{Name: "has_complete", err: errors.New(`entschema: missing required field "UrlCrawl.has_complete"`)}
	}
	return nil
}

func (ucc *UrlCrawlCreate) sqlSave(ctx context.Context) (*UrlCrawl, error) {
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (ucc *UrlCrawlCreate) createSpec() (*UrlCrawl, *sqlgraph.CreateSpec) {
	var (
		_node = &UrlCrawl{config: ucc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: urlcrawl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: urlcrawl.FieldID,
			},
		}
	)
	_spec.OnConflict = ucc.conflict
	if id, ok := ucc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ucc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urlcrawl.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := ucc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: urlcrawl.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ucc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: urlcrawl.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ucc.mutation.DeactivatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: urlcrawl.FieldDeactivatedAt,
		})
		_node.DeactivatedAt = &value
	}
	if value, ok := ucc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urlcrawl.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := ucc.mutation.CommunityUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urlcrawl.FieldCommunityUID,
		})
		_node.CommunityUID = value
	}
	if value, ok := ucc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: urlcrawl.FieldCount,
		})
		_node.Count = value
	}
	if value, ok := ucc.mutation.HasComplete(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: urlcrawl.FieldHasComplete,
		})
		_node.HasComplete = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UrlCrawl.Create().
//		SetUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UrlCrawlUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (ucc *UrlCrawlCreate) OnConflict(opts ...sql.ConflictOption) *UrlCrawlUpsertOne {
	ucc.conflict = opts
	return &UrlCrawlUpsertOne{
		create: ucc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UrlCrawl.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ucc *UrlCrawlCreate) OnConflictColumns(columns ...string) *UrlCrawlUpsertOne {
	ucc.conflict = append(ucc.conflict, sql.ConflictColumns(columns...))
	return &UrlCrawlUpsertOne{
		create: ucc,
	}
}

type (
	// UrlCrawlUpsertOne is the builder for "upsert"-ing
	//  one UrlCrawl node.
	UrlCrawlUpsertOne struct {
		create *UrlCrawlCreate
	}

	// UrlCrawlUpsert is the "OnConflict" setter.
	UrlCrawlUpsert struct {
		*sql.UpdateSet
	}
)

// SetUID sets the "uid" field.
func (u *UrlCrawlUpsert) SetUID(v string) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldUID, v)
	return u
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateUID() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldUID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UrlCrawlUpsert) SetCreatedAt(v time.Time) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateCreatedAt() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UrlCrawlUpsert) SetUpdatedAt(v time.Time) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateUpdatedAt() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldUpdatedAt)
	return u
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *UrlCrawlUpsert) SetDeactivatedAt(v time.Time) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldDeactivatedAt, v)
	return u
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateDeactivatedAt() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldDeactivatedAt)
	return u
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *UrlCrawlUpsert) ClearDeactivatedAt() *UrlCrawlUpsert {
	u.SetNull(urlcrawl.FieldDeactivatedAt)
	return u
}

// SetURL sets the "url" field.
func (u *UrlCrawlUpsert) SetURL(v string) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateURL() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldURL)
	return u
}

// SetCommunityUID sets the "community_uid" field.
func (u *UrlCrawlUpsert) SetCommunityUID(v string) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldCommunityUID, v)
	return u
}

// UpdateCommunityUID sets the "community_uid" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateCommunityUID() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldCommunityUID)
	return u
}

// ClearCommunityUID clears the value of the "community_uid" field.
func (u *UrlCrawlUpsert) ClearCommunityUID() *UrlCrawlUpsert {
	u.SetNull(urlcrawl.FieldCommunityUID)
	return u
}

// SetCount sets the "count" field.
func (u *UrlCrawlUpsert) SetCount(v int) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldCount, v)
	return u
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateCount() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldCount)
	return u
}

// AddCount adds v to the "count" field.
func (u *UrlCrawlUpsert) AddCount(v int) *UrlCrawlUpsert {
	u.Add(urlcrawl.FieldCount, v)
	return u
}

// SetHasComplete sets the "has_complete" field.
func (u *UrlCrawlUpsert) SetHasComplete(v bool) *UrlCrawlUpsert {
	u.Set(urlcrawl.FieldHasComplete, v)
	return u
}

// UpdateHasComplete sets the "has_complete" field to the value that was provided on create.
func (u *UrlCrawlUpsert) UpdateHasComplete() *UrlCrawlUpsert {
	u.SetExcluded(urlcrawl.FieldHasComplete)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UrlCrawl.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(urlcrawl.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UrlCrawlUpsertOne) UpdateNewValues() *UrlCrawlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(urlcrawl.FieldID)
		}
		if _, exists := u.create.mutation.UID(); exists {
			s.SetIgnore(urlcrawl.FieldUID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(urlcrawl.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UrlCrawl.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UrlCrawlUpsertOne) Ignore() *UrlCrawlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UrlCrawlUpsertOne) DoNothing() *UrlCrawlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UrlCrawlCreate.OnConflict
// documentation for more info.
func (u *UrlCrawlUpsertOne) Update(set func(*UrlCrawlUpsert)) *UrlCrawlUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UrlCrawlUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *UrlCrawlUpsertOne) SetUID(v string) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateUID() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UrlCrawlUpsertOne) SetCreatedAt(v time.Time) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateCreatedAt() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UrlCrawlUpsertOne) SetUpdatedAt(v time.Time) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateUpdatedAt() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *UrlCrawlUpsertOne) SetDeactivatedAt(v time.Time) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateDeactivatedAt() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *UrlCrawlUpsertOne) ClearDeactivatedAt() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetURL sets the "url" field.
func (u *UrlCrawlUpsertOne) SetURL(v string) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateURL() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateURL()
	})
}

// SetCommunityUID sets the "community_uid" field.
func (u *UrlCrawlUpsertOne) SetCommunityUID(v string) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCommunityUID(v)
	})
}

// UpdateCommunityUID sets the "community_uid" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateCommunityUID() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCommunityUID()
	})
}

// ClearCommunityUID clears the value of the "community_uid" field.
func (u *UrlCrawlUpsertOne) ClearCommunityUID() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.ClearCommunityUID()
	})
}

// SetCount sets the "count" field.
func (u *UrlCrawlUpsertOne) SetCount(v int) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *UrlCrawlUpsertOne) AddCount(v int) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateCount() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCount()
	})
}

// SetHasComplete sets the "has_complete" field.
func (u *UrlCrawlUpsertOne) SetHasComplete(v bool) *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetHasComplete(v)
	})
}

// UpdateHasComplete sets the "has_complete" field to the value that was provided on create.
func (u *UrlCrawlUpsertOne) UpdateHasComplete() *UrlCrawlUpsertOne {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateHasComplete()
	})
}

// Exec executes the query.
func (u *UrlCrawlUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for UrlCrawlCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UrlCrawlUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UrlCrawlUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UrlCrawlUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UrlCrawlCreateBulk is the builder for creating many UrlCrawl entities in bulk.
type UrlCrawlCreateBulk struct {
	config
	builders []*UrlCrawlCreate
	conflict []sql.ConflictOption
}

// Save creates the UrlCrawl entities in the database.
func (uccb *UrlCrawlCreateBulk) Save(ctx context.Context) ([]*UrlCrawl, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UrlCrawl, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UrlCrawlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UrlCrawlCreateBulk) SaveX(ctx context.Context) []*UrlCrawl {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UrlCrawlCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UrlCrawlCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UrlCrawl.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UrlCrawlUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (uccb *UrlCrawlCreateBulk) OnConflict(opts ...sql.ConflictOption) *UrlCrawlUpsertBulk {
	uccb.conflict = opts
	return &UrlCrawlUpsertBulk{
		create: uccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UrlCrawl.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uccb *UrlCrawlCreateBulk) OnConflictColumns(columns ...string) *UrlCrawlUpsertBulk {
	uccb.conflict = append(uccb.conflict, sql.ConflictColumns(columns...))
	return &UrlCrawlUpsertBulk{
		create: uccb,
	}
}

// UrlCrawlUpsertBulk is the builder for "upsert"-ing
// a bulk of UrlCrawl nodes.
type UrlCrawlUpsertBulk struct {
	create *UrlCrawlCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UrlCrawl.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(urlcrawl.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UrlCrawlUpsertBulk) UpdateNewValues() *UrlCrawlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(urlcrawl.FieldID)
				return
			}
			if _, exists := b.mutation.UID(); exists {
				s.SetIgnore(urlcrawl.FieldUID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(urlcrawl.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UrlCrawl.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UrlCrawlUpsertBulk) Ignore() *UrlCrawlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UrlCrawlUpsertBulk) DoNothing() *UrlCrawlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UrlCrawlCreateBulk.OnConflict
// documentation for more info.
func (u *UrlCrawlUpsertBulk) Update(set func(*UrlCrawlUpsert)) *UrlCrawlUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UrlCrawlUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *UrlCrawlUpsertBulk) SetUID(v string) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateUID() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UrlCrawlUpsertBulk) SetCreatedAt(v time.Time) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateCreatedAt() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UrlCrawlUpsertBulk) SetUpdatedAt(v time.Time) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateUpdatedAt() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *UrlCrawlUpsertBulk) SetDeactivatedAt(v time.Time) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateDeactivatedAt() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *UrlCrawlUpsertBulk) ClearDeactivatedAt() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetURL sets the "url" field.
func (u *UrlCrawlUpsertBulk) SetURL(v string) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateURL() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateURL()
	})
}

// SetCommunityUID sets the "community_uid" field.
func (u *UrlCrawlUpsertBulk) SetCommunityUID(v string) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCommunityUID(v)
	})
}

// UpdateCommunityUID sets the "community_uid" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateCommunityUID() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCommunityUID()
	})
}

// ClearCommunityUID clears the value of the "community_uid" field.
func (u *UrlCrawlUpsertBulk) ClearCommunityUID() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.ClearCommunityUID()
	})
}

// SetCount sets the "count" field.
func (u *UrlCrawlUpsertBulk) SetCount(v int) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *UrlCrawlUpsertBulk) AddCount(v int) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateCount() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateCount()
	})
}

// SetHasComplete sets the "has_complete" field.
func (u *UrlCrawlUpsertBulk) SetHasComplete(v bool) *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.SetHasComplete(v)
	})
}

// UpdateHasComplete sets the "has_complete" field to the value that was provided on create.
func (u *UrlCrawlUpsertBulk) UpdateHasComplete() *UrlCrawlUpsertBulk {
	return u.Update(func(s *UrlCrawlUpsert) {
		s.UpdateHasComplete()
	})
}

// Exec executes the query.
func (u *UrlCrawlUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entschema: OnConflict was set for builder %d. Set it on the UrlCrawlCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for UrlCrawlCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UrlCrawlUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
