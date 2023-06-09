// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/collaboration"
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollaborationUpdate is the builder for updating Collaboration entities.
type CollaborationUpdate struct {
	config
	hooks    []Hook
	mutation *CollaborationMutation
}

// Where appends a list predicates to the CollaborationUpdate builder.
func (cu *CollaborationUpdate) Where(ps ...predicate.Collaboration) *CollaborationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CollaborationUpdate) SetUpdatedAt(t time.Time) *CollaborationUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (cu *CollaborationUpdate) SetDeactivatedAt(t time.Time) *CollaborationUpdate {
	cu.mutation.SetDeactivatedAt(t)
	return cu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (cu *CollaborationUpdate) SetNillableDeactivatedAt(t *time.Time) *CollaborationUpdate {
	if t != nil {
		cu.SetDeactivatedAt(*t)
	}
	return cu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (cu *CollaborationUpdate) ClearDeactivatedAt() *CollaborationUpdate {
	cu.mutation.ClearDeactivatedAt()
	return cu
}

// SetShardUID sets the "shard_uid" field.
func (cu *CollaborationUpdate) SetShardUID(s string) *CollaborationUpdate {
	cu.mutation.SetShardUID(s)
	return cu
}

// SetFolderUID sets the "folder_uid" field.
func (cu *CollaborationUpdate) SetFolderUID(s string) *CollaborationUpdate {
	cu.mutation.SetFolderUID(s)
	return cu
}

// SetUserUID sets the "user_uid" field.
func (cu *CollaborationUpdate) SetUserUID(s string) *CollaborationUpdate {
	cu.mutation.SetUserUID(s)
	return cu
}

// SetAuthority sets the "authority" field.
func (cu *CollaborationUpdate) SetAuthority(s string) *CollaborationUpdate {
	cu.mutation.SetAuthority(s)
	return cu
}

// SetFolderNumber sets the "folder_number" field.
func (cu *CollaborationUpdate) SetFolderNumber(s string) *CollaborationUpdate {
	cu.mutation.SetFolderNumber(s)
	return cu
}

// Mutation returns the CollaborationMutation object of the builder.
func (cu *CollaborationUpdate) Mutation() *CollaborationMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollaborationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollaborationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollaborationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollaborationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollaborationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CollaborationUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := collaboration.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CollaborationUpdate) check() error {
	if v, ok := cu.mutation.FolderNumber(); ok {
		if err := collaboration.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "Collaboration.folder_number": %w`, err)}
		}
	}
	return nil
}

func (cu *CollaborationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collaboration.Table,
			Columns: collaboration.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: collaboration.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collaboration.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collaboration.FieldDeactivatedAt,
		})
	}
	if cu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collaboration.FieldDeactivatedAt,
		})
	}
	if value, ok := cu.mutation.ShardUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldShardUID,
		})
	}
	if value, ok := cu.mutation.FolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldFolderUID,
		})
	}
	if value, ok := cu.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldUserUID,
		})
	}
	if value, ok := cu.mutation.Authority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldAuthority,
		})
	}
	if value, ok := cu.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldFolderNumber,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collaboration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CollaborationUpdateOne is the builder for updating a single Collaboration entity.
type CollaborationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollaborationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CollaborationUpdateOne) SetUpdatedAt(t time.Time) *CollaborationUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (cuo *CollaborationUpdateOne) SetDeactivatedAt(t time.Time) *CollaborationUpdateOne {
	cuo.mutation.SetDeactivatedAt(t)
	return cuo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (cuo *CollaborationUpdateOne) SetNillableDeactivatedAt(t *time.Time) *CollaborationUpdateOne {
	if t != nil {
		cuo.SetDeactivatedAt(*t)
	}
	return cuo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (cuo *CollaborationUpdateOne) ClearDeactivatedAt() *CollaborationUpdateOne {
	cuo.mutation.ClearDeactivatedAt()
	return cuo
}

// SetShardUID sets the "shard_uid" field.
func (cuo *CollaborationUpdateOne) SetShardUID(s string) *CollaborationUpdateOne {
	cuo.mutation.SetShardUID(s)
	return cuo
}

// SetFolderUID sets the "folder_uid" field.
func (cuo *CollaborationUpdateOne) SetFolderUID(s string) *CollaborationUpdateOne {
	cuo.mutation.SetFolderUID(s)
	return cuo
}

// SetUserUID sets the "user_uid" field.
func (cuo *CollaborationUpdateOne) SetUserUID(s string) *CollaborationUpdateOne {
	cuo.mutation.SetUserUID(s)
	return cuo
}

// SetAuthority sets the "authority" field.
func (cuo *CollaborationUpdateOne) SetAuthority(s string) *CollaborationUpdateOne {
	cuo.mutation.SetAuthority(s)
	return cuo
}

// SetFolderNumber sets the "folder_number" field.
func (cuo *CollaborationUpdateOne) SetFolderNumber(s string) *CollaborationUpdateOne {
	cuo.mutation.SetFolderNumber(s)
	return cuo
}

// Mutation returns the CollaborationMutation object of the builder.
func (cuo *CollaborationUpdateOne) Mutation() *CollaborationMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollaborationUpdateOne) Select(field string, fields ...string) *CollaborationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collaboration entity.
func (cuo *CollaborationUpdateOne) Save(ctx context.Context) (*Collaboration, error) {
	var (
		err  error
		node *Collaboration
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollaborationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollaborationUpdateOne) SaveX(ctx context.Context) *Collaboration {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollaborationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollaborationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CollaborationUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := collaboration.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CollaborationUpdateOne) check() error {
	if v, ok := cuo.mutation.FolderNumber(); ok {
		if err := collaboration.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "Collaboration.folder_number": %w`, err)}
		}
	}
	return nil
}

func (cuo *CollaborationUpdateOne) sqlSave(ctx context.Context) (_node *Collaboration, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collaboration.Table,
			Columns: collaboration.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: collaboration.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "Collaboration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collaboration.FieldID)
		for _, f := range fields {
			if !collaboration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != collaboration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collaboration.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collaboration.FieldDeactivatedAt,
		})
	}
	if cuo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collaboration.FieldDeactivatedAt,
		})
	}
	if value, ok := cuo.mutation.ShardUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldShardUID,
		})
	}
	if value, ok := cuo.mutation.FolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldFolderUID,
		})
	}
	if value, ok := cuo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldUserUID,
		})
	}
	if value, ok := cuo.mutation.Authority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldAuthority,
		})
	}
	if value, ok := cuo.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: collaboration.FieldFolderNumber,
		})
	}
	_node = &Collaboration{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collaboration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
