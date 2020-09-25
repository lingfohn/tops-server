// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/lingfohn/lime/ent/permission"
	"github.com/lingfohn/lime/ent/predicate"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks      []Hook
	mutation   *PermissionMutation
	predicates []predicate.Permission
}

// Where adds a new predicate for the builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetMethod sets the method field.
func (pu *PermissionUpdate) SetMethod(s string) *PermissionUpdate {
	pu.mutation.SetMethod(s)
	return pu
}

// SetFullpath sets the fullpath field.
func (pu *PermissionUpdate) SetFullpath(s string) *PermissionUpdate {
	pu.mutation.SetFullpath(s)
	return pu
}

// SetAction sets the action field.
func (pu *PermissionUpdate) SetAction(s string) *PermissionUpdate {
	pu.mutation.SetAction(s)
	return pu
}

// SetSummary sets the summary field.
func (pu *PermissionUpdate) SetSummary(s string) *PermissionUpdate {
	pu.mutation.SetSummary(s)
	return pu
}

// SetControlLevel sets the controlLevel field.
func (pu *PermissionUpdate) SetControlLevel(i int) *PermissionUpdate {
	pu.mutation.ResetControlLevel()
	pu.mutation.SetControlLevel(i)
	return pu
}

// AddControlLevel adds i to controlLevel.
func (pu *PermissionUpdate) AddControlLevel(i int) *PermissionUpdate {
	pu.mutation.AddControlLevel(i)
	return pu
}

// SetStatus sets the status field.
func (pu *PermissionUpdate) SetStatus(i int) *PermissionUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableStatus(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to status.
func (pu *PermissionUpdate) AddStatus(i int) *PermissionUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// SetUpdatedAt sets the updatedAt field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldMethod,
		})
	}
	if value, ok := pu.mutation.Fullpath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldFullpath,
		})
	}
	if value, ok := pu.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldAction,
		})
	}
	if value, ok := pu.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSummary,
		})
	}
	if value, ok := pu.mutation.ControlLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
	}
	if value, ok := pu.mutation.AddedControlLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// SetMethod sets the method field.
func (puo *PermissionUpdateOne) SetMethod(s string) *PermissionUpdateOne {
	puo.mutation.SetMethod(s)
	return puo
}

// SetFullpath sets the fullpath field.
func (puo *PermissionUpdateOne) SetFullpath(s string) *PermissionUpdateOne {
	puo.mutation.SetFullpath(s)
	return puo
}

// SetAction sets the action field.
func (puo *PermissionUpdateOne) SetAction(s string) *PermissionUpdateOne {
	puo.mutation.SetAction(s)
	return puo
}

// SetSummary sets the summary field.
func (puo *PermissionUpdateOne) SetSummary(s string) *PermissionUpdateOne {
	puo.mutation.SetSummary(s)
	return puo
}

// SetControlLevel sets the controlLevel field.
func (puo *PermissionUpdateOne) SetControlLevel(i int) *PermissionUpdateOne {
	puo.mutation.ResetControlLevel()
	puo.mutation.SetControlLevel(i)
	return puo
}

// AddControlLevel adds i to controlLevel.
func (puo *PermissionUpdateOne) AddControlLevel(i int) *PermissionUpdateOne {
	puo.mutation.AddControlLevel(i)
	return puo
}

// SetStatus sets the status field.
func (puo *PermissionUpdateOne) SetStatus(i int) *PermissionUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableStatus(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to status.
func (puo *PermissionUpdateOne) AddStatus(i int) *PermissionUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// SetUpdatedAt sets the updatedAt field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	var (
		err  error
		node *Permission
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (pe *Permission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Permission.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldMethod,
		})
	}
	if value, ok := puo.mutation.Fullpath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldFullpath,
		})
	}
	if value, ok := puo.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldAction,
		})
	}
	if value, ok := puo.mutation.Summary(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSummary,
		})
	}
	if value, ok := puo.mutation.ControlLevel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
	}
	if value, ok := puo.mutation.AddedControlLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
	}
	pe = &Permission{config: puo.config}
	_spec.Assign = pe.assignValues
	_spec.ScanValues = pe.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pe, nil
}