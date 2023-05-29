// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/focus"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FocusCreate is the builder for creating a Focus entity.
type FocusCreate struct {
	config
	mutation *FocusMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUID sets the "uid" field.
func (fc *FocusCreate) SetUID(s string) *FocusCreate {
	fc.mutation.SetUID(s)
	return fc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (fc *FocusCreate) SetNillableUID(s *string) *FocusCreate {
	if s != nil {
		fc.SetUID(*s)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FocusCreate) SetCreatedAt(t time.Time) *FocusCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FocusCreate) SetNillableCreatedAt(t *time.Time) *FocusCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FocusCreate) SetUpdatedAt(t time.Time) *FocusCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FocusCreate) SetNillableUpdatedAt(t *time.Time) *FocusCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (fc *FocusCreate) SetDeactivatedAt(t time.Time) *FocusCreate {
	fc.mutation.SetDeactivatedAt(t)
	return fc
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (fc *FocusCreate) SetNillableDeactivatedAt(t *time.Time) *FocusCreate {
	if t != nil {
		fc.SetDeactivatedAt(*t)
	}
	return fc
}

// SetFollowerUID sets the "follower_uid" field.
func (fc *FocusCreate) SetFollowerUID(s string) *FocusCreate {
	fc.mutation.SetFollowerUID(s)
	return fc
}

// SetFolloweeUID sets the "followee_uid" field.
func (fc *FocusCreate) SetFolloweeUID(s string) *FocusCreate {
	fc.mutation.SetFolloweeUID(s)
	return fc
}

// SetStatus sets the "status" field.
func (fc *FocusCreate) SetStatus(b bool) *FocusCreate {
	fc.mutation.SetStatus(b)
	return fc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fc *FocusCreate) SetNillableStatus(b *bool) *FocusCreate {
	if b != nil {
		fc.SetStatus(*b)
	}
	return fc
}

// Mutation returns the FocusMutation object of the builder.
func (fc *FocusCreate) Mutation() *FocusMutation {
	return fc.mutation
}

// Save creates the Focus in the database.
func (fc *FocusCreate) Save(ctx context.Context) (*Focus, error) {
	var (
		err  error
		node *Focus
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FocusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FocusCreate) SaveX(ctx context.Context) *Focus {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FocusCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FocusCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FocusCreate) defaults() {
	if _, ok := fc.mutation.UID(); !ok {
		v := focus.DefaultUID()
		fc.mutation.SetUID(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := focus.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := focus.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.Status(); !ok {
		v := focus.DefaultStatus
		fc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FocusCreate) check() error {
	if _, ok := fc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`entschema: missing required field "Focus.uid"`)}
	}
	if v, ok := fc.mutation.UID(); ok {
		if err := focus.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf(`entschema: validator failed for field "Focus.uid": %w`, err)}
		}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entschema: missing required field "Focus.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`entschema: missing required field "Focus.updated_at"`)}
	}
	if _, ok := fc.mutation.FollowerUID(); !ok {
		return &ValidationError{Name: "follower_uid", err: errors.New(`entschema: missing required field "Focus.follower_uid"`)}
	}
	if _, ok := fc.mutation.FolloweeUID(); !ok {
		return &ValidationError{Name: "followee_uid", err: errors.New(`entschema: missing required field "Focus.followee_uid"`)}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`entschema: missing required field "Focus.status"`)}
	}
	return nil
}

func (fc *FocusCreate) sqlSave(ctx context.Context) (*Focus, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FocusCreate) createSpec() (*Focus, *sqlgraph.CreateSpec) {
	var (
		_node = &Focus{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: focus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: focus.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: focus.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: focus.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: focus.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.DeactivatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: focus.FieldDeactivatedAt,
		})
		_node.DeactivatedAt = &value
	}
	if value, ok := fc.mutation.FollowerUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: focus.FieldFollowerUID,
		})
		_node.FollowerUID = value
	}
	if value, ok := fc.mutation.FolloweeUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: focus.FieldFolloweeUID,
		})
		_node.FolloweeUID = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: focus.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Focus.Create().
//		SetUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FocusUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (fc *FocusCreate) OnConflict(opts ...sql.ConflictOption) *FocusUpsertOne {
	fc.conflict = opts
	return &FocusUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Focus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fc *FocusCreate) OnConflictColumns(columns ...string) *FocusUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FocusUpsertOne{
		create: fc,
	}
}

type (
	// FocusUpsertOne is the builder for "upsert"-ing
	//  one Focus node.
	FocusUpsertOne struct {
		create *FocusCreate
	}

	// FocusUpsert is the "OnConflict" setter.
	FocusUpsert struct {
		*sql.UpdateSet
	}
)

