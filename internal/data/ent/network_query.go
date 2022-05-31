// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/menta2l/go-hwc/internal/data/ent/host"
	"github.com/menta2l/go-hwc/internal/data/ent/network"
	"github.com/menta2l/go-hwc/internal/data/ent/predicate"
)

// NetworkQuery is the builder for querying Network entities.
type NetworkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Network
	// eager-loading edges.
	withHostID *HostQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NetworkQuery builder.
func (nq *NetworkQuery) Where(ps ...predicate.Network) *NetworkQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NetworkQuery) Limit(limit int) *NetworkQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NetworkQuery) Offset(offset int) *NetworkQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NetworkQuery) Unique(unique bool) *NetworkQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NetworkQuery) Order(o ...OrderFunc) *NetworkQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryHostID chains the current query on the "host_id" edge.
func (nq *NetworkQuery) QueryHostID() *HostQuery {
	query := &HostQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(network.Table, network.FieldID, selector),
			sqlgraph.To(host.Table, host.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, network.HostIDTable, network.HostIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Network entity from the query.
// Returns a *NotFoundError when no Network was found.
func (nq *NetworkQuery) First(ctx context.Context) (*Network, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{network.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NetworkQuery) FirstX(ctx context.Context) *Network {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Network ID from the query.
// Returns a *NotFoundError when no Network ID was found.
func (nq *NetworkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{network.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NetworkQuery) FirstIDX(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Network entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Network entity is found.
// Returns a *NotFoundError when no Network entities are found.
func (nq *NetworkQuery) Only(ctx context.Context) (*Network, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{network.Label}
	default:
		return nil, &NotSingularError{network.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NetworkQuery) OnlyX(ctx context.Context) *Network {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Network ID in the query.
// Returns a *NotSingularError when more than one Network ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NetworkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = &NotSingularError{network.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NetworkQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Networks.
func (nq *NetworkQuery) All(ctx context.Context) ([]*Network, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NetworkQuery) AllX(ctx context.Context) []*Network {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Network IDs.
func (nq *NetworkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nq.Select(network.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NetworkQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NetworkQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NetworkQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NetworkQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NetworkQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NetworkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NetworkQuery) Clone() *NetworkQuery {
	if nq == nil {
		return nil
	}
	return &NetworkQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.Network{}, nq.predicates...),
		withHostID: nq.withHostID.Clone(),
		// clone intermediate query.
		sql:    nq.sql.Clone(),
		path:   nq.path,
		unique: nq.unique,
	}
}

// WithHostID tells the query-builder to eager-load the nodes that are connected to
// the "host_id" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NetworkQuery) WithHostID(opts ...func(*HostQuery)) *NetworkQuery {
	query := &HostQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withHostID = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Idx int `json:"idx,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Network.Query().
//		GroupBy(network.FieldIdx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nq *NetworkQuery) GroupBy(field string, fields ...string) *NetworkGroupBy {
	group := &NetworkGroupBy{config: nq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Idx int `json:"idx,omitempty"`
//	}
//
//	client.Network.Query().
//		Select(network.FieldIdx).
//		Scan(ctx, &v)
//
func (nq *NetworkQuery) Select(fields ...string) *NetworkSelect {
	nq.fields = append(nq.fields, fields...)
	return &NetworkSelect{NetworkQuery: nq}
}

func (nq *NetworkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !network.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NetworkQuery) sqlAll(ctx context.Context) ([]*Network, error) {
	var (
		nodes       = []*Network{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withHostID != nil,
		}
	)
	if nq.withHostID != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, network.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Network{config: nq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nq.withHostID; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Network)
		for i := range nodes {
			if nodes[i].host_network_id == nil {
				continue
			}
			fk := *nodes[i].host_network_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(host.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "host_network_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HostID = n
			}
		}
	}

	return nodes, nil
}

func (nq *NetworkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.fields
	if len(nq.fields) > 0 {
		_spec.Unique = nq.unique != nil && *nq.unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NetworkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nq *NetworkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   network.Table,
			Columns: network.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: network.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, network.FieldID)
		for i := range fields {
			if fields[i] != network.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NetworkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(network.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = network.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.unique != nil && *nq.unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NetworkGroupBy is the group-by builder for Network entities.
type NetworkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NetworkGroupBy) Aggregate(fns ...AggregateFunc) *NetworkGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NetworkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ngb *NetworkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NetworkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ngb *NetworkGroupBy) StringsX(ctx context.Context) []string {
	v, err := ngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ngb *NetworkGroupBy) StringX(ctx context.Context) string {
	v, err := ngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NetworkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ngb *NetworkGroupBy) IntsX(ctx context.Context) []int {
	v, err := ngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ngb *NetworkGroupBy) IntX(ctx context.Context) int {
	v, err := ngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NetworkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ngb *NetworkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ngb *NetworkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NetworkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ngb *NetworkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NetworkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ngb *NetworkGroupBy) BoolX(ctx context.Context) bool {
	v, err := ngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ngb *NetworkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ngb.fields {
		if !network.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NetworkGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NetworkSelect is the builder for selecting fields of Network entities.
type NetworkSelect struct {
	*NetworkQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NetworkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NetworkQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ns *NetworkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NetworkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ns *NetworkSelect) StringsX(ctx context.Context) []string {
	v, err := ns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ns *NetworkSelect) StringX(ctx context.Context) string {
	v, err := ns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NetworkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ns *NetworkSelect) IntsX(ctx context.Context) []int {
	v, err := ns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ns *NetworkSelect) IntX(ctx context.Context) int {
	v, err := ns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NetworkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ns *NetworkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ns *NetworkSelect) Float64X(ctx context.Context) float64 {
	v, err := ns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NetworkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ns *NetworkSelect) BoolsX(ctx context.Context) []bool {
	v, err := ns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ns *NetworkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{network.Label}
	default:
		err = fmt.Errorf("ent: NetworkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ns *NetworkSelect) BoolX(ctx context.Context) bool {
	v, err := ns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ns *NetworkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
