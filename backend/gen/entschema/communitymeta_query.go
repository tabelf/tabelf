// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"tabelf/backend/gen/entschema/communitymeta"
	"tabelf/backend/gen/entschema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityMetaQuery is the builder for querying CommunityMeta entities.
type CommunityMetaQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CommunityMeta
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommunityMetaQuery builder.
func (cmq *CommunityMetaQuery) Where(ps ...predicate.CommunityMeta) *CommunityMetaQuery {
	cmq.predicates = append(cmq.predicates, ps...)
	return cmq
}

// Limit adds a limit step to the query.
func (cmq *CommunityMetaQuery) Limit(limit int) *CommunityMetaQuery {
	cmq.limit = &limit
	return cmq
}

// Offset adds an offset step to the query.
func (cmq *CommunityMetaQuery) Offset(offset int) *CommunityMetaQuery {
	cmq.offset = &offset
	return cmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cmq *CommunityMetaQuery) Unique(unique bool) *CommunityMetaQuery {
	cmq.unique = &unique
	return cmq
}

// Order adds an order step to the query.
func (cmq *CommunityMetaQuery) Order(o ...OrderFunc) *CommunityMetaQuery {
	cmq.order = append(cmq.order, o...)
	return cmq
}

// First returns the first CommunityMeta entity from the query.
// Returns a *NotFoundError when no CommunityMeta was found.
func (cmq *CommunityMetaQuery) First(ctx context.Context) (*CommunityMeta, error) {
	nodes, err := cmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{communitymeta.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cmq *CommunityMetaQuery) FirstX(ctx context.Context) *CommunityMeta {
	node, err := cmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommunityMeta ID from the query.
// Returns a *NotFoundError when no CommunityMeta ID was found.
func (cmq *CommunityMetaQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{communitymeta.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cmq *CommunityMetaQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := cmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommunityMeta entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CommunityMeta entity is not found.
// Returns a *NotFoundError when no CommunityMeta entities are found.
func (cmq *CommunityMetaQuery) Only(ctx context.Context) (*CommunityMeta, error) {
	nodes, err := cmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{communitymeta.Label}
	default:
		return nil, &NotSingularError{communitymeta.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cmq *CommunityMetaQuery) OnlyX(ctx context.Context) *CommunityMeta {
	node, err := cmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommunityMeta ID in the query.
// Returns a *NotSingularError when exactly one CommunityMeta ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cmq *CommunityMetaQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = cmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = &NotSingularError{communitymeta.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cmq *CommunityMetaQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := cmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommunityMetaSlice.
func (cmq *CommunityMetaQuery) All(ctx context.Context) ([]*CommunityMeta, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cmq *CommunityMetaQuery) AllX(ctx context.Context) []*CommunityMeta {
	nodes, err := cmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommunityMeta IDs.
func (cmq *CommunityMetaQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := cmq.Select(communitymeta.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cmq *CommunityMetaQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := cmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cmq *CommunityMetaQuery) Count(ctx context.Context) (int, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cmq *CommunityMetaQuery) CountX(ctx context.Context) int {
	count, err := cmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cmq *CommunityMetaQuery) Exist(ctx context.Context) (bool, error) {
	if err := cmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cmq *CommunityMetaQuery) ExistX(ctx context.Context) bool {
	exist, err := cmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommunityMetaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cmq *CommunityMetaQuery) Clone() *CommunityMetaQuery {
	if cmq == nil {
		return nil
	}
	return &CommunityMetaQuery{
		config:     cmq.config,
		limit:      cmq.limit,
		offset:     cmq.offset,
		order:      append([]OrderFunc{}, cmq.order...),
		predicates: append([]predicate.CommunityMeta{}, cmq.predicates...),
		// clone intermediate query.
		sql:  cmq.sql.Clone(),
		path: cmq.path,
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
//	client.CommunityMeta.Query().
//		GroupBy(communitymeta.FieldUID).
//		Aggregate(entschema.Count()).
//		Scan(ctx, &v)
//
func (cmq *CommunityMetaQuery) GroupBy(field string, fields ...string) *CommunityMetaGroupBy {
	group := &CommunityMetaGroupBy{config: cmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cmq.sqlQuery(ctx), nil
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
//	client.CommunityMeta.Query().
//		Select(communitymeta.FieldUID).
//		Scan(ctx, &v)
//
func (cmq *CommunityMetaQuery) Select(fields ...string) *CommunityMetaSelect {
	cmq.fields = append(cmq.fields, fields...)
	return &CommunityMetaSelect{CommunityMetaQuery: cmq}
}

func (cmq *CommunityMetaQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cmq.fields {
		if !communitymeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entschema: invalid field %q for query", f)}
		}
	}
	if cmq.path != nil {
		prev, err := cmq.path(ctx)
		if err != nil {
			return err
		}
		cmq.sql = prev
	}
	return nil
}

func (cmq *CommunityMetaQuery) sqlAll(ctx context.Context) ([]*CommunityMeta, error) {
	var (
		nodes = []*CommunityMeta{}
		_spec = cmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CommunityMeta{config: cmq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cmq *CommunityMetaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cmq.querySpec()
	_spec.Node.Columns = cmq.fields
	if len(cmq.fields) > 0 {
		_spec.Unique = cmq.unique != nil && *cmq.unique
	}
	return sqlgraph.CountNodes(ctx, cmq.driver, _spec)
}

func (cmq *CommunityMetaQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entschema: check existence: %w", err)
	}
	return n > 0, nil
}

func (cmq *CommunityMetaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   communitymeta.Table,
			Columns: communitymeta.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: communitymeta.FieldID,
			},
		},
		From:   cmq.sql,
		Unique: true,
	}
	if unique := cmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, communitymeta.FieldID)
		for i := range fields {
			if fields[i] != communitymeta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cmq *CommunityMetaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cmq.driver.Dialect())
	t1 := builder.Table(communitymeta.Table)
	columns := cmq.fields
	if len(columns) == 0 {
		columns = communitymeta.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cmq.sql != nil {
		selector = cmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cmq.unique != nil && *cmq.unique {
		selector.Distinct()
	}
	for _, p := range cmq.predicates {
		p(selector)
	}
	for _, p := range cmq.order {
		p(selector)
	}
	if offset := cmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommunityMetaGroupBy is the group-by builder for CommunityMeta entities.
type CommunityMetaGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cmgb *CommunityMetaGroupBy) Aggregate(fns ...AggregateFunc) *CommunityMetaGroupBy {
	cmgb.fns = append(cmgb.fns, fns...)
	return cmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cmgb *CommunityMetaGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cmgb.path(ctx)
	if err != nil {
		return err
	}
	cmgb.sql = query
	return cmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) StringsX(ctx context.Context) []string {
	v, err := cmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) StringX(ctx context.Context) string {
	v, err := cmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) IntsX(ctx context.Context) []int {
	v, err := cmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) IntX(ctx context.Context) int {
	v, err := cmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cmgb.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cmgb *CommunityMetaGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cmgb *CommunityMetaGroupBy) BoolX(ctx context.Context) bool {
	v, err := cmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cmgb *CommunityMetaGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cmgb.fields {
		if !communitymeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cmgb *CommunityMetaGroupBy) sqlQuery() *sql.Selector {
	selector := cmgb.sql.Select()
	aggregation := make([]string, 0, len(cmgb.fns))
	for _, fn := range cmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cmgb.fields)+len(cmgb.fns))
		for _, f := range cmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cmgb.fields...)...)
}

// CommunityMetaSelect is the builder for selecting fields of CommunityMeta entities.
type CommunityMetaSelect struct {
	*CommunityMetaQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cms *CommunityMetaSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cms.prepareQuery(ctx); err != nil {
		return err
	}
	cms.sql = cms.CommunityMetaQuery.sqlQuery(ctx)
	return cms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cms *CommunityMetaSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cms *CommunityMetaSelect) StringsX(ctx context.Context) []string {
	v, err := cms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cms *CommunityMetaSelect) StringX(ctx context.Context) string {
	v, err := cms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cms *CommunityMetaSelect) IntsX(ctx context.Context) []int {
	v, err := cms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cms *CommunityMetaSelect) IntX(ctx context.Context) int {
	v, err := cms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cms *CommunityMetaSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cms *CommunityMetaSelect) Float64X(ctx context.Context) float64 {
	v, err := cms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cms.fields) > 1 {
		return nil, errors.New("entschema: CommunityMetaSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cms *CommunityMetaSelect) BoolsX(ctx context.Context) []bool {
	v, err := cms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cms *CommunityMetaSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communitymeta.Label}
	default:
		err = fmt.Errorf("entschema: CommunityMetaSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cms *CommunityMetaSelect) BoolX(ctx context.Context) bool {
	v, err := cms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cms *CommunityMetaSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cms.sql.Query()
	if err := cms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}