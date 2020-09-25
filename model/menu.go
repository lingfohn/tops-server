package model

import (
	"context"

	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/menu"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/storage"
)

type MenuManager struct {
	db *ent.Client
}

func NewMenuManager() *MenuManager {
	return &MenuManager{db: storage.GetDB()}
}

func (p *MenuManager) CreateMenu(m ent.Menu) (menu *ent.Menu, err error) {
	mc := p.db.Menu.Create()
	mc.SetLevel(m.Level).
		SetName(m.Name).
		SetComponent(m.Component).
		SetPath(m.Path).
		SetRedirect(m.Redirect).
		SetTitle(m.Title).
		SetIcon(m.Icon).
		SetWeight(m.Weight).
		SetNillableShow(m.Show).
		SetKeepAlive(m.KeepAlive).
		SetParentId(m.ParentId)
	if m.ParentId != 0 {
		mc.SetParentID(m.ParentId)
	}
	return mc.Save(context.TODO())
}

func (p *MenuManager) QueryAll(q QueryMap) (count int, menus ent.Menus, err error) {
	menuQuery := p.db.Menu.Query()
	menuQuery = menuQuery.Where(p.predicate(q)...)
	if val, ok := q.GetBool("withParent"); ok && val {
		menuQuery = menuQuery.WithParent()
	}
	if val, ok := q.GetBool("withSubmenus"); ok && val {
		menuQuery = menuQuery.WithSubmenus()
	}
	count, err = menuQuery.Count(context.TODO())
	menus, err = menuQuery.Offset(q.OffSet()).Limit(q.Limit()).Order(q.Order()...).All(context.TODO())
	return
}

func (p *MenuManager) predicate(q QueryMap) []predicate.Menu {
	var query []predicate.Menu
	if val, ok := q.GetInt(menu.FieldLevel); ok {
		query = append(query, menu.LevelEQ(val))
	}
	return query
}

func (p *MenuManager) Delete(id int) (err error) {
	return p.db.Menu.DeleteOneID(id).Exec(context.TODO())
}

func (p *MenuManager) Update(id int, m ent.Menu) (memu *ent.Menu, err error) {
	return p.db.Menu.UpdateOneID(id).
		SetName(m.Name).
		SetComponent(m.Component).
		SetPath(m.Path).
		SetRedirect(m.Redirect).
		SetTitle(m.Title).
		SetIcon(m.Icon).
		SetWeight(m.Weight).
		SetNillableShow(m.Show).
		SetKeepAlive(m.KeepAlive).
		SetParentId(m.ParentId).
		Save(context.TODO())

}
