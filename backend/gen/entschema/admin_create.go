// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/admin"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUID sets the "uid" field.
func (ac *AdminCreate) SetUID(s string) *AdminCreate {
	ac.mutation.SetUID(s)
	return ac
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUID(s *string) *AdminCreate {
	if s != nil {
		ac.SetUID(*s)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminCreate) SetCreatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminCreate) SetUpdatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (ac *AdminCreate) SetDeactivatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetDeactivatedAt(t)
	return ac
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableDeactivatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetDeactivatedAt(*t)
	}
	return ac
}

// SetUserUID sets the "user_uid" field.
func (ac *AdminCreate) SetUserUID(s string) *AdminCreate {
	ac.mutation.SetUserUID(s)
	return ac
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	var (
		err  error
		node *Admin
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.UID(); !ok {
		v := admin.DefaultUID()
		ac.mutation.SetUID(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admin.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admin.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`entschema: missing required field "Admin.uid"`)}
	}
	if v, ok := ac.mutation.UID(); ok {
		if err := admin.UIDValidator(v); err != nil {
			return &ValidationError{Name: "uid", err: fmt.Errorf(`entschema: validator failed for field "Admin.uid": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entschema: missing required field "Admin.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`entschema: missing required field "Admin.updated_at"`)}
	}
	if _, ok := ac.mutation.UserUID(); !ok {
		return &ValidationError{Name: "user_uid", err: errors.New(`entschema: missing required field "Admin.user_uid"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: admin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: admin.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.UID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUID,
		})
		_node.UID = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeactivatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldDeactivatedAt,
		})
		_node.DeactivatedAt = &value
	}
	if value, ok := ac.mutation.UserUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUserUID,
		})
		_node.UserUID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.Create().
//		SetUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AdminCreate) OnConflict(opts ...sql.ConflictOption) *AdminUpsertOne {
	ac.conflict = opts
	return &AdminUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AdminCreate) OnConflictColumns(columns ...string) *AdminUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertOne{
		create: ac,
	}
}

type (
	// AdminUpsertOne is the builder for "upsert"-ing
	//  one Admin node.
	AdminUpsertOne struct {
		create *AdminCreate
	}

	// AdminUpsert is the "OnConflict" setter.
	AdminUpsert struct {
		*sql.UpdateSet
	}
)

// SetUID sets the "uid" field.
func (u *AdminUpsert) SetUID(v string) *AdminUpsert {
	u.Set(admin.FieldUID, v)
	return u
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUID() *AdminUpsert {
	u.SetExcluded(admin.FieldUID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsert) SetCreatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateCreatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsert) SetUpdatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUpdatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldUpdatedAt)
	return u
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *AdminUpsert) SetDeactivatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldDeactivatedAt, v)
	return u
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateDeactivatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldDeactivatedAt)
	return u
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *AdminUpsert) ClearDeactivatedAt() *AdminUpsert {
	u.SetNull(admin.FieldDeactivatedAt)
	return u
}

// SetUserUID sets the "user_uid" field.
func (u *AdminUpsert) SetUserUID(v string) *AdminUpsert {
	u.Set(admin.FieldUserUID, v)
	return u
}

// UpdateUserUID sets the "user_uid" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUserUID() *AdminUpsert {
	u.SetExcluded(admin.FieldUserUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AdminUpsertOne) UpdateNewValues() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.UID(); exists {
			s.SetIgnore(admin.FieldUID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(admin.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Admin.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AdminUpsertOne) Ignore() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertOne) DoNothing() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreate.OnConflict
// documentation for more info.
func (u *AdminUpsertOne) Update(set func(*AdminUpsert)) *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *AdminUpsertOne) SetUID(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUID() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsertOne) SetCreatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateCreatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertOne) SetUpdatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUpdatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *AdminUpsertOne) SetDeactivatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateDeactivatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *AdminUpsertOne) ClearDeactivatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetUserUID sets the "user_uid" field.
func (u *AdminUpsertOne) SetUserUID(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUserUID(v)
	})
}

// UpdateUserUID sets the "user_uid" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUserUID() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUserUID()
	})
}

// Exec executes the query.
func (u *AdminUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for AdminCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AdminUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AdminUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	builders []*AdminCreate
	conflict []sql.ConflictOption
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AdminCreateBulk) OnConflict(opts ...sql.ConflictOption) *AdminUpsertBulk {
	acb.conflict = opts
	return &AdminUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AdminCreateBulk) OnConflictColumns(columns ...string) *AdminUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertBulk{
		create: acb,
	}
}

// AdminUpsertBulk is the builder for "upsert"-ing
// a bulk of Admin nodes.
type AdminUpsertBulk struct {
	create *AdminCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AdminUpsertBulk) UpdateNewValues() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.UID(); exists {
				s.SetIgnore(admin.FieldUID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(admin.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AdminUpsertBulk) Ignore() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertBulk) DoNothing() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreateBulk.OnConflict
// documentation for more info.
func (u *AdminUpsertBulk) Update(set func(*AdminUpsert)) *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetUID sets the "uid" field.
func (u *AdminUpsertBulk) SetUID(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUID(v)
	})
}

// UpdateUID sets the "uid" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUID() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AdminUpsertBulk) SetCreatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateCreatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertBulk) SetUpdatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUpdatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (u *AdminUpsertBulk) SetDeactivatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeactivatedAt(v)
	})
}

// UpdateDeactivatedAt sets the "deactivated_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateDeactivatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeactivatedAt()
	})
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (u *AdminUpsertBulk) ClearDeactivatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.ClearDeactivatedAt()
	})
}

// SetUserUID sets the "user_uid" field.
func (u *AdminUpsertBulk) SetUserUID(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUserUID(v)
	})
}

// UpdateUserUID sets the "user_uid" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUserUID() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUserUID()
	})
}

// Exec executes the query.
func (u *AdminUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entschema: OnConflict was set for builder %d. Set it on the AdminCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entschema: missing options for AdminCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}