// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/teamgroup"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamGroupUpdate is the builder for updating TeamGroup entities.
type TeamGroupUpdate struct {
	config
	hooks    []Hook
	mutation *TeamGroupMutation
}

// Where appends a list predicates to the TeamGroupUpdate builder.
func (tgu *TeamGroupUpdate) Where(ps ...predicate.TeamGroup) *TeamGroupUpdate {
	tgu.mutation.Where(ps...)
	return tgu
}

// SetUpdatedAt sets the "updated_at" field.
func (tgu *TeamGroupUpdate) SetUpdatedAt(t time.Time) *TeamGroupUpdate {
	tgu.mutation.SetUpdatedAt(t)
	return tgu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (tgu *TeamGroupUpdate) SetDeactivatedAt(t time.Time) *TeamGroupUpdate {
	tgu.mutation.SetDeactivatedAt(t)
	return tgu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (tgu *TeamGroupUpdate) SetNillableDeactivatedAt(t *time.Time) *TeamGroupUpdate {
	if t != nil {
		tgu.SetDeactivatedAt(*t)
	}
	return tgu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (tgu *TeamGroupUpdate) ClearDeactivatedAt() *TeamGroupUpdate {
	tgu.mutation.ClearDeactivatedAt()
	return tgu
}

// SetUserUID sets the "user_uid" field.
func (tgu *TeamGroupUpdate) SetUserUID(s string) *TeamGroupUpdate {
	tgu.mutation.SetUserUID(s)
	return tgu
}

// SetTeamUID sets the "team_uid" field.
func (tgu *TeamGroupUpdate) SetTeamUID(s string) *TeamGroupUpdate {
	tgu.mutation.SetTeamUID(s)
	return tgu
}

// Mutation returns the TeamGroupMutation object of the builder.
func (tgu *TeamGroupUpdate) Mutation() *TeamGroupMutation {
	return tgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tgu *TeamGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tgu.defaults()
	if len(tgu.hooks) == 0 {
		affected, err = tgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tgu.mutation = mutation
			affected, err = tgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tgu.hooks) - 1; i >= 0; i-- {
			if tgu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tgu *TeamGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := tgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tgu *TeamGroupUpdate) Exec(ctx context.Context) error {
	_, err := tgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tgu *TeamGroupUpdate) ExecX(ctx context.Context) {
	if err := tgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tgu *TeamGroupUpdate) defaults() {
	if _, ok := tgu.mutation.UpdatedAt(); !ok {
		v := teamgroup.UpdateDefaultUpdatedAt()
		tgu.mutation.SetUpdatedAt(v)
	}
}

func (tgu *TeamGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teamgroup.Table,
			Columns: teamgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: teamgroup.FieldID,
			},
		},
	}
	if ps := tgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teamgroup.FieldUpdatedAt,
		})
	}
	if value, ok := tgu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teamgroup.FieldDeactivatedAt,
		})
	}
	if tgu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teamgroup.FieldDeactivatedAt,
		})
	}
	if value, ok := tgu.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teamgroup.FieldUserUID,
		})
	}
	if value, ok := tgu.mutation.TeamUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teamgroup.FieldTeamUID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TeamGroupUpdateOne is the builder for updating a single TeamGroup entity.
type TeamGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamGroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tguo *TeamGroupUpdateOne) SetUpdatedAt(t time.Time) *TeamGroupUpdateOne {
	tguo.mutation.SetUpdatedAt(t)
	return tguo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (tguo *TeamGroupUpdateOne) SetDeactivatedAt(t time.Time) *TeamGroupUpdateOne {
	tguo.mutation.SetDeactivatedAt(t)
	return tguo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (tguo *TeamGroupUpdateOne) SetNillableDeactivatedAt(t *time.Time) *TeamGroupUpdateOne {
	if t != nil {
		tguo.SetDeactivatedAt(*t)
	}
	return tguo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (tguo *TeamGroupUpdateOne) ClearDeactivatedAt() *TeamGroupUpdateOne {
	tguo.mutation.ClearDeactivatedAt()
	return tguo
}

// SetUserUID sets the "user_uid" field.
func (tguo *TeamGroupUpdateOne) SetUserUID(s string) *TeamGroupUpdateOne {
	tguo.mutation.SetUserUID(s)
	return tguo
}

// SetTeamUID sets the "team_uid" field.
func (tguo *TeamGroupUpdateOne) SetTeamUID(s string) *TeamGroupUpdateOne {
	tguo.mutation.SetTeamUID(s)
	return tguo
}

// Mutation returns the TeamGroupMutation object of the builder.
func (tguo *TeamGroupUpdateOne) Mutation() *TeamGroupMutation {
	return tguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tguo *TeamGroupUpdateOne) Select(field string, fields ...string) *TeamGroupUpdateOne {
	tguo.fields = append([]string{field}, fields...)
	return tguo
}

// Save executes the query and returns the updated TeamGroup entity.
func (tguo *TeamGroupUpdateOne) Save(ctx context.Context) (*TeamGroup, error) {
	var (
		err  error
		node *TeamGroup
	)
	tguo.defaults()
	if len(tguo.hooks) == 0 {
		node, err = tguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tguo.mutation = mutation
			node, err = tguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tguo.hooks) - 1; i >= 0; i-- {
			if tguo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tguo *TeamGroupUpdateOne) SaveX(ctx context.Context) *TeamGroup {
	node, err := tguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tguo *TeamGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := tguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tguo *TeamGroupUpdateOne) ExecX(ctx context.Context) {
	if err := tguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tguo *TeamGroupUpdateOne) defaults() {
	if _, ok := tguo.mutation.UpdatedAt(); !ok {
		v := teamgroup.UpdateDefaultUpdatedAt()
		tguo.mutation.SetUpdatedAt(v)
	}
}

func (tguo *TeamGroupUpdateOne) sqlSave(ctx context.Context) (_node *TeamGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teamgroup.Table,
			Columns: teamgroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: teamgroup.FieldID,
			},
		},
	}
	id, ok := tguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "TeamGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teamgroup.FieldID)
		for _, f := range fields {
			if !teamgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != teamgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teamgroup.FieldUpdatedAt,
		})
	}
	if value, ok := tguo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: teamgroup.FieldDeactivatedAt,
		})
	}
	if tguo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: teamgroup.FieldDeactivatedAt,
		})
	}
	if value, ok := tguo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teamgroup.FieldUserUID,
		})
	}
	if value, ok := tguo.mutation.TeamUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teamgroup.FieldTeamUID,
		})
	}
	_node = &TeamGroup{config: tguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teamgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
