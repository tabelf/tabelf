// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/notice"
	"tabelf/backend/gen/entschema/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NoticeUpdate is the builder for updating Notice entities.
type NoticeUpdate struct {
	config
	hooks    []Hook
	mutation *NoticeMutation
}

// Where appends a list predicates to the NoticeUpdate builder.
func (nu *NoticeUpdate) Where(ps ...predicate.Notice) *NoticeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NoticeUpdate) SetUpdatedAt(t time.Time) *NoticeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (nu *NoticeUpdate) SetDeactivatedAt(t time.Time) *NoticeUpdate {
	nu.mutation.SetDeactivatedAt(t)
	return nu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (nu *NoticeUpdate) SetNillableDeactivatedAt(t *time.Time) *NoticeUpdate {
	if t != nil {
		nu.SetDeactivatedAt(*t)
	}
	return nu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (nu *NoticeUpdate) ClearDeactivatedAt() *NoticeUpdate {
	nu.mutation.ClearDeactivatedAt()
	return nu
}

// SetContent sets the "content" field.
func (nu *NoticeUpdate) SetContent(s string) *NoticeUpdate {
	nu.mutation.SetContent(s)
	return nu
}

// SetProcess sets the "process" field.
func (nu *NoticeUpdate) SetProcess(b bool) *NoticeUpdate {
	nu.mutation.SetProcess(b)
	return nu
}

// SetNillableProcess sets the "process" field if the given value is not nil.
func (nu *NoticeUpdate) SetNillableProcess(b *bool) *NoticeUpdate {
	if b != nil {
		nu.SetProcess(*b)
	}
	return nu
}

// Mutation returns the NoticeMutation object of the builder.
func (nu *NoticeUpdate) Mutation() *NoticeMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NoticeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nu.defaults()
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NoticeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NoticeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NoticeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NoticeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NoticeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := notice.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NoticeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notice.Table,
			Columns: notice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notice.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notice.FieldUpdatedAt,
		})
	}
	if value, ok := nu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notice.FieldDeactivatedAt,
		})
	}
	if nu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: notice.FieldDeactivatedAt,
		})
	}
	if value, ok := nu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notice.FieldContent,
		})
	}
	if value, ok := nu.mutation.Process(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notice.FieldProcess,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NoticeUpdateOne is the builder for updating a single Notice entity.
type NoticeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NoticeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NoticeUpdateOne) SetUpdatedAt(t time.Time) *NoticeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (nuo *NoticeUpdateOne) SetDeactivatedAt(t time.Time) *NoticeUpdateOne {
	nuo.mutation.SetDeactivatedAt(t)
	return nuo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (nuo *NoticeUpdateOne) SetNillableDeactivatedAt(t *time.Time) *NoticeUpdateOne {
	if t != nil {
		nuo.SetDeactivatedAt(*t)
	}
	return nuo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (nuo *NoticeUpdateOne) ClearDeactivatedAt() *NoticeUpdateOne {
	nuo.mutation.ClearDeactivatedAt()
	return nuo
}

// SetContent sets the "content" field.
func (nuo *NoticeUpdateOne) SetContent(s string) *NoticeUpdateOne {
	nuo.mutation.SetContent(s)
	return nuo
}

// SetProcess sets the "process" field.
func (nuo *NoticeUpdateOne) SetProcess(b bool) *NoticeUpdateOne {
	nuo.mutation.SetProcess(b)
	return nuo
}

// SetNillableProcess sets the "process" field if the given value is not nil.
func (nuo *NoticeUpdateOne) SetNillableProcess(b *bool) *NoticeUpdateOne {
	if b != nil {
		nuo.SetProcess(*b)
	}
	return nuo
}

// Mutation returns the NoticeMutation object of the builder.
func (nuo *NoticeUpdateOne) Mutation() *NoticeMutation {
	return nuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NoticeUpdateOne) Select(field string, fields ...string) *NoticeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notice entity.
func (nuo *NoticeUpdateOne) Save(ctx context.Context) (*Notice, error) {
	var (
		err  error
		node *Notice
	)
	nuo.defaults()
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NoticeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NoticeUpdateOne) SaveX(ctx context.Context) *Notice {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NoticeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NoticeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NoticeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := notice.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NoticeUpdateOne) sqlSave(ctx context.Context) (_node *Notice, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notice.Table,
			Columns: notice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: notice.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "Notice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notice.FieldID)
		for _, f := range fields {
			if !notice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != notice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notice.FieldUpdatedAt,
		})
	}
	if value, ok := nuo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: notice.FieldDeactivatedAt,
		})
	}
	if nuo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: notice.FieldDeactivatedAt,
		})
	}
	if value, ok := nuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: notice.FieldContent,
		})
	}
	if value, ok := nuo.mutation.Process(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: notice.FieldProcess,
		})
	}
	_node = &Notice{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}