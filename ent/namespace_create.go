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
	"github.com/lingfohn/lime/ent/k8scluster"
	"github.com/lingfohn/lime/ent/namespace"
)

// NamespaceCreate is the builder for creating a Namespace entity.
type NamespaceCreate struct {
	config
	mutation *NamespaceMutation
	hooks    []Hook
}

// SetName sets the name field.
func (nc *NamespaceCreate) SetName(s string) *NamespaceCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetDockerRepo sets the dockerRepo field.
func (nc *NamespaceCreate) SetDockerRepo(s string) *NamespaceCreate {
	nc.mutation.SetDockerRepo(s)
	return nc
}

// SetRepoNamespace sets the repoNamespace field.
func (nc *NamespaceCreate) SetRepoNamespace(s string) *NamespaceCreate {
	nc.mutation.SetRepoNamespace(s)
	return nc
}

// SetActive sets the active field.
func (nc *NamespaceCreate) SetActive(s string) *NamespaceCreate {
	nc.mutation.SetActive(s)
	return nc
}

// SetCreatedAt sets the createdAt field.
func (nc *NamespaceCreate) SetCreatedAt(t time.Time) *NamespaceCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableCreatedAt(t *time.Time) *NamespaceCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the updatedAt field.
func (nc *NamespaceCreate) SetUpdatedAt(t time.Time) *NamespaceCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (nc *NamespaceCreate) SetNillableUpdatedAt(t *time.Time) *NamespaceCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetClusterID sets the cluster edge to K8sCluster by id.
func (nc *NamespaceCreate) SetClusterID(id int) *NamespaceCreate {
	nc.mutation.SetClusterID(id)
	return nc
}

// SetNillableClusterID sets the cluster edge to K8sCluster by id if the given value is not nil.
func (nc *NamespaceCreate) SetNillableClusterID(id *int) *NamespaceCreate {
	if id != nil {
		nc = nc.SetClusterID(*id)
	}
	return nc
}

// SetCluster sets the cluster edge to K8sCluster.
func (nc *NamespaceCreate) SetCluster(k *K8sCluster) *NamespaceCreate {
	return nc.SetClusterID(k.ID)
}

// AddApplicationIDs adds the applications edge to Application by ids.
func (nc *NamespaceCreate) AddApplicationIDs(ids ...int) *NamespaceCreate {
	nc.mutation.AddApplicationIDs(ids...)
	return nc
}

// AddApplications adds the applications edges to Application.
func (nc *NamespaceCreate) AddApplications(a ...*Application) *NamespaceCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return nc.AddApplicationIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nc *NamespaceCreate) Mutation() *NamespaceMutation {
	return nc.mutation
}

// Save creates the Namespace in the database.
func (nc *NamespaceCreate) Save(ctx context.Context) (*Namespace, error) {
	var (
		err  error
		node *Namespace
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NamespaceCreate) SaveX(ctx context.Context) *Namespace {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (nc *NamespaceCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := namespace.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := namespace.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NamespaceCreate) check() error {
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := nc.mutation.DockerRepo(); !ok {
		return &ValidationError{Name: "dockerRepo", err: errors.New("ent: missing required field \"dockerRepo\"")}
	}
	if _, ok := nc.mutation.RepoNamespace(); !ok {
		return &ValidationError{Name: "repoNamespace", err: errors.New("ent: missing required field \"repoNamespace\"")}
	}
	if _, ok := nc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New("ent: missing required field \"active\"")}
	}
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New("ent: missing required field \"createdAt\"")}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (nc *NamespaceCreate) sqlSave(ctx context.Context) (*Namespace, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NamespaceCreate) createSpec() (*Namespace, *sqlgraph.CreateSpec) {
	var (
		_node = &Namespace{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: namespace.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespace.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespace.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nc.mutation.DockerRepo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespace.FieldDockerRepo,
		})
		_node.DockerRepo = value
	}
	if value, ok := nc.mutation.RepoNamespace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespace.FieldRepoNamespace,
		})
		_node.RepoNamespace = value
	}
	if value, ok := nc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespace.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: namespace.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: namespace.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := nc.mutation.ClusterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   namespace.ClusterTable,
			Columns: []string{namespace.ClusterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: k8scluster.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ApplicationsTable,
			Columns: []string{namespace.ApplicationsColumn},
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
	return _node, _spec
}

// NamespaceCreateBulk is the builder for creating a bulk of Namespace entities.
type NamespaceCreateBulk struct {
	config
	builders []*NamespaceCreate
}

// Save creates the Namespace entities in the database.
func (ncb *NamespaceCreateBulk) Save(ctx context.Context) ([]*Namespace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Namespace, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NamespaceMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ncb *NamespaceCreateBulk) SaveX(ctx context.Context) []*Namespace {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
