// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"tabelf/backend/gen/entschema/payorder"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/spec/schema"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PayOrderUpdate is the builder for updating PayOrder entities.
type PayOrderUpdate struct {
	config
	hooks    []Hook
	mutation *PayOrderMutation
}

// Where appends a list predicates to the PayOrderUpdate builder.
func (pou *PayOrderUpdate) Where(ps ...predicate.PayOrder) *PayOrderUpdate {
	pou.mutation.Where(ps...)
	return pou
}

// SetUpdatedAt sets the "updated_at" field.
func (pou *PayOrderUpdate) SetUpdatedAt(t time.Time) *PayOrderUpdate {
	pou.mutation.SetUpdatedAt(t)
	return pou
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (pou *PayOrderUpdate) SetDeactivatedAt(t time.Time) *PayOrderUpdate {
	pou.mutation.SetDeactivatedAt(t)
	return pou
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableDeactivatedAt(t *time.Time) *PayOrderUpdate {
	if t != nil {
		pou.SetDeactivatedAt(*t)
	}
	return pou
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (pou *PayOrderUpdate) ClearDeactivatedAt() *PayOrderUpdate {
	pou.mutation.ClearDeactivatedAt()
	return pou
}

// SetOrderNumber sets the "order_number" field.
func (pou *PayOrderUpdate) SetOrderNumber(s string) *PayOrderUpdate {
	pou.mutation.SetOrderNumber(s)
	return pou
}

// SetOrderType sets the "order_type" field.
func (pou *PayOrderUpdate) SetOrderType(s string) *PayOrderUpdate {
	pou.mutation.SetOrderType(s)
	return pou
}

// SetUserUID sets the "user_uid" field.
func (pou *PayOrderUpdate) SetUserUID(s string) *PayOrderUpdate {
	pou.mutation.SetUserUID(s)
	return pou
}

// SetPayMethod sets the "pay_method" field.
func (pou *PayOrderUpdate) SetPayMethod(s string) *PayOrderUpdate {
	pou.mutation.SetPayMethod(s)
	return pou
}

// SetPaymentAmount sets the "payment_amount" field.
func (pou *PayOrderUpdate) SetPaymentAmount(s string) *PayOrderUpdate {
	pou.mutation.SetPaymentAmount(s)
	return pou
}

// SetTotalPrice sets the "total_price" field.
func (pou *PayOrderUpdate) SetTotalPrice(s string) *PayOrderUpdate {
	pou.mutation.SetTotalPrice(s)
	return pou
}

// SetRechargeRecord sets the "recharge_record" field.
func (pou *PayOrderUpdate) SetRechargeRecord(sr schema.RechargeRecord) *PayOrderUpdate {
	pou.mutation.SetRechargeRecord(sr)
	return pou
}

// SetStatus sets the "status" field.
func (pou *PayOrderUpdate) SetStatus(s string) *PayOrderUpdate {
	pou.mutation.SetStatus(s)
	return pou
}

// SetTransactionNumber sets the "transaction_number" field.
func (pou *PayOrderUpdate) SetTransactionNumber(s string) *PayOrderUpdate {
	pou.mutation.SetTransactionNumber(s)
	return pou
}

// SetNillableTransactionNumber sets the "transaction_number" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableTransactionNumber(s *string) *PayOrderUpdate {
	if s != nil {
		pou.SetTransactionNumber(*s)
	}
	return pou
}

// ClearTransactionNumber clears the value of the "transaction_number" field.
func (pou *PayOrderUpdate) ClearTransactionNumber() *PayOrderUpdate {
	pou.mutation.ClearTransactionNumber()
	return pou
}

// SetThirdpartyNumber sets the "thirdparty_number" field.
func (pou *PayOrderUpdate) SetThirdpartyNumber(s string) *PayOrderUpdate {
	pou.mutation.SetThirdpartyNumber(s)
	return pou
}

// SetNillableThirdpartyNumber sets the "thirdparty_number" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableThirdpartyNumber(s *string) *PayOrderUpdate {
	if s != nil {
		pou.SetThirdpartyNumber(*s)
	}
	return pou
}

// ClearThirdpartyNumber clears the value of the "thirdparty_number" field.
func (pou *PayOrderUpdate) ClearThirdpartyNumber() *PayOrderUpdate {
	pou.mutation.ClearThirdpartyNumber()
	return pou
}

// SetMchID sets the "mch_id" field.
func (pou *PayOrderUpdate) SetMchID(s string) *PayOrderUpdate {
	pou.mutation.SetMchID(s)
	return pou
}

// SetNillableMchID sets the "mch_id" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableMchID(s *string) *PayOrderUpdate {
	if s != nil {
		pou.SetMchID(*s)
	}
	return pou
}

