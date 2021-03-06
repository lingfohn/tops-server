// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/lingfohn/lime/ent/permission"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetMethod sets the method field.
func (pc *PermissionCreate) SetMethod(s string) *PermissionCreate {
	pc.mutation.SetMethod(s)
	return pc
}

// SetFullpath sets the fullpath field.
func (pc *PermissionCreate) SetFullpath(s string) *PermissionCreate {
	pc.mutation.SetFullpath(s)
	return pc
}

// SetAction sets the action field.
func (pc *PermissionCreate) SetAction(s string) *PermissionCreate {
	pc.mutation.SetAction(s)
	return pc
}

// SetSummary sets the summary field.
func (pc *PermissionCreate) SetSummary(s string) *PermissionCreate {
	pc.mutation.SetSummary(s)
	return pc
}

// SetControlLevel sets the controlLevel field.
func (pc *PermissionCreate) SetControlLevel(i int) *PermissionCreate {
	pc.mutation.SetControlLevel(i)
	return pc
}

// SetStatus sets the status field.
func (pc *PermissionCreate) SetStatus(i int) *PermissionCreate {
	pc.mutation.SetStatus(i)
	return pc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (pc *PermissionCreate) SetNillableStatus(i *int) *PermissionCreate {
	if i != nil {
		pc.SetStatus(*i)
	}
	return pc
}

// SetCreatedAt sets the createdAt field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updatedAt field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	var (
		err  error
		node *Permission
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.Status(); !ok {
		v := permission.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := permission.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New("ent: missing required field \"method\"")}
	}
	if _, ok := pc.mutation.Fullpath(); !ok {
		return &ValidationError{Name: "fullpath", err: errors.New("ent: missing required field \"fullpath\"")}
	}
	if _, ok := pc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New("ent: missing required field \"action\"")}
	}
	if _, ok := pc.mutation.Summary(); !ok {
		return &ValidationError{Name: "summary", err: errors.New("ent: missing required field \"summary\"")}
	}
	if _, ok := pc.mutation.ControlLevel(); !ok {
		return &ValidationError{Name: "controlLevel", err: errors.New("ent: missing required field \"controlLevel\"")}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New("ent: missing required field \"createdAt\"")}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: permission.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permission.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := pc.mutation.Fullpath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldFullpath,
		})
		_node.Fullpath = value
	}
	if value, ok := pc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldAction,
		})
		_node.Action = value
	}
	if value, ok := pc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSummary,
		})
		_node.Summary = value
	}
	if value, ok := pc.mutation.ControlLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
		_node.ControlLevel = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// PermissionCreateBulk is the builder for creating a bulk of Permission entities.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
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

// SaveX calls Save and panics if Save returns an error.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
