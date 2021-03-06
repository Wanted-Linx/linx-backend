// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlog"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlogparticipant"
)

// ProjectLogParticipantUpdate is the builder for updating ProjectLogParticipant entities.
type ProjectLogParticipantUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectLogParticipantMutation
}

// Where appends a list predicates to the ProjectLogParticipantUpdate builder.
func (plpu *ProjectLogParticipantUpdate) Where(ps ...predicate.ProjectLogParticipant) *ProjectLogParticipantUpdate {
	plpu.mutation.Where(ps...)
	return plpu
}

// SetName sets the "name" field.
func (plpu *ProjectLogParticipantUpdate) SetName(s string) *ProjectLogParticipantUpdate {
	plpu.mutation.SetName(s)
	return plpu
}

// SetProjectLogID sets the "project_log" edge to the ProjectLog entity by ID.
func (plpu *ProjectLogParticipantUpdate) SetProjectLogID(id int) *ProjectLogParticipantUpdate {
	plpu.mutation.SetProjectLogID(id)
	return plpu
}

// SetProjectLog sets the "project_log" edge to the ProjectLog entity.
func (plpu *ProjectLogParticipantUpdate) SetProjectLog(p *ProjectLog) *ProjectLogParticipantUpdate {
	return plpu.SetProjectLogID(p.ID)
}

// Mutation returns the ProjectLogParticipantMutation object of the builder.
func (plpu *ProjectLogParticipantUpdate) Mutation() *ProjectLogParticipantMutation {
	return plpu.mutation
}

// ClearProjectLog clears the "project_log" edge to the ProjectLog entity.
func (plpu *ProjectLogParticipantUpdate) ClearProjectLog() *ProjectLogParticipantUpdate {
	plpu.mutation.ClearProjectLog()
	return plpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (plpu *ProjectLogParticipantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(plpu.hooks) == 0 {
		if err = plpu.check(); err != nil {
			return 0, err
		}
		affected, err = plpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLogParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plpu.check(); err != nil {
				return 0, err
			}
			plpu.mutation = mutation
			affected, err = plpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(plpu.hooks) - 1; i >= 0; i-- {
			if plpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (plpu *ProjectLogParticipantUpdate) SaveX(ctx context.Context) int {
	affected, err := plpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (plpu *ProjectLogParticipantUpdate) Exec(ctx context.Context) error {
	_, err := plpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plpu *ProjectLogParticipantUpdate) ExecX(ctx context.Context) {
	if err := plpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plpu *ProjectLogParticipantUpdate) check() error {
	if _, ok := plpu.mutation.ProjectLogID(); plpu.mutation.ProjectLogCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_log\"")
	}
	return nil
}

func (plpu *ProjectLogParticipantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlogparticipant.Table,
			Columns: projectlogparticipant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectlogparticipant.FieldID,
			},
		},
	}
	if ps := plpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plpu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlogparticipant.FieldName,
		})
	}
	if plpu.mutation.ProjectLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlogparticipant.ProjectLogTable,
			Columns: []string{projectlogparticipant.ProjectLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plpu.mutation.ProjectLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlogparticipant.ProjectLogTable,
			Columns: []string{projectlogparticipant.ProjectLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, plpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlogparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProjectLogParticipantUpdateOne is the builder for updating a single ProjectLogParticipant entity.
type ProjectLogParticipantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectLogParticipantMutation
}

// SetName sets the "name" field.
func (plpuo *ProjectLogParticipantUpdateOne) SetName(s string) *ProjectLogParticipantUpdateOne {
	plpuo.mutation.SetName(s)
	return plpuo
}

// SetProjectLogID sets the "project_log" edge to the ProjectLog entity by ID.
func (plpuo *ProjectLogParticipantUpdateOne) SetProjectLogID(id int) *ProjectLogParticipantUpdateOne {
	plpuo.mutation.SetProjectLogID(id)
	return plpuo
}

// SetProjectLog sets the "project_log" edge to the ProjectLog entity.
func (plpuo *ProjectLogParticipantUpdateOne) SetProjectLog(p *ProjectLog) *ProjectLogParticipantUpdateOne {
	return plpuo.SetProjectLogID(p.ID)
}

// Mutation returns the ProjectLogParticipantMutation object of the builder.
func (plpuo *ProjectLogParticipantUpdateOne) Mutation() *ProjectLogParticipantMutation {
	return plpuo.mutation
}

// ClearProjectLog clears the "project_log" edge to the ProjectLog entity.
func (plpuo *ProjectLogParticipantUpdateOne) ClearProjectLog() *ProjectLogParticipantUpdateOne {
	plpuo.mutation.ClearProjectLog()
	return plpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (plpuo *ProjectLogParticipantUpdateOne) Select(field string, fields ...string) *ProjectLogParticipantUpdateOne {
	plpuo.fields = append([]string{field}, fields...)
	return plpuo
}

// Save executes the query and returns the updated ProjectLogParticipant entity.
func (plpuo *ProjectLogParticipantUpdateOne) Save(ctx context.Context) (*ProjectLogParticipant, error) {
	var (
		err  error
		node *ProjectLogParticipant
	)
	if len(plpuo.hooks) == 0 {
		if err = plpuo.check(); err != nil {
			return nil, err
		}
		node, err = plpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLogParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plpuo.check(); err != nil {
				return nil, err
			}
			plpuo.mutation = mutation
			node, err = plpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(plpuo.hooks) - 1; i >= 0; i-- {
			if plpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (plpuo *ProjectLogParticipantUpdateOne) SaveX(ctx context.Context) *ProjectLogParticipant {
	node, err := plpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (plpuo *ProjectLogParticipantUpdateOne) Exec(ctx context.Context) error {
	_, err := plpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plpuo *ProjectLogParticipantUpdateOne) ExecX(ctx context.Context) {
	if err := plpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plpuo *ProjectLogParticipantUpdateOne) check() error {
	if _, ok := plpuo.mutation.ProjectLogID(); plpuo.mutation.ProjectLogCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"project_log\"")
	}
	return nil
}

func (plpuo *ProjectLogParticipantUpdateOne) sqlSave(ctx context.Context) (_node *ProjectLogParticipant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   projectlogparticipant.Table,
			Columns: projectlogparticipant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectlogparticipant.FieldID,
			},
		},
	}
	id, ok := plpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProjectLogParticipant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := plpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectlogparticipant.FieldID)
		for _, f := range fields {
			if !projectlogparticipant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectlogparticipant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := plpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := plpuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlogparticipant.FieldName,
		})
	}
	if plpuo.mutation.ProjectLogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlogparticipant.ProjectLogTable,
			Columns: []string{projectlogparticipant.ProjectLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := plpuo.mutation.ProjectLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectlogparticipant.ProjectLogTable,
			Columns: []string{projectlogparticipant.ProjectLogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectLogParticipant{config: plpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, plpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectlogparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
