// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/lingfohn/lime/ent/menu"
	"github.com/lingfohn/lime/ent/predicate"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// Where adds a new predicate for the builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetPath sets the path field.
func (mu *MenuUpdate) SetPath(s string) *MenuUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetName sets the name field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetComponent sets the component field.
func (mu *MenuUpdate) SetComponent(s string) *MenuUpdate {
	mu.mutation.SetComponent(s)
	return mu
}

// SetParentId sets the parentId field.
func (mu *MenuUpdate) SetParentId(i int) *MenuUpdate {
	mu.mutation.ResetParentId()
	mu.mutation.SetParentId(i)
	return mu
}

// AddParentId adds i to parentId.
func (mu *MenuUpdate) AddParentId(i int) *MenuUpdate {
	mu.mutation.AddParentId(i)
	return mu
}

// SetRedirect sets the redirect field.
func (mu *MenuUpdate) SetRedirect(s string) *MenuUpdate {
	mu.mutation.SetRedirect(s)
	return mu
}

// SetNillableRedirect sets the redirect field if the given value is not nil.
func (mu *MenuUpdate) SetNillableRedirect(s *string) *MenuUpdate {
	if s != nil {
		mu.SetRedirect(*s)
	}
	return mu
}

// ClearRedirect clears the value of redirect.
func (mu *MenuUpdate) ClearRedirect() *MenuUpdate {
	mu.mutation.ClearRedirect()
	return mu
}

// SetWeight sets the weight field.
func (mu *MenuUpdate) SetWeight(i int) *MenuUpdate {
	mu.mutation.ResetWeight()
	mu.mutation.SetWeight(i)
	return mu
}

// SetNillableWeight sets the weight field if the given value is not nil.
func (mu *MenuUpdate) SetNillableWeight(i *int) *MenuUpdate {
	if i != nil {
		mu.SetWeight(*i)
	}
	return mu
}

// AddWeight adds i to weight.
func (mu *MenuUpdate) AddWeight(i int) *MenuUpdate {
	mu.mutation.AddWeight(i)
	return mu
}

// ClearWeight clears the value of weight.
func (mu *MenuUpdate) ClearWeight() *MenuUpdate {
	mu.mutation.ClearWeight()
	return mu
}

// SetLevel sets the level field.
func (mu *MenuUpdate) SetLevel(i int) *MenuUpdate {
	mu.mutation.ResetLevel()
	mu.mutation.SetLevel(i)
	return mu
}

// AddLevel adds i to level.
func (mu *MenuUpdate) AddLevel(i int) *MenuUpdate {
	mu.mutation.AddLevel(i)
	return mu
}

// SetTitle sets the title field.
func (mu *MenuUpdate) SetTitle(s string) *MenuUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetNillableTitle sets the title field if the given value is not nil.
func (mu *MenuUpdate) SetNillableTitle(s *string) *MenuUpdate {
	if s != nil {
		mu.SetTitle(*s)
	}
	return mu
}

// ClearTitle clears the value of title.
func (mu *MenuUpdate) ClearTitle() *MenuUpdate {
	mu.mutation.ClearTitle()
	return mu
}

// SetIcon sets the icon field.
func (mu *MenuUpdate) SetIcon(s string) *MenuUpdate {
	mu.mutation.SetIcon(s)
	return mu
}

// SetNillableIcon sets the icon field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIcon(s *string) *MenuUpdate {
	if s != nil {
		mu.SetIcon(*s)
	}
	return mu
}

// ClearIcon clears the value of icon.
func (mu *MenuUpdate) ClearIcon() *MenuUpdate {
	mu.mutation.ClearIcon()
	return mu
}

// SetTarget sets the target field.
func (mu *MenuUpdate) SetTarget(s string) *MenuUpdate {
	mu.mutation.SetTarget(s)
	return mu
}

// SetNillableTarget sets the target field if the given value is not nil.
func (mu *MenuUpdate) SetNillableTarget(s *string) *MenuUpdate {
	if s != nil {
		mu.SetTarget(*s)
	}
	return mu
}