// ClearMchID clears the value of the "mch_id" field.
func (pou *PayOrderUpdate) ClearMchID() *PayOrderUpdate {
	pou.mutation.ClearMchID()
	return pou
}

// SetOpenID sets the "open_id" field.
func (pou *PayOrderUpdate) SetOpenID(s string) *PayOrderUpdate {
	pou.mutation.SetOpenID(s)
	return pou
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableOpenID(s *string) *PayOrderUpdate {
	if s != nil {
		pou.SetOpenID(*s)
	}
	return pou
}

// ClearOpenID clears the value of the "open_id" field.
func (pou *PayOrderUpdate) ClearOpenID() *PayOrderUpdate {
	pou.mutation.ClearOpenID()
	return pou
}

// SetMemberExpired sets the "member_expired" field.
func (pou *PayOrderUpdate) SetMemberExpired(t time.Time) *PayOrderUpdate {
	pou.mutation.SetMemberExpired(t)
	return pou
}

// SetNillableMemberExpired sets the "member_expired" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableMemberExpired(t *time.Time) *PayOrderUpdate {
	if t != nil {
		pou.SetMemberExpired(*t)
	}
	return pou
}

// ClearMemberExpired clears the value of the "member_expired" field.
func (pou *PayOrderUpdate) ClearMemberExpired() *PayOrderUpdate {
	pou.mutation.ClearMemberExpired()
	return pou
}

// SetCancelEvent sets the "cancel_event" field.
func (pou *PayOrderUpdate) SetCancelEvent(s string) *PayOrderUpdate {
	pou.mutation.SetCancelEvent(s)
	return pou
}

// SetNillableCancelEvent sets the "cancel_event" field if the given value is not nil.
func (pou *PayOrderUpdate) SetNillableCancelEvent(s *string) *PayOrderUpdate {
	if s != nil {
		pou.SetCancelEvent(*s)
	}
	return pou
}

// ClearCancelEvent clears the value of the "cancel_event" field.
func (pou *PayOrderUpdate) ClearCancelEvent() *PayOrderUpdate {
	pou.mutation.ClearCancelEvent()
	return pou
}

