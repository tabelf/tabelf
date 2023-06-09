// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/sharelink"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShareLinkUpdate is the builder for updating ShareLink entities.
type ShareLinkUpdate struct {
	config
	hooks    []Hook
	mutation *ShareLinkMutation
}

// Where appends a list predicates to the ShareLinkUpdate builder.
func (slu *ShareLinkUpdate) Where(ps ...predicate.ShareLink) *ShareLinkUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetUpdatedAt sets the "updated_at" field.
func (slu *ShareLinkUpdate) SetUpdatedAt(t time.Time) *ShareLinkUpdate {
	slu.mutation.SetUpdatedAt(t)
	return slu
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (slu *ShareLinkUpdate) SetDeactivatedAt(t time.Time) *ShareLinkUpdate {
	slu.mutation.SetDeactivatedAt(t)
	return slu
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (slu *ShareLinkUpdate) SetNillableDeactivatedAt(t *time.Time) *ShareLinkUpdate {
	if t != nil {
		slu.SetDeactivatedAt(*t)
	}
	return slu
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (slu *ShareLinkUpdate) ClearDeactivatedAt() *ShareLinkUpdate {
	slu.mutation.ClearDeactivatedAt()
	return slu
}

// SetFolderUID sets the "folder_uid" field.
func (slu *ShareLinkUpdate) SetFolderUID(s string) *ShareLinkUpdate {
	slu.mutation.SetFolderUID(s)
	return slu
}

// SetUserUID sets the "user_uid" field.
func (slu *ShareLinkUpdate) SetUserUID(s string) *ShareLinkUpdate {
	slu.mutation.SetUserUID(s)
	return slu
}

// SetAuthority sets the "authority" field.
func (slu *ShareLinkUpdate) SetAuthority(s string) *ShareLinkUpdate {
	slu.mutation.SetAuthority(s)
	return slu
}

// SetValidDay sets the "valid_day" field.
func (slu *ShareLinkUpdate) SetValidDay(i int) *ShareLinkUpdate {
	slu.mutation.ResetValidDay()
	slu.mutation.SetValidDay(i)
	return slu
}

// AddValidDay adds i to the "valid_day" field.
func (slu *ShareLinkUpdate) AddValidDay(i int) *ShareLinkUpdate {
	slu.mutation.AddValidDay(i)
	return slu
}

// SetExpiredAt sets the "expired_at" field.
func (slu *ShareLinkUpdate) SetExpiredAt(t time.Time) *ShareLinkUpdate {
	slu.mutation.SetExpiredAt(t)
	return slu
}

// SetRecentAt sets the "recent_at" field.
func (slu *ShareLinkUpdate) SetRecentAt(t time.Time) *ShareLinkUpdate {
	slu.mutation.SetRecentAt(t)
	return slu
}

// SetFolderNumber sets the "folder_number" field.
func (slu *ShareLinkUpdate) SetFolderNumber(s string) *ShareLinkUpdate {
	slu.mutation.SetFolderNumber(s)
	return slu
}

// Mutation returns the ShareLinkMutation object of the builder.
func (slu *ShareLinkUpdate) Mutation() *ShareLinkMutation {
	return slu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *ShareLinkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	slu.defaults()
	if len(slu.hooks) == 0 {
		if err = slu.check(); err != nil {
			return 0, err
		}
		affected, err = slu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShareLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = slu.check(); err != nil {
				return 0, err
			}
			slu.mutation = mutation
			affected, err = slu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(slu.hooks) - 1; i >= 0; i-- {
			if slu.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = slu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, slu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (slu *ShareLinkUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *ShareLinkUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *ShareLinkUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slu *ShareLinkUpdate) defaults() {
	if _, ok := slu.mutation.UpdatedAt(); !ok {
		v := sharelink.UpdateDefaultUpdatedAt()
		slu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slu *ShareLinkUpdate) check() error {
	if v, ok := slu.mutation.FolderNumber(); ok {
		if err := sharelink.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "ShareLink.folder_number": %w`, err)}
		}
	}
	return nil
}

func (slu *ShareLinkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sharelink.Table,
			Columns: sharelink.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sharelink.FieldID,
			},
		},
	}
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldUpdatedAt,
		})
	}
	if value, ok := slu.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldDeactivatedAt,
		})
	}
	if slu.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sharelink.FieldDeactivatedAt,
		})
	}
	if value, ok := slu.mutation.FolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldFolderUID,
		})
	}
	if value, ok := slu.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldUserUID,
		})
	}
	if value, ok := slu.mutation.Authority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldAuthority,
		})
	}
	if value, ok := slu.mutation.ValidDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sharelink.FieldValidDay,
		})
	}
	if value, ok := slu.mutation.AddedValidDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sharelink.FieldValidDay,
		})
	}
	if value, ok := slu.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldExpiredAt,
		})
	}
	if value, ok := slu.mutation.RecentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldRecentAt,
		})
	}
	if value, ok := slu.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldFolderNumber,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sharelink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ShareLinkUpdateOne is the builder for updating a single ShareLink entity.
type ShareLinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShareLinkMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (sluo *ShareLinkUpdateOne) SetUpdatedAt(t time.Time) *ShareLinkUpdateOne {
	sluo.mutation.SetUpdatedAt(t)
	return sluo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (sluo *ShareLinkUpdateOne) SetDeactivatedAt(t time.Time) *ShareLinkUpdateOne {
	sluo.mutation.SetDeactivatedAt(t)
	return sluo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (sluo *ShareLinkUpdateOne) SetNillableDeactivatedAt(t *time.Time) *ShareLinkUpdateOne {
	if t != nil {
		sluo.SetDeactivatedAt(*t)
	}
	return sluo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (sluo *ShareLinkUpdateOne) ClearDeactivatedAt() *ShareLinkUpdateOne {
	sluo.mutation.ClearDeactivatedAt()
	return sluo
}

// SetFolderUID sets the "folder_uid" field.
func (sluo *ShareLinkUpdateOne) SetFolderUID(s string) *ShareLinkUpdateOne {
	sluo.mutation.SetFolderUID(s)
	return sluo
}

// SetUserUID sets the "user_uid" field.
func (sluo *ShareLinkUpdateOne) SetUserUID(s string) *ShareLinkUpdateOne {
	sluo.mutation.SetUserUID(s)
	return sluo
}

// SetAuthority sets the "authority" field.
func (sluo *ShareLinkUpdateOne) SetAuthority(s string) *ShareLinkUpdateOne {
	sluo.mutation.SetAuthority(s)
	return sluo
}

// SetValidDay sets the "valid_day" field.
func (sluo *ShareLinkUpdateOne) SetValidDay(i int) *ShareLinkUpdateOne {
	sluo.mutation.ResetValidDay()
	sluo.mutation.SetValidDay(i)
	return sluo
}

// AddValidDay adds i to the "valid_day" field.
func (sluo *ShareLinkUpdateOne) AddValidDay(i int) *ShareLinkUpdateOne {
	sluo.mutation.AddValidDay(i)
	return sluo
}

// SetExpiredAt sets the "expired_at" field.
func (sluo *ShareLinkUpdateOne) SetExpiredAt(t time.Time) *ShareLinkUpdateOne {
	sluo.mutation.SetExpiredAt(t)
	return sluo
}

// SetRecentAt sets the "recent_at" field.
func (sluo *ShareLinkUpdateOne) SetRecentAt(t time.Time) *ShareLinkUpdateOne {
	sluo.mutation.SetRecentAt(t)
	return sluo
}

// SetFolderNumber sets the "folder_number" field.
func (sluo *ShareLinkUpdateOne) SetFolderNumber(s string) *ShareLinkUpdateOne {
	sluo.mutation.SetFolderNumber(s)
	return sluo
}

// Mutation returns the ShareLinkMutation object of the builder.
func (sluo *ShareLinkUpdateOne) Mutation() *ShareLinkMutation {
	return sluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *ShareLinkUpdateOne) Select(field string, fields ...string) *ShareLinkUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated ShareLink entity.
func (sluo *ShareLinkUpdateOne) Save(ctx context.Context) (*ShareLink, error) {
	var (
		err  error
		node *ShareLink
	)
	sluo.defaults()
	if len(sluo.hooks) == 0 {
		if err = sluo.check(); err != nil {
			return nil, err
		}
		node, err = sluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShareLinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sluo.check(); err != nil {
				return nil, err
			}
			sluo.mutation = mutation
			node, err = sluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sluo.hooks) - 1; i >= 0; i-- {
			if sluo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = sluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *ShareLinkUpdateOne) SaveX(ctx context.Context) *ShareLink {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *ShareLinkUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *ShareLinkUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sluo *ShareLinkUpdateOne) defaults() {
	if _, ok := sluo.mutation.UpdatedAt(); !ok {
		v := sharelink.UpdateDefaultUpdatedAt()
		sluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sluo *ShareLinkUpdateOne) check() error {
	if v, ok := sluo.mutation.FolderNumber(); ok {
		if err := sharelink.FolderNumberValidator(v); err != nil {
			return &ValidationError{Name: "folder_number", err: fmt.Errorf(`entschema: validator failed for field "ShareLink.folder_number": %w`, err)}
		}
	}
	return nil
}

func (sluo *ShareLinkUpdateOne) sqlSave(ctx context.Context) (_node *ShareLink, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sharelink.Table,
			Columns: sharelink.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sharelink.FieldID,
			},
		},
	}
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "ShareLink.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sharelink.FieldID)
		for _, f := range fields {
			if !sharelink.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != sharelink.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldUpdatedAt,
		})
	}
	if value, ok := sluo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldDeactivatedAt,
		})
	}
	if sluo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sharelink.FieldDeactivatedAt,
		})
	}
	if value, ok := sluo.mutation.FolderUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldFolderUID,
		})
	}
	if value, ok := sluo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldUserUID,
		})
	}
	if value, ok := sluo.mutation.Authority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldAuthority,
		})
	}
	if value, ok := sluo.mutation.ValidDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sharelink.FieldValidDay,
		})
	}
	if value, ok := sluo.mutation.AddedValidDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: sharelink.FieldValidDay,
		})
	}
	if value, ok := sluo.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldExpiredAt,
		})
	}
	if value, ok := sluo.mutation.RecentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sharelink.FieldRecentAt,
		})
	}
	if value, ok := sluo.mutation.FolderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sharelink.FieldFolderNumber,
		})
	}
	_node = &ShareLink{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sharelink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