// ClearTarget clears the value of target.
func (mu *MenuUpdate) ClearTarget() *MenuUpdate {
	mu.mutation.ClearTarget()
	return mu
}

// SetKeepAlive sets the keepAlive field.
func (mu *MenuUpdate) SetKeepAlive(b bool) *MenuUpdate {
	mu.mutation.SetKeepAlive(b)
	return mu
}

// SetShow sets the show field.
func (mu *MenuUpdate) SetShow(b bool) *MenuUpdate {
	mu.mutation.SetShow(b)
	return mu
}

// SetNillableShow sets the show field if the given value is not nil.
func (mu *MenuUpdate) SetNillableShow(b *bool) *MenuUpdate {
	if b != nil {
		mu.SetShow(*b)
	}
	return mu
}

// ClearShow clears the value of show.
func (mu *MenuUpdate) ClearShow() *MenuUpdate {
	mu.mutation.ClearShow()
	return mu
}

// SetUpdatedAt sets the updatedAt field.
func (mu *MenuUpdate) SetUpdatedAt(t time.Time) *MenuUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetParentID sets the parent edge to Menu by id.
func (mu *MenuUpdate) SetParentID(id int) *MenuUpdate {
	mu.mutation.SetParentID(id)
	return mu
}

// SetNillableParentID sets the parent edge to Menu by id if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(id *int) *MenuUpdate {
	if id != nil {
		mu = mu.SetParentID(*id)
	}
	return mu
}

// SetParent sets the parent edge to Menu.
func (mu *MenuUpdate) SetParent(m *Menu) *MenuUpdate {
	return mu.SetParentID(m.ID)
}

// AddSubmenuIDs adds the submenus edge to Menu by ids.
func (mu *MenuUpdate) AddSubmenuIDs(ids ...int) *MenuUpdate {
	mu.mutation.AddSubmenuIDs(ids...)
	return mu
}

// AddSubmenus adds the submenus edges to Menu.
func (mu *MenuUpdate) AddSubmenus(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddSubmenuIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// ClearParent clears the "parent" edge to type Menu.
func (mu *MenuUpdate) ClearParent() *MenuUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearSubmenus clears all "submenus" edges to type Menu.
func (mu *MenuUpdate) ClearSubmenus() *MenuUpdate {
	mu.mutation.ClearSubmenus()
	return mu
}

// RemoveSubmenuIDs removes the submenus edge to Menu by ids.
func (mu *MenuUpdate) RemoveSubmenuIDs(ids ...int) *MenuUpdate {
	mu.mutation.RemoveSubmenuIDs(ids...)
	return mu
}

// RemoveSubmenus removes submenus edges to Menu.
func (mu *MenuUpdate) RemoveSubmenus(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveSubmenuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MenuUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldPath,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if value, ok := mu.mutation.Component(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldComponent,
		})
	}
	if value, ok := mu.mutation.ParentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldParentId,
		})
	}
	if value, ok := mu.mutation.AddedParentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldParentId,
		})
	}
	if value, ok := mu.mutation.Redirect(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldRedirect,
		})
	}
	if mu.mutation.RedirectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldRedirect,
		})
	}
	if value, ok := mu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldWeight,
		})
	}
	if value, ok := mu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldWeight,
		})
	}
	if mu.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: menu.FieldWeight,
		})
	}
	if value, ok := mu.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldLevel,
		})
	}
	if value, ok := mu.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldLevel,
		})
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTitle,
		})
	}
	if mu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTitle,
		})
	}
	if value, ok := mu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if mu.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := mu.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTarget,
		})
	}
	if mu.mutation.TargetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTarget,
		})
	}
	if value, ok := mu.mutation.KeepAlive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldKeepAlive,
		})
	}
	if value, ok := mu.mutation.Show(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldShow,
		})
	}
	if mu.mutation.ShowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: menu.FieldShow,
		})
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldUpdatedAt,
		})
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.SubmenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedSubmenusIDs(); len(nodes) > 0 && !mu.mutation.SubmenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.SubmenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// SetPath sets the path field.
func (muo *MenuUpdateOne) SetPath(s string) *MenuUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetName sets the name field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetComponent sets the component field.
func (muo *MenuUpdateOne) SetComponent(s string) *MenuUpdateOne {
	muo.mutation.SetComponent(s)
	return muo
}

