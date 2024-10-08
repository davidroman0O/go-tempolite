// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/davidroman0O/go-tempolite/ent/executioncontext"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
)

// ExecutionContextQuery is the builder for querying ExecutionContext entities.
type ExecutionContextQuery struct {
	config
	ctx        *QueryContext
	order      []executioncontext.OrderOption
	inters     []Interceptor
	predicates []predicate.ExecutionContext
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExecutionContextQuery builder.
func (ecq *ExecutionContextQuery) Where(ps ...predicate.ExecutionContext) *ExecutionContextQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *ExecutionContextQuery) Limit(limit int) *ExecutionContextQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *ExecutionContextQuery) Offset(offset int) *ExecutionContextQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *ExecutionContextQuery) Unique(unique bool) *ExecutionContextQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *ExecutionContextQuery) Order(o ...executioncontext.OrderOption) *ExecutionContextQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// First returns the first ExecutionContext entity from the query.
// Returns a *NotFoundError when no ExecutionContext was found.
func (ecq *ExecutionContextQuery) First(ctx context.Context) (*ExecutionContext, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{executioncontext.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *ExecutionContextQuery) FirstX(ctx context.Context) *ExecutionContext {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExecutionContext ID from the query.
// Returns a *NotFoundError when no ExecutionContext ID was found.
func (ecq *ExecutionContextQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{executioncontext.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *ExecutionContextQuery) FirstIDX(ctx context.Context) string {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExecutionContext entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExecutionContext entity is found.
// Returns a *NotFoundError when no ExecutionContext entities are found.
func (ecq *ExecutionContextQuery) Only(ctx context.Context) (*ExecutionContext, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{executioncontext.Label}
	default:
		return nil, &NotSingularError{executioncontext.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *ExecutionContextQuery) OnlyX(ctx context.Context) *ExecutionContext {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExecutionContext ID in the query.
// Returns a *NotSingularError when more than one ExecutionContext ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *ExecutionContextQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{executioncontext.Label}
	default:
		err = &NotSingularError{executioncontext.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *ExecutionContextQuery) OnlyIDX(ctx context.Context) string {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExecutionContexts.
func (ecq *ExecutionContextQuery) All(ctx context.Context) ([]*ExecutionContext, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryAll)
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExecutionContext, *ExecutionContextQuery]()
	return withInterceptors[[]*ExecutionContext](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *ExecutionContextQuery) AllX(ctx context.Context) []*ExecutionContext {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExecutionContext IDs.
func (ecq *ExecutionContextQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryIDs)
	if err = ecq.Select(executioncontext.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *ExecutionContextQuery) IDsX(ctx context.Context) []string {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *ExecutionContextQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryCount)
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*ExecutionContextQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *ExecutionContextQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *ExecutionContextQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, ent.OpQueryExist)
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *ExecutionContextQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExecutionContextQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *ExecutionContextQuery) Clone() *ExecutionContextQuery {
	if ecq == nil {
		return nil
	}
	return &ExecutionContextQuery{
		config:     ecq.config,
		ctx:        ecq.ctx.Clone(),
		order:      append([]executioncontext.OrderOption{}, ecq.order...),
		inters:     append([]Interceptor{}, ecq.inters...),
		predicates: append([]predicate.ExecutionContext{}, ecq.predicates...),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ecq *ExecutionContextQuery) GroupBy(field string, fields ...string) *ExecutionContextGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExecutionContextGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = executioncontext.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ecq *ExecutionContextQuery) Select(fields ...string) *ExecutionContextSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &ExecutionContextSelect{ExecutionContextQuery: ecq}
	sbuild.label = executioncontext.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExecutionContextSelect configured with the given aggregations.
func (ecq *ExecutionContextQuery) Aggregate(fns ...AggregateFunc) *ExecutionContextSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *ExecutionContextQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !executioncontext.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *ExecutionContextQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExecutionContext, error) {
	var (
		nodes = []*ExecutionContext{}
		_spec = ecq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExecutionContext).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExecutionContext{config: ecq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ecq *ExecutionContextQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *ExecutionContextQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(executioncontext.Table, executioncontext.Columns, sqlgraph.NewFieldSpec(executioncontext.FieldID, field.TypeString))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, executioncontext.FieldID)
		for i := range fields {
			if fields[i] != executioncontext.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *ExecutionContextQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(executioncontext.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = executioncontext.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExecutionContextGroupBy is the group-by builder for ExecutionContext entities.
type ExecutionContextGroupBy struct {
	selector
	build *ExecutionContextQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *ExecutionContextGroupBy) Aggregate(fns ...AggregateFunc) *ExecutionContextGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *ExecutionContextGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, ent.OpQueryGroupBy)
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExecutionContextQuery, *ExecutionContextGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *ExecutionContextGroupBy) sqlScan(ctx context.Context, root *ExecutionContextQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExecutionContextSelect is the builder for selecting fields of ExecutionContext entities.
type ExecutionContextSelect struct {
	*ExecutionContextQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *ExecutionContextSelect) Aggregate(fns ...AggregateFunc) *ExecutionContextSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *ExecutionContextSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, ent.OpQuerySelect)
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExecutionContextQuery, *ExecutionContextSelect](ctx, ecs.ExecutionContextQuery, ecs, ecs.inters, v)
}

func (ecs *ExecutionContextSelect) sqlScan(ctx context.Context, root *ExecutionContextQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
