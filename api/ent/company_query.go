// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/company"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
	"github.com/Wanted-Linx/linx-backend/api/ent/tasktype"
	"github.com/Wanted-Linx/linx-backend/api/ent/user"
)

// CompanyQuery is the builder for querying Company entities.
type CompanyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Company
	// eager-loading edges.
	withUser     *UserQuery
	withProject  *ProjectQuery
	withTaskType *TaskTypeQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanyQuery builder.
func (cq *CompanyQuery) Where(ps ...predicate.Company) *CompanyQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *CompanyQuery) Limit(limit int) *CompanyQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *CompanyQuery) Offset(offset int) *CompanyQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CompanyQuery) Unique(unique bool) *CompanyQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *CompanyQuery) Order(o ...OrderFunc) *CompanyQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryUser chains the current query on the "user" edge.
func (cq *CompanyQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, company.UserTable, company.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProject chains the current query on the "project" edge.
func (cq *CompanyQuery) QueryProject() *ProjectQuery {
	query := &ProjectQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.ProjectTable, company.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTaskType chains the current query on the "task_type" edge.
func (cq *CompanyQuery) QueryTaskType() *TaskTypeQuery {
	query := &TaskTypeQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(company.Table, company.FieldID, selector),
			sqlgraph.To(tasktype.Table, tasktype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, company.TaskTypeTable, company.TaskTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Company entity from the query.
// Returns a *NotFoundError when no Company was found.
func (cq *CompanyQuery) First(ctx context.Context) (*Company, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{company.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CompanyQuery) FirstX(ctx context.Context) *Company {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Company ID from the query.
// Returns a *NotFoundError when no Company ID was found.
func (cq *CompanyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{company.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CompanyQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Company entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Company entity is not found.
// Returns a *NotFoundError when no Company entities are found.
func (cq *CompanyQuery) Only(ctx context.Context) (*Company, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{company.Label}
	default:
		return nil, &NotSingularError{company.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CompanyQuery) OnlyX(ctx context.Context) *Company {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Company ID in the query.
// Returns a *NotSingularError when exactly one Company ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cq *CompanyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = &NotSingularError{company.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CompanyQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Companies.
func (cq *CompanyQuery) All(ctx context.Context) ([]*Company, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *CompanyQuery) AllX(ctx context.Context) []*Company {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Company IDs.
func (cq *CompanyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(company.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CompanyQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CompanyQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CompanyQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CompanyQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CompanyQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CompanyQuery) Clone() *CompanyQuery {
	if cq == nil {
		return nil
	}
	return &CompanyQuery{
		config:       cq.config,
		limit:        cq.limit,
		offset:       cq.offset,
		order:        append([]OrderFunc{}, cq.order...),
		predicates:   append([]predicate.Company{}, cq.predicates...),
		withUser:     cq.withUser.Clone(),
		withProject:  cq.withProject.Clone(),
		withTaskType: cq.withTaskType.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithUser(opts ...func(*UserQuery)) *CompanyQuery {
	query := &UserQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withUser = query
	return cq
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithProject(opts ...func(*ProjectQuery)) *CompanyQuery {
	query := &ProjectQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withProject = query
	return cq
}

// WithTaskType tells the query-builder to eager-load the nodes that are connected to
// the "task_type" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CompanyQuery) WithTaskType(opts ...func(*TaskTypeQuery)) *CompanyQuery {
	query := &TaskTypeQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withTaskType = query
	return cq
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
//	client.Company.Query().
//		GroupBy(company.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CompanyQuery) GroupBy(field string, fields ...string) *CompanyGroupBy {
	group := &CompanyGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	return group
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
//	client.Company.Query().
//		Select(company.FieldName).
//		Scan(ctx, &v)
//
func (cq *CompanyQuery) Select(fields ...string) *CompanySelect {
	cq.fields = append(cq.fields, fields...)
	return &CompanySelect{CompanyQuery: cq}
}

func (cq *CompanyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !company.ValidColumn(f) {
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

func (cq *CompanyQuery) sqlAll(ctx context.Context) ([]*Company, error) {
	var (
		nodes       = []*Company{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withUser != nil,
			cq.withProject != nil,
			cq.withTaskType != nil,
		}
	)
	if cq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, company.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Company{config: cq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Company)
		for i := range nodes {
			if nodes[i].user_company == nil {
				continue
			}
			fk := *nodes[i].user_company
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_company" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := cq.withProject; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Company)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Project = []*Project{}
		}
		query.withFKs = true
		query.Where(predicate.Project(func(s *sql.Selector) {
			s.Where(sql.InValues(company.ProjectColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.company_project
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "company_project" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "company_project" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Project = append(node.Edges.Project, n)
		}
	}

	if query := cq.withTaskType; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Company)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TaskType = []*TaskType{}
		}
		query.withFKs = true
		query.Where(predicate.TaskType(func(s *sql.Selector) {
			s.Where(sql.InValues(company.TaskTypeColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.company_task_type
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "company_task_type" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "company_task_type" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.TaskType = append(node.Edges.TaskType, n)
		}
	}

	return nodes, nil
}

func (cq *CompanyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CompanyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cq *CompanyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, company.FieldID)
		for i := range fields {
			if fields[i] != company.FieldID {
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
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
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

func (cq *CompanyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(company.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = company.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompanyGroupBy is the group-by builder for Company entities.
type CompanyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CompanyGroupBy) Aggregate(fns ...AggregateFunc) *CompanyGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *CompanyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *CompanyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CompanyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *CompanyGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cgb *CompanyGroupBy) StringX(ctx context.Context) string {
	v, err := cgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CompanyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *CompanyGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cgb *CompanyGroupBy) IntX(ctx context.Context) int {
	v, err := cgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CompanyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *CompanyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cgb *CompanyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: CompanyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *CompanyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *CompanyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cgb *CompanyGroupBy) BoolX(ctx context.Context) bool {
	v, err := cgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *CompanyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !company.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *CompanyGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// CompanySelect is the builder for selecting fields of Company entities.
type CompanySelect struct {
	*CompanyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CompanySelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.CompanyQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *CompanySelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CompanySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *CompanySelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cs *CompanySelect) StringX(ctx context.Context) string {
	v, err := cs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CompanySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *CompanySelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cs *CompanySelect) IntX(ctx context.Context) int {
	v, err := cs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CompanySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *CompanySelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cs *CompanySelect) Float64X(ctx context.Context) float64 {
	v, err := cs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: CompanySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *CompanySelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cs *CompanySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{company.Label}
	default:
		err = fmt.Errorf("ent: CompanySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cs *CompanySelect) BoolX(ctx context.Context) bool {
	v, err := cs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *CompanySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
