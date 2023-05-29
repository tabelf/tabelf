// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema/communitycategory"
	"tabelf/backend/gen/entschema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityCategoryDelete is the builder for deleting a CommunityCategory entity.
type CommunityCategoryDelete struct {
	config
	hooks    []Hook
	mutation *CommunityCategoryMutation
}

// Where appends a list predicates to the CommunityCategoryDelete builder.
func (ccd *CommunityCategoryDelete) Where(ps ...predicate.CommunityCategory) *CommunityCategoryDelete {
	ccd.mutation.Where(ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *CommunityCategoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccd.hooks) == 0 {
		affected, err = ccd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunityCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccd.mutation = mutation
			affected, err = ccd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccd.hooks) - 1; i >= 0; i-- {
			if ccd.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = ccd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *CommunityCategoryDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *CommunityCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: communitycategory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: communitycategory.FieldID,
			},
		},
	}
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
}

// CommunityCategoryDeleteOne is the builder for deleting a single CommunityCategory entity.
type CommunityCategoryDeleteOne struct {
	ccd *CommunityCategoryDelete
}

// Exec executes the deletion query.
func (ccdo *CommunityCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{communitycategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *CommunityCategoryDeleteOne) ExecX(ctx context.Context) {
	ccdo.ccd.ExecX(ctx)
}
