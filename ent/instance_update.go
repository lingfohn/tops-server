// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/build"
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/instance"
	"github.com/lingfohn/lime/ent/predicate"
)

// InstanceUpdate is the builder for updating Instance entities.
type InstanceUpdate struct {
	config
	hooks    []Hook
	mutation *InstanceMutation
}

// Where adds a new predicate for the builder.
func (iu *InstanceUpdate) Where(ps ...predicate.Instance) *InstanceUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetName sets the name field.
func (iu *InstanceUpdate) SetName(s string) *InstanceUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetApplicationId sets the applicationId field.
func (iu *InstanceUpdate) SetApplicationId(i int) *InstanceUpdate {
	iu.mutation.ResetApplicationId()
	iu.mutation.SetApplicationId(i)
	return iu
}

// AddApplicationId adds i to applicationId.
func (iu *InstanceUpdate) AddApplicationId(i int) *InstanceUpdate {
	iu.mutation.AddApplicationId(i)
	return iu
}

// SetUpdatedAt sets the updatedAt field.
func (iu *InstanceUpdate) SetUpdatedAt(t time.Time) *InstanceUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetApplicationID sets the application edge to Application by id.
func (iu *InstanceUpdate) SetApplicationID(id int) *InstanceUpdate {
	iu.mutation.SetApplicationID(id)
	return iu
}

// SetNillableApplicationID sets the application edge to Application by id if the given value is not nil.
func (iu *InstanceUpdate) SetNillableApplicationID(id *int) *InstanceUpdate {
	if id != nil {
		iu = iu.SetApplicationID(*id)
	}
	return iu
}

// SetApplication sets the application edge to Application.
func (iu *InstanceUpdate) SetApplication(a *Application) *InstanceUpdate {
	return iu.SetApplicationID(a.ID)
}

// AddBuildIDs adds the builds edge to Build by ids.
func (iu *InstanceUpdate) AddBuildIDs(ids ...int) *InstanceUpdate {
	iu.mutation.AddBuildIDs(ids...)
	return iu
}

// AddBuilds adds the builds edges to Build.
func (iu *InstanceUpdate) AddBuilds(b ...*Build) *InstanceUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.AddBuildIDs(ids...)
}

// SetConfigID sets the config edge to HelmConfig by id.
func (iu *InstanceUpdate) SetConfigID(id int) *InstanceUpdate {
	iu.mutation.SetConfigID(id)
	return iu
}

// SetNillableConfigID sets the config edge to HelmConfig by id if the given value is not nil.
func (iu *InstanceUpdate) SetNillableConfigID(id *int) *InstanceUpdate {
	if id != nil {
		iu = iu.SetConfigID(*id)
	}
	return iu
}

// SetConfig sets the config edge to HelmConfig.
func (iu *InstanceUpdate) SetConfig(h *HelmConfig) *InstanceUpdate {
	return iu.SetConfigID(h.ID)
}

// Mutation returns the InstanceMutation object of the builder.
func (iu *InstanceUpdate) Mutation() *InstanceMutation {
	return iu.mutation
}

// ClearApplication clears the "application" edge to type Application.
func (iu *InstanceUpdate) ClearApplication() *InstanceUpdate {
	iu.mutation.ClearApplication()
	return iu
}

// ClearBuilds clears all "builds" edges to type Build.
func (iu *InstanceUpdate) ClearBuilds() *InstanceUpdate {
	iu.mutation.ClearBuilds()
	return iu
}

// RemoveBuildIDs removes the builds edge to Build by ids.
func (iu *InstanceUpdate) RemoveBuildIDs(ids ...int) *InstanceUpdate {
	iu.mutation.RemoveBuildIDs(ids...)
	return iu
}

// RemoveBuilds removes builds edges to Build.
func (iu *InstanceUpdate) RemoveBuilds(b ...*Build) *InstanceUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iu.RemoveBuildIDs(ids...)
}

// ClearConfig clears the "config" edge to type HelmConfig.
func (iu *InstanceUpdate) ClearConfig() *InstanceUpdate {
	iu.mutation.ClearConfig()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InstanceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InstanceUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InstanceUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InstanceUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InstanceUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := instance.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

func (iu *InstanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instance.Table,
			Columns: instance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instance.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldName,
		})
	}
	if value, ok := iu.mutation.ApplicationId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instance.FieldApplicationId,
		})
	}
	if value, ok := iu.mutation.AddedApplicationId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instance.FieldApplicationId,
		})
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldUpdatedAt,
		})
	}
	if iu.mutation.ApplicationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ApplicationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.BuildsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedBuildsIDs(); len(nodes) > 0 && !iu.mutation.BuildsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.BuildsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ConfigCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ConfigIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InstanceUpdateOne is the builder for updating a single Instance entity.
type InstanceUpdateOne struct {
	config
	hooks    []Hook
	mutation *InstanceMutation
}

