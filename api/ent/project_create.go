// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/company"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectclub"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectlog"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *ProjectCreate) SetContent(s string) *ProjectCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetStartDate sets the "start_date" field.
func (pc *ProjectCreate) SetStartDate(s string) *ProjectCreate {
	pc.mutation.SetStartDate(s)
	return pc
}

// SetEndDate sets the "end_date" field.
func (pc *ProjectCreate) SetEndDate(s string) *ProjectCreate {
	pc.mutation.SetEndDate(s)
	return pc
}

// SetApplyingStartDate sets the "applying_start_date" field.
func (pc *ProjectCreate) SetApplyingStartDate(s string) *ProjectCreate {
	pc.mutation.SetApplyingStartDate(s)
	return pc
}

// SetApplyingEndDate sets the "applying_end_date" field.
func (pc *ProjectCreate) SetApplyingEndDate(s string) *ProjectCreate {
	pc.mutation.SetApplyingEndDate(s)
	return pc
}

// SetQualification sets the "qualification" field.
func (pc *ProjectCreate) SetQualification(s string) *ProjectCreate {
	pc.mutation.SetQualification(s)
	return pc
}

// SetProfileImage sets the "profile_image" field.
func (pc *ProjectCreate) SetProfileImage(s string) *ProjectCreate {
	pc.mutation.SetProfileImage(s)
	return pc
}

// SetNillableProfileImage sets the "profile_image" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableProfileImage(s *string) *ProjectCreate {
	if s != nil {
		pc.SetProfileImage(*s)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetSponsorFee sets the "sponsor_fee" field.
func (pc *ProjectCreate) SetSponsorFee(i int) *ProjectCreate {
	pc.mutation.SetSponsorFee(i)
	return pc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pc *ProjectCreate) SetCompanyID(id int) *ProjectCreate {
	pc.mutation.SetCompanyID(id)
	return pc
}

// SetCompany sets the "company" edge to the Company entity.
func (pc *ProjectCreate) SetCompany(c *Company) *ProjectCreate {
	return pc.SetCompanyID(c.ID)
}

// SetClubID sets the "club" edge to the Club entity by ID.
func (pc *ProjectCreate) SetClubID(id int) *ProjectCreate {
	pc.mutation.SetClubID(id)
	return pc
}

// SetNillableClubID sets the "club" edge to the Club entity by ID if the given value is not nil.
func (pc *ProjectCreate) SetNillableClubID(id *int) *ProjectCreate {
	if id != nil {
		pc = pc.SetClubID(*id)
	}
	return pc
}

// SetClub sets the "club" edge to the Club entity.
func (pc *ProjectCreate) SetClub(c *Club) *ProjectCreate {
	return pc.SetClubID(c.ID)
}

// AddProjectClubIDs adds the "project_club" edge to the ProjectClub entity by IDs.
func (pc *ProjectCreate) AddProjectClubIDs(ids ...int) *ProjectCreate {
	pc.mutation.AddProjectClubIDs(ids...)
	return pc
}

// AddProjectClub adds the "project_club" edges to the ProjectClub entity.
func (pc *ProjectCreate) AddProjectClub(p ...*ProjectClub) *ProjectCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProjectClubIDs(ids...)
}

// AddProjectLogIDs adds the "project_log" edge to the ProjectLog entity by IDs.
func (pc *ProjectCreate) AddProjectLogIDs(ids ...int) *ProjectCreate {
	pc.mutation.AddProjectLogIDs(ids...)
	return pc
}

// AddProjectLog adds the "project_log" edges to the ProjectLog entity.
func (pc *ProjectCreate) AddProjectLog(p ...*ProjectLog) *ProjectCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProjectLogIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if _, ok := pc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "start_date"`)}
	}
	if _, ok := pc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New(`ent: missing required field "end_date"`)}
	}
	if _, ok := pc.mutation.ApplyingStartDate(); !ok {
		return &ValidationError{Name: "applying_start_date", err: errors.New(`ent: missing required field "applying_start_date"`)}
	}
	if _, ok := pc.mutation.ApplyingEndDate(); !ok {
		return &ValidationError{Name: "applying_end_date", err: errors.New(`ent: missing required field "applying_end_date"`)}
	}
	if _, ok := pc.mutation.Qualification(); !ok {
		return &ValidationError{Name: "qualification", err: errors.New(`ent: missing required field "qualification"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := pc.mutation.SponsorFee(); !ok {
		return &ValidationError{Name: "sponsor_fee", err: errors.New(`ent: missing required field "sponsor_fee"`)}
	}
	if _, ok := pc.mutation.CompanyID(); !ok {
		return &ValidationError{Name: "company", err: errors.New("ent: missing required edge \"company\"")}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: project.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := pc.mutation.StartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldStartDate,
		})
		_node.StartDate = value
	}
	if value, ok := pc.mutation.EndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldEndDate,
		})
		_node.EndDate = value
	}
	if value, ok := pc.mutation.ApplyingStartDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldApplyingStartDate,
		})
		_node.ApplyingStartDate = value
	}
	if value, ok := pc.mutation.ApplyingEndDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldApplyingEndDate,
		})
		_node.ApplyingEndDate = value
	}
	if value, ok := pc.mutation.Qualification(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldQualification,
		})
		_node.Qualification = value
	}
	if value, ok := pc.mutation.ProfileImage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldProfileImage,
		})
		_node.ProfileImage = &value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.SponsorFee(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: project.FieldSponsorFee,
		})
		_node.SponsorFee = value
	}
	if nodes := pc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.CompanyTable,
			Columns: []string{project.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_project = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.ClubTable,
			Columns: []string{project.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.club_project = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ProjectClubTable,
			Columns: []string{project.ProjectClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: projectclub.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProjectLogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ProjectLogTable,
			Columns: []string{project.ProjectLogColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
