// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/Wanted-Linx/linx-backend/api/ent/clubmember"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
	"github.com/Wanted-Linx/linx-backend/api/ent/tasktype"
	"github.com/Wanted-Linx/linx-backend/api/ent/user"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetUniversity sets the "university" field.
func (su *StudentUpdate) SetUniversity(s string) *StudentUpdate {
	su.mutation.SetUniversity(s)
	return su
}

// SetDescription sets the "description" field.
func (su *StudentUpdate) SetDescription(s string) *StudentUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *StudentUpdate) SetNillableDescription(s *string) *StudentUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *StudentUpdate) ClearDescription() *StudentUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetProfileLink sets the "profile_link" field.
func (su *StudentUpdate) SetProfileLink(s string) *StudentUpdate {
	su.mutation.SetProfileLink(s)
	return su
}

// SetNillableProfileLink sets the "profile_link" field if the given value is not nil.
func (su *StudentUpdate) SetNillableProfileLink(s *string) *StudentUpdate {
	if s != nil {
		su.SetProfileLink(*s)
	}
	return su
}

// ClearProfileLink clears the value of the "profile_link" field.
func (su *StudentUpdate) ClearProfileLink() *StudentUpdate {
	su.mutation.ClearProfileLink()
	return su
}

// SetProfileImage sets the "profile_image" field.
func (su *StudentUpdate) SetProfileImage(s string) *StudentUpdate {
	su.mutation.SetProfileImage(s)
	return su
}

// SetNillableProfileImage sets the "profile_image" field if the given value is not nil.
func (su *StudentUpdate) SetNillableProfileImage(s *string) *StudentUpdate {
	if s != nil {
		su.SetProfileImage(*s)
	}
	return su
}

// ClearProfileImage clears the value of the "profile_image" field.
func (su *StudentUpdate) ClearProfileImage() *StudentUpdate {
	su.mutation.ClearProfileImage()
	return su
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *StudentUpdate) SetUserID(id int) *StudentUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *StudentUpdate) SetUser(u *User) *StudentUpdate {
	return su.SetUserID(u.ID)
}

// AddClubIDs adds the "club" edge to the Club entity by IDs.
func (su *StudentUpdate) AddClubIDs(ids ...int) *StudentUpdate {
	su.mutation.AddClubIDs(ids...)
	return su
}

// AddClub adds the "club" edges to the Club entity.
func (su *StudentUpdate) AddClub(c ...*Club) *StudentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddClubIDs(ids...)
}

// AddClubMemberIDs adds the "club_member" edge to the ClubMember entity by IDs.
func (su *StudentUpdate) AddClubMemberIDs(ids ...int) *StudentUpdate {
	su.mutation.AddClubMemberIDs(ids...)
	return su
}

// AddClubMember adds the "club_member" edges to the ClubMember entity.
func (su *StudentUpdate) AddClubMember(c ...*ClubMember) *StudentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddClubMemberIDs(ids...)
}

// AddTaskTypeIDs adds the "task_type" edge to the TaskType entity by IDs.
func (su *StudentUpdate) AddTaskTypeIDs(ids ...int) *StudentUpdate {
	su.mutation.AddTaskTypeIDs(ids...)
	return su
}

// AddTaskType adds the "task_type" edges to the TaskType entity.
func (su *StudentUpdate) AddTaskType(t ...*TaskType) *StudentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.AddTaskTypeIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (su *StudentUpdate) ClearUser() *StudentUpdate {
	su.mutation.ClearUser()
	return su
}

// ClearClub clears all "club" edges to the Club entity.
func (su *StudentUpdate) ClearClub() *StudentUpdate {
	su.mutation.ClearClub()
	return su
}

// RemoveClubIDs removes the "club" edge to Club entities by IDs.
func (su *StudentUpdate) RemoveClubIDs(ids ...int) *StudentUpdate {
	su.mutation.RemoveClubIDs(ids...)
	return su
}

// RemoveClub removes "club" edges to Club entities.
func (su *StudentUpdate) RemoveClub(c ...*Club) *StudentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveClubIDs(ids...)
}

