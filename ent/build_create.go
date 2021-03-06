// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/lingfohn/lime/ent/build"
	"github.com/lingfohn/lime/ent/instance"
)

// BuildCreate is the builder for creating a Build entity.
type BuildCreate struct {
	config
	mutation *BuildMutation
	hooks    []Hook
}

// SetTag sets the tag field.
func (bc *BuildCreate) SetTag(s string) *BuildCreate {
	bc.mutation.SetTag(s)
	return bc
}

// SetName sets the name field.
func (bc *BuildCreate) SetName(s string) *BuildCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetStatus sets the status field.
func (bc *BuildCreate) SetStatus(i int) *BuildCreate {
	bc.mutation.SetStatus(i)
	return bc
}

// SetNillableStatus sets the status field if the given value is not nil.
func (bc *BuildCreate) SetNillableStatus(i *int) *BuildCreate {
	if i != nil {
		bc.SetStatus(*i)
	}
	return bc
}

// SetCommitId sets the commitId field.
func (bc *BuildCreate) SetCommitId(s string) *BuildCreate {
	bc.mutation.SetCommitId(s)
	return bc
}

// SetShortId sets the shortId field.
func (bc *BuildCreate) SetShortId(s string) *BuildCreate {
	bc.mutation.SetShortId(s)
	return bc
}

// SetMessage sets the message field.
func (bc *BuildCreate) SetMessage(s string) *BuildCreate {
	bc.mutation.SetMessage(s)
	return bc
}

// SetCommitterName sets the committerName field.
func (bc *BuildCreate) SetCommitterName(s string) *BuildCreate {
	bc.mutation.SetCommitterName(s)
	return bc
}

// SetCommittedData sets the committedData field.
func (bc *BuildCreate) SetCommittedData(t time.Time) *BuildCreate {
	bc.mutation.SetCommittedData(t)
	return bc
}

// SetBuildTime sets the buildTime field.
func (bc *BuildCreate) SetBuildTime(t time.Time) *BuildCreate {
	bc.mutation.SetBuildTime(t)
	return bc
}

// SetJenkinsQueueId sets the jenkinsQueueId field.
func (bc *BuildCreate) SetJenkinsQueueId(i int) *BuildCreate {
	bc.mutation.SetJenkinsQueueId(i)
	return bc
}

// SetJenkinsBuildId sets the jenkinsBuildId field.
func (bc *BuildCreate) SetJenkinsBuildId(i int) *BuildCreate {
	bc.mutation.SetJenkinsBuildId(i)
	return bc
}

// SetNillableJenkinsBuildId sets the jenkinsBuildId field if the given value is not nil.
func (bc *BuildCreate) SetNillableJenkinsBuildId(i *int) *BuildCreate {
	if i != nil {
		bc.SetJenkinsBuildId(*i)
	}
	return bc
}

// SetCreatedAt sets the createdAt field.
func (bc *BuildCreate) SetCreatedAt(t time.Time) *BuildCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (bc *BuildCreate) SetNillableCreatedAt(t *time.Time) *BuildCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the updatedAt field.
func (bc *BuildCreate) SetUpdatedAt(t time.Time) *BuildCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (bc *BuildCreate) SetNillableUpdatedAt(t *time.Time) *BuildCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetInstanceID sets the instance edge to Instance by id.
func (bc *BuildCreate) SetInstanceID(id int) *BuildCreate {
	bc.mutation.SetInstanceID(id)
	return bc
}

// SetNillableInstanceID sets the instance edge to Instance by id if the given value is not nil.
func (bc *BuildCreate) SetNillableInstanceID(id *int) *BuildCreate {
	if id != nil {
		bc = bc.SetInstanceID(*id)
	}
	return bc
}

// SetInstance sets the instance edge to Instance.
func (bc *BuildCreate) SetInstance(i *Instance) *BuildCreate {
	return bc.SetInstanceID(i.ID)
}

// Mutation returns the BuildMutation object of the builder.
func (bc *BuildCreate) Mutation() *BuildMutation {
	return bc.mutation
}

// Save creates the Build in the database.
func (bc *BuildCreate) Save(ctx context.Context) (*Build, error) {
	var (
		err  error
		node *Build
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BuildCreate) SaveX(ctx context.Context) *Build {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bc *BuildCreate) defaults() {
	if _, ok := bc.mutation.Status(); !ok {
		v := build.DefaultStatus
		bc.mutation.SetStatus(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := build.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := build.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BuildCreate) check() error {
	if _, ok := bc.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New("ent: missing required field \"tag\"")}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New("ent: missing required field \"status\"")}
	}
	if _, ok := bc.mutation.CommitId(); !ok {
		return &ValidationError{Name: "commitId", err: errors.New("ent: missing required field \"commitId\"")}
	}
	if _, ok := bc.mutation.ShortId(); !ok {
		return &ValidationError{Name: "shortId", err: errors.New("ent: missing required field \"shortId\"")}
	}
	if _, ok := bc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New("ent: missing required field \"message\"")}
	}
	if _, ok := bc.mutation.CommitterName(); !ok {
		return &ValidationError{Name: "committerName", err: errors.New("ent: missing required field \"committerName\"")}
	}
	if _, ok := bc.mutation.CommittedData(); !ok {
		return &ValidationError{Name: "committedData", err: errors.New("ent: missing required field \"committedData\"")}
	}
	if _, ok := bc.mutation.BuildTime(); !ok {
		return &ValidationError{Name: "buildTime", err: errors.New("ent: missing required field \"buildTime\"")}
	}
	if _, ok := bc.mutation.JenkinsQueueId(); !ok {
		return &ValidationError{Name: "jenkinsQueueId", err: errors.New("ent: missing required field \"jenkinsQueueId\"")}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New("ent: missing required field \"createdAt\"")}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	return nil
}

func (bc *BuildCreate) sqlSave(ctx context.Context) (*Build, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BuildCreate) createSpec() (*Build, *sqlgraph.CreateSpec) {
	var (
		_node = &Build{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: build.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: build.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Tag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldTag,
		})
		_node.Tag = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := bc.mutation.CommitId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldCommitId,
		})
		_node.CommitId = value
	}
	if value, ok := bc.mutation.ShortId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldShortId,
		})
		_node.ShortId = value
	}
	if value, ok := bc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := bc.mutation.CommitterName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: build.FieldCommitterName,
		})
		_node.CommitterName = value
	}
	if value, ok := bc.mutation.CommittedData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: build.FieldCommittedData,
		})
		_node.CommittedData = value
	}
	if value, ok := bc.mutation.BuildTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: build.FieldBuildTime,
		})
		_node.BuildTime = value
	}
	if value, ok := bc.mutation.JenkinsQueueId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldJenkinsQueueId,
		})
		_node.JenkinsQueueId = value
	}
	if value, ok := bc.mutation.JenkinsBuildId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: build.FieldJenkinsBuildId,
		})
		_node.JenkinsBuildId = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: build.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: build.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   build.InstanceTable,
			Columns: []string{build.InstanceColumn},
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
	return _node, _spec
}

// BuildCreateBulk is the builder for creating a bulk of Build entities.
type BuildCreateBulk struct {
	config
	builders []*BuildCreate
}

// Save creates the Build entities in the database.
func (bcb *BuildCreateBulk) Save(ctx context.Context) ([]*Build, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Build, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BuildMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (bcb *BuildCreateBulk) SaveX(ctx context.Context) []*Build {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
