// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PersonalFolderUpdate is the builder for updating PersonalFolder entities.
type PersonalFolderUpdate struct {
	config
	hooks    []Hook
	mutation *PersonalFolderMutation
}

// Where appends a list predicates to the PersonalFolderUpdate builder.
func (pfu *PersonalFolderUpdate) Where(ps ...predicate.PersonalFolder) *PersonalFolderUpdate {
	pfu.mutation.Where(ps...)
	return pfu
}

// SetUpdatedAt sets the "updated_at" field.
func (pfu *PersonalFolderUpdate) SetUpdatedAt(t time.Time) *PersonalFolderUpdate {
	pfu.mutation.SetUpdatedAt(t)
	return pfu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (pfu *PersonalFolderUpdate) SetDeactivatedAt(t time.Time) *PersonalFolderUpdate {
	pfu.mutation.SetDeactivatedAt(t)
	return pfu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (pfu *PersonalFolderUpdate) SetNillableDeactivatedAt(t *time.Time) *PersonalFolderUpdate {
	if t != nil {
		pfu.SetDeactivatedAt(*t)
	}
	return pfu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (pfu *PersonalFolderUpdate) ClearDeactivatedAt() *PersonalFolderUpdate {
	pfu.mutation.ClearDeactivatedAt()
	return pfu
}

// SetUserUID sets the "user_uid" field.
func (pfu *PersonalFolderUpdate) SetUserUID(s string) *PersonalFolderUpdate {
	pfu.mutation.SetUserUID(s)
	return pfu
}

// SetFolderName sets the "folder_name" field.
func (pfu *PersonalFolderUpdate) SetFolderName(s string) *PersonalFolderUpdate {
	pfu.mutation.SetFolderName(s)
	return pfu
}

// SetFolderNumber sets the "folder_number" field.
func (pfu *PersonalFolderUpdate) SetFolderNumber(s string) *PersonalFolderUpdate {
	pfu.mutation.SetFolderNumber(s)
	return pfu
}

// SetHasOpen sets the "has_open" field.
func (pfu *PersonalFolderUpdate) SetHasOpen(b bool) *PersonalFolderUpdate {
	pfu.mutation.SetHasOpen(b)
	return pfu
}

// SetNillableHasOpen sets the "has_open" field if the given value is not nil.
func (pfu *PersonalFolderUpdate) SetNillableHasOpen(b *bool) *PersonalFolderUpdate {
	if b != nil {
		pfu.SetHasOpen(*b)
	}
	return pfu
}

// Mutation returns the PersonalFolderMutation object of the builder.
func (pfu *PersonalFolderUpdate) Mutation() *PersonalFolderMutation {
	return pfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pfu *PersonalFolderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pfu.defaults()
	if len(pfu.hooks) == 0 {
		if err = pfu.check(); err != nil {
			return 0, err
		}
		affected, err = pfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalFolderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pfu.check(); err != nil {
				return 0, err
			}
			pfu.mutation = mutation
			affected, err = pfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pfu.hooks) - 1; i >= 0; i-- {
			if pfu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = pfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pfu *PersonalFolderUpdate) SaveX(ctx context.Context) int {
	affected, err := pfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pfu *PersonalFolderUpdate) Exec(ctx context.Context) error {
	_, err := pfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfu *PersonalFolderUpdate) ExecX(ctx context.Context) {
	if err := pfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pfu *PersonalFolderUpdate) defaults() {
	if _, ok := pfu.mutation.UpdatedAt(); !ok {
		v := personalfolder.UpdateDefaultUpdatedAt()
		pfu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfu *PersonalFolderUpdate) check() error {
	if v, ok := pfu.mutation.FolderNumber(); ok {
		if err := personalfolder.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "PersonalFolder.folder_number": %w`, err)}
		}
	}
	return nil
}

func (pfu *PersonalFolderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   personalfolder.Table,
			Columns: personalfolder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: personalfolder.FieldID,
			},
		},
	}
	if ps := pfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pfu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalfolder.FieldUpdatedAt,
		})
	}
	if value, ok := pfu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalfolder.FieldDeactivatedAt,
		})
	}
	if pfu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: personalfolder.FieldDeactivatedAt,
		})
	}
	if value, ok := pfu.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldUserUID,
		})
	}
	if value, ok := pfu.mutation.FolderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldFolderName,
		})
	}
	if value, ok := pfu.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldFolderNumber,
		})
	}
	if value, ok := pfu.mutation.HasOpen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: personalfolder.FieldHasOpen,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalfolder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PersonalFolderUpdateOne is the builder for updating a single PersonalFolder entity.
type PersonalFolderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonalFolderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pfuo *PersonalFolderUpdateOne) SetUpdatedAt(t time.Time) *PersonalFolderUpdateOne {
	pfuo.mutation.SetUpdatedAt(t)
	return pfuo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (pfuo *PersonalFolderUpdateOne) SetDeactivatedAt(t time.Time) *PersonalFolderUpdateOne {
	pfuo.mutation.SetDeactivatedAt(t)
	return pfuo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (pfuo *PersonalFolderUpdateOne) SetNillableDeactivatedAt(t *time.Time) *PersonalFolderUpdateOne {
	if t != nil {
		pfuo.SetDeactivatedAt(*t)
	}
	return pfuo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (pfuo *PersonalFolderUpdateOne) ClearDeactivatedAt() *PersonalFolderUpdateOne {
	pfuo.mutation.ClearDeactivatedAt()
	return pfuo
}

// SetUserUID sets the "user_uid" field.
func (pfuo *PersonalFolderUpdateOne) SetUserUID(s string) *PersonalFolderUpdateOne {
	pfuo.mutation.SetUserUID(s)
	return pfuo
}

// SetFolderName sets the "folder_name" field.
func (pfuo *PersonalFolderUpdateOne) SetFolderName(s string) *PersonalFolderUpdateOne {
	pfuo.mutation.SetFolderName(s)
	return pfuo
}

// SetFolderNumber sets the "folder_number" field.
func (pfuo *PersonalFolderUpdateOne) SetFolderNumber(s string) *PersonalFolderUpdateOne {
	pfuo.mutation.SetFolderNumber(s)
	return pfuo
}

// SetHasOpen sets the "has_open" field.
func (pfuo *PersonalFolderUpdateOne) SetHasOpen(b bool) *PersonalFolderUpdateOne {
	pfuo.mutation.SetHasOpen(b)
	return pfuo
}

// SetNillableHasOpen sets the "has_open" field if the given value is not nil.
func (pfuo *PersonalFolderUpdateOne) SetNillableHasOpen(b *bool) *PersonalFolderUpdateOne {
	if b != nil {
		pfuo.SetHasOpen(*b)
	}
	return pfuo
}

// Mutation returns the PersonalFolderMutation object of the builder.
func (pfuo *PersonalFolderUpdateOne) Mutation() *PersonalFolderMutation {
	return pfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pfuo *PersonalFolderUpdateOne) Select(field string, fields ...string) *PersonalFolderUpdateOne {
	pfuo.fields = append([]string{field}, fields...)
	return pfuo
}

// Save executes the query and returns the updated PersonalFolder entity.
func (pfuo *PersonalFolderUpdateOne) Save(ctx context.Context) (*PersonalFolder, error) {
	var (
		err  error
		node *PersonalFolder
	)
	pfuo.defaults()
	if len(pfuo.hooks) == 0 {
		if err = pfuo.check(); err != nil {
			return nil, err
		}
		node, err = pfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonalFolderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pfuo.check(); err != nil {
				return nil, err
			}
			pfuo.mutation = mutation
			node, err = pfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pfuo.hooks) - 1; i >= 0; i-- {
			if pfuo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = pfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pfuo *PersonalFolderUpdateOne) SaveX(ctx context.Context) *PersonalFolder {
	node, err := pfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pfuo *PersonalFolderUpdateOne) Exec(ctx context.Context) error {
	_, err := pfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfuo *PersonalFolderUpdateOne) ExecX(ctx context.Context) {
	if err := pfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pfuo *PersonalFolderUpdateOne) defaults() {
	if _, ok := pfuo.mutation.UpdatedAt(); !ok {
		v := personalfolder.UpdateDefaultUpdatedAt()
		pfuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfuo *PersonalFolderUpdateOne) check() error {
	if v, ok := pfuo.mutation.FolderNumber(); ok {
		if err := personalfolder.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "PersonalFolder.folder_number": %w`, err)}
		}
	}
	return nil
}

func (pfuo *PersonalFolderUpdateOne) sqlSave(ctx context.Context) (_node *PersonalFolder, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   personalfolder.Table,
			Columns: personalfolder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: personalfolder.FieldID,
			},
		},
	}
	id, ok := pfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "PersonalFolder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personalfolder.FieldID)
		for _, f := range fields {
			if !personalfolder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != personalfolder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pfuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalfolder.FieldUpdatedAt,
		})
	}
	if value, ok := pfuo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: personalfolder.FieldDeactivatedAt,
		})
	}
	if pfuo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: personalfolder.FieldDeactivatedAt,
		})
	}
	if value, ok := pfuo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldUserUID,
		})
	}
	if value, ok := pfuo.mutation.FolderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldFolderName,
		})
	}
	if value, ok := pfuo.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personalfolder.FieldFolderNumber,
		})
	}
	if value, ok := pfuo.mutation.HasOpen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: personalfolder.FieldHasOpen,
		})
	}
	_node = &PersonalFolder{config: pfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalfolder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
