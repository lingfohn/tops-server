// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
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

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	if _, ok := pc.mutation.Method(); !ok {
		return nil, errors.New("ent: missing required field \"method\"")
	}
	if _, ok := pc.mutation.Fullpath(); !ok {
		return nil, errors.New("ent: missing required field \"fullpath\"")
	}
	if _, ok := pc.mutation.Action(); !ok {
		return nil, errors.New("ent: missing required field \"action\"")
	}
	if _, ok := pc.mutation.Summary(); !ok {
		return nil, errors.New("ent: missing required field \"summary\"")
	}
	if _, ok := pc.mutation.ControlLevel(); !ok {
		return nil, errors.New("ent: missing required field \"controlLevel\"")
	}
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
	var (
		err  error
		node *Permission
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
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

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	var (
		pe    = &Permission{config: pc.config}
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
		pe.Method = value
	}
	if value, ok := pc.mutation.Fullpath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldFullpath,
		})
		pe.Fullpath = value
	}
	if value, ok := pc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldAction,
		})
		pe.Action = value
	}
	if value, ok := pc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: permission.FieldSummary,
		})
		pe.Summary = value
	}
	if value, ok := pc.mutation.ControlLevel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldControlLevel,
		})
		pe.ControlLevel = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: permission.FieldStatus,
		})
		pe.Status = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldCreatedAt,
		})
		pe.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: permission.FieldUpdatedAt,
		})
		pe.UpdatedAt = value
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pe.ID = int(id)
	return pe, nil
}