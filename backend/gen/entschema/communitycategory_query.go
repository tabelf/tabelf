// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tabelf/backend/gen/entschema/communitycategory"
	"tabelf/backend/gen/entschema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityCategoryQuery is the builder for querying CommunityCategory entities.
type CommunityCategoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CommunityCategory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommunityCategoryQuery builder.
func (ccq *CommunityCategoryQuery) Where(ps ...predicate.CommunityCategory) *CommunityCategoryQuery {
	ccq.predicates = append(ccq.predicates, ps...)
	return ccq
}

// Limit adds a limit step to the query.
func (ccq *CommunityCategoryQuery) Limit(limit int) *CommunityCategoryQuery {
	ccq.limit = &limit
	return ccq
}

// Offset adds an offset step to the query.
func (ccq *CommunityCategoryQuery) Offset(offset int) *CommunityCategoryQuery {
	ccq.offset = &offset
	return ccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ccq *CommunityCategoryQuery) Unique(unique bool) *CommunityCategoryQuery {
	ccq.unique = &unique
	return ccq
}

// Order adds an order step to the query.
func (ccq *CommunityCategoryQuery) Order(o ...OrderFunc) *CommunityCategoryQuery {
	ccq.order = append(ccq.order, o...)
	return ccq
}

// First returns the first CommunityCategory entity from the query.
// Returns a *NotFoundError when no CommunityCategory was found.
func (ccq *CommunityCategoryQuery) First(ctx context.Context) (*CommunityCategory, error) {
	nodes, err := ccq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{communitycategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) FirstX(ctx context.Context) *CommunityCategory {
	node, err := ccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommunityCategory ID from the query.
// Returns a *NotFoundError when no CommunityCategory ID was found.
func (ccq *CommunityCategoryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ccq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{communitycategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommunityCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CommunityCategory entity is not found.
// Returns a *NotFoundError when no CommunityCategory entities are found.
func (ccq *CommunityCategoryQuery) Only(ctx context.Context) (*CommunityCategory, error) {
	nodes, err := ccq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{communitycategory.Label}
	default:
		return nil, &NotSingularError{communitycategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) OnlyX(ctx context.Context) *CommunityCategory {
	node, err := ccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommunityCategory ID in the query.
// Returns a *NotSingularError when exactly one CommunityCategory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ccq *CommunityCategoryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ccq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = &NotSingularError{communitycategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommunityCategories.
func (ccq *CommunityCategoryQuery) All(ctx context.Context) ([]*CommunityCategory, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ccq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) AllX(ctx context.Context) []*CommunityCategory {
	nodes, err := ccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommunityCategory IDs.
func (ccq *CommunityCategoryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := ccq.Select(communitycategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ccq *CommunityCategoryQuery) Count(ctx context.Context) (int, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ccq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) CountX(ctx context.Context) int {
	count, err := ccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ccq *CommunityCategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := ccq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ccq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ccq *CommunityCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommunityCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ccq *CommunityCategoryQuery) Clone() *CommunityCategoryQuery {
	if ccq == nil {
		return nil
	}
	return &CommunityCategoryQuery{
		config:     ccq.config,
		limit:      ccq.limit,
		offset:     ccq.offset,
		order:      append([]OrderFunc{}, ccq.order...),
		predicates: append([]predicate.CommunityCategory{}, ccq.predicates...),
		// clone intermediate query.
		sql:  ccq.sql.Clone(),
		path: ccq.path,
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
//	client.CommunityCategory.Query().
//		GroupBy(communitycategory.FieldUID).
//		Aggregate(entschema.Count()).
//		Scan(ctx, &v)
//
func (ccq *CommunityCategoryQuery) GroupBy(field string, fields ...string) *CommunityCategoryGroupBy {
	group := &CommunityCategoryGroupBy{config: ccq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ccq.sqlQuery(ctx), nil
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
//	client.CommunityCategory.Query().
//		Select(communitycategory.FieldUID).
//		Scan(ctx, &v)
//
func (ccq *CommunityCategoryQuery) Select(fields ...string) *CommunityCategorySelect {
	ccq.fields = append(ccq.fields, fields...)
	return &CommunityCategorySelect{CommunityCategoryQuery: ccq}
}

func (ccq *CommunityCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ccq.fields {
		if !communitycategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
		}
	}
	if ccq.path != nil {
		prev, err := ccq.path(ctx)
		if err != nil {
			return err
		}
		ccq.sql = prev
	}
	return nil
}

func (ccq *CommunityCategoryQuery) sqlAll(ctx context.Context) ([]*CommunityCategory, error) {
	var (
		nodes = []*CommunityCategory{}
		_spec = ccq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CommunityCategory{config: ccq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ccq *CommunityCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ccq.querySpec()
	_spec.Node.Columns = ccq.fields
	if len(ccq.fields) > 0 {
		_spec.Unique = ccq.unique != nil && *ccq.unique
	}
	return sqlgraph.CountNodes(ctx, ccq.driver, _spec)
}

func (ccq *CommunityCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ccq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entschema: check existence: %w", err)
	}
	return n > 0, nil
}

func (ccq *CommunityCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   communitycategory.Table,
			Columns: communitycategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: communitycategory.FieldID,
			},
		},
		From:   ccq.sql,
		Unique: true,
	}
	if unique := ccq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ccq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, communitycategory.FieldID)
		for i := range fields {
			if fields[i] != communitycategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ccq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ccq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ccq *CommunityCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ccq.driver.Dialect())
	t1 := builder.Table(communitycategory.Table)
	columns := ccq.fields
	if len(columns) == 0 {
		columns = communitycategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ccq.sql != nil {
		selector = ccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ccq.unique != nil && *ccq.unique {
		selector.Distinct()
	}
	for _, p := range ccq.predicates {
		p(selector)
	}
	for _, p := range ccq.order {
		p(selector)
	}
	if offset := ccq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ccq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommunityCategoryGroupBy is the group-by builder for CommunityCategory entities.
type CommunityCategoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ccgb *CommunityCategoryGroupBy) Aggregate(fns ...AggregateFunc) *CommunityCategoryGroupBy {
	ccgb.fns = append(ccgb.fns, fns...)
	return ccgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ccgb *CommunityCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ccgb.path(ctx)
	if err != nil {
		return err
	}
	ccgb.sql = query
	return ccgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ccgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := ccgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) StringX(ctx context.Context) string {
	v, err := ccgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := ccgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) IntX(ctx context.Context) int {
	v, err := ccgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ccgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ccgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ccgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ccgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ccgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ccgb *CommunityCategoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccgb *CommunityCategoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := ccgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccgb *CommunityCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ccgb.fields {
		if !communitycategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ccgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ccgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ccgb *CommunityCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := ccgb.sql.Select()
	aggregation := make([]string, 0, len(ccgb.fns))
	for _, fn := range ccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ccgb.fields)+len(ccgb.fns))
		for _, f := range ccgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ccgb.fields...)...)
}

// CommunityCategorySelect is the builder for selecting fields of CommunityCategory entities.
type CommunityCategorySelect struct {
	*CommunityCategoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ccs *CommunityCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ccs.prepareQuery(ctx); err != nil {
		return err
	}
	ccs.sql = ccs.CommunityCategoryQuery.sqlQuery(ctx)
	return ccs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ccs *CommunityCategorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := ccs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ccs *CommunityCategorySelect) StringsX(ctx context.Context) []string {
	v, err := ccs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ccs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ccs *CommunityCategorySelect) StringX(ctx context.Context) string {
	v, err := ccs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ccs *CommunityCategorySelect) IntsX(ctx context.Context) []int {
	v, err := ccs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ccs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ccs *CommunityCategorySelect) IntX(ctx context.Context) int {
	v, err := ccs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ccs *CommunityCategorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := ccs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ccs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ccs *CommunityCategorySelect) Float64X(ctx context.Context) float64 {
	v, err := ccs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ccs.fields) > 1 {
		return nil, errors.New("entschema: CommunityCategorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ccs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ccs *CommunityCategorySelect) BoolsX(ctx context.Context) []bool {
	v, err := ccs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ccs *CommunityCategorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ccs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitycategory.Label}
	default:
		err = fmt.Errorf("entschema: CommunityCategorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ccs *CommunityCategorySelect) BoolX(ctx context.Context) bool {
	v, err := ccs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ccs *CommunityCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ccs.sql.Query()
	if err := ccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
