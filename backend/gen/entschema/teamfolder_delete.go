// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/teamfolder"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeamFolderDelete is the builder for deleting a TeamFolder entity.
type TeamFolderDelete struct {
	config
	hooks    []Hook
	mutation *TeamFolderMutation
}

// Where appends a list predicates to the TeamFolderDelete builder.
func (tfd *TeamFolderDelete) Where(ps ...predicate.TeamFolder) *TeamFolderDelete {
	tfd.mutation.Where(ps...)
	return tfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tfd *TeamFolderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tfd.hooks) == 0 {
		affected, err = tfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamFolderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tfd.mutation = mutation
			affected, err = tfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tfd.hooks) - 1; i >= 0; i-- {
			if tfd.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = tfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfd *TeamFolderDelete) ExecX(ctx context.Context) int {
	n, err := tfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tfd *TeamFolderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: teamfolder.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: teamfolder.FieldID,
			},
		},
	}
	if ps := tfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tfd.driver, _spec)
}

// TeamFolderDeleteOne is the builder for deleting a single TeamFolder entity.
type TeamFolderDeleteOne struct {
	tfd *TeamFolderDelete
}

// Exec executes the deletion query.
func (tfdo *TeamFolderDeleteOne) Exec(ctx context.Context) error {
	n, err := tfdo.tfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{teamfolder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tfdo *TeamFolderDeleteOne) ExecX(ctx context.Context) {
	tfdo.tfd.ExecX(ctx)
}
