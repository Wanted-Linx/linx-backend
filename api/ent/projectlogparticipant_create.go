// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlog"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlogparticipant"
)

// ProjectLogParticipantCreate is the builder for creating a ProjectLogParticipant entity.
type ProjectLogParticipantCreate struct {
	config
	mutation *ProjectLogParticipantMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (plpc *ProjectLogParticipantCreate) SetName(s string) *ProjectLogParticipantCreate {
	plpc.mutation.SetName(s)
	return plpc
}

// SetProjectLogID sets the "project_log" edge to the ProjectLog entity by ID.
func (plpc *ProjectLogParticipantCreate) SetProjectLogID(id int) *ProjectLogParticipantCreate {
	plpc.mutation.SetProjectLogID(id)
	return plpc
}

// SetProjectLog sets the "project_log" edge to the ProjectLog entity.
func (plpc *ProjectLogParticipantCreate) SetProjectLog(p *ProjectLog) *ProjectLogParticipantCreate {
	return plpc.SetProjectLogID(p.ID)
}

// Mutation returns the ProjectLogParticipantMutation object of the builder.
func (plpc *ProjectLogParticipantCreate) Mutation() *ProjectLogParticipantMutation {
	return plpc.mutation
}

// Save creates the ProjectLogParticipant in the database.
func (plpc *ProjectLogParticipantCreate) Save(ctx context.Context) (*ProjectLogParticipant, error) {
	var (
		err  error
		node *ProjectLogParticipant
	)
	if len(plpc.hooks) == 0 {
		if err = plpc.check(); err != nil {
			return nil, err
		}
		node, err = plpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectLogParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = plpc.check(); err != nil {
				return nil, err
			}
			plpc.mutation = mutation
			if node, err = plpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(plpc.hooks) - 1; i >= 0; i-- {
			if plpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = plpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, plpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (plpc *ProjectLogParticipantCreate) SaveX(ctx context.Context) *ProjectLogParticipant {
	v, err := plpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (plpc *ProjectLogParticipantCreate) Exec(ctx context.Context) error {
	_, err := plpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plpc *ProjectLogParticipantCreate) ExecX(ctx context.Context) {
	if err := plpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (plpc *ProjectLogParticipantCreate) check() error {
	if _, ok := plpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := plpc.mutation.ProjectLogID(); !ok {
		return &ValidationError{Name: "project_log", err: errors.New("ent: missing required edge \"project_log\"")}
	}
	return nil
}

func (plpc *ProjectLogParticipantCreate) sqlSave(ctx context.Context) (*ProjectLogParticipant, error) {
	_node, _spec := plpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, plpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (plpc *ProjectLogParticipantCreate) createSpec() (*ProjectLogParticipant, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectLogParticipant{config: plpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: projectlogparticipant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: projectlogparticipant.FieldID,
			},
		}
	)
	if value, ok := plpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: projectlogparticipant.FieldName,
		})
		_node.Name = value
	}
	if nodes := plpc.mutation.ProjectLogIDs(); len(nodes) > 0 {
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
		_node.project_log_project_log_participant = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectLogParticipantCreateBulk is the builder for creating many ProjectLogParticipant entities in bulk.
type ProjectLogParticipantCreateBulk struct {
	config
	builders []*ProjectLogParticipantCreate
}

// Save creates the ProjectLogParticipant entities in the database.
func (plpcb *ProjectLogParticipantCreateBulk) Save(ctx context.Context) ([]*ProjectLogParticipant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(plpcb.builders))
	nodes := make([]*ProjectLogParticipant, len(plpcb.builders))
	mutators := make([]Mutator, len(plpcb.builders))
	for i := range plpcb.builders {
		func(i int, root context.Context) {
			builder := plpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectLogParticipantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, plpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, plpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, plpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (plpcb *ProjectLogParticipantCreateBulk) SaveX(ctx context.Context) []*ProjectLogParticipant {
	v, err := plpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (plpcb *ProjectLogParticipantCreateBulk) Exec(ctx context.Context) error {
	_, err := plpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (plpcb *ProjectLogParticipantCreateBulk) ExecX(ctx context.Context) {
	if err := plpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
