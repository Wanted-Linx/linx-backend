// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

// The ClubFunc type is an adapter to allow the use of ordinary
// function as Club mutator.
type ClubFunc func(context.Context, *ent.ClubMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClubFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ClubMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClubMutation", m)
	}
	return f(ctx, mv)
}

// The ClubMemberFunc type is an adapter to allow the use of ordinary
// function as ClubMember mutator.
type ClubMemberFunc func(context.Context, *ent.ClubMemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ClubMemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ClubMemberMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ClubMemberMutation", m)
	}
	return f(ctx, mv)
}

// The CompanyFunc type is an adapter to allow the use of ordinary
// function as Company mutator.
type CompanyFunc func(context.Context, *ent.CompanyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompanyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *ent.ProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectClubFunc type is an adapter to allow the use of ordinary
// function as ProjectClub mutator.
type ProjectClubFunc func(context.Context, *ent.ProjectClubMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectClubFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectClubMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectClubMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectLogFunc type is an adapter to allow the use of ordinary
// function as ProjectLog mutator.
type ProjectLogFunc func(context.Context, *ent.ProjectLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectLogMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectLogMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectLogFeedbackFunc type is an adapter to allow the use of ordinary
// function as ProjectLogFeedback mutator.
type ProjectLogFeedbackFunc func(context.Context, *ent.ProjectLogFeedbackMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectLogFeedbackFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectLogFeedbackMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectLogFeedbackMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectLogParticipantFunc type is an adapter to allow the use of ordinary
// function as ProjectLogParticipant mutator.
type ProjectLogParticipantFunc func(context.Context, *ent.ProjectLogParticipantMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectLogParticipantFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectLogParticipantMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectLogParticipantMutation", m)
	}
	return f(ctx, mv)
}

// The StudentFunc type is an adapter to allow the use of ordinary
// function as Student mutator.
type StudentFunc func(context.Context, *ent.StudentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StudentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StudentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StudentMutation", m)
	}
	return f(ctx, mv)
}

// The TaskTypeFunc type is an adapter to allow the use of ordinary
// function as TaskType mutator.
type TaskTypeFunc func(context.Context, *ent.TaskTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskTypeMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
