// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/build"
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/instance"
)

// InstanceCreate is the builder for creating a Instance entity.
type InstanceCreate struct {
	config
	mutation *InstanceMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ic *InstanceCreate) SetName(s string) *InstanceCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetApplicationId sets the applicationId field.
func (ic *InstanceCreate) SetApplicationId(i int) *InstanceCreate {
	ic.mutation.SetApplicationId(i)
	return ic
}

// SetCreatedAt sets the createdAt field.
func (ic *InstanceCreate) SetCreatedAt(t time.Time) *InstanceCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (ic *InstanceCreate) SetNillableCreatedAt(t *time.Time) *InstanceCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the updatedAt field.
func (ic *InstanceCreate) SetUpdatedAt(t time.Time) *InstanceCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (ic *InstanceCreate) SetNillableUpdatedAt(t *time.Time) *InstanceCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetApplicationID sets the application edge to Application by id.
func (ic *InstanceCreate) SetApplicationID(id int) *InstanceCreate {
	ic.mutation.SetApplicationID(id)
	return ic
}

// SetNillableApplicationID sets the application edge to Application by id if the given value is not nil.
func (ic *InstanceCreate) SetNillableApplicationID(id *int) *InstanceCreate {
	if id != nil {
		ic = ic.SetApplicationID(*id)
	}
	return ic
}

// SetApplication sets the application edge to Application.
func (ic *InstanceCreate) SetApplication(a *Application) *InstanceCreate {
	return ic.SetApplicationID(a.ID)
}

// AddBuildIDs adds the builds edge to Build by ids.
func (ic *InstanceCreate) AddBuildIDs(ids ...int) *InstanceCreate {
	ic.mutation.AddBuildIDs(ids...)
	return ic
}

// AddBuilds adds the builds edges to Build.
func (ic *InstanceCreate) AddBuilds(b ...*Build) *InstanceCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ic.AddBuildIDs(ids...)
}

// SetConfigID sets the config edge to HelmConfig by id.
func (ic *InstanceCreate) SetConfigID(id int) *InstanceCreate {
	ic.mutation.SetConfigID(id)
	return ic
}

// SetNillableConfigID sets the config edge to HelmConfig by id if the given value is not nil.
func (ic *InstanceCreate) SetNillableConfigID(id *int) *InstanceCreate {
	if id != nil {
		ic = ic.SetConfigID(*id)
	}
	return ic
}

// SetConfig sets the config edge to HelmConfig.
func (ic *InstanceCreate) SetConfig(h *HelmConfig) *InstanceCreate {
	return ic.SetConfigID(h.ID)
}

// Mutation returns the InstanceMutation object of the builder.
func (ic *InstanceCreate) Mutation() *InstanceMutation {
	return ic.mutation
}

// Save creates the Instance in the database.
func (ic *InstanceCreate) Save(ctx context.Context) (*Instance, error) {
	var (
		err  error
		node *Instance
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			node, err = ic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstanceCreate) SaveX(ctx context.Context) *Instance {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ic *InstanceCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := instance.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := instance.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InstanceCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ic.mutation.ApplicationId(); !ok {
		return &ValidationError{Name: "applicationId", err: errors.New("ent: missing required field \"applicationId\"")}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New("ent: missing required field \"createdAt\"")}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (ic *InstanceCreate) sqlSave(ctx context.Context) (*Instance, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *InstanceCreate) createSpec() (*Instance, *sqlgraph.CreateSpec) {
	var (
		_node = &Instance{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instance.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ic.mutation.ApplicationId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instance.FieldApplicationId,
		})
		_node.ApplicationId = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ic.mutation.ApplicationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instance.ApplicationTable,
			Columns: []string{instance.ApplicationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: application.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.BuildsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instance.BuildsTable,
			Columns: []string{instance.BuildsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   instance.ConfigTable,
			Columns: []string{instance.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: helmconfig.FieldID,
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

// InstanceCreateBulk is the builder for creating a bulk of Instance entities.
type InstanceCreateBulk struct {
	config
	builders []*InstanceCreate
}

// Save creates the Instance entities in the database.
func (icb *InstanceCreateBulk) Save(ctx context.Context) ([]*Instance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Instance, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstanceMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (icb *InstanceCreateBulk) SaveX(ctx context.Context) []*Instance {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
