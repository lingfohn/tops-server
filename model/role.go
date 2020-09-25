package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/ent/role"
	"github.com/lingfohn/lime/storage"
)

type RoleManager struct {
	db *ent.Client
}

func NewRoleManager() *RoleManager {
	return &RoleManager{db: storage.GetDB()}
}

func (r *RoleManager) CreateRole(ro *ent.Role) (role *ent.Role, err error) {
	return r.db.Role.Create().SetName(ro.Name).SetDescription(ro.Description).SetStatus(ro.Status).Save(context.TODO())
}

func (r *RoleManager) Query(q QueryMap) (count int, roles ent.Roles, err error) {
	roleQuery := r.db.Role.Query().Where(r.predicate(q)...)
	if count, err = roleQuery.Count(context.TODO()); err != nil {
		return
	}
	roles, err = roleQuery.
		Order(q.Order()...).
		Offset(q.OffSet()).
		Limit(q.Limit()).
		All(context.TODO())
	return
}

func (r *RoleManager) predicate(q QueryMap) []predicate.Role {
	var query []predicate.Role
	if val, ok := q.GetString(role.FieldName); ok {
		query = append(query, role.NameContains(val))
	}
	if val, ok := q.GetString(role.FieldDescription); ok {
		query = append(query, role.DescriptionContains(val))
	}
	if val, ok := q.GetInt(role.FieldStatus); ok {
		query = append(query, role.StatusEQ(val))
	}
	return query
}

func (r *RoleManager) Delete(id int) (err error) {
	return r.db.Role.DeleteOneID(id).Exec(context.TODO())
}

func (r *RoleManager) Update(id int, ro ent.Role) (role *ent.Role, err error) {
	return r.db.Role.UpdateOneID(id).
		SetStatus(ro.Status).
		SetName(ro.Name).
		SetDescription(ro.Description).
		Save(context.TODO())
}
