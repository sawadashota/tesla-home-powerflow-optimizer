// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargecommandhistory"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/predicate"
)

// ChargeCommandHistoryDelete is the builder for deleting a ChargeCommandHistory entity.
type ChargeCommandHistoryDelete struct {
	config
	hooks    []Hook
	mutation *ChargeCommandHistoryMutation
}

// Where appends a list predicates to the ChargeCommandHistoryDelete builder.
func (cchd *ChargeCommandHistoryDelete) Where(ps ...predicate.ChargeCommandHistory) *ChargeCommandHistoryDelete {
	cchd.mutation.Where(ps...)
	return cchd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cchd *ChargeCommandHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cchd.sqlExec, cchd.mutation, cchd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cchd *ChargeCommandHistoryDelete) ExecX(ctx context.Context) int {
	n, err := cchd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cchd *ChargeCommandHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chargecommandhistory.Table, sqlgraph.NewFieldSpec(chargecommandhistory.FieldID, field.TypeInt))
	if ps := cchd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cchd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cchd.mutation.done = true
	return affected, err
}

// ChargeCommandHistoryDeleteOne is the builder for deleting a single ChargeCommandHistory entity.
type ChargeCommandHistoryDeleteOne struct {
	cchd *ChargeCommandHistoryDelete
}

// Where appends a list predicates to the ChargeCommandHistoryDelete builder.
func (cchdo *ChargeCommandHistoryDeleteOne) Where(ps ...predicate.ChargeCommandHistory) *ChargeCommandHistoryDeleteOne {
	cchdo.cchd.mutation.Where(ps...)
	return cchdo
}

// Exec executes the deletion query.
func (cchdo *ChargeCommandHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := cchdo.cchd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chargecommandhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cchdo *ChargeCommandHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := cchdo.Exec(ctx); err != nil {
		panic(err)
	}
}
