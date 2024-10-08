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
	"github.com/davidroman0O/go-tempolite/ent/node"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/sagatask"
)

// SagaTaskQuery is the builder for querying SagaTask entities.
type SagaTaskQuery struct {
	config
	ctx        *QueryContext
	order      []sagatask.OrderOption
	inters     []Interceptor
	predicates []predicate.SagaTask
	withNode   *NodeQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SagaTaskQuery builder.
func (stq *SagaTaskQuery) Where(ps ...predicate.SagaTask) *SagaTaskQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit the number of records to be returned by this query.
func (stq *SagaTaskQuery) Limit(limit int) *SagaTaskQuery {
	stq.ctx.Limit = &limit
	return stq
}

// Offset to start from.
func (stq *SagaTaskQuery) Offset(offset int) *SagaTaskQuery {
	stq.ctx.Offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SagaTaskQuery) Unique(unique bool) *SagaTaskQuery {
	stq.ctx.Unique = &unique
	return stq
}

// Order specifies how the records should be ordered.
func (stq *SagaTaskQuery) Order(o ...sagatask.OrderOption) *SagaTaskQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// QueryNode chains the current query on the "node" edge.
func (stq *SagaTaskQuery) QueryNode() *NodeQuery {
	query := (&NodeClient{config: stq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sagatask.Table, sagatask.FieldID, selector),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, sagatask.NodeTable, sagatask.NodeColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SagaTask entity from the query.
// Returns a *NotFoundError when no SagaTask was found.
func (stq *SagaTaskQuery) First(ctx context.Context) (*SagaTask, error) {
	nodes, err := stq.Limit(1).All(setContextOp(ctx, stq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sagatask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SagaTaskQuery) FirstX(ctx context.Context) *SagaTask {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SagaTask ID from the query.
// Returns a *NotFoundError when no SagaTask ID was found.
func (stq *SagaTaskQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(setContextOp(ctx, stq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sagatask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SagaTaskQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SagaTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SagaTask entity is found.
// Returns a *NotFoundError when no SagaTask entities are found.
func (stq *SagaTaskQuery) Only(ctx context.Context) (*SagaTask, error) {
	nodes, err := stq.Limit(2).All(setContextOp(ctx, stq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sagatask.Label}
	default:
		return nil, &NotSingularError{sagatask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SagaTaskQuery) OnlyX(ctx context.Context) *SagaTask {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SagaTask ID in the query.
// Returns a *NotSingularError when more than one SagaTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SagaTaskQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(setContextOp(ctx, stq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sagatask.Label}
	default:
		err = &NotSingularError{sagatask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SagaTaskQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SagaTasks.
func (stq *SagaTaskQuery) All(ctx context.Context) ([]*SagaTask, error) {
	ctx = setContextOp(ctx, stq.ctx, ent.OpQueryAll)
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SagaTask, *SagaTaskQuery]()
	return withInterceptors[[]*SagaTask](ctx, stq, qr, stq.inters)
}

// AllX is like All, but panics if an error occurs.
func (stq *SagaTaskQuery) AllX(ctx context.Context) []*SagaTask {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SagaTask IDs.
func (stq *SagaTaskQuery) IDs(ctx context.Context) (ids []int, err error) {
	if stq.ctx.Unique == nil && stq.path != nil {
		stq.Unique(true)
	}
	ctx = setContextOp(ctx, stq.ctx, ent.OpQueryIDs)
	if err = stq.Select(sagatask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SagaTaskQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SagaTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, stq.ctx, ent.OpQueryCount)
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, stq, querierCount[*SagaTaskQuery](), stq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SagaTaskQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SagaTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, stq.ctx, ent.OpQueryExist)
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SagaTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SagaTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SagaTaskQuery) Clone() *SagaTaskQuery {
	if stq == nil {
		return nil
	}
	return &SagaTaskQuery{
		config:     stq.config,
		ctx:        stq.ctx.Clone(),
		order:      append([]sagatask.OrderOption{}, stq.order...),
		inters:     append([]Interceptor{}, stq.inters...),
		predicates: append([]predicate.SagaTask{}, stq.predicates...),
		withNode:   stq.withNode.Clone(),
		// clone intermediate query.
		sql:  stq.sql.Clone(),
		path: stq.path,
	}
}

// WithNode tells the query-builder to eager-load the nodes that are connected to
// the "node" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *SagaTaskQuery) WithNode(opts ...func(*NodeQuery)) *SagaTaskQuery {
	query := (&NodeClient{config: stq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	stq.withNode = query
	return stq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (stq *SagaTaskQuery) GroupBy(field string, fields ...string) *SagaTaskGroupBy {
	stq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SagaTaskGroupBy{build: stq}
	grbuild.flds = &stq.ctx.Fields
	grbuild.label = sagatask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (stq *SagaTaskQuery) Select(fields ...string) *SagaTaskSelect {
	stq.ctx.Fields = append(stq.ctx.Fields, fields...)
	sbuild := &SagaTaskSelect{SagaTaskQuery: stq}
	sbuild.label = sagatask.Label
	sbuild.flds, sbuild.scan = &stq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SagaTaskSelect configured with the given aggregations.
func (stq *SagaTaskQuery) Aggregate(fns ...AggregateFunc) *SagaTaskSelect {
	return stq.Select().Aggregate(fns...)
}

func (stq *SagaTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range stq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, stq); err != nil {
				return err
			}
		}
	}
	for _, f := range stq.ctx.Fields {
		if !sagatask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *SagaTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SagaTask, error) {
	var (
		nodes       = []*SagaTask{}
		withFKs     = stq.withFKs
		_spec       = stq.querySpec()
		loadedTypes = [1]bool{
			stq.withNode != nil,
		}
	)
	if stq.withNode != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sagatask.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SagaTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SagaTask{config: stq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := stq.withNode; query != nil {
		if err := stq.loadNode(ctx, query, nodes, nil,
			func(n *SagaTask, e *Node) { n.Edges.Node = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (stq *SagaTaskQuery) loadNode(ctx context.Context, query *NodeQuery, nodes []*SagaTask, init func(*SagaTask), assign func(*SagaTask, *Node)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*SagaTask)
	for i := range nodes {
		if nodes[i].node_saga_step_task == nil {
			continue
		}
		fk := *nodes[i].node_saga_step_task
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(node.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "node_saga_step_task" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (stq *SagaTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.ctx.Fields
	if len(stq.ctx.Fields) > 0 {
		_spec.Unique = stq.ctx.Unique != nil && *stq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SagaTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sagatask.Table, sagatask.Columns, sqlgraph.NewFieldSpec(sagatask.FieldID, field.TypeInt))
	_spec.From = stq.sql
	if unique := stq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if stq.path != nil {
		_spec.Unique = true
	}
	if fields := stq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sagatask.FieldID)
		for i := range fields {
			if fields[i] != sagatask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SagaTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(sagatask.Table)
	columns := stq.ctx.Fields
	if len(columns) == 0 {
		columns = sagatask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.ctx.Unique != nil && *stq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SagaTaskGroupBy is the group-by builder for SagaTask entities.
type SagaTaskGroupBy struct {
	selector
	build *SagaTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SagaTaskGroupBy) Aggregate(fns ...AggregateFunc) *SagaTaskGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the selector query and scans the result into the given value.
func (stgb *SagaTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, stgb.build.ctx, ent.OpQueryGroupBy)
	if err := stgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SagaTaskQuery, *SagaTaskGroupBy](ctx, stgb.build, stgb, stgb.build.inters, v)
}

func (stgb *SagaTaskGroupBy) sqlScan(ctx context.Context, root *SagaTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*stgb.flds)+len(stgb.fns))
		for _, f := range *stgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*stgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SagaTaskSelect is the builder for selecting fields of SagaTask entities.
type SagaTaskSelect struct {
	*SagaTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sts *SagaTaskSelect) Aggregate(fns ...AggregateFunc) *SagaTaskSelect {
	sts.fns = append(sts.fns, fns...)
	return sts
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SagaTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sts.ctx, ent.OpQuerySelect)
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SagaTaskQuery, *SagaTaskSelect](ctx, sts.SagaTaskQuery, sts, sts.inters, v)
}

func (sts *SagaTaskSelect) sqlScan(ctx context.Context, root *SagaTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sts.fns))
	for _, fn := range sts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}