// SetUID sets the "uid" field.
func (u *FocusUpsert) SetUID(v string) *FocusUpsert {
	u.Set(focus.FieldUID, v)
	return u
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *FocusUpsert) UpdateUID() *FocusUpsert {
	u.SetExcluded(focus.FieldUID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FocusUpsert) SetCreatedAt(v time.Time) *FocusUpsert {
	u.Set(focus.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FocusUpsert) UpdateCreatedAt() *FocusUpsert {
	u.SetExcluded(focus.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FocusUpsert) SetUpdatedAt(v time.Time) *FocusUpsert {
	u.Set(focus.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FocusUpsert) UpdateUpdatedAt() *FocusUpsert {
	u.SetExcluded(focus.FieldUpdatedAt)
	return u
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *FocusUpsert) SetDeactivatedAt(v time.Time) *FocusUpsert {
	u.Set(focus.FieldDeactivatedAt, v)
	return u
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *FocusUpsert) UpdateDeactivatedAt() *FocusUpsert {
	u.SetExcluded(focus.FieldDeactivatedAt)
	return u
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *FocusUpsert) ClearDeactivatedAt() *FocusUpsert {
	u.SetNull(focus.FieldDeactivatedAt)
	return u
}

// SetFollowerUID sets the "follower_uid" field.
func (u *FocusUpsert) SetFollowerUID(v string) *FocusUpsert {
	u.Set(focus.FieldFollowerUID, v)
	return u
}

// UpdateFollowerUID sets the "follower_uid" field to the value that was provided on create.
func (u *FocusUpsert) UpdateFollowerUID() *FocusUpsert {
	u.SetExcluded(focus.FieldFollowerUID)
	return u
}

// SetFolloweeUID sets the "followee_uid" field.
func (u *FocusUpsert) SetFolloweeUID(v string) *FocusUpsert {
	u.Set(focus.FieldFolloweeUID, v)
	return u
}

// UpdateFolloweeUID sets the "followee_uid" field to the value that was provided on create.
func (u *FocusUpsert) UpdateFolloweeUID() *FocusUpsert {
	u.SetExcluded(focus.FieldFolloweeUID)
	return u
}

// SetStatus sets the "status" field.
func (u *FocusUpsert) SetStatus(v bool) *FocusUpsert {
	u.Set(focus.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FocusUpsert) UpdateStatus() *FocusUpsert {
	u.SetExcluded(focus.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Focus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *FocusUpsertOne) UpdateNewValues() *FocusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.UID(); exists {
			s.SetIgnore(focus.FieldUID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(focus.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Focus.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FocusUpsertOne) Ignore() *FocusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FocusUpsertOne) DoNothing() *FocusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FocusCreate.OnConflict
// documentation for more info.
func (u *FocusUpsertOne) Update(set func(*FocusUpsert)) *FocusUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FocusUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *FocusUpsertOne) SetUID(v string) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateUID() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FocusUpsertOne) SetCreatedAt(v time.Time) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateCreatedAt() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FocusUpsertOne) SetUpdatedAt(v time.Time) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateUpdatedAt() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *FocusUpsertOne) SetDeactivatedAt(v time.Time) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateDeactivatedAt() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *FocusUpsertOne) ClearDeactivatedAt() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetFollowerUID sets the "follower_uid" field.
func (u *FocusUpsertOne) SetFollowerUID(v string) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetFollowerUID(v)
	})
}

// UpdateFollowerUID sets the "follower_uid" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateFollowerUID() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateFollowerUID()
	})
}

// SetFolloweeUID sets the "followee_uid" field.
func (u *FocusUpsertOne) SetFolloweeUID(v string) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetFolloweeUID(v)
	})
}

// UpdateFolloweeUID sets the "followee_uid" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateFolloweeUID() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateFolloweeUID()
	})
}

// SetStatus sets the "status" field.
func (u *FocusUpsertOne) SetStatus(v bool) *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FocusUpsertOne) UpdateStatus() *FocusUpsertOne {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *FocusUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for FocusCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FocusUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FocusUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FocusUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FocusCreateBulk is the builder for creating many Focus entities in bulk.
type FocusCreateBulk struct {
	config
	builders []*FocusCreate
	conflict []sql.ConflictOption
}

// Save creates the Focus entities in the database.
func (fcb *FocusCreateBulk) Save(ctx context.Context) ([]*Focus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Focus, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FocusMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FocusCreateBulk) SaveX(ctx context.Context) []*Focus {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FocusCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FocusCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Focus.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FocusUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (fcb *FocusCreateBulk) OnConflict(opts ...sql.ConflictOption) *FocusUpsertBulk {
	fcb.conflict = opts
	return &FocusUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Focus.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fcb *FocusCreateBulk) OnConflictColumns(columns ...string) *FocusUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FocusUpsertBulk{
		create: fcb,
	}
}

// FocusUpsertBulk is the builder for "upsert"-ing
// a bulk of Focus nodes.
type FocusUpsertBulk struct {
	create *FocusCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Focus.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *FocusUpsertBulk) UpdateNewValues() *FocusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.UID(); exists {
				s.SetIgnore(focus.FieldUID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(focus.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Focus.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FocusUpsertBulk) Ignore() *FocusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FocusUpsertBulk) DoNothing() *FocusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FocusCreateBulk.OnConflict
// documentation for more info.
func (u *FocusUpsertBulk) Update(set func(*FocusUpsert)) *FocusUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FocusUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *FocusUpsertBulk) SetUID(v string) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateUID() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FocusUpsertBulk) SetCreatedAt(v time.Time) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateCreatedAt() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FocusUpsertBulk) SetUpdatedAt(v time.Time) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateUpdatedAt() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *FocusUpsertBulk) SetDeactivatedAt(v time.Time) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateDeactivatedAt() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *FocusUpsertBulk) ClearDeactivatedAt() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetFollowerUID sets the "follower_uid" field.
func (u *FocusUpsertBulk) SetFollowerUID(v string) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetFollowerUID(v)
	})
}

// UpdateFollowerUID sets the "follower_uid" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateFollowerUID() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateFollowerUID()
	})
}

// SetFolloweeUID sets the "followee_uid" field.
func (u *FocusUpsertBulk) SetFolloweeUID(v string) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetFolloweeUID(v)
	})
}

// UpdateFolloweeUID sets the "followee_uid" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateFolloweeUID() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateFolloweeUID()
	})
}

// SetStatus sets the "status" field.
func (u *FocusUpsertBulk) SetStatus(v bool) *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *FocusUpsertBulk) UpdateStatus() *FocusUpsertBulk {
	return u.Update(func(s *FocusUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *FocusUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entschema: OnConflict was set for builder %d. Set it on the FocusCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for FocusCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FocusUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}