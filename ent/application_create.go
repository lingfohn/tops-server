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
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/instance"
	"github.com/lingfohn/lime/ent/namespace"
	"github.com/lingfohn/lime/ent/project"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
}

// SetName sets the name field.
func (ac *ApplicationCreate) SetName(s string) *ApplicationCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetMulti sets the multi field.
func (ac *ApplicationCreate) SetMulti(b bool) *ApplicationCreate {
	ac.mutation.SetMulti(b)
	return ac
}

// SetNillableMulti sets the multi field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableMulti(b *bool) *ApplicationCreate {
	if b != nil {
		ac.SetMulti(*b)
	}
	return ac
}

// SetProjectId sets the projectId field.
func (ac *ApplicationCreate) SetProjectId(i int) *ApplicationCreate {
	ac.mutation.SetProjectId(i)
	return ac
}

// SetNamespaceId sets the namespaceId field.
func (ac *ApplicationCreate) SetNamespaceId(i int) *ApplicationCreate {
	ac.mutation.SetNamespaceId(i)
	return ac
}

// SetCreatedAt sets the createdAt field.
func (ac *ApplicationCreate) SetCreatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the updatedAt field.
func (ac *ApplicationCreate) SetUpdatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableUpdatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetNamespaceID sets the namespace edge to Namespace by id.
func (ac *ApplicationCreate) SetNamespaceID(id int) *ApplicationCreate {
	ac.mutation.SetNamespaceID(id)
	return ac
}

// SetNillableNamespaceID sets the namespace edge to Namespace by id if the given value is not nil.
func (ac *ApplicationCreate) SetNillableNamespaceID(id *int) *ApplicationCreate {
	if id != nil {
		ac = ac.SetNamespaceID(*id)
	}
	return ac
}

// SetNamespace sets the namespace edge to Namespace.
func (ac *ApplicationCreate) SetNamespace(n *Namespace) *ApplicationCreate {
	return ac.SetNamespaceID(n.ID)
}

// SetProjectID sets the project edge to Project by id.
func (ac *ApplicationCreate) SetProjectID(id int) *ApplicationCreate {
	ac.mutation.SetProjectID(id)
	return ac
}

// SetNillableProjectID sets the project edge to Project by id if the given value is not nil.
func (ac *ApplicationCreate) SetNillableProjectID(id *int) *ApplicationCreate {
	if id != nil {
		ac = ac.SetProjectID(*id)
	}
	return ac
}

// SetProject sets the project edge to Project.
func (ac *ApplicationCreate) SetProject(p *Project) *ApplicationCreate {
	return ac.SetProjectID(p.ID)
}

// AddInstanceIDs adds the instances edge to Instance by ids.
func (ac *ApplicationCreate) AddInstanceIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddInstanceIDs(ids...)
	return ac
}

// AddInstances adds the instances edges to Instance.
func (ac *ApplicationCreate) AddInstances(i ...*Instance) *ApplicationCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ac.AddInstanceIDs(ids...)
}

// SetConfigID sets the config edge to HelmConfig by id.
func (ac *ApplicationCreate) SetConfigID(id int) *ApplicationCreate {
	ac.mutation.SetConfigID(id)
	return ac
}

// SetNillableConfigID sets the config edge to HelmConfig by id if the given value is not nil.
func (ac *ApplicationCreate) SetNillableConfigID(id *int) *ApplicationCreate {
	if id != nil {
		ac = ac.SetConfigID(*id)
	}
	return ac
}

// SetConfig sets the config edge to HelmConfig.
func (ac *ApplicationCreate) SetConfig(h *HelmConfig) *ApplicationCreate {
	return ac.SetConfigID(h.ID)
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	var (
		err  error
		node *Application
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ac *ApplicationCreate) defaults() {
	if _, ok := ac.mutation.Multi(); !ok {
		v := application.DefaultMulti
		ac.mutation.SetMulti(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := application.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := application.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApplicationCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := ac.mutation.Multi(); !ok {
		return &ValidationError{Name: "multi", err: errors.New("ent: missing required field \"multi\"")}
	}
	if _, ok := ac.mutation.ProjectId(); !ok {
		return &ValidationError{Name: "projectId", err: errors.New("ent: missing required field \"projectId\"")}
	}
	if _, ok := ac.mutation.NamespaceId(); !ok {
		return &ValidationError{Name: "namespaceId", err: errors.New("ent: missing required field \"namespaceId\"")}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New("ent: missing required field \"createdAt\"")}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		_node = &Application{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: application.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: application.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Multi(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: application.FieldMulti,
		})
		_node.Multi = value
	}
	if value, ok := ac.mutation.ProjectId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: application.FieldProjectId,
		})
		_node.ProjectId = value
	}
	if value, ok := ac.mutation.NamespaceId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: application.FieldNamespaceId,
		})
		_node.NamespaceId = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.NamespaceTable,
			Columns: []string{application.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.ProjectTable,
			Columns: []string{application.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.InstancesTable,
			Columns: []string{application.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   application.ConfigTable,
			Columns: []string{application.ConfigColumn},
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

// ApplicationCreateBulk is the builder for creating a bulk of Application entities.
type ApplicationCreateBulk struct {
	config
	builders []*ApplicationCreate
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