// Mutation returns the PayOrderMutation object of the builder.
func (pou *PayOrderUpdate) Mutation() *PayOrderMutation {
	return pou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pou *PayOrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pou.defaults()
	if len(pou.hooks) == 0 {
		if err = pou.check(); err != nil {
			return 0, err
		}
		affected, err = pou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pou.check(); err != nil {
				return 0, err
			}
			pou.mutation = mutation
			affected, err = pou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pou.hooks) - 1; i >= 0; i-- {
			if pou.hooks[i] == nil {
				return 0, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = pou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pou *PayOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := pou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pou *PayOrderUpdate) Exec(ctx context.Context) error {
	_, err := pou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pou *PayOrderUpdate) ExecX(ctx context.Context) {
	if err := pou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pou *PayOrderUpdate) defaults() {
	if _, ok := pou.mutation.UpdatedAt(); !ok {
		v := payorder.UpdateDefaultUpdatedAt()
		pou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pou *PayOrderUpdate) check() error {
	if v, ok := pou.mutation.Status(); ok {
		if err := payorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entschema: validator failed for field "PayOrder.status": %w`, err)}
		}
	}
	if v, ok := pou.mutation.CancelEvent(); ok {
		if err := payorder.CancelEventValidator(v); err != nil {
			return &ValidationError{Name: "cancel_event", err: fmt.Errorf(`entschema: validator failed for field "PayOrder.cancel_event": %w`, err)}
		}
	}
	return nil
}

func (pou *PayOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payorder.Table,
			Columns: payorder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: payorder.FieldID,
			},
		},
	}
	if ps := pou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldUpdatedAt,
		})
	}
	if value, ok := pou.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldDeactivatedAt,
		})
	}
	if pou.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: payorder.FieldDeactivatedAt,
		})
	}
	if value, ok := pou.mutation.OrderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOrderNumber,
		})
	}
	if value, ok := pou.mutation.OrderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOrderType,
		})
	}
	if value, ok := pou.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldUserUID,
		})
	}
	if value, ok := pou.mutation.PayMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldPayMethod,
		})
	}
	if value, ok := pou.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldPaymentAmount,
		})
	}
	if value, ok := pou.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldTotalPrice,
		})
	}
	if value, ok := pou.mutation.RechargeRecord(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: payorder.FieldRechargeRecord,
		})
	}
	if value, ok := pou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldStatus,
		})
	}
	if value, ok := pou.mutation.TransactionNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldTransactionNumber,
		})
	}
	if pou.mutation.TransactionNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldTransactionNumber,
		})
	}
	if value, ok := pou.mutation.ThirdpartyNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldThirdpartyNumber,
		})
	}
	if pou.mutation.ThirdpartyNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldThirdpartyNumber,
		})
	}
	if value, ok := pou.mutation.MchID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldMchID,
		})
	}
	if pou.mutation.MchIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldMchID,
		})
	}
	if value, ok := pou.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOpenID,
		})
	}
	if pou.mutation.OpenIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldOpenID,
		})
	}
	if value, ok := pou.mutation.MemberExpired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldMemberExpired,
		})
	}
	if pou.mutation.MemberExpiredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: payorder.FieldMemberExpired,
		})
	}
	if value, ok := pou.mutation.CancelEvent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldCancelEvent,
		})
	}
	if pou.mutation.CancelEventCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldCancelEvent,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PayOrderUpdateOne is the builder for updating a single PayOrder entity.
type PayOrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PayOrderMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pouo *PayOrderUpdateOne) SetUpdatedAt(t time.Time) *PayOrderUpdateOne {
	pouo.mutation.SetUpdatedAt(t)
	return pouo
}

// SetDeactivatedAt sets the "deactivated_at" field.
func (pouo *PayOrderUpdateOne) SetDeactivatedAt(t time.Time) *PayOrderUpdateOne {
	pouo.mutation.SetDeactivatedAt(t)
	return pouo
}

// SetNillableDeactivatedAt sets the "deactivated_at" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableDeactivatedAt(t *time.Time) *PayOrderUpdateOne {
	if t != nil {
		pouo.SetDeactivatedAt(*t)
	}
	return pouo
}

// ClearDeactivatedAt clears the value of the "deactivated_at" field.
func (pouo *PayOrderUpdateOne) ClearDeactivatedAt() *PayOrderUpdateOne {
	pouo.mutation.ClearDeactivatedAt()
	return pouo
}

// SetOrderNumber sets the "order_number" field.
func (pouo *PayOrderUpdateOne) SetOrderNumber(s string) *PayOrderUpdateOne {
	pouo.mutation.SetOrderNumber(s)
	return pouo
}

// SetOrderType sets the "order_type" field.
func (pouo *PayOrderUpdateOne) SetOrderType(s string) *PayOrderUpdateOne {
	pouo.mutation.SetOrderType(s)
	return pouo
}

// SetUserUID sets the "user_uid" field.
func (pouo *PayOrderUpdateOne) SetUserUID(s string) *PayOrderUpdateOne {
	pouo.mutation.SetUserUID(s)
	return pouo
}

// SetPayMethod sets the "pay_method" field.
func (pouo *PayOrderUpdateOne) SetPayMethod(s string) *PayOrderUpdateOne {
	pouo.mutation.SetPayMethod(s)
	return pouo
}

// SetPaymentAmount sets the "payment_amount" field.
func (pouo *PayOrderUpdateOne) SetPaymentAmount(s string) *PayOrderUpdateOne {
	pouo.mutation.SetPaymentAmount(s)
	return pouo
}

// SetTotalPrice sets the "total_price" field.
func (pouo *PayOrderUpdateOne) SetTotalPrice(s string) *PayOrderUpdateOne {
	pouo.mutation.SetTotalPrice(s)
	return pouo
}

// SetRechargeRecord sets the "recharge_record" field.
func (pouo *PayOrderUpdateOne) SetRechargeRecord(sr schema.RechargeRecord) *PayOrderUpdateOne {
	pouo.mutation.SetRechargeRecord(sr)
	return pouo
}

// SetStatus sets the "status" field.
func (pouo *PayOrderUpdateOne) SetStatus(s string) *PayOrderUpdateOne {
	pouo.mutation.SetStatus(s)
	return pouo
}

// SetTransactionNumber sets the "transaction_number" field.
func (pouo *PayOrderUpdateOne) SetTransactionNumber(s string) *PayOrderUpdateOne {
	pouo.mutation.SetTransactionNumber(s)
	return pouo
}

// SetNillableTransactionNumber sets the "transaction_number" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableTransactionNumber(s *string) *PayOrderUpdateOne {
	if s != nil {
		pouo.SetTransactionNumber(*s)
	}
	return pouo
}

// ClearTransactionNumber clears the value of the "transaction_number" field.
func (pouo *PayOrderUpdateOne) ClearTransactionNumber() *PayOrderUpdateOne {
	pouo.mutation.ClearTransactionNumber()
	return pouo
}

// SetThirdpartyNumber sets the "thirdparty_number" field.
func (pouo *PayOrderUpdateOne) SetThirdpartyNumber(s string) *PayOrderUpdateOne {
	pouo.mutation.SetThirdpartyNumber(s)
	return pouo
}

// SetNillableThirdpartyNumber sets the "thirdparty_number" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableThirdpartyNumber(s *string) *PayOrderUpdateOne {
	if s != nil {
		pouo.SetThirdpartyNumber(*s)
	}
	return pouo
}

// ClearThirdpartyNumber clears the value of the "thirdparty_number" field.
func (pouo *PayOrderUpdateOne) ClearThirdpartyNumber() *PayOrderUpdateOne {
	pouo.mutation.ClearThirdpartyNumber()
	return pouo
}

// SetMchID sets the "mch_id" field.
func (pouo *PayOrderUpdateOne) SetMchID(s string) *PayOrderUpdateOne {
	pouo.mutation.SetMchID(s)
	return pouo
}

// SetNillableMchID sets the "mch_id" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableMchID(s *string) *PayOrderUpdateOne {
	if s != nil {
		pouo.SetMchID(*s)
	}
	return pouo
}

// ClearMchID clears the value of the "mch_id" field.
func (pouo *PayOrderUpdateOne) ClearMchID() *PayOrderUpdateOne {
	pouo.mutation.ClearMchID()
	return pouo
}

// SetOpenID sets the "open_id" field.
func (pouo *PayOrderUpdateOne) SetOpenID(s string) *PayOrderUpdateOne {
	pouo.mutation.SetOpenID(s)
	return pouo
}

// SetNillableOpenID sets the "open_id" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableOpenID(s *string) *PayOrderUpdateOne {
	if s != nil {
		pouo.SetOpenID(*s)
	}
	return pouo
}

// ClearOpenID clears the value of the "open_id" field.
func (pouo *PayOrderUpdateOne) ClearOpenID() *PayOrderUpdateOne {
	pouo.mutation.ClearOpenID()
	return pouo
}

// SetMemberExpired sets the "member_expired" field.
func (pouo *PayOrderUpdateOne) SetMemberExpired(t time.Time) *PayOrderUpdateOne {
	pouo.mutation.SetMemberExpired(t)
	return pouo
}

// SetNillableMemberExpired sets the "member_expired" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableMemberExpired(t *time.Time) *PayOrderUpdateOne {
	if t != nil {
		pouo.SetMemberExpired(*t)
	}
	return pouo
}

// ClearMemberExpired clears the value of the "member_expired" field.
func (pouo *PayOrderUpdateOne) ClearMemberExpired() *PayOrderUpdateOne {
	pouo.mutation.ClearMemberExpired()
	return pouo
}

// SetCancelEvent sets the "cancel_event" field.
func (pouo *PayOrderUpdateOne) SetCancelEvent(s string) *PayOrderUpdateOne {
	pouo.mutation.SetCancelEvent(s)
	return pouo
}

// SetNillableCancelEvent sets the "cancel_event" field if the given value is not nil.
func (pouo *PayOrderUpdateOne) SetNillableCancelEvent(s *string) *PayOrderUpdateOne {
	if s != nil {
		pouo.SetCancelEvent(*s)
	}
	return pouo
}

// ClearCancelEvent clears the value of the "cancel_event" field.
func (pouo *PayOrderUpdateOne) ClearCancelEvent() *PayOrderUpdateOne {
	pouo.mutation.ClearCancelEvent()
	return pouo
}

// Mutation returns the PayOrderMutation object of the builder.
func (pouo *PayOrderUpdateOne) Mutation() *PayOrderMutation {
	return pouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pouo *PayOrderUpdateOne) Select(field string, fields ...string) *PayOrderUpdateOne {
	pouo.fields = append([]string{field}, fields...)
	return pouo
}

// Save executes the query and returns the updated PayOrder entity.
func (pouo *PayOrderUpdateOne) Save(ctx context.Context) (*PayOrder, error) {
	var (
		err  error
		node *PayOrder
	)
	pouo.defaults()
	if len(pouo.hooks) == 0 {
		if err = pouo.check(); err != nil {
			return nil, err
		}
		node, err = pouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pouo.check(); err != nil {
				return nil, err
			}
			pouo.mutation = mutation
			node, err = pouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pouo.hooks) - 1; i >= 0; i-- {
			if pouo.hooks[i] == nil {
				return nil, fmt.Errorf("entschema: uninitialized hook (forgotten import entschema/runtime?)")
			}
			mut = pouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pouo *PayOrderUpdateOne) SaveX(ctx context.Context) *PayOrder {
	node, err := pouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pouo *PayOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := pouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pouo *PayOrderUpdateOne) ExecX(ctx context.Context) {
	if err := pouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pouo *PayOrderUpdateOne) defaults() {
	if _, ok := pouo.mutation.UpdatedAt(); !ok {
		v := payorder.UpdateDefaultUpdatedAt()
		pouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pouo *PayOrderUpdateOne) check() error {
	if v, ok := pouo.mutation.Status(); ok {
		if err := payorder.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`entschema: validator failed for field "PayOrder.status": %w`, err)}
		}
	}
	if v, ok := pouo.mutation.CancelEvent(); ok {
		if err := payorder.CancelEventValidator(v); err != nil {
			return &ValidationError{Name: "cancel_event", err: fmt.Errorf(`entschema: validator failed for field "PayOrder.cancel_event": %w`, err)}
		}
	}
	return nil
}

func (pouo *PayOrderUpdateOne) sqlSave(ctx context.Context) (_node *PayOrder, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payorder.Table,
			Columns: payorder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: payorder.FieldID,
			},
		},
	}
	id, ok := pouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entschema: missing "PayOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payorder.FieldID)
		for _, f := range fields {
			if !payorder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
			}
			if f != payorder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldUpdatedAt,
		})
	}
	if value, ok := pouo.mutation.DeactivatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldDeactivatedAt,
		})
	}
	if pouo.mutation.DeactivatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: payorder.FieldDeactivatedAt,
		})
	}
	if value, ok := pouo.mutation.OrderNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOrderNumber,
		})
	}
	if value, ok := pouo.mutation.OrderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOrderType,
		})
	}
	if value, ok := pouo.mutation.UserUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldUserUID,
		})
	}
	if value, ok := pouo.mutation.PayMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldPayMethod,
		})
	}
	if value, ok := pouo.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldPaymentAmount,
		})
	}
	if value, ok := pouo.mutation.TotalPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldTotalPrice,
		})
	}
	if value, ok := pouo.mutation.RechargeRecord(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: payorder.FieldRechargeRecord,
		})
	}
	if value, ok := pouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldStatus,
		})
	}
	if value, ok := pouo.mutation.TransactionNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldTransactionNumber,
		})
	}
	if pouo.mutation.TransactionNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldTransactionNumber,
		})
	}
	if value, ok := pouo.mutation.ThirdpartyNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldThirdpartyNumber,
		})
	}
	if pouo.mutation.ThirdpartyNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldThirdpartyNumber,
		})
	}
	if value, ok := pouo.mutation.MchID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldMchID,
		})
	}
	if pouo.mutation.MchIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldMchID,
		})
	}
	if value, ok := pouo.mutation.OpenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldOpenID,
		})
	}
	if pouo.mutation.OpenIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldOpenID,
		})
	}
	if value, ok := pouo.mutation.MemberExpired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: payorder.FieldMemberExpired,
		})
	}
	if pouo.mutation.MemberExpiredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: payorder.FieldMemberExpired,
		})
	}
	if value, ok := pouo.mutation.CancelEvent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payorder.FieldCancelEvent,
		})
	}
	if pouo.mutation.CancelEventCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: payorder.FieldCancelEvent,
		})
	}
	_node = &PayOrder{config: pouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payorder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
