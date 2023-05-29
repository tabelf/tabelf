// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tabelf/backend/gen/entschema/goodarticlecategory"
	"tabelf/backend/gen/entschema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GoodArticleCategoryQuery is the builder for querying GoodArticleCategory entities.
type GoodArticleCategoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodArticleCategory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodArticleCategoryQuery builder.
func (gacq *GoodArticleCategoryQuery) Where(ps ...predicate.GoodArticleCategory) *GoodArticleCategoryQuery {
	gacq.predicates = append(gacq.predicates, ps...)
	return gacq
}

// Limit adds a limit step to the query.
func (gacq *GoodArticleCategoryQuery) Limit(limit int) *GoodArticleCategoryQuery {
	gacq.limit = &limit
	return gacq
}

// Offset adds an offset step to the query.
func (gacq *GoodArticleCategoryQuery) Offset(offset int) *GoodArticleCategoryQuery {
	gacq.offset = &offset
	return gacq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gacq *GoodArticleCategoryQuery) Unique(unique bool) *GoodArticleCategoryQuery {
	gacq.unique = &unique
	return gacq
}

// Order adds an order step to the query.
func (gacq *GoodArticleCategoryQuery) Order(o ...OrderFunc) *GoodArticleCategoryQuery {
	gacq.order = append(gacq.order, o...)
	return gacq
}

