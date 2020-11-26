// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/project"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
}

// SetProjectName sets the projectName field.
func (pc *ProjectCreate) SetProjectName(s string) *ProjectCreate {
	pc.mutation.SetProjectName(s)
	return pc
}

// SetProType sets the proType field.
func (pc *ProjectCreate) SetProType(s string) *ProjectCreate {
	pc.mutation.SetProType(s)
	return pc
}

// SetDescription sets the description field.
func (pc *ProjectCreate) SetDescription(s string) *ProjectCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetGitlab sets the gitlab field.
func (pc *ProjectCreate) SetGitlab(s string) *ProjectCreate {
	pc.mutation.SetGitlab(s)
	return pc
}

// SetPort sets the port field.
func (pc *ProjectCreate) SetPort(i int) *ProjectCreate {
	pc.mutation.SetPort(i)
	return pc
}

// SetDebugPort sets the debugPort field.
func (pc *ProjectCreate) SetDebugPort(i int) *ProjectCreate {
	pc.mutation.SetDebugPort(i)
	return pc
}

// SetCreatedAt sets the createdAt field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updatedAt field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// AddApplicationIDs adds the applications edge to Application by ids.
func (pc *ProjectCreate) AddApplicationIDs(ids ...int) *ProjectCreate {
	pc.mutation.AddApplicationIDs(ids...)
	return pc
}

// AddApplications adds the applications edges to Application.
func (pc *ProjectCreate) AddApplications(a ...*Application) *ProjectCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddApplicationIDs(ids...)
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	if _, ok := pc.mutation.ProjectName(); !ok {
		return nil, errors.New("ent: missing required field \"projectName\"")
	}
	if _, ok := pc.mutation.ProType(); !ok {
		return nil, errors.New("ent: missing required field \"proType\"")
	}
	if _, ok := pc.mutation.Gitlab(); !ok {
		return nil, errors.New("ent: missing required field \"gitlab\"")
	}
	if _, ok := pc.mutation.Port(); !ok {
		return nil, errors.New("ent: missing required field \"port\"")
	}
	if _, ok := pc.mutation.DebugPort(); !ok {
		return nil, errors.New("ent: missing required field \"debugPort\"")
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := project.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	var (
		err  error
		node *Project
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
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
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	var (
		pr    = &Project{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: project.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.ProjectName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldProjectName,
		})
		pr.ProjectName = value
	}
	if value, ok := pc.mutation.ProType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldProType,
		})
		pr.ProType = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldDescription,
		})
		pr.Description = value
	}
	if value, ok := pc.mutation.Gitlab(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldGitlab,
		})
		pr.Gitlab = value
	}
	if value, ok := pc.mutation.Port(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: project.FieldPort,
		})
		pr.Port = value
	}
	if value, ok := pc.mutation.DebugPort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: project.FieldDebugPort,
		})
		pr.DebugPort = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldCreatedAt,
		})
		pr.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
		pr.UpdatedAt = value
	}
	if nodes := pc.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ApplicationsTable,
			Columns: []string{project.ApplicationsColumn},
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
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}