// SetParentId sets the parentId field.
func (muo *MenuUpdateOne) SetParentId(i int) *MenuUpdateOne {
	muo.mutation.ResetParentId()
	muo.mutation.SetParentId(i)
	return muo
}

// AddParentId adds i to parentId.
func (muo *MenuUpdateOne) AddParentId(i int) *MenuUpdateOne {
	muo.mutation.AddParentId(i)
	return muo
}

// SetRedirect sets the redirect field.
func (muo *MenuUpdateOne) SetRedirect(s string) *MenuUpdateOne {
	muo.mutation.SetRedirect(s)
	return muo
}

// SetNillableRedirect sets the redirect field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableRedirect(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetRedirect(*s)
	}
	return muo
}

// ClearRedirect clears the value of redirect.
func (muo *MenuUpdateOne) ClearRedirect() *MenuUpdateOne {
	muo.mutation.ClearRedirect()
	return muo
}

// SetWeight sets the weight field.
func (muo *MenuUpdateOne) SetWeight(i int) *MenuUpdateOne {
	muo.mutation.ResetWeight()
	muo.mutation.SetWeight(i)
	return muo
}

// SetNillableWeight sets the weight field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableWeight(i *int) *MenuUpdateOne {
	if i != nil {
		muo.SetWeight(*i)
	}
	return muo
}

// AddWeight adds i to weight.
func (muo *MenuUpdateOne) AddWeight(i int) *MenuUpdateOne {
	muo.mutation.AddWeight(i)
	return muo
}

// ClearWeight clears the value of weight.
func (muo *MenuUpdateOne) ClearWeight() *MenuUpdateOne {
	muo.mutation.ClearWeight()
	return muo
}

// SetLevel sets the level field.
func (muo *MenuUpdateOne) SetLevel(i int) *MenuUpdateOne {
	muo.mutation.ResetLevel()
	muo.mutation.SetLevel(i)
	return muo
}

// AddLevel adds i to level.
func (muo *MenuUpdateOne) AddLevel(i int) *MenuUpdateOne {
	muo.mutation.AddLevel(i)
	return muo
}

// SetTitle sets the title field.
func (muo *MenuUpdateOne) SetTitle(s string) *MenuUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetNillableTitle sets the title field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableTitle(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetTitle(*s)
	}
	return muo
}

// ClearTitle clears the value of title.
func (muo *MenuUpdateOne) ClearTitle() *MenuUpdateOne {
	muo.mutation.ClearTitle()
	return muo
}

// SetIcon sets the icon field.
func (muo *MenuUpdateOne) SetIcon(s string) *MenuUpdateOne {
	muo.mutation.SetIcon(s)
	return muo
}

// SetNillableIcon sets the icon field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIcon(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetIcon(*s)
	}
	return muo
}

// ClearIcon clears the value of icon.
func (muo *MenuUpdateOne) ClearIcon() *MenuUpdateOne {
	muo.mutation.ClearIcon()
	return muo
}

// SetTarget sets the target field.
func (muo *MenuUpdateOne) SetTarget(s string) *MenuUpdateOne {
	muo.mutation.SetTarget(s)
	return muo
}

// SetNillableTarget sets the target field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableTarget(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetTarget(*s)
	}
	return muo
}

// ClearTarget clears the value of target.
func (muo *MenuUpdateOne) ClearTarget() *MenuUpdateOne {
	muo.mutation.ClearTarget()
	return muo
}

