// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/sharelink"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShareLinkDelete is the builder for deleting a ShareLink entity.
type ShareLinkDelete struct {
	config
	hooks    []Hook
	mutation *ShareLinkMutation
}

// Where appends a list predicates to the ShareLinkDelete builder.
func (sld *ShareLinkDelete) Where(ps ...predicate.ShareLink) *ShareLinkDelete {
	sld.mutation.Where(ps...)
	return sld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sld *ShareLinkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sld.hooks) == 0 {
		affected, err = sld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShareLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sld.mutation = mutation
			affected, err = sld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sld.hooks) - 1; i >= 0; i-- {
			if sld.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = sld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sld *ShareLinkDelete) ExecX(ctx context.Context) int {
	n, err := sld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sld *ShareLinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sharelink.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sharelink.FieldID,
			},
		},
	}
	if ps := sld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sld.driver, _spec)
}

// ShareLinkDeleteOne is the builder for deleting a single ShareLink entity.
type ShareLinkDeleteOne struct {
	sld *ShareLinkDelete
}

// Exec executes the deletion query.
func (sldo *ShareLinkDeleteOne) Exec(ctx context.Context) error {
	n, err := sldo.sld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sharelink.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sldo *ShareLinkDeleteOne) ExecX(ctx context.Context) {
	sldo.sld.ExecX(ctx)
}