// First returns the first GoodArticleCategory entity from the query.
// Returns a *NotFoundError when no GoodArticleCategory was found.
func (gacq *GoodArticleCategoryQuery) First(ctx context.Context) (*GoodArticleCategory, error) {
	nodes, err := gacq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodarticlecategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) FirstX(ctx context.Context) *GoodArticleCategory {
	node, err := gacq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodArticleCategory ID from the query.
// Returns a *NotFoundError when no GoodArticleCategory ID was found.
func (gacq *GoodArticleCategoryQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gacq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodarticlecategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := gacq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodArticleCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodArticleCategory entity is not found.
// Returns a *NotFoundError when no GoodArticleCategory entities are found.
func (gacq *GoodArticleCategoryQuery) Only(ctx context.Context) (*GoodArticleCategory, error) {
	nodes, err := gacq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodarticlecategory.Label}
	default:
		return nil, &NotSingularError{goodarticlecategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) OnlyX(ctx context.Context) *GoodArticleCategory {
	node, err := gacq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodArticleCategory ID in the query.
// Returns a *NotSingularError when exactly one GoodArticleCategory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gacq *GoodArticleCategoryQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gacq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = &NotSingularError{goodarticlecategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := gacq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodArticleCategories.
func (gacq *GoodArticleCategoryQuery) All(ctx context.Context) ([]*GoodArticleCategory, error) {
	if err := gacq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gacq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) AllX(ctx context.Context) []*GoodArticleCategory {
	nodes, err := gacq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodArticleCategory IDs.
func (gacq *GoodArticleCategoryQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := gacq.Select(goodarticlecategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := gacq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gacq *GoodArticleCategoryQuery) Count(ctx context.Context) (int, error) {
	if err := gacq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gacq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) CountX(ctx context.Context) int {
	count, err := gacq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gacq *GoodArticleCategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := gacq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gacq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gacq *GoodArticleCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := gacq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodArticleCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gacq *GoodArticleCategoryQuery) Clone() *GoodArticleCategoryQuery {
	if gacq == nil {
		return nil
	}
	return &GoodArticleCategoryQuery{
		config:     gacq.config,
		limit:      gacq.limit,
		offset:     gacq.offset,
		order:      append([]OrderFunc{}, gacq.order...),
		predicates: append([]predicate.GoodArticleCategory{}, gacq.predicates...),
		// clone intermediate query.
		sql:  gacq.sql.Clone(),
		path: gacq.path,
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
//	client.GoodArticleCategory.Query().
//		GroupBy(goodarticlecategory.FieldUID).
//		Aggregate(entschema.Count()).
//		Scan(ctx, &v)
//
func (gacq *GoodArticleCategoryQuery) GroupBy(field string, fields ...string) *GoodArticleCategoryGroupBy {
	group := &GoodArticleCategoryGroupBy{config: gacq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gacq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gacq.sqlQuery(ctx), nil
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
//	client.GoodArticleCategory.Query().
//		Select(goodarticlecategory.FieldUID).
//		Scan(ctx, &v)
//
func (gacq *GoodArticleCategoryQuery) Select(fields ...string) *GoodArticleCategorySelect {
	gacq.fields = append(gacq.fields, fields...)
	return &GoodArticleCategorySelect{GoodArticleCategoryQuery: gacq}
}

func (gacq *GoodArticleCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gacq.fields {
		if !goodarticlecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
		}
	}
	if gacq.path != nil {
		prev, err := gacq.path(ctx)
		if err != nil {
			return err
		}
		gacq.sql = prev
	}
	return nil
}

func (gacq *GoodArticleCategoryQuery) sqlAll(ctx context.Context) ([]*GoodArticleCategory, error) {
	var (
		nodes = []*GoodArticleCategory{}
		_spec = gacq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodArticleCategory{config: gacq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gacq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gacq *GoodArticleCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gacq.querySpec()
	_spec.Node.Columns = gacq.fields
	if len(gacq.fields) > 0 {
		_spec.Unique = gacq.unique != nil && *gacq.unique
	}
	return sqlgraph.CountNodes(ctx, gacq.driver, _spec)
}

func (gacq *GoodArticleCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gacq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entschema: check existence: %w", err)
	}
	return n > 0, nil
}

func (gacq *GoodArticleCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodarticlecategory.Table,
			Columns: goodarticlecategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: goodarticlecategory.FieldID,
			},
		},
		From:   gacq.sql,
		Unique: true,
	}
	if unique := gacq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gacq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodarticlecategory.FieldID)
		for i := range fields {
			if fields[i] != goodarticlecategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gacq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gacq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gacq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gacq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gacq *GoodArticleCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gacq.driver.Dialect())
	t1 := builder.Table(goodarticlecategory.Table)
	columns := gacq.fields
	if len(columns) == 0 {
		columns = goodarticlecategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gacq.sql != nil {
		selector = gacq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gacq.unique != nil && *gacq.unique {
		selector.Distinct()
	}
	for _, p := range gacq.predicates {
		p(selector)
	}
	for _, p := range gacq.order {
		p(selector)
	}
	if offset := gacq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gacq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodArticleCategoryGroupBy is the group-by builder for GoodArticleCategory entities.
type GoodArticleCategoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gacgb *GoodArticleCategoryGroupBy) Aggregate(fns ...AggregateFunc) *GoodArticleCategoryGroupBy {
	gacgb.fns = append(gacgb.fns, fns...)
	return gacgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gacgb *GoodArticleCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gacgb.path(ctx)
	if err != nil {
		return err
	}
	gacgb.sql = query
	return gacgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gacgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gacgb.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gacgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := gacgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gacgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) StringX(ctx context.Context) string {
	v, err := gacgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gacgb.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gacgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := gacgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gacgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) IntX(ctx context.Context) int {
	v, err := gacgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gacgb.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gacgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gacgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gacgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gacgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gacgb.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gacgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gacgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gacgb *GoodArticleCategoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gacgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gacgb *GoodArticleCategoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := gacgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gacgb *GoodArticleCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gacgb.fields {
		if !goodarticlecategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gacgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gacgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gacgb *GoodArticleCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := gacgb.sql.Select()
	aggregation := make([]string, 0, len(gacgb.fns))
	for _, fn := range gacgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gacgb.fields)+len(gacgb.fns))
		for _, f := range gacgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gacgb.fields...)...)
}

// GoodArticleCategorySelect is the builder for selecting fields of GoodArticleCategory entities.
type GoodArticleCategorySelect struct {
	*GoodArticleCategoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gacs *GoodArticleCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := gacs.prepareQuery(ctx); err != nil {
		return err
	}
	gacs.sql = gacs.GoodArticleCategoryQuery.sqlQuery(ctx)
	return gacs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := gacs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(gacs.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gacs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) StringsX(ctx context.Context) []string {
	v, err := gacs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gacs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) StringX(ctx context.Context) string {
	v, err := gacs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(gacs.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gacs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) IntsX(ctx context.Context) []int {
	v, err := gacs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gacs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) IntX(ctx context.Context) int {
	v, err := gacs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gacs.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gacs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := gacs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gacs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) Float64X(ctx context.Context) float64 {
	v, err := gacs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gacs.fields) > 1 {
		return nil, errors.New("entschema: GoodArticleCategorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gacs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) BoolsX(ctx context.Context) []bool {
	v, err := gacs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gacs *GoodArticleCategorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gacs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodarticlecategory.Label}
	default:
		err = fmt.Errorf("entschema: GoodArticleCategorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gacs *GoodArticleCategorySelect) BoolX(ctx context.Context) bool {
	v, err := gacs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gacs *GoodArticleCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gacs.sql.Query()
	if err := gacs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
