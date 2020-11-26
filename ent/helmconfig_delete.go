// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/predicate"
)

// HelmConfigDelete is the builder for deleting a HelmConfig entity.
type HelmConfigDelete struct {
	config
	hooks      []Hook
	mutation   *HelmConfigMutation
	predicates []predicate.HelmConfig
}

// Where adds a new predicate to the delete builder.
func (hcd *HelmConfigDelete) Where(ps ...predicate.HelmConfig) *HelmConfigDelete {
	hcd.predicates = append(hcd.predicates, ps...)
	return hcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hcd *HelmConfigDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hcd.hooks) == 0 {
		affected, err = hcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HelmConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hcd.mutation = mutation
			affected, err = hcd.sqlExec(ctx)
			return affected, err
		})
		for i := len(hcd.hooks) - 1; i >= 0; i-- {
			mut = hcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcd *HelmConfigDelete) ExecX(ctx context.Context) int {
	n, err := hcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hcd *HelmConfigDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: helmconfig.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: helmconfig.FieldID,
			},
		},
	}
	if ps := hcd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hcd.driver, _spec)
}

// HelmConfigDeleteOne is the builder for deleting a single HelmConfig entity.
type HelmConfigDeleteOne struct {
	hcd *HelmConfigDelete
}

// Exec executes the deletion query.
func (hcdo *HelmConfigDeleteOne) Exec(ctx context.Context) error {
	n, err := hcdo.hcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{helmconfig.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hcdo *HelmConfigDeleteOne) ExecX(ctx context.Context) {
	hcdo.hcd.ExecX(ctx)
}