// ClearClubMember clears all "club_member" edges to the ClubMember entity.
func (su *StudentUpdate) ClearClubMember() *StudentUpdate {
	su.mutation.ClearClubMember()
	return su
}

// RemoveClubMemberIDs removes the "club_member" edge to ClubMember entities by IDs.
func (su *StudentUpdate) RemoveClubMemberIDs(ids ...int) *StudentUpdate {
	su.mutation.RemoveClubMemberIDs(ids...)
	return su
}

// RemoveClubMember removes "club_member" edges to ClubMember entities.
func (su *StudentUpdate) RemoveClubMember(c ...*ClubMember) *StudentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveClubMemberIDs(ids...)
}

// ClearTaskType clears all "task_type" edges to the TaskType entity.
func (su *StudentUpdate) ClearTaskType() *StudentUpdate {
	su.mutation.ClearTaskType()
	return su
}

// RemoveTaskTypeIDs removes the "task_type" edge to TaskType entities by IDs.
func (su *StudentUpdate) RemoveTaskTypeIDs(ids ...int) *StudentUpdate {
	su.mutation.RemoveTaskTypeIDs(ids...)
	return su
}

// RemoveTaskType removes "task_type" edges to TaskType entities.
func (su *StudentUpdate) RemoveTaskType(t ...*TaskType) *StudentUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return su.RemoveTaskTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StudentUpdate) check() error {
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   student.Table,
			Columns: student.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: student.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldName,
		})
	}
	if value, ok := su.mutation.University(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldUniversity,
		})
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldDescription,
		})
	}
	if su.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldDescription,
		})
	}
	if value, ok := su.mutation.ProfileLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldProfileLink,
		})
	}
	if su.mutation.ProfileLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldProfileLink,
		})
	}
	if value, ok := su.mutation.ProfileImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldProfileImage,
		})
	}
	if su.mutation.ProfileImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldProfileImage,
		})
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedClubIDs(); len(nodes) > 0 && !su.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedClubMemberIDs(); len(nodes) > 0 && !su.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ClubMemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.TaskTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedTaskTypeIDs(); len(nodes) > 0 && !su.mutation.TaskTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.TaskTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetUniversity sets the "university" field.
func (suo *StudentUpdateOne) SetUniversity(s string) *StudentUpdateOne {
	suo.mutation.SetUniversity(s)
	return suo
}

// SetDescription sets the "description" field.
func (suo *StudentUpdateOne) SetDescription(s string) *StudentUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableDescription(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *StudentUpdateOne) ClearDescription() *StudentUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetProfileLink sets the "profile_link" field.
func (suo *StudentUpdateOne) SetProfileLink(s string) *StudentUpdateOne {
	suo.mutation.SetProfileLink(s)
	return suo
}

// SetNillableProfileLink sets the "profile_link" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableProfileLink(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetProfileLink(*s)
	}
	return suo
}

// ClearProfileLink clears the value of the "profile_link" field.
func (suo *StudentUpdateOne) ClearProfileLink() *StudentUpdateOne {
	suo.mutation.ClearProfileLink()
	return suo
}

// SetProfileImage sets the "profile_image" field.
func (suo *StudentUpdateOne) SetProfileImage(s string) *StudentUpdateOne {
	suo.mutation.SetProfileImage(s)
	return suo
}

// SetNillableProfileImage sets the "profile_image" field if the given value is not nil.
func (suo *StudentUpdateOne) SetNillableProfileImage(s *string) *StudentUpdateOne {
	if s != nil {
		suo.SetProfileImage(*s)
	}
	return suo
}

// ClearProfileImage clears the value of the "profile_image" field.
func (suo *StudentUpdateOne) ClearProfileImage() *StudentUpdateOne {
	suo.mutation.ClearProfileImage()
	return suo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *StudentUpdateOne) SetUserID(id int) *StudentUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *StudentUpdateOne) SetUser(u *User) *StudentUpdateOne {
	return suo.SetUserID(u.ID)
}

// AddClubIDs adds the "club" edge to the Club entity by IDs.
func (suo *StudentUpdateOne) AddClubIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.AddClubIDs(ids...)
	return suo
}

