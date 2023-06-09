// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tabelf/backend/gen/entschema/predicate"
	"tabelf/backend/gen/entschema/sharelink"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShareLinkQuery is the builder for querying ShareLink entities.
type ShareLinkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShareLink
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShareLinkQuery builder.
func (slq *ShareLinkQuery) Where(ps ...predicate.ShareLink) *ShareLinkQuery {
	slq.predicates = append(slq.predicates, ps...)
	return slq
}

// Limit adds a limit step to the query.
func (slq *ShareLinkQuery) Limit(limit int) *ShareLinkQuery {
	slq.limit = &limit
	return slq
}

// Offset adds an offset step to the query.
func (slq *ShareLinkQuery) Offset(offset int) *ShareLinkQuery {
	slq.offset = &offset
	return slq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (slq *ShareLinkQuery) Unique(unique bool) *ShareLinkQuery {
	slq.unique = &unique
	return slq
}

// Order adds an order step to the query.
func (slq *ShareLinkQuery) Order(o ...OrderFunc) *ShareLinkQuery {
	slq.order = append(slq.order, o...)
	return slq
}

// First returns the first ShareLink entity from the query.
// Returns a *NotFoundError when no ShareLink was found.
func (slq *ShareLinkQuery) First(ctx context.Context) (*ShareLink, error) {
	nodes, err := slq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sharelink.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (slq *ShareLinkQuery) FirstX(ctx context.Context) *ShareLink {
	node, err := slq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShareLink ID from the query.
// Returns a *NotFoundError when no ShareLink ID was found.
func (slq *ShareLinkQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = slq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sharelink.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (slq *ShareLinkQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := slq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShareLink entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShareLink entity is not found.
// Returns a *NotFoundError when no ShareLink entities are found.
func (slq *ShareLinkQuery) Only(ctx context.Context) (*ShareLink, error) {
	nodes, err := slq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sharelink.Label}
	default:
		return nil, &NotSingularError{sharelink.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (slq *ShareLinkQuery) OnlyX(ctx context.Context) *ShareLink {
	node, err := slq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShareLink ID in the query.
// Returns a *NotSingularError when exactly one ShareLink ID is not found.
// Returns a *NotFoundError when no entities are found.
func (slq *ShareLinkQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = slq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = &NotSingularError{sharelink.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (slq *ShareLinkQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := slq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShareLinks.
func (slq *ShareLinkQuery) All(ctx context.Context) ([]*ShareLink, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return slq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (slq *ShareLinkQuery) AllX(ctx context.Context) []*ShareLink {
	nodes, err := slq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShareLink IDs.
func (slq *ShareLinkQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := slq.Select(sharelink.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (slq *ShareLinkQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := slq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (slq *ShareLinkQuery) Count(ctx context.Context) (int, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return slq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (slq *ShareLinkQuery) CountX(ctx context.Context) int {
	count, err := slq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (slq *ShareLinkQuery) Exist(ctx context.Context) (bool, error) {
	if err := slq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return slq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (slq *ShareLinkQuery) ExistX(ctx context.Context) bool {
	exist, err := slq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShareLinkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (slq *ShareLinkQuery) Clone() *ShareLinkQuery {
	if slq == nil {
		return nil
	}
	return &ShareLinkQuery{
		config:     slq.config,
		limit:      slq.limit,
		offset:     slq.offset,
		order:      append([]OrderFunc{}, slq.order...),
		predicates: append([]predicate.ShareLink{}, slq.predicates...),
		// clone intermediate query.
		sql:  slq.sql.Clone(),
		path: slq.path,
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
//	client.ShareLink.Query().
//		GroupBy(sharelink.FieldUID).
//		Aggregate(entschema.Count()).
//		Scan(ctx, &v)
//
func (slq *ShareLinkQuery) GroupBy(field string, fields ...string) *ShareLinkGroupBy {
	group := &ShareLinkGroupBy{config: slq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := slq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return slq.sqlQuery(ctx), nil
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
//	client.ShareLink.Query().
//		Select(sharelink.FieldUID).
//		Scan(ctx, &v)
//
func (slq *ShareLinkQuery) Select(fields ...string) *ShareLinkSelect {
	slq.fields = append(slq.fields, fields...)
	return &ShareLinkSelect{ShareLinkQuery: slq}
}

func (slq *ShareLinkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range slq.fields {
		if !sharelink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
		}
	}
	if slq.path != nil {
		prev, err := slq.path(ctx)
		if err != nil {
			return err
		}
		slq.sql = prev
	}
	return nil
}

func (slq *ShareLinkQuery) sqlAll(ctx context.Context) ([]*ShareLink, error) {
	var (
		nodes = []*ShareLink{}
		_spec = slq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShareLink{config: slq.config}
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
	if err := sqlgraph.QueryNodes(ctx, slq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (slq *ShareLinkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := slq.querySpec()
	_spec.Node.Columns = slq.fields
	if len(slq.fields) > 0 {
		_spec.Unique = slq.unique != nil && *slq.unique
	}
	return sqlgraph.CountNodes(ctx, slq.driver, _spec)
}

func (slq *ShareLinkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := slq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entschema: check existence: %w", err)
	}
	return n > 0, nil
}

func (slq *ShareLinkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sharelink.Table,
			Columns: sharelink.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: sharelink.FieldID,
			},
		},
		From:   slq.sql,
		Unique: true,
	}
	if unique := slq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := slq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sharelink.FieldID)
		for i := range fields {
			if fields[i] != sharelink.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := slq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := slq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := slq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := slq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (slq *ShareLinkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(slq.driver.Dialect())
	t1 := builder.Table(sharelink.Table)
	columns := slq.fields
	if len(columns) == 0 {
		columns = sharelink.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if slq.sql != nil {
		selector = slq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if slq.unique != nil && *slq.unique {
		selector.Distinct()
	}
	for _, p := range slq.predicates {
		p(selector)
	}
	for _, p := range slq.order {
		p(selector)
	}
	if offset := slq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := slq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShareLinkGroupBy is the group-by builder for ShareLink entities.
type ShareLinkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (slgb *ShareLinkGroupBy) Aggregate(fns ...AggregateFunc) *ShareLinkGroupBy {
	slgb.fns = append(slgb.fns, fns...)
	return slgb
}

// Scan applies the group-by query and scans the result into the given value.
func (slgb *ShareLinkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := slgb.path(ctx)
	if err != nil {
		return err
	}
	slgb.sql = query
	return slgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := slgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(slgb.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := slgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) StringsX(ctx context.Context) []string {
	v, err := slgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = slgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) StringX(ctx context.Context) string {
	v, err := slgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(slgb.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := slgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) IntsX(ctx context.Context) []int {
	v, err := slgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = slgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) IntX(ctx context.Context) int {
	v, err := slgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(slgb.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := slgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := slgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = slgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := slgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(slgb.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := slgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := slgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (slgb *ShareLinkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = slgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (slgb *ShareLinkGroupBy) BoolX(ctx context.Context) bool {
	v, err := slgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (slgb *ShareLinkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range slgb.fields {
		if !sharelink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := slgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := slgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (slgb *ShareLinkGroupBy) sqlQuery() *sql.Selector {
	selector := slgb.sql.Select()
	aggregation := make([]string, 0, len(slgb.fns))
	for _, fn := range slgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(slgb.fields)+len(slgb.fns))
		for _, f := range slgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(slgb.fields...)...)
}

// ShareLinkSelect is the builder for selecting fields of ShareLink entities.
type ShareLinkSelect struct {
	*ShareLinkQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sls *ShareLinkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sls.prepareQuery(ctx); err != nil {
		return err
	}
	sls.sql = sls.ShareLinkQuery.sqlQuery(ctx)
	return sls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sls *ShareLinkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sls.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sls *ShareLinkSelect) StringsX(ctx context.Context) []string {
	v, err := sls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sls *ShareLinkSelect) StringX(ctx context.Context) string {
	v, err := sls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sls.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sls *ShareLinkSelect) IntsX(ctx context.Context) []int {
	v, err := sls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sls *ShareLinkSelect) IntX(ctx context.Context) int {
	v, err := sls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sls.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sls *ShareLinkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sls *ShareLinkSelect) Float64X(ctx context.Context) float64 {
	v, err := sls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sls.fields) > 1 {
		return nil, errors.New("entschema: ShareLinkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sls *ShareLinkSelect) BoolsX(ctx context.Context) []bool {
	v, err := sls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sls *ShareLinkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sharelink.Label}
	default:
		err = fmt.Errorf("entschema: ShareLinkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sls *ShareLinkSelect) BoolX(ctx context.Context) bool {
	v, err := sls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sls *ShareLinkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sls.sql.Query()
	if err := sls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
