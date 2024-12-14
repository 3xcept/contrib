// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/3xcept/contrib/entgql/internal/todofed/ent/category"
	"github.com/3xcept/contrib/entgql/internal/todofed/ent/predicate"
	"github.com/3xcept/contrib/entgql/internal/todofed/ent/todo"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryQuery is the builder for querying Category entities.
type CategoryQuery struct {
	config
	ctx            *QueryContext
	order          []category.OrderOption
	inters         []Interceptor
	predicates     []predicate.Category
	withTodos      *TodoQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*Category) error
	withNamedTodos map[string]*TodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CategoryQuery builder.
func (cq *CategoryQuery) Where(ps ...predicate.Category) *CategoryQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CategoryQuery) Limit(limit int) *CategoryQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CategoryQuery) Offset(offset int) *CategoryQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CategoryQuery) Unique(unique bool) *CategoryQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CategoryQuery) Order(o ...category.OrderOption) *CategoryQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryTodos chains the current query on the "todos" edge.
func (cq *CategoryQuery) QueryTodos() *TodoQuery {
	query := (&TodoClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.TodosTable, category.TodosColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Category entity from the query.
// Returns a *NotFoundError when no Category was found.
func (cq *CategoryQuery) First(ctx context.Context) (*Category, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{category.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CategoryQuery) FirstX(ctx context.Context) *Category {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Category ID from the query.
// Returns a *NotFoundError when no Category ID was found.
func (cq *CategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{category.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Category entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Category entity is found.
// Returns a *NotFoundError when no Category entities are found.
func (cq *CategoryQuery) Only(ctx context.Context) (*Category, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{category.Label}
	default:
		return nil, &NotSingularError{category.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CategoryQuery) OnlyX(ctx context.Context) *Category {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Category ID in the query.
// Returns a *NotSingularError when more than one Category ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{category.Label}
	default:
		err = &NotSingularError{category.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Categories.
func (cq *CategoryQuery) All(ctx context.Context) ([]*Category, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Category, *CategoryQuery]()
	return withInterceptors[[]*Category](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CategoryQuery) AllX(ctx context.Context) []*Category {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Category IDs.
func (cq *CategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(category.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CategoryQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CategoryQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CategoryQuery) Clone() *CategoryQuery {
	if cq == nil {
		return nil
	}
	return &CategoryQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]category.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Category{}, cq.predicates...),
		withTodos:  cq.withTodos.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithTodos tells the query-builder to eager-load the nodes that are connected to
// the "todos" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CategoryQuery) WithTodos(opts ...func(*TodoQuery)) *CategoryQuery {
	query := (&TodoClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withTodos = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Category.Query().
//		GroupBy(category.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CategoryQuery) GroupBy(field string, fields ...string) *CategoryGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CategoryGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = category.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Category.Query().
//		Select(category.FieldText).
//		Scan(ctx, &v)
func (cq *CategoryQuery) Select(fields ...string) *CategorySelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CategorySelect{CategoryQuery: cq}
	sbuild.label = category.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CategorySelect configured with the given aggregations.
func (cq *CategoryQuery) Aggregate(fns ...AggregateFunc) *CategorySelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !category.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Category, error) {
	var (
		nodes       = []*Category{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withTodos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Category).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Category{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withTodos; query != nil {
		if err := cq.loadTodos(ctx, query, nodes,
			func(n *Category) { n.Edges.Todos = []*Todo{} },
			func(n *Category, e *Todo) { n.Edges.Todos = append(n.Edges.Todos, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedTodos {
		if err := cq.loadTodos(ctx, query, nodes,
			func(n *Category) { n.appendNamedTodos(name) },
			func(n *Category, e *Todo) { n.appendNamedTodos(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CategoryQuery) loadTodos(ctx context.Context, query *TodoQuery, nodes []*Category, init func(*Category), assign func(*Category, *Todo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Category)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(category.TodosColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.category_todos
		if fk == nil {
			return fmt.Errorf(`foreign-key "category_todos" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "category_todos" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for i := range fields {
			if fields[i] != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(category.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = category.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTodos tells the query-builder to eager-load the nodes that are connected to the "todos"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CategoryQuery) WithNamedTodos(name string, opts ...func(*TodoQuery)) *CategoryQuery {
	query := (&TodoClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedTodos == nil {
		cq.withNamedTodos = make(map[string]*TodoQuery)
	}
	cq.withNamedTodos[name] = query
	return cq
}

// CategoryGroupBy is the group-by builder for Category entities.
type CategoryGroupBy struct {
	selector
	build *CategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CategoryGroupBy) Aggregate(fns ...AggregateFunc) *CategoryGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryQuery, *CategoryGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CategoryGroupBy) sqlScan(ctx context.Context, root *CategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CategorySelect is the builder for selecting fields of Category entities.
type CategorySelect struct {
	*CategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CategorySelect) Aggregate(fns ...AggregateFunc) *CategorySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryQuery, *CategorySelect](ctx, cs.CategoryQuery, cs, cs.inters, v)
}

func (cs *CategorySelect) sqlScan(ctx context.Context, root *CategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