// SetName sets the name field.
func (iuo *InstanceUpdateOne) SetName(s string) *InstanceUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetApplicationId sets the applicationId field.
func (iuo *InstanceUpdateOne) SetApplicationId(i int) *InstanceUpdateOne {
	iuo.mutation.ResetApplicationId()
	iuo.mutation.SetApplicationId(i)
	return iuo
}

// AddApplicationId adds i to applicationId.
func (iuo *InstanceUpdateOne) AddApplicationId(i int) *InstanceUpdateOne {
	iuo.mutation.AddApplicationId(i)
	return iuo
}

// SetUpdatedAt sets the updatedAt field.
func (iuo *InstanceUpdateOne) SetUpdatedAt(t time.Time) *InstanceUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetApplicationID sets the application edge to Application by id.
func (iuo *InstanceUpdateOne) SetApplicationID(id int) *InstanceUpdateOne {
	iuo.mutation.SetApplicationID(id)
	return iuo
}

// SetNillableApplicationID sets the application edge to Application by id if the given value is not nil.
func (iuo *InstanceUpdateOne) SetNillableApplicationID(id *int) *InstanceUpdateOne {
	if id != nil {
		iuo = iuo.SetApplicationID(*id)
	}
	return iuo
}

// SetApplication sets the application edge to Application.
func (iuo *InstanceUpdateOne) SetApplication(a *Application) *InstanceUpdateOne {
	return iuo.SetApplicationID(a.ID)
}

// AddBuildIDs adds the builds edge to Build by ids.
func (iuo *InstanceUpdateOne) AddBuildIDs(ids ...int) *InstanceUpdateOne {
	iuo.mutation.AddBuildIDs(ids...)
	return iuo
}

// AddBuilds adds the builds edges to Build.
func (iuo *InstanceUpdateOne) AddBuilds(b ...*Build) *InstanceUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.AddBuildIDs(ids...)
}

// SetConfigID sets the config edge to HelmConfig by id.
func (iuo *InstanceUpdateOne) SetConfigID(id int) *InstanceUpdateOne {
	iuo.mutation.SetConfigID(id)
	return iuo
}

// SetNillableConfigID sets the config edge to HelmConfig by id if the given value is not nil.
func (iuo *InstanceUpdateOne) SetNillableConfigID(id *int) *InstanceUpdateOne {
	if id != nil {
		iuo = iuo.SetConfigID(*id)
	}
	return iuo
}

// SetConfig sets the config edge to HelmConfig.
func (iuo *InstanceUpdateOne) SetConfig(h *HelmConfig) *InstanceUpdateOne {
	return iuo.SetConfigID(h.ID)
}

// Mutation returns the InstanceMutation object of the builder.
func (iuo *InstanceUpdateOne) Mutation() *InstanceMutation {
	return iuo.mutation
}

// ClearApplication clears the "application" edge to type Application.
func (iuo *InstanceUpdateOne) ClearApplication() *InstanceUpdateOne {
	iuo.mutation.ClearApplication()
	return iuo
}

// ClearBuilds clears all "builds" edges to type Build.
func (iuo *InstanceUpdateOne) ClearBuilds() *InstanceUpdateOne {
	iuo.mutation.ClearBuilds()
	return iuo
}

// RemoveBuildIDs removes the builds edge to Build by ids.
func (iuo *InstanceUpdateOne) RemoveBuildIDs(ids ...int) *InstanceUpdateOne {
	iuo.mutation.RemoveBuildIDs(ids...)
	return iuo
}

// RemoveBuilds removes builds edges to Build.
func (iuo *InstanceUpdateOne) RemoveBuilds(b ...*Build) *InstanceUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return iuo.RemoveBuildIDs(ids...)
}

// ClearConfig clears the "config" edge to type HelmConfig.
func (iuo *InstanceUpdateOne) ClearConfig() *InstanceUpdateOne {
	iuo.mutation.ClearConfig()
	return iuo
}

// Save executes the query and returns the updated entity.
func (iuo *InstanceUpdateOne) Save(ctx context.Context) (*Instance, error) {
	var (
		err  error
		node *Instance
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InstanceUpdateOne) SaveX(ctx context.Context) *Instance {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InstanceUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InstanceUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InstanceUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := instance.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

func (iuo *InstanceUpdateOne) sqlSave(ctx context.Context) (_node *Instance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instance.Table,
			Columns: instance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instance.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Instance.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldName,
		})
	}
	if value, ok := iuo.mutation.ApplicationId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instance.FieldApplicationId,
		})
	}
	if value, ok := iuo.mutation.AddedApplicationId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: instance.FieldApplicationId,
		})
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldUpdatedAt,
		})
	}
	if iuo.mutation.ApplicationCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ApplicationIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.BuildsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedBuildsIDs(); len(nodes) > 0 && !iuo.mutation.BuildsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.BuildsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ConfigCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ConfigIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Instance{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instance.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