// AddClub adds the "club" edges to the Club entity.
func (suo *StudentUpdateOne) AddClub(c ...*Club) *StudentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddClubIDs(ids...)
}

// AddClubMemberIDs adds the "club_member" edge to the ClubMember entity by IDs.
func (suo *StudentUpdateOne) AddClubMemberIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.AddClubMemberIDs(ids...)
	return suo
}

// AddClubMember adds the "club_member" edges to the ClubMember entity.
func (suo *StudentUpdateOne) AddClubMember(c ...*ClubMember) *StudentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddClubMemberIDs(ids...)
}

// AddTaskTypeIDs adds the "task_type" edge to the TaskType entity by IDs.
func (suo *StudentUpdateOne) AddTaskTypeIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.AddTaskTypeIDs(ids...)
	return suo
}

// AddTaskType adds the "task_type" edges to the TaskType entity.
func (suo *StudentUpdateOne) AddTaskType(t ...*TaskType) *StudentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.AddTaskTypeIDs(ids...)
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (suo *StudentUpdateOne) ClearUser() *StudentUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// ClearClub clears all "club" edges to the Club entity.
func (suo *StudentUpdateOne) ClearClub() *StudentUpdateOne {
	suo.mutation.ClearClub()
	return suo
}

// RemoveClubIDs removes the "club" edge to Club entities by IDs.
func (suo *StudentUpdateOne) RemoveClubIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.RemoveClubIDs(ids...)
	return suo
}

// RemoveClub removes "club" edges to Club entities.
func (suo *StudentUpdateOne) RemoveClub(c ...*Club) *StudentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveClubIDs(ids...)
}

// ClearClubMember clears all "club_member" edges to the ClubMember entity.
func (suo *StudentUpdateOne) ClearClubMember() *StudentUpdateOne {
	suo.mutation.ClearClubMember()
	return suo
}

// RemoveClubMemberIDs removes the "club_member" edge to ClubMember entities by IDs.
func (suo *StudentUpdateOne) RemoveClubMemberIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.RemoveClubMemberIDs(ids...)
	return suo
}

// RemoveClubMember removes "club_member" edges to ClubMember entities.
func (suo *StudentUpdateOne) RemoveClubMember(c ...*ClubMember) *StudentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveClubMemberIDs(ids...)
}

// ClearTaskType clears all "task_type" edges to the TaskType entity.
func (suo *StudentUpdateOne) ClearTaskType() *StudentUpdateOne {
	suo.mutation.ClearTaskType()
	return suo
}

// RemoveTaskTypeIDs removes the "task_type" edge to TaskType entities by IDs.
func (suo *StudentUpdateOne) RemoveTaskTypeIDs(ids ...int) *StudentUpdateOne {
	suo.mutation.RemoveTaskTypeIDs(ids...)
	return suo
}

// RemoveTaskType removes "task_type" edges to TaskType entities.
func (suo *StudentUpdateOne) RemoveTaskType(t ...*TaskType) *StudentUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return suo.RemoveTaskTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	var (
		err  error
		node *Student
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StudentUpdateOne) check() error {
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   student.Table,
			Columns: student.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: student.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Student.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldName,
		})
	}
	if value, ok := suo.mutation.University(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldUniversity,
		})
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldDescription,
		})
	}
	if suo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldDescription,
		})
	}
	if value, ok := suo.mutation.ProfileLink(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldProfileLink,
		})
	}
	if suo.mutation.ProfileLinkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldProfileLink,
		})
	}
	if value, ok := suo.mutation.ProfileImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: student.FieldProfileImage,
		})
	}
	if suo.mutation.ProfileImageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: student.FieldProfileImage,
		})
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.UserTable,
			Columns: []string{student.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedClubIDs(); len(nodes) > 0 && !suo.mutation.ClubCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubTable,
			Columns: []string{student.ClubColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedClubMemberIDs(); len(nodes) > 0 && !suo.mutation.ClubMemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ClubMemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ClubMemberTable,
			Columns: []string{student.ClubMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: clubmember.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.TaskTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedTaskTypeIDs(); len(nodes) > 0 && !suo.mutation.TaskTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.TaskTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.TaskTypeTable,
			Columns: []string{student.TaskTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
