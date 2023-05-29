// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/workspace"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceUpdate is the builder for updating Workspace entities.
type WorkspaceUpdate struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where appends a list predicates to the WorkspaceUpdate builder.
func (wu *WorkspaceUpdate) Where(ps ...predicate.Workspace) *WorkspaceUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WorkspaceUpdate) SetUpdatedAt(t time.Time) *WorkspaceUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (wu *WorkspaceUpdate) SetDeactivatedAt(t time.Time) *WorkspaceUpdate {
	wu.mutation.SetDeactivatedAt(t)
	return wu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillableDeactivatedAt(t *time.Time) *WorkspaceUpdate {
	if t != nil {
		wu.SetDeactivatedAt(*t)
	}
	return wu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (wu *WorkspaceUpdate) ClearDeactivatedAt() *WorkspaceUpdate {
	wu.mutation.ClearDeactivatedAt()
	return wu
}

// SetName sets the "name" field.
func (wu *WorkspaceUpdate) SetName(s string) *WorkspaceUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetType sets the "type" field.
func (wu *WorkspaceUpdate) SetType(i int) *WorkspaceUpdate {
	wu.mutation.ResetType()
	wu.mutation.SetType(i)
	return wu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillableType(i *int) *WorkspaceUpdate {
	if i != nil {
		wu.SetType(*i)
	}
	return wu
}

// AddType adds i to the "type" field.
func (wu *WorkspaceUpdate) AddType(i int) *WorkspaceUpdate {
	wu.mutation.AddType(i)
	return wu
}

// SetUserUID sets the "user_uid" field.
func (wu *WorkspaceUpdate) SetUserUID(s string) *WorkspaceUpdate {
	wu.mutation.SetUserUID(s)
	return wu
}

// SetPersonalFolderUID sets the "personal_folder_uid" field.
func (wu *WorkspaceUpdate) SetPersonalFolderUID(s string) *WorkspaceUpdate {
	wu.mutation.SetPersonalFolderUID(s)
	return wu
}

// SetNillablePersonalFolderUID sets the "personal_folder_uid" field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillablePersonalFolderUID(s *string) *WorkspaceUpdate {
	if s != nil {
		wu.SetPersonalFolderUID(*s)
	}
	return wu
}

// ClearPersonalFolderUID clears the value of the "personal_folder_uid" field.
func (wu *WorkspaceUpdate) ClearPersonalFolderUID() *WorkspaceUpdate {
	wu.mutation.ClearPersonalFolderUID()
	return wu
}

// SetTeamFolderUID sets the "team_folder_uid" field.
func (wu *WorkspaceUpdate) SetTeamFolderUID(s string) *WorkspaceUpdate {
	wu.mutation.SetTeamFolderUID(s)
	return wu
}

// SetNillableTeamFolderUID sets the "team_folder_uid" field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillableTeamFolderUID(s *string) *WorkspaceUpdate {
	if s != nil {
		wu.SetTeamFolderUID(*s)
	}
	return wu
}

// ClearTeamFolderUID clears the value of the "team_folder_uid" field.
func (wu *WorkspaceUpdate) ClearTeamFolderUID() *WorkspaceUpdate {
	wu.mutation.ClearTeamFolderUID()
	return wu
}

// SetIsOpen sets the "is_open" field.
func (wu *WorkspaceUpdate) SetIsOpen(b bool) *WorkspaceUpdate {
	wu.mutation.SetIsOpen(b)
	return wu
}

// SetNillableIsOpen sets the "is_open" field if the given value is not nil.
func (wu *WorkspaceUpdate) SetNillableIsOpen(b *bool) *WorkspaceUpdate {
	if b != nil {
		wu.SetIsOpen(*b)
	}
	return wu
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wu *WorkspaceUpdate) Mutation() *WorkspaceMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkspaceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkspaceUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkspaceUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkspaceUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WorkspaceUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := workspace.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

func (wu *WorkspaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: workspace.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
	}
	if value, ok := wu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldDeactivatedAt,
		})
	}
	if wu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workspace.FieldDeactivatedAt,
		})
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workspace.FieldType,
		})
	}
	if value, ok := wu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workspace.FieldType,
		})
	}
	if value, ok := wu.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldUserUID,
		})
	}
	if value, ok := wu.mutation.PersonalFolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldPersonalFolderUID,
		})
	}
	if wu.mutation.PersonalFolderUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldPersonalFolderUID,
		})
	}
	if value, ok := wu.mutation.TeamFolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldTeamFolderUID,
		})
	}
	if wu.mutation.TeamFolderUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldTeamFolderUID,
		})
	}
	if value, ok := wu.mutation.IsOpen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workspace.FieldIsOpen,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkspaceUpdateOne is the builder for updating a single Workspace entity.
type WorkspaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkspaceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WorkspaceUpdateOne) SetUpdatedAt(t time.Time) *WorkspaceUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (wuo *WorkspaceUpdateOne) SetDeactivatedAt(t time.Time) *WorkspaceUpdateOne {
	wuo.mutation.SetDeactivatedAt(t)
	return wuo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillableDeactivatedAt(t *time.Time) *WorkspaceUpdateOne {
	if t != nil {
		wuo.SetDeactivatedAt(*t)
	}
	return wuo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (wuo *WorkspaceUpdateOne) ClearDeactivatedAt() *WorkspaceUpdateOne {
	wuo.mutation.ClearDeactivatedAt()
	return wuo
}

// SetName sets the "name" field.
func (wuo *WorkspaceUpdateOne) SetName(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetType sets the "type" field.
func (wuo *WorkspaceUpdateOne) SetType(i int) *WorkspaceUpdateOne {
	wuo.mutation.ResetType()
	wuo.mutation.SetType(i)
	return wuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillableType(i *int) *WorkspaceUpdateOne {
	if i != nil {
		wuo.SetType(*i)
	}
	return wuo
}

// AddType adds i to the "type" field.
func (wuo *WorkspaceUpdateOne) AddType(i int) *WorkspaceUpdateOne {
	wuo.mutation.AddType(i)
	return wuo
}

// SetUserUID sets the "user_uid" field.
func (wuo *WorkspaceUpdateOne) SetUserUID(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetUserUID(s)
	return wuo
}

// SetPersonalFolderUID sets the "personal_folder_uid" field.
func (wuo *WorkspaceUpdateOne) SetPersonalFolderUID(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetPersonalFolderUID(s)
	return wuo
}

// SetNillablePersonalFolderUID sets the "personal_folder_uid" field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillablePersonalFolderUID(s *string) *WorkspaceUpdateOne {
	if s != nil {
		wuo.SetPersonalFolderUID(*s)
	}
	return wuo
}

// ClearPersonalFolderUID clears the value of the "personal_folder_uid" field.
func (wuo *WorkspaceUpdateOne) ClearPersonalFolderUID() *WorkspaceUpdateOne {
	wuo.mutation.ClearPersonalFolderUID()
	return wuo
}

// SetTeamFolderUID sets the "team_folder_uid" field.
func (wuo *WorkspaceUpdateOne) SetTeamFolderUID(s string) *WorkspaceUpdateOne {
	wuo.mutation.SetTeamFolderUID(s)
	return wuo
}

// SetNillableTeamFolderUID sets the "team_folder_uid" field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillableTeamFolderUID(s *string) *WorkspaceUpdateOne {
	if s != nil {
		wuo.SetTeamFolderUID(*s)
	}
	return wuo
}

// ClearTeamFolderUID clears the value of the "team_folder_uid" field.
func (wuo *WorkspaceUpdateOne) ClearTeamFolderUID() *WorkspaceUpdateOne {
	wuo.mutation.ClearTeamFolderUID()
	return wuo
}

// SetIsOpen sets the "is_open" field.
func (wuo *WorkspaceUpdateOne) SetIsOpen(b bool) *WorkspaceUpdateOne {
	wuo.mutation.SetIsOpen(b)
	return wuo
}

// SetNillableIsOpen sets the "is_open" field if the given value is not nil.
func (wuo *WorkspaceUpdateOne) SetNillableIsOpen(b *bool) *WorkspaceUpdateOne {
	if b != nil {
		wuo.SetIsOpen(*b)
	}
	return wuo
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wuo *WorkspaceUpdateOne) Mutation() *WorkspaceMutation {
	return wuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkspaceUpdateOne) Select(field string, fields ...string) *WorkspaceUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Workspace entity.
func (wuo *WorkspaceUpdateOne) Save(ctx context.Context) (*Workspace, error) {
	var (
		err  error
		node *Workspace
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkspaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) SaveX(ctx context.Context) *Workspace {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkspaceUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkspaceUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WorkspaceUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := workspace.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

func (wuo *WorkspaceUpdateOne) sqlSave(ctx context.Context) (_node *Workspace, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workspace.Table,
			Columns: workspace.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: workspace.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "Workspace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workspace.FieldID)
		for _, f := range fields {
			if !workspace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != workspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldUpdatedAt,
		})
	}
	if value, ok := wuo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workspace.FieldDeactivatedAt,
		})
	}
	if wuo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workspace.FieldDeactivatedAt,
		})
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldName,
		})
	}
	if value, ok := wuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workspace.FieldType,
		})
	}
	if value, ok := wuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workspace.FieldType,
		})
	}
	if value, ok := wuo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldUserUID,
		})
	}
	if value, ok := wuo.mutation.PersonalFolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldPersonalFolderUID,
		})
	}
	if wuo.mutation.PersonalFolderUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldPersonalFolderUID,
		})
	}
	if value, ok := wuo.mutation.TeamFolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workspace.FieldTeamFolderUID,
		})
	}
	if wuo.mutation.TeamFolderUIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workspace.FieldTeamFolderUID,
		})
	}
	if value, ok := wuo.mutation.IsOpen(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workspace.FieldIsOpen,
		})
	}
	_node = &Workspace{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workspace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
