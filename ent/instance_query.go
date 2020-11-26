// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/build"
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/instance"
	"github.com/lingfohn/lime/ent/predicate"
)

// InstanceQuery is the builder for querying Instance entities.
type InstanceQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Instance
	// eager-loading edges.
	withApplication *ApplicationQuery
	withBuilds      *BuildQuery
	withConfig      *HelmConfigQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (iq *InstanceQuery) Where(ps ...predicate.Instance) *InstanceQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InstanceQuery) Limit(limit int) *InstanceQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InstanceQuery) Offset(offset int) *InstanceQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *InstanceQuery) Order(o ...Order) *InstanceQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryApplication chains the current query on the application edge.
func (iq *InstanceQuery) QueryApplication() *ApplicationQuery {
	query := &ApplicationQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, iq.sqlQuery()),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instance.ApplicationTable, instance.ApplicationColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBuilds chains the current query on the builds edge.
func (iq *InstanceQuery) QueryBuilds() *BuildQuery {
	query := &BuildQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, iq.sqlQuery()),
			sqlgraph.To(build.Table, build.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instance.BuildsTable, instance.BuildsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConfig chains the current query on the config edge.
func (iq *InstanceQuery) QueryConfig() *HelmConfigQuery {
	query := &HelmConfigQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, iq.sqlQuery()),
			sqlgraph.To(helmconfig.Table, helmconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, instance.ConfigTable, instance.ConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Instance entity in the query. Returns *NotFoundError when no instance was found.
func (iq *InstanceQuery) First(ctx context.Context) (*Instance, error) {
	is, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(is) == 0 {
		return nil, &NotFoundError{instance.Label}
	}
	return is[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstanceQuery) FirstX(ctx context.Context) *Instance {
	i, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return i
}

// FirstID returns the first Instance id in the query. Returns *NotFoundError when no id was found.
func (iq *InstanceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instance.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (iq *InstanceQuery) FirstXID(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Instance entity in the query, returns an error if not exactly one entity was returned.
func (iq *InstanceQuery) Only(ctx context.Context) (*Instance, error) {
	is, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(is) {
	case 1:
		return is[0], nil
	case 0:
		return nil, &NotFoundError{instance.Label}
	default:
		return nil, &NotSingularError{instance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstanceQuery) OnlyX(ctx context.Context) *Instance {
	i, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// OnlyID returns the only Instance id in the query, returns an error if not exactly one id was returned.
func (iq *InstanceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = &NotSingularError{instance.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (iq *InstanceQuery) OnlyXID(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Instances.
func (iq *InstanceQuery) All(ctx context.Context) ([]*Instance, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstanceQuery) AllX(ctx context.Context) []*Instance {
	is, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return is
}

// IDs executes the query and returns a list of Instance ids.
func (iq *InstanceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(instance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstanceQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstanceQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstanceQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstanceQuery) Clone() *InstanceQuery {
	return &InstanceQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]Order{}, iq.order...),
		unique:     append([]string{}, iq.unique...),
		predicates: append([]predicate.Instance{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

//  WithApplication tells the query-builder to eager-loads the nodes that are connected to
// the "application" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InstanceQuery) WithApplication(opts ...func(*ApplicationQuery)) *InstanceQuery {
	query := &ApplicationQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withApplication = query
	return iq
}

//  WithBuilds tells the query-builder to eager-loads the nodes that are connected to
// the "builds" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InstanceQuery) WithBuilds(opts ...func(*BuildQuery)) *InstanceQuery {
	query := &BuildQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withBuilds = query
	return iq
}

//  WithConfig tells the query-builder to eager-loads the nodes that are connected to
// the "config" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InstanceQuery) WithConfig(opts ...func(*HelmConfigQuery)) *InstanceQuery {
	query := &HelmConfigQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withConfig = query
	return iq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Instance.Query().
//		GroupBy(instance.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InstanceQuery) GroupBy(field string, fields ...string) *InstanceGroupBy {
	group := &InstanceGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Instance.Query().
//		Select(instance.FieldName).
//		Scan(ctx, &v)
//
func (iq *InstanceQuery) Select(field string, fields ...string) *InstanceSelect {
	selector := &InstanceSelect{config: iq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(), nil
	}
	return selector
}

func (iq *InstanceQuery) prepareQuery(ctx context.Context) error {
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstanceQuery) sqlAll(ctx context.Context) ([]*Instance, error) {
	var (
		nodes       = []*Instance{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [3]bool{
			iq.withApplication != nil,
			iq.withBuilds != nil,
			iq.withConfig != nil,
		}
	)
	if iq.withApplication != nil || iq.withConfig != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, instance.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Instance{config: iq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withApplication; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Instance)
		for i := range nodes {
			if fk := nodes[i].application_instances; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(application.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "application_instances" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Application = n
			}
		}
	}

	if query := iq.withBuilds; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Instance)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Build(func(s *sql.Selector) {
			s.Where(sql.InValues(instance.BuildsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.instance_builds
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "instance_builds" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instance_builds" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Builds = append(node.Edges.Builds, n)
		}
	}

	if query := iq.withConfig; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Instance)
		for i := range nodes {
			if fk := nodes[i].instance_config; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(helmconfig.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "instance_config" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Config = n
			}
		}
	}

	return nodes, nil
}

func (iq *InstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iq *InstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instance.Table,
			Columns: instance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instance.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstanceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(instance.Table)
	selector := builder.Select(t1.Columns(instance.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(instance.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstanceGroupBy is the builder for group-by Instance entities.
type InstanceGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstanceGroupBy) Aggregate(fns ...Aggregate) *InstanceGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scan the result into the given value.
func (igb *InstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *InstanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (igb *InstanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *InstanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (igb *InstanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *InstanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (igb *InstanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *InstanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (igb *InstanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *InstanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *InstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := igb.sqlQuery().Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InstanceGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

// InstanceSelect is the builder for select fields of Instance entities.
type InstanceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (is *InstanceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := is.path(ctx)
	if err != nil {
		return err
	}
	is.sql = query
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *InstanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *InstanceSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *InstanceSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *InstanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *InstanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *InstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *InstanceSelect) sqlQuery() sql.Querier {
	selector := is.sql
	selector.Select(selector.Columns(is.fields...)...)
	return selector
}
