// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/chargestatecache"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/ent/predicate"
)

// ChargeStateCacheDelete is the builder for deleting a ChargeStateCache entity.
type ChargeStateCacheDelete struct {
	config
	hooks    []Hook
	mutation *ChargeStateCacheMutation
}

// Where appends a list predicates to the ChargeStateCacheDelete builder.
func (cscd *ChargeStateCacheDelete) Where(ps ...predicate.ChargeStateCache) *ChargeStateCacheDelete {
	cscd.mutation.Where(ps...)
	return cscd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cscd *ChargeStateCacheDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cscd.sqlExec, cscd.mutation, cscd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cscd *ChargeStateCacheDelete) ExecX(ctx context.Context) int {
	n, err := cscd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cscd *ChargeStateCacheDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chargestatecache.Table, sqlgraph.NewFieldSpec(chargestatecache.FieldID, field.TypeInt))
	if ps := cscd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cscd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cscd.mutation.done = true
	return affected, err
}

// ChargeStateCacheDeleteOne is the builder for deleting a single ChargeStateCache entity.
type ChargeStateCacheDeleteOne struct {
	cscd *ChargeStateCacheDelete
}

// Where appends a list predicates to the ChargeStateCacheDelete builder.
func (cscdo *ChargeStateCacheDeleteOne) Where(ps ...predicate.ChargeStateCache) *ChargeStateCacheDeleteOne {
	cscdo.cscd.mutation.Where(ps...)
	return cscdo
}

// Exec executes the deletion query.
func (cscdo *ChargeStateCacheDeleteOne) Exec(ctx context.Context) error {
	n, err := cscdo.cscd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chargestatecache.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cscdo *ChargeStateCacheDeleteOne) ExecX(ctx context.Context) {
	if err := cscdo.Exec(ctx); err != nil {
		panic(err)
	}
}
