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
	"github.com/davidroman0O/go-tempolite/ent/executionunit"
	"github.com/davidroman0O/go-tempolite/ent/predicate"
	"github.com/davidroman0O/go-tempolite/ent/task"
)

// TaskQuery is the builder for querying Task entities.
type TaskQuery struct {
	config
	ctx               *QueryContext
	order             []task.OrderOption
	inters            []Interceptor
	predicates        []predicate.Task
	withExecutionUnit *ExecutionUnitQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TaskQuery builder.
func (tq *TaskQuery) Where(ps ...predicate.Task) *TaskQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TaskQuery) Limit(limit int) *TaskQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TaskQuery) Offset(offset int) *TaskQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TaskQuery) Unique(unique bool) *TaskQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TaskQuery) Order(o ...task.OrderOption) *TaskQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryExecutionUnit chains the current query on the "execution_unit" edge.
func (tq *TaskQuery) QueryExecutionUnit() *ExecutionUnitQuery {
	query := (&ExecutionUnitClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, selector),
			sqlgraph.To(executionunit.Table, executionunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, task.ExecutionUnitTable, task.ExecutionUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Task entity from the query.
// Returns a *NotFoundError when no Task was found.
func (tq *TaskQuery) First(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{task.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TaskQuery) FirstX(ctx context.Context) *Task {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Task ID from the query.
// Returns a *NotFoundError when no Task ID was found.
func (tq *TaskQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{task.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TaskQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Task entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Task entity is found.
// Returns a *NotFoundError when no Task entities are found.
func (tq *TaskQuery) Only(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{task.Label}
	default:
		return nil, &NotSingularError{task.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TaskQuery) OnlyX(ctx context.Context) *Task {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Task ID in the query.
// Returns a *NotSingularError when more than one Task ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TaskQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{task.Label}
	default:
		err = &NotSingularError{task.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TaskQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tasks.
func (tq *TaskQuery) All(ctx context.Context) ([]*Task, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Task, *TaskQuery]()
	return withInterceptors[[]*Task](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TaskQuery) AllX(ctx context.Context) []*Task {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Task IDs.
func (tq *TaskQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(task.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TaskQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TaskQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TaskQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TaskQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TaskQuery) Clone() *TaskQuery {
	if tq == nil {
		return nil
	}
	return &TaskQuery{
		config:            tq.config,
		ctx:               tq.ctx.Clone(),
		order:             append([]task.OrderOption{}, tq.order...),
		inters:            append([]Interceptor{}, tq.inters...),
		predicates:        append([]predicate.Task{}, tq.predicates...),
		withExecutionUnit: tq.withExecutionUnit.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithExecutionUnit tells the query-builder to eager-load the nodes that are connected to
// the "execution_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TaskQuery) WithExecutionUnit(opts ...func(*ExecutionUnitQuery)) *TaskQuery {
	query := (&ExecutionUnitClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withExecutionUnit = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type task.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Task.Query().
//		GroupBy(task.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TaskQuery) GroupBy(field string, fields ...string) *TaskGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaskGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = task.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type task.Type `json:"type,omitempty"`
//	}
//
//	client.Task.Query().
//		Select(task.FieldType).
//		Scan(ctx, &v)
func (tq *TaskQuery) Select(fields ...string) *TaskSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TaskSelect{TaskQuery: tq}
	sbuild.label = task.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaskSelect configured with the given aggregations.
func (tq *TaskQuery) Aggregate(fns ...AggregateFunc) *TaskSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !task.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Task, error) {
	var (
		nodes       = []*Task{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withExecutionUnit != nil,
		}
	)
	if tq.withExecutionUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, task.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Task).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Task{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withExecutionUnit; query != nil {
		if err := tq.loadExecutionUnit(ctx, query, nodes, nil,
			func(n *Task, e *ExecutionUnit) { n.Edges.ExecutionUnit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TaskQuery) loadExecutionUnit(ctx context.Context, query *ExecutionUnitQuery, nodes []*Task, init func(*Task), assign func(*Task, *ExecutionUnit)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Task)
	for i := range nodes {
		if nodes[i].execution_unit_tasks == nil {
			continue
		}
		fk := *nodes[i].execution_unit_tasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(executionunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "execution_unit_tasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeString))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for i := range fields {
			if fields[i] != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(task.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = task.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TaskGroupBy is the group-by builder for Task entities.
type TaskGroupBy struct {
	selector
	build *TaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TaskGroupBy) Aggregate(fns ...AggregateFunc) *TaskGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TaskGroupBy) sqlScan(ctx context.Context, root *TaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TaskSelect is the builder for selecting fields of Task entities.
type TaskSelect struct {
	*TaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TaskSelect) Aggregate(fns ...AggregateFunc) *TaskSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskSelect](ctx, ts.TaskQuery, ts, ts.inters, v)
}

func (ts *TaskSelect) sqlScan(ctx context.Context, root *TaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
