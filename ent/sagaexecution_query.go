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
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/saga"
	"github.com/davidroman0O/go-tempolite/ent/sagaexecution"
)

// SagaExecutionQuery is the builder for querying SagaExecution entities.
type SagaExecutionQuery struct {
	config
	ctx        *QueryContext
	order      []sagaexecution.OrderOption
	inters     []Interceptor
	predicates []predicate.SagaExecution
	withSaga   *SagaQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SagaExecutionQuery builder.
func (seq *SagaExecutionQuery) Where(ps ...predicate.SagaExecution) *SagaExecutionQuery {
	seq.predicates = append(seq.predicates, ps...)
	return seq
}

// Limit the number of records to be returned by this query.
func (seq *SagaExecutionQuery) Limit(limit int) *SagaExecutionQuery {
	seq.ctx.Limit = &limit
	return seq
}

// Offset to start from.
func (seq *SagaExecutionQuery) Offset(offset int) *SagaExecutionQuery {
	seq.ctx.Offset = &offset
	return seq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (seq *SagaExecutionQuery) Unique(unique bool) *SagaExecutionQuery {
	seq.ctx.Unique = &unique
	return seq
}

// Order specifies how the records should be ordered.
func (seq *SagaExecutionQuery) Order(o ...sagaexecution.OrderOption) *SagaExecutionQuery {
	seq.order = append(seq.order, o...)
	return seq
}

// QuerySaga chains the current query on the "saga" edge.
func (seq *SagaExecutionQuery) QuerySaga() *SagaQuery {
	query := (&SagaClient{config: seq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := seq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := seq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sagaexecution.Table, sagaexecution.FieldID, selector),
			sqlgraph.To(saga.Table, saga.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sagaexecution.SagaTable, sagaexecution.SagaColumn),
		)
		fromU = sqlgraph.SetNeighbors(seq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SagaExecution entity from the query.
// Returns a *NotFoundError when no SagaExecution was found.
func (seq *SagaExecutionQuery) First(ctx context.Context) (*SagaExecution, error) {
	nodes, err := seq.Limit(1).All(setContextOp(ctx, seq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sagaexecution.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (seq *SagaExecutionQuery) FirstX(ctx context.Context) *SagaExecution {
	node, err := seq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SagaExecution ID from the query.
// Returns a *NotFoundError when no SagaExecution ID was found.
func (seq *SagaExecutionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = seq.Limit(1).IDs(setContextOp(ctx, seq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sagaexecution.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (seq *SagaExecutionQuery) FirstIDX(ctx context.Context) string {
	id, err := seq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SagaExecution entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SagaExecution entity is found.
// Returns a *NotFoundError when no SagaExecution entities are found.
func (seq *SagaExecutionQuery) Only(ctx context.Context) (*SagaExecution, error) {
	nodes, err := seq.Limit(2).All(setContextOp(ctx, seq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sagaexecution.Label}
	default:
		return nil, &NotSingularError{sagaexecution.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (seq *SagaExecutionQuery) OnlyX(ctx context.Context) *SagaExecution {
	node, err := seq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SagaExecution ID in the query.
// Returns a *NotSingularError when more than one SagaExecution ID is found.
// Returns a *NotFoundError when no entities are found.
func (seq *SagaExecutionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = seq.Limit(2).IDs(setContextOp(ctx, seq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sagaexecution.Label}
	default:
		err = &NotSingularError{sagaexecution.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (seq *SagaExecutionQuery) OnlyIDX(ctx context.Context) string {
	id, err := seq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SagaExecutions.
func (seq *SagaExecutionQuery) All(ctx context.Context) ([]*SagaExecution, error) {
	ctx = setContextOp(ctx, seq.ctx, ent.OpQueryAll)
	if err := seq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SagaExecution, *SagaExecutionQuery]()
	return withInterceptors[[]*SagaExecution](ctx, seq, qr, seq.inters)
}

// AllX is like All, but panics if an error occurs.
func (seq *SagaExecutionQuery) AllX(ctx context.Context) []*SagaExecution {
	nodes, err := seq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SagaExecution IDs.
func (seq *SagaExecutionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if seq.ctx.Unique == nil && seq.path != nil {
		seq.Unique(true)
	}
	ctx = setContextOp(ctx, seq.ctx, ent.OpQueryIDs)
	if err = seq.Select(sagaexecution.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (seq *SagaExecutionQuery) IDsX(ctx context.Context) []string {
	ids, err := seq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (seq *SagaExecutionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, seq.ctx, ent.OpQueryCount)
	if err := seq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, seq, querierCount[*SagaExecutionQuery](), seq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (seq *SagaExecutionQuery) CountX(ctx context.Context) int {
	count, err := seq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (seq *SagaExecutionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, seq.ctx, ent.OpQueryExist)
	switch _, err := seq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (seq *SagaExecutionQuery) ExistX(ctx context.Context) bool {
	exist, err := seq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SagaExecutionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (seq *SagaExecutionQuery) Clone() *SagaExecutionQuery {
	if seq == nil {
		return nil
	}
	return &SagaExecutionQuery{
		config:     seq.config,
		ctx:        seq.ctx.Clone(),
		order:      append([]sagaexecution.OrderOption{}, seq.order...),
		inters:     append([]Interceptor{}, seq.inters...),
		predicates: append([]predicate.SagaExecution{}, seq.predicates...),
		withSaga:   seq.withSaga.Clone(),
		// clone intermediate query.
		sql:  seq.sql.Clone(),
		path: seq.path,
	}
}

// WithSaga tells the query-builder to eager-load the nodes that are connected to
// the "saga" edge. The optional arguments are used to configure the query builder of the edge.
func (seq *SagaExecutionQuery) WithSaga(opts ...func(*SagaQuery)) *SagaExecutionQuery {
	query := (&SagaClient{config: seq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	seq.withSaga = query
	return seq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HandlerName string `json:"handler_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SagaExecution.Query().
//		GroupBy(sagaexecution.FieldHandlerName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (seq *SagaExecutionQuery) GroupBy(field string, fields ...string) *SagaExecutionGroupBy {
	seq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SagaExecutionGroupBy{build: seq}
	grbuild.flds = &seq.ctx.Fields
	grbuild.label = sagaexecution.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HandlerName string `json:"handler_name,omitempty"`
//	}
//
//	client.SagaExecution.Query().
//		Select(sagaexecution.FieldHandlerName).
//		Scan(ctx, &v)
func (seq *SagaExecutionQuery) Select(fields ...string) *SagaExecutionSelect {
	seq.ctx.Fields = append(seq.ctx.Fields, fields...)
	sbuild := &SagaExecutionSelect{SagaExecutionQuery: seq}
	sbuild.label = sagaexecution.Label
	sbuild.flds, sbuild.scan = &seq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SagaExecutionSelect configured with the given aggregations.
func (seq *SagaExecutionQuery) Aggregate(fns ...AggregateFunc) *SagaExecutionSelect {
	return seq.Select().Aggregate(fns...)
}

func (seq *SagaExecutionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range seq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, seq); err != nil {
				return err
			}
		}
	}
	for _, f := range seq.ctx.Fields {
		if !sagaexecution.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if seq.path != nil {
		prev, err := seq.path(ctx)
		if err != nil {
			return err
		}
		seq.sql = prev
	}
	return nil
}

func (seq *SagaExecutionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SagaExecution, error) {
	var (
		nodes       = []*SagaExecution{}
		withFKs     = seq.withFKs
		_spec       = seq.querySpec()
		loadedTypes = [1]bool{
			seq.withSaga != nil,
		}
	)
	if seq.withSaga != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sagaexecution.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SagaExecution).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SagaExecution{config: seq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, seq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := seq.withSaga; query != nil {
		if err := seq.loadSaga(ctx, query, nodes, nil,
			func(n *SagaExecution, e *Saga) { n.Edges.Saga = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (seq *SagaExecutionQuery) loadSaga(ctx context.Context, query *SagaQuery, nodes []*SagaExecution, init func(*SagaExecution), assign func(*SagaExecution, *Saga)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*SagaExecution)
	for i := range nodes {
		if nodes[i].saga_steps == nil {
			continue
		}
		fk := *nodes[i].saga_steps
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(saga.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "saga_steps" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (seq *SagaExecutionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := seq.querySpec()
	_spec.Node.Columns = seq.ctx.Fields
	if len(seq.ctx.Fields) > 0 {
		_spec.Unique = seq.ctx.Unique != nil && *seq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, seq.driver, _spec)
}

func (seq *SagaExecutionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sagaexecution.Table, sagaexecution.Columns, sqlgraph.NewFieldSpec(sagaexecution.FieldID, field.TypeString))
	_spec.From = seq.sql
	if unique := seq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if seq.path != nil {
		_spec.Unique = true
	}
	if fields := seq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sagaexecution.FieldID)
		for i := range fields {
			if fields[i] != sagaexecution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := seq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := seq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := seq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := seq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (seq *SagaExecutionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(seq.driver.Dialect())
	t1 := builder.Table(sagaexecution.Table)
	columns := seq.ctx.Fields
	if len(columns) == 0 {
		columns = sagaexecution.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if seq.sql != nil {
		selector = seq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if seq.ctx.Unique != nil && *seq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range seq.predicates {
		p(selector)
	}
	for _, p := range seq.order {
		p(selector)
	}
	if offset := seq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := seq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SagaExecutionGroupBy is the group-by builder for SagaExecution entities.
type SagaExecutionGroupBy struct {
	selector
	build *SagaExecutionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (segb *SagaExecutionGroupBy) Aggregate(fns ...AggregateFunc) *SagaExecutionGroupBy {
	segb.fns = append(segb.fns, fns...)
	return segb
}

// Scan applies the selector query and scans the result into the given value.
func (segb *SagaExecutionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, segb.build.ctx, ent.OpQueryGroupBy)
	if err := segb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SagaExecutionQuery, *SagaExecutionGroupBy](ctx, segb.build, segb, segb.build.inters, v)
}

func (segb *SagaExecutionGroupBy) sqlScan(ctx context.Context, root *SagaExecutionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(segb.fns))
	for _, fn := range segb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*segb.flds)+len(segb.fns))
		for _, f := range *segb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*segb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := segb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SagaExecutionSelect is the builder for selecting fields of SagaExecution entities.
type SagaExecutionSelect struct {
	*SagaExecutionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ses *SagaExecutionSelect) Aggregate(fns ...AggregateFunc) *SagaExecutionSelect {
	ses.fns = append(ses.fns, fns...)
	return ses
}

// Scan applies the selector query and scans the result into the given value.
func (ses *SagaExecutionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ses.ctx, ent.OpQuerySelect)
	if err := ses.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SagaExecutionQuery, *SagaExecutionSelect](ctx, ses.SagaExecutionQuery, ses, ses.inters, v)
}

func (ses *SagaExecutionSelect) sqlScan(ctx context.Context, root *SagaExecutionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ses.fns))
	for _, fn := range ses.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ses.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ses.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
