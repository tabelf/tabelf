// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/teamgroup"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamGroupDelete is the builder for deleting a TeamGroup entity.
type TeamGroupDelete struct {
	config
	hooks    []Hook
	mutation *TeamGroupMutation
}

// Where appends a list predicates to the TeamGroupDelete builder.
func (tgd *TeamGroupDelete) Where(ps ...predicate.TeamGroup) *TeamGroupDelete {
	tgd.mutation.Where(ps...)
	return tgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tgd *TeamGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tgd.hooks) == 0 {
		affected, err = tgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tgd.mutation = mutation
			affected, err = tgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tgd.hooks) - 1; i >= 0; i-- {
			if tgd.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tgd *TeamGroupDelete) ExecX(ctx context.Context) int {
	n, err := tgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tgd *TeamGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: teamgroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: teamgroup.FieldID,
			},
		},
	}
	if ps := tgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tgd.driver, _spec)
}

// TeamGroupDeleteOne is the builder for deleting a single TeamGroup entity.
type TeamGroupDeleteOne struct {
	tgd *TeamGroupDelete
}

// Exec executes the deletion query.
func (tgdo *TeamGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := tgdo.tgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teamgroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tgdo *TeamGroupDeleteOne) ExecX(ctx context.Context) {
	tgdo.tgd.ExecX(ctx)
}
