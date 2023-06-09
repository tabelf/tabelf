// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/recharge"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RechargeUpdate is the builder for updating Recharge entities.
type RechargeUpdate struct {
	config
	hooks    []Hook
	mutation *RechargeMutation
}

// Where appends a list predicates to the RechargeUpdate builder.
func (ru *RechargeUpdate) Where(ps ...predicate.Recharge) *RechargeUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RechargeUpdate) SetUpdatedAt(t time.Time) *RechargeUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (ru *RechargeUpdate) SetDeactivatedAt(t time.Time) *RechargeUpdate {
	ru.mutation.SetDeactivatedAt(t)
	return ru
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (ru *RechargeUpdate) SetNillableDeactivatedAt(t *time.Time) *RechargeUpdate {
	if t != nil {
		ru.SetDeactivatedAt(*t)
	}
	return ru
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (ru *RechargeUpdate) ClearDeactivatedAt() *RechargeUpdate {
	ru.mutation.ClearDeactivatedAt()
	return ru
}

// SetTitle sets the "title" field.
func (ru *RechargeUpdate) SetTitle(s string) *RechargeUpdate {
	ru.mutation.SetTitle(s)
	return ru
}

// SetOriginAmount sets the "origin_amount" field.
func (ru *RechargeUpdate) SetOriginAmount(s string) *RechargeUpdate {
	ru.mutation.SetOriginAmount(s)
	return ru
}

// SetAmount sets the "amount" field.
func (ru *RechargeUpdate) SetAmount(s string) *RechargeUpdate {
	ru.mutation.SetAmount(s)
	return ru
}

// SetDescriptions sets the "descriptions" field.
func (ru *RechargeUpdate) SetDescriptions(s []string) *RechargeUpdate {
	ru.mutation.SetDescriptions(s)
	return ru
}

// SetThemeColor sets the "theme_color" field.
func (ru *RechargeUpdate) SetThemeColor(s string) *RechargeUpdate {
	ru.mutation.SetThemeColor(s)
	return ru
}

// SetNillableThemeColor sets the "theme_color" field if the given value is not nil.
func (ru *RechargeUpdate) SetNillableThemeColor(s *string) *RechargeUpdate {
	if s != nil {
		ru.SetThemeColor(*s)
	}
	return ru
}

// ClearThemeColor clears the value of the "theme_color" field.
func (ru *RechargeUpdate) ClearThemeColor() *RechargeUpdate {
	ru.mutation.ClearThemeColor()
	return ru
}

// SetYear sets the "year" field.
func (ru *RechargeUpdate) SetYear(i int) *RechargeUpdate {
	ru.mutation.ResetYear()
	ru.mutation.SetYear(i)
	return ru
}

// AddYear adds i to the "year" field.
func (ru *RechargeUpdate) AddYear(i int) *RechargeUpdate {
	ru.mutation.AddYear(i)
	return ru
}

// SetMonth sets the "month" field.
func (ru *RechargeUpdate) SetMonth(i int) *RechargeUpdate {
	ru.mutation.ResetMonth()
	ru.mutation.SetMonth(i)
	return ru
}

// AddMonth adds i to the "month" field.
func (ru *RechargeUpdate) AddMonth(i int) *RechargeUpdate {
	ru.mutation.AddMonth(i)
	return ru
}

// SetDefault sets the "default" field.
func (ru *RechargeUpdate) SetDefault(b bool) *RechargeUpdate {
	ru.mutation.SetDefault(b)
	return ru
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (ru *RechargeUpdate) SetNillableDefault(b *bool) *RechargeUpdate {
	if b != nil {
		ru.SetDefault(*b)
	}
	return ru
}

// Mutation returns the RechargeMutation object of the builder.
func (ru *RechargeUpdate) Mutation() *RechargeMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RechargeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RechargeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RechargeUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RechargeUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RechargeUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RechargeUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := recharge.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *RechargeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recharge.Table,
			Columns: recharge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: recharge.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recharge.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recharge.FieldDeactivatedAt,
		})
	}
	if ru.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recharge.FieldDeactivatedAt,
		})
	}
	if value, ok := ru.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldTitle,
		})
	}
	if value, ok := ru.mutation.OriginAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldOriginAmount,
		})
	}
	if value, ok := ru.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldAmount,
		})
	}
	if value, ok := ru.mutation.Descriptions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: recharge.FieldDescriptions,
		})
	}
	if value, ok := ru.mutation.ThemeColor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldThemeColor,
		})
	}
	if ru.mutation.ThemeColorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: recharge.FieldThemeColor,
		})
	}
	if value, ok := ru.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldYear,
		})
	}
	if value, ok := ru.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldYear,
		})
	}
	if value, ok := ru.mutation.Month(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldMonth,
		})
	}
	if value, ok := ru.mutation.AddedMonth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldMonth,
		})
	}
	if value, ok := ru.mutation.Default(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: recharge.FieldDefault,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recharge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RechargeUpdateOne is the builder for updating a single Recharge entity.
type RechargeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RechargeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RechargeUpdateOne) SetUpdatedAt(t time.Time) *RechargeUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (ruo *RechargeUpdateOne) SetDeactivatedAt(t time.Time) *RechargeUpdateOne {
	ruo.mutation.SetDeactivatedAt(t)
	return ruo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (ruo *RechargeUpdateOne) SetNillableDeactivatedAt(t *time.Time) *RechargeUpdateOne {
	if t != nil {
		ruo.SetDeactivatedAt(*t)
	}
	return ruo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (ruo *RechargeUpdateOne) ClearDeactivatedAt() *RechargeUpdateOne {
	ruo.mutation.ClearDeactivatedAt()
	return ruo
}

// SetTitle sets the "title" field.
func (ruo *RechargeUpdateOne) SetTitle(s string) *RechargeUpdateOne {
	ruo.mutation.SetTitle(s)
	return ruo
}

// SetOriginAmount sets the "origin_amount" field.
func (ruo *RechargeUpdateOne) SetOriginAmount(s string) *RechargeUpdateOne {
	ruo.mutation.SetOriginAmount(s)
	return ruo
}

// SetAmount sets the "amount" field.
func (ruo *RechargeUpdateOne) SetAmount(s string) *RechargeUpdateOne {
	ruo.mutation.SetAmount(s)
	return ruo
}

// SetDescriptions sets the "descriptions" field.
func (ruo *RechargeUpdateOne) SetDescriptions(s []string) *RechargeUpdateOne {
	ruo.mutation.SetDescriptions(s)
	return ruo
}

// SetThemeColor sets the "theme_color" field.
func (ruo *RechargeUpdateOne) SetThemeColor(s string) *RechargeUpdateOne {
	ruo.mutation.SetThemeColor(s)
	return ruo
}

// SetNillableThemeColor sets the "theme_color" field if the given value is not nil.
func (ruo *RechargeUpdateOne) SetNillableThemeColor(s *string) *RechargeUpdateOne {
	if s != nil {
		ruo.SetThemeColor(*s)
	}
	return ruo
}

// ClearThemeColor clears the value of the "theme_color" field.
func (ruo *RechargeUpdateOne) ClearThemeColor() *RechargeUpdateOne {
	ruo.mutation.ClearThemeColor()
	return ruo
}

// SetYear sets the "year" field.
func (ruo *RechargeUpdateOne) SetYear(i int) *RechargeUpdateOne {
	ruo.mutation.ResetYear()
	ruo.mutation.SetYear(i)
	return ruo
}

// AddYear adds i to the "year" field.
func (ruo *RechargeUpdateOne) AddYear(i int) *RechargeUpdateOne {
	ruo.mutation.AddYear(i)
	return ruo
}

// SetMonth sets the "month" field.
func (ruo *RechargeUpdateOne) SetMonth(i int) *RechargeUpdateOne {
	ruo.mutation.ResetMonth()
	ruo.mutation.SetMonth(i)
	return ruo
}

// AddMonth adds i to the "month" field.
func (ruo *RechargeUpdateOne) AddMonth(i int) *RechargeUpdateOne {
	ruo.mutation.AddMonth(i)
	return ruo
}

// SetDefault sets the "default" field.
func (ruo *RechargeUpdateOne) SetDefault(b bool) *RechargeUpdateOne {
	ruo.mutation.SetDefault(b)
	return ruo
}

// SetNillableDefault sets the "default" field if the given value is not nil.
func (ruo *RechargeUpdateOne) SetNillableDefault(b *bool) *RechargeUpdateOne {
	if b != nil {
		ruo.SetDefault(*b)
	}
	return ruo
}

// Mutation returns the RechargeMutation object of the builder.
func (ruo *RechargeUpdateOne) Mutation() *RechargeMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RechargeUpdateOne) Select(field string, fields ...string) *RechargeUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recharge entity.
func (ruo *RechargeUpdateOne) Save(ctx context.Context) (*Recharge, error) {
	var (
		err  error
		node *Recharge
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RechargeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RechargeUpdateOne) SaveX(ctx context.Context) *Recharge {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RechargeUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RechargeUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RechargeUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := recharge.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *RechargeUpdateOne) sqlSave(ctx context.Context) (_node *Recharge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recharge.Table,
			Columns: recharge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: recharge.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "Recharge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recharge.FieldID)
		for _, f := range fields {
			if !recharge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != recharge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recharge.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: recharge.FieldDeactivatedAt,
		})
	}
	if ruo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: recharge.FieldDeactivatedAt,
		})
	}
	if value, ok := ruo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldTitle,
		})
	}
	if value, ok := ruo.mutation.OriginAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldOriginAmount,
		})
	}
	if value, ok := ruo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldAmount,
		})
	}
	if value, ok := ruo.mutation.Descriptions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: recharge.FieldDescriptions,
		})
	}
	if value, ok := ruo.mutation.ThemeColor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recharge.FieldThemeColor,
		})
	}
	if ruo.mutation.ThemeColorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: recharge.FieldThemeColor,
		})
	}
	if value, ok := ruo.mutation.Year(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldYear,
		})
	}
	if value, ok := ruo.mutation.AddedYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldYear,
		})
	}
	if value, ok := ruo.mutation.Month(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldMonth,
		})
	}
	if value, ok := ruo.mutation.AddedMonth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: recharge.FieldMonth,
		})
	}
	if value, ok := ruo.mutation.Default(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: recharge.FieldDefault,
		})
	}
	_node = &Recharge{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recharge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
