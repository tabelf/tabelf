// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/stationmeta"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StationMetaQuery is the builder for querying StationMeta entities.
type StationMetaQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StationMeta
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StationMetaQuery builder.
func (smq *StationMetaQuery) Where(ps ...predicate.StationMeta) *StationMetaQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit adds a limit step to the query.
func (smq *StationMetaQuery) Limit(limit int) *StationMetaQuery {
	smq.limit = &limit
	return smq
}

// Offset adds an offset step to the query.
func (smq *StationMetaQuery) Offset(offset int) *StationMetaQuery {
	smq.offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *StationMetaQuery) Unique(unique bool) *StationMetaQuery {
	smq.unique = &unique
	return smq
}

// Order adds an order step to the query.
func (smq *StationMetaQuery) Order(o ...OrderFunc) *StationMetaQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// First returns the first StationMeta entity from the query.
// Returns a *NotFoundError when no StationMeta was found.
func (smq *StationMetaQuery) First(ctx context.Context) (*StationMeta, error) {
	nodes, err := smq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{stationmeta.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *StationMetaQuery) FirstX(ctx context.Context) *StationMeta {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StationMeta ID from the query.
// Returns a *NotFoundError when no StationMeta ID was found.
func (smq *StationMetaQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = smq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{stationmeta.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *StationMetaQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StationMeta entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one StationMeta entity is not found.
// Returns a *NotFoundError when no StationMeta entities are found.
func (smq *StationMetaQuery) Only(ctx context.Context) (*StationMeta, error) {
	nodes, err := smq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{stationmeta.Label}
	default:
		return nil, &NotSingularError{stationmeta.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *StationMetaQuery) OnlyX(ctx context.Context) *StationMeta {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StationMeta ID in the query.
// Returns a *NotSingularError when exactly one StationMeta ID is not found.
// Returns a *NotFoundError when no entities are found.
func (smq *StationMetaQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = smq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = &NotSingularError{stationmeta.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *StationMetaQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StationMetaSlice.
func (smq *StationMetaQuery) All(ctx context.Context) ([]*StationMeta, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (smq *StationMetaQuery) AllX(ctx context.Context) []*StationMeta {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StationMeta IDs.
func (smq *StationMetaQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := smq.Select(stationmeta.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *StationMetaQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *StationMetaQuery) Count(ctx context.Context) (int, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (smq *StationMetaQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *StationMetaQuery) Exist(ctx context.Context) (bool, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *StationMetaQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StationMetaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *StationMetaQuery) Clone() *StationMetaQuery {
	if smq == nil {
		return nil
	}
	return &StationMetaQuery{
		config:     smq.config,
		limit:      smq.limit,
		offset:     smq.offset,
		order:      append([]OrderFunc{}, smq.order...),
		predicates: append([]predicate.StationMeta{}, smq.predicates...),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UID string `json:"uid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StationMeta.Query().
//		GroupBy(stationmeta.FieldUID).
//		Aggregate(entschema.Count()).
//		Scan(ctx, &v)
//
func (smq *StationMetaQuery) GroupBy(field string, fields ...string) *StationMetaGroupBy {
	group := &StationMetaGroupBy{config: smq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UID string `json:"uid,omitempty"`
//	}
//
//	client.StationMeta.Query().
//		Select(stationmeta.FieldUID).
//		Scan(ctx, &v)
//
func (smq *StationMetaQuery) Select(fields ...string) *StationMetaSelect {
	smq.fields = append(smq.fields, fields...)
	return &StationMetaSelect{StationMetaQuery: smq}
}

func (smq *StationMetaQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smq.fields {
		if !stationmeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *StationMetaQuery) sqlAll(ctx context.Context) ([]*StationMeta, error) {
	var (
		nodes = []*StationMeta{}
		_spec = smq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StationMeta{config: smq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("entschema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (smq *StationMetaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	_spec.Node.Columns = smq.fields
	if len(smq.fields) > 0 {
		_spec.Unique = smq.unique != nil && *smq.unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *StationMetaQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := smq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entschema: check existence: %w", err)
	}
	return n > 0, nil
}

func (smq *StationMetaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stationmeta.Table,
			Columns: stationmeta.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: stationmeta.FieldID,
			},
		},
		From:   smq.sql,
		Unique: true,
	}
	if unique := smq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := smq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stationmeta.FieldID)
		for i := range fields {
			if fields[i] != stationmeta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *StationMetaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(stationmeta.Table)
	columns := smq.fields
	if len(columns) == 0 {
		columns = stationmeta.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.unique != nil && *smq.unique {
		selector.Distinct()
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StationMetaGroupBy is the group-by builder for StationMeta entities.
type StationMetaGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *StationMetaGroupBy) Aggregate(fns ...AggregateFunc) *StationMetaGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the group-by query and scans the result into the given value.
func (smgb *StationMetaGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := smgb.path(ctx)
	if err != nil {
		return err
	}
	smgb.sql = query
	return smgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smgb *StationMetaGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := smgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("entschema: StationMetaGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smgb *StationMetaGroupBy) StringsX(ctx context.Context) []string {
	v, err := smgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smgb *StationMetaGroupBy) StringX(ctx context.Context) string {
	v, err := smgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("entschema: StationMetaGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smgb *StationMetaGroupBy) IntsX(ctx context.Context) []int {
	v, err := smgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smgb *StationMetaGroupBy) IntX(ctx context.Context) int {
	v, err := smgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("entschema: StationMetaGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smgb *StationMetaGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := smgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smgb *StationMetaGroupBy) Float64X(ctx context.Context) float64 {
	v, err := smgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("entschema: StationMetaGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smgb *StationMetaGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := smgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *StationMetaGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smgb *StationMetaGroupBy) BoolX(ctx context.Context) bool {
	v, err := smgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smgb *StationMetaGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range smgb.fields {
		if !stationmeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := smgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smgb *StationMetaGroupBy) sqlQuery() *sql.Selector {
	selector := smgb.sql.Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(smgb.fields)+len(smgb.fns))
		for _, f := range smgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(smgb.fields...)...)
}

// StationMetaSelect is the builder for selecting fields of StationMeta entities.
type StationMetaSelect struct {
	*StationMetaQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sms *StationMetaSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	sms.sql = sms.StationMetaQuery.sqlQuery(ctx)
	return sms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sms *StationMetaSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("entschema: StationMetaSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sms *StationMetaSelect) StringsX(ctx context.Context) []string {
	v, err := sms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sms *StationMetaSelect) StringX(ctx context.Context) string {
	v, err := sms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("entschema: StationMetaSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sms *StationMetaSelect) IntsX(ctx context.Context) []int {
	v, err := sms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sms *StationMetaSelect) IntX(ctx context.Context) int {
	v, err := sms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("entschema: StationMetaSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sms *StationMetaSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sms *StationMetaSelect) Float64X(ctx context.Context) float64 {
	v, err := sms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("entschema: StationMetaSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sms *StationMetaSelect) BoolsX(ctx context.Context) []bool {
	v, err := sms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sms *StationMetaSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{stationmeta.Label}
	default:
		err = fmt.Errorf("entschema: StationMetaSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sms *StationMetaSelect) BoolX(ctx context.Context) bool {
	v, err := sms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sms *StationMetaSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sms.sql.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
