// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbipuser"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/predicate"
)

// CmdbIpUserQuery is the builder for querying CmdbIpUser entities.
type CmdbIpUserQuery struct {
	config
	ctx        *QueryContext
	order      []cmdbipuser.OrderOption
	inters     []Interceptor
	predicates []predicate.CmdbIpUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CmdbIpUserQuery builder.
func (ciuq *CmdbIpUserQuery) Where(ps ...predicate.CmdbIpUser) *CmdbIpUserQuery {
	ciuq.predicates = append(ciuq.predicates, ps...)
	return ciuq
}

// Limit the number of records to be returned by this query.
func (ciuq *CmdbIpUserQuery) Limit(limit int) *CmdbIpUserQuery {
	ciuq.ctx.Limit = &limit
	return ciuq
}

// Offset to start from.
func (ciuq *CmdbIpUserQuery) Offset(offset int) *CmdbIpUserQuery {
	ciuq.ctx.Offset = &offset
	return ciuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciuq *CmdbIpUserQuery) Unique(unique bool) *CmdbIpUserQuery {
	ciuq.ctx.Unique = &unique
	return ciuq
}

// Order specifies how the records should be ordered.
func (ciuq *CmdbIpUserQuery) Order(o ...cmdbipuser.OrderOption) *CmdbIpUserQuery {
	ciuq.order = append(ciuq.order, o...)
	return ciuq
}

// First returns the first CmdbIpUser entity from the query.
// Returns a *NotFoundError when no CmdbIpUser was found.
func (ciuq *CmdbIpUserQuery) First(ctx context.Context) (*CmdbIpUser, error) {
	nodes, err := ciuq.Limit(1).All(setContextOp(ctx, ciuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cmdbipuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) FirstX(ctx context.Context) *CmdbIpUser {
	node, err := ciuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CmdbIpUser ID from the query.
// Returns a *NotFoundError when no CmdbIpUser ID was found.
func (ciuq *CmdbIpUserQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ciuq.Limit(1).IDs(setContextOp(ctx, ciuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cmdbipuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ciuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CmdbIpUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CmdbIpUser entity is found.
// Returns a *NotFoundError when no CmdbIpUser entities are found.
func (ciuq *CmdbIpUserQuery) Only(ctx context.Context) (*CmdbIpUser, error) {
	nodes, err := ciuq.Limit(2).All(setContextOp(ctx, ciuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cmdbipuser.Label}
	default:
		return nil, &NotSingularError{cmdbipuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) OnlyX(ctx context.Context) *CmdbIpUser {
	node, err := ciuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CmdbIpUser ID in the query.
// Returns a *NotSingularError when more than one CmdbIpUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciuq *CmdbIpUserQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ciuq.Limit(2).IDs(setContextOp(ctx, ciuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cmdbipuser.Label}
	default:
		err = &NotSingularError{cmdbipuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ciuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CmdbIpUsers.
func (ciuq *CmdbIpUserQuery) All(ctx context.Context) ([]*CmdbIpUser, error) {
	ctx = setContextOp(ctx, ciuq.ctx, "All")
	if err := ciuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CmdbIpUser, *CmdbIpUserQuery]()
	return withInterceptors[[]*CmdbIpUser](ctx, ciuq, qr, ciuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) AllX(ctx context.Context) []*CmdbIpUser {
	nodes, err := ciuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CmdbIpUser IDs.
func (ciuq *CmdbIpUserQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if ciuq.ctx.Unique == nil && ciuq.path != nil {
		ciuq.Unique(true)
	}
	ctx = setContextOp(ctx, ciuq.ctx, "IDs")
	if err = ciuq.Select(cmdbipuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ciuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciuq *CmdbIpUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ciuq.ctx, "Count")
	if err := ciuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ciuq, querierCount[*CmdbIpUserQuery](), ciuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) CountX(ctx context.Context) int {
	count, err := ciuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciuq *CmdbIpUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ciuq.ctx, "Exist")
	switch _, err := ciuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ciuq *CmdbIpUserQuery) ExistX(ctx context.Context) bool {
	exist, err := ciuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CmdbIpUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciuq *CmdbIpUserQuery) Clone() *CmdbIpUserQuery {
	if ciuq == nil {
		return nil
	}
	return &CmdbIpUserQuery{
		config:     ciuq.config,
		ctx:        ciuq.ctx.Clone(),
		order:      append([]cmdbipuser.OrderOption{}, ciuq.order...),
		inters:     append([]Interceptor{}, ciuq.inters...),
		predicates: append([]predicate.CmdbIpUser{}, ciuq.predicates...),
		// clone intermediate query.
		sql:  ciuq.sql.Clone(),
		path: ciuq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status uint8 `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CmdbIpUser.Query().
//		GroupBy(cmdbipuser.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ciuq *CmdbIpUserQuery) GroupBy(field string, fields ...string) *CmdbIpUserGroupBy {
	ciuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CmdbIpUserGroupBy{build: ciuq}
	grbuild.flds = &ciuq.ctx.Fields
	grbuild.label = cmdbipuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status uint8 `json:"status,omitempty"`
//	}
//
//	client.CmdbIpUser.Query().
//		Select(cmdbipuser.FieldStatus).
//		Scan(ctx, &v)
func (ciuq *CmdbIpUserQuery) Select(fields ...string) *CmdbIpUserSelect {
	ciuq.ctx.Fields = append(ciuq.ctx.Fields, fields...)
	sbuild := &CmdbIpUserSelect{CmdbIpUserQuery: ciuq}
	sbuild.label = cmdbipuser.Label
	sbuild.flds, sbuild.scan = &ciuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CmdbIpUserSelect configured with the given aggregations.
func (ciuq *CmdbIpUserQuery) Aggregate(fns ...AggregateFunc) *CmdbIpUserSelect {
	return ciuq.Select().Aggregate(fns...)
}

func (ciuq *CmdbIpUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ciuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ciuq); err != nil {
				return err
			}
		}
	}
	for _, f := range ciuq.ctx.Fields {
		if !cmdbipuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciuq.path != nil {
		prev, err := ciuq.path(ctx)
		if err != nil {
			return err
		}
		ciuq.sql = prev
	}
	return nil
}

func (ciuq *CmdbIpUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CmdbIpUser, error) {
	var (
		nodes = []*CmdbIpUser{}
		_spec = ciuq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CmdbIpUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CmdbIpUser{config: ciuq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ciuq *CmdbIpUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciuq.querySpec()
	_spec.Node.Columns = ciuq.ctx.Fields
	if len(ciuq.ctx.Fields) > 0 {
		_spec.Unique = ciuq.ctx.Unique != nil && *ciuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ciuq.driver, _spec)
}

func (ciuq *CmdbIpUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cmdbipuser.Table, cmdbipuser.Columns, sqlgraph.NewFieldSpec(cmdbipuser.FieldID, field.TypeUint64))
	_spec.From = ciuq.sql
	if unique := ciuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ciuq.path != nil {
		_spec.Unique = true
	}
	if fields := ciuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cmdbipuser.FieldID)
		for i := range fields {
			if fields[i] != cmdbipuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciuq *CmdbIpUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciuq.driver.Dialect())
	t1 := builder.Table(cmdbipuser.Table)
	columns := ciuq.ctx.Fields
	if len(columns) == 0 {
		columns = cmdbipuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciuq.sql != nil {
		selector = ciuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciuq.ctx.Unique != nil && *ciuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ciuq.predicates {
		p(selector)
	}
	for _, p := range ciuq.order {
		p(selector)
	}
	if offset := ciuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CmdbIpUserGroupBy is the group-by builder for CmdbIpUser entities.
type CmdbIpUserGroupBy struct {
	selector
	build *CmdbIpUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ciugb *CmdbIpUserGroupBy) Aggregate(fns ...AggregateFunc) *CmdbIpUserGroupBy {
	ciugb.fns = append(ciugb.fns, fns...)
	return ciugb
}

// Scan applies the selector query and scans the result into the given value.
func (ciugb *CmdbIpUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ciugb.build.ctx, "GroupBy")
	if err := ciugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CmdbIpUserQuery, *CmdbIpUserGroupBy](ctx, ciugb.build, ciugb, ciugb.build.inters, v)
}

func (ciugb *CmdbIpUserGroupBy) sqlScan(ctx context.Context, root *CmdbIpUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ciugb.fns))
	for _, fn := range ciugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ciugb.flds)+len(ciugb.fns))
		for _, f := range *ciugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ciugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ciugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CmdbIpUserSelect is the builder for selecting fields of CmdbIpUser entities.
type CmdbIpUserSelect struct {
	*CmdbIpUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cius *CmdbIpUserSelect) Aggregate(fns ...AggregateFunc) *CmdbIpUserSelect {
	cius.fns = append(cius.fns, fns...)
	return cius
}

// Scan applies the selector query and scans the result into the given value.
func (cius *CmdbIpUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cius.ctx, "Select")
	if err := cius.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CmdbIpUserQuery, *CmdbIpUserSelect](ctx, cius.CmdbIpUserQuery, cius, cius.inters, v)
}

func (cius *CmdbIpUserSelect) sqlScan(ctx context.Context, root *CmdbIpUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cius.fns))
	for _, fn := range cius.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cius.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cius.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
