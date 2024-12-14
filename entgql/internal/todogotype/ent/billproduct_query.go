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
	"fmt"
	"math"

	"github.com/3xcept/contrib/entgql/internal/todogotype/ent/billproduct"
	"github.com/3xcept/contrib/entgql/internal/todogotype/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillProductQuery is the builder for querying BillProduct entities.
type BillProductQuery struct {
	config
	ctx        *QueryContext
	order      []billproduct.OrderOption
	inters     []Interceptor
	predicates []predicate.BillProduct
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*BillProduct) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillProductQuery builder.
func (bpq *BillProductQuery) Where(ps ...predicate.BillProduct) *BillProductQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit the number of records to be returned by this query.
func (bpq *BillProductQuery) Limit(limit int) *BillProductQuery {
	bpq.ctx.Limit = &limit
	return bpq
}

// Offset to start from.
func (bpq *BillProductQuery) Offset(offset int) *BillProductQuery {
	bpq.ctx.Offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BillProductQuery) Unique(unique bool) *BillProductQuery {
	bpq.ctx.Unique = &unique
	return bpq
}

// Order specifies how the records should be ordered.
func (bpq *BillProductQuery) Order(o ...billproduct.OrderOption) *BillProductQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// First returns the first BillProduct entity from the query.
// Returns a *NotFoundError when no BillProduct was found.
func (bpq *BillProductQuery) First(ctx context.Context) (*BillProduct, error) {
	nodes, err := bpq.Limit(1).All(setContextOp(ctx, bpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BillProductQuery) FirstX(ctx context.Context) *BillProduct {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillProduct ID from the query.
// Returns a *NotFoundError when no BillProduct ID was found.
func (bpq *BillProductQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bpq.Limit(1).IDs(setContextOp(ctx, bpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BillProductQuery) FirstIDX(ctx context.Context) string {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillProduct entity is found.
// Returns a *NotFoundError when no BillProduct entities are found.
func (bpq *BillProductQuery) Only(ctx context.Context) (*BillProduct, error) {
	nodes, err := bpq.Limit(2).All(setContextOp(ctx, bpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billproduct.Label}
	default:
		return nil, &NotSingularError{billproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BillProductQuery) OnlyX(ctx context.Context) *BillProduct {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillProduct ID in the query.
// Returns a *NotSingularError when more than one BillProduct ID is found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BillProductQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bpq.Limit(2).IDs(setContextOp(ctx, bpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billproduct.Label}
	default:
		err = &NotSingularError{billproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BillProductQuery) OnlyIDX(ctx context.Context) string {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillProducts.
func (bpq *BillProductQuery) All(ctx context.Context) ([]*BillProduct, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryAll)
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillProduct, *BillProductQuery]()
	return withInterceptors[[]*BillProduct](ctx, bpq, qr, bpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BillProductQuery) AllX(ctx context.Context) []*BillProduct {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillProduct IDs.
func (bpq *BillProductQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bpq.ctx.Unique == nil && bpq.path != nil {
		bpq.Unique(true)
	}
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryIDs)
	if err = bpq.Select(billproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BillProductQuery) IDsX(ctx context.Context) []string {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BillProductQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryCount)
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bpq, querierCount[*BillProductQuery](), bpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BillProductQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BillProductQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryExist)
	switch _, err := bpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BillProductQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BillProductQuery) Clone() *BillProductQuery {
	if bpq == nil {
		return nil
	}
	return &BillProductQuery{
		config:     bpq.config,
		ctx:        bpq.ctx.Clone(),
		order:      append([]billproduct.OrderOption{}, bpq.order...),
		inters:     append([]Interceptor{}, bpq.inters...),
		predicates: append([]predicate.BillProduct{}, bpq.predicates...),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillProduct.Query().
//		GroupBy(billproduct.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bpq *BillProductQuery) GroupBy(field string, fields ...string) *BillProductGroupBy {
	bpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillProductGroupBy{build: bpq}
	grbuild.flds = &bpq.ctx.Fields
	grbuild.label = billproduct.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.BillProduct.Query().
//		Select(billproduct.FieldName).
//		Scan(ctx, &v)
func (bpq *BillProductQuery) Select(fields ...string) *BillProductSelect {
	bpq.ctx.Fields = append(bpq.ctx.Fields, fields...)
	sbuild := &BillProductSelect{BillProductQuery: bpq}
	sbuild.label = billproduct.Label
	sbuild.flds, sbuild.scan = &bpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillProductSelect configured with the given aggregations.
func (bpq *BillProductQuery) Aggregate(fns ...AggregateFunc) *BillProductSelect {
	return bpq.Select().Aggregate(fns...)
}

func (bpq *BillProductQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bpq); err != nil {
				return err
			}
		}
	}
	for _, f := range bpq.ctx.Fields {
		if !billproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BillProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillProduct, error) {
	var (
		nodes = []*BillProduct{}
		_spec = bpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillProduct).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillProduct{config: bpq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bpq.modifiers) > 0 {
		_spec.Modifiers = bpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range bpq.loadTotal {
		if err := bpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bpq *BillProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	if len(bpq.modifiers) > 0 {
		_spec.Modifiers = bpq.modifiers
	}
	_spec.Node.Columns = bpq.ctx.Fields
	if len(bpq.ctx.Fields) > 0 {
		_spec.Unique = bpq.ctx.Unique != nil && *bpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BillProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billproduct.Table, billproduct.Columns, sqlgraph.NewFieldSpec(billproduct.FieldID, field.TypeString))
	_spec.From = bpq.sql
	if unique := bpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bpq.path != nil {
		_spec.Unique = true
	}
	if fields := bpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billproduct.FieldID)
		for i := range fields {
			if fields[i] != billproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BillProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(billproduct.Table)
	columns := bpq.ctx.Fields
	if len(columns) == 0 {
		columns = billproduct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.ctx.Unique != nil && *bpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BillProductGroupBy is the group-by builder for BillProduct entities.
type BillProductGroupBy struct {
	selector
	build *BillProductQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BillProductGroupBy) Aggregate(fns ...AggregateFunc) *BillProductGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the selector query and scans the result into the given value.
func (bpgb *BillProductGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bpgb.build.ctx, ent.OpQueryGroupBy)
	if err := bpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillProductQuery, *BillProductGroupBy](ctx, bpgb.build, bpgb, bpgb.build.inters, v)
}

func (bpgb *BillProductGroupBy) sqlScan(ctx context.Context, root *BillProductQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bpgb.flds)+len(bpgb.fns))
		for _, f := range *bpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillProductSelect is the builder for selecting fields of BillProduct entities.
type BillProductSelect struct {
	*BillProductQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bps *BillProductSelect) Aggregate(fns ...AggregateFunc) *BillProductSelect {
	bps.fns = append(bps.fns, fns...)
	return bps
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BillProductSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bps.ctx, ent.OpQuerySelect)
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillProductQuery, *BillProductSelect](ctx, bps.BillProductQuery, bps, bps.inters, v)
}

func (bps *BillProductSelect) sqlScan(ctx context.Context, root *BillProductQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bps.fns))
	for _, fn := range bps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