// SetKeepAlive sets the keepAlive field.
func (muo *MenuUpdateOne) SetKeepAlive(b bool) *MenuUpdateOne {
	muo.mutation.SetKeepAlive(b)
	return muo
}

// SetShow sets the show field.
func (muo *MenuUpdateOne) SetShow(b bool) *MenuUpdateOne {
	muo.mutation.SetShow(b)
	return muo
}

// SetNillableShow sets the show field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableShow(b *bool) *MenuUpdateOne {
	if b != nil {
		muo.SetShow(*b)
	}
	return muo
}

// ClearShow clears the value of show.
func (muo *MenuUpdateOne) ClearShow() *MenuUpdateOne {
	muo.mutation.ClearShow()
	return muo
}

// SetUpdatedAt sets the updatedAt field.
func (muo *MenuUpdateOne) SetUpdatedAt(t time.Time) *MenuUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetParentID sets the parent edge to Menu by id.
func (muo *MenuUpdateOne) SetParentID(id int) *MenuUpdateOne {
	muo.mutation.SetParentID(id)
	return muo
}

// SetNillableParentID sets the parent edge to Menu by id if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(id *int) *MenuUpdateOne {
	if id != nil {
		muo = muo.SetParentID(*id)
	}
	return muo
}

// SetParent sets the parent edge to Menu.
func (muo *MenuUpdateOne) SetParent(m *Menu) *MenuUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddSubmenuIDs adds the submenus edge to Menu by ids.
func (muo *MenuUpdateOne) AddSubmenuIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.AddSubmenuIDs(ids...)
	return muo
}

// AddSubmenus adds the submenus edges to Menu.
func (muo *MenuUpdateOne) AddSubmenus(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddSubmenuIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// ClearParent clears the "parent" edge to type Menu.
func (muo *MenuUpdateOne) ClearParent() *MenuUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearSubmenus clears all "submenus" edges to type Menu.
func (muo *MenuUpdateOne) ClearSubmenus() *MenuUpdateOne {
	muo.mutation.ClearSubmenus()
	return muo
}

// RemoveSubmenuIDs removes the submenus edge to Menu by ids.
func (muo *MenuUpdateOne) RemoveSubmenuIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.RemoveSubmenuIDs(ids...)
	return muo
}

// RemoveSubmenus removes submenus edges to Menu.
func (muo *MenuUpdateOne) RemoveSubmenus(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveSubmenuIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	var (
		err  error
		node *Menu
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MenuUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Menu.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldPath,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if value, ok := muo.mutation.Component(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldComponent,
		})
	}
	if value, ok := muo.mutation.ParentId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldParentId,
		})
	}
	if value, ok := muo.mutation.AddedParentId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldParentId,
		})
	}
	if value, ok := muo.mutation.Redirect(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldRedirect,
		})
	}
	if muo.mutation.RedirectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldRedirect,
		})
	}
	if value, ok := muo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldWeight,
		})
	}
	if value, ok := muo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldWeight,
		})
	}
	if muo.mutation.WeightCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: menu.FieldWeight,
		})
	}
	if value, ok := muo.mutation.Level(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldLevel,
		})
	}
	if value, ok := muo.mutation.AddedLevel(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldLevel,
		})
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTitle,
		})
	}
	if muo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTitle,
		})
	}
	if value, ok := muo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
	}
	if muo.mutation.IconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldIcon,
		})
	}
	if value, ok := muo.mutation.Target(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTarget,
		})
	}
	if muo.mutation.TargetCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: menu.FieldTarget,
		})
	}
	if value, ok := muo.mutation.KeepAlive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldKeepAlive,
		})
	}
	if value, ok := muo.mutation.Show(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldShow,
		})
	}
	if muo.mutation.ShowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: menu.FieldShow,
		})
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldUpdatedAt,
		})
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.SubmenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedSubmenusIDs(); len(nodes) > 0 && !muo.mutation.SubmenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.SubmenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.SubmenusTable,
			Columns: []string{menu.SubmenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
