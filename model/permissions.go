package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/permission"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/storage"
)

type PermissionManager struct {
	db *ent.Client
}

func NewPermissionManager() *PermissionManager {
	return &PermissionManager{db: storage.GetDB()}
}

func (p *PermissionManager) CreatePermission(perm ent.Permission) (permission *ent.Permission, err error) {
	return p.db.Permission.Create().
		SetMethod(perm.Method).
		SetFullpath(perm.Fullpath).
		SetAction(perm.Action).
		SetSummary(perm.Summary).
		SetControlLevel(perm.ControlLevel).
		SetStatus(perm.Status).
		Save(context.TODO())
}

func (p *PermissionManager) predicate(qm QueryMap) []predicate.Permission {
	var query []predicate.Permission
	if val, ok := qm.GetString(permission.FieldMethod); ok {
		query = append(query, permission.ActionEQ(val))
	}
	if val, ok := qm.GetString(permission.FieldFullpath); ok {
		query = append(query, permission.FullpathContains(val))
	}
	if val, ok := qm.GetString(permission.FieldMethod); ok {
		query = append(query, permission.MethodEQ(val))
	}
	if val, ok := qm.GetString(permission.FieldSummary); ok {
		query = append(query, permission.SummaryContains(val))
	}
	if val, ok := qm.GetInt(permission.FieldStatus); ok {
		query = append(query, permission.StatusEQ(val))
	}
	if val, ok := qm.GetInt(permission.FieldControlLevel); ok {
		query = append(query, permission.ControlLevelEQ(val))
	}
	return query
}

func (p *PermissionManager) Query(q QueryMap) (count int, perms ent.Permissions, err error) {
	permQuery := p.db.Permission.
		Query().
		Where(p.predicate(q)...)
	if count, err = permQuery.Count(context.TODO()); err != nil {
		return
	}
	perms, err = permQuery.
		Order(q.Order()...).
		Offset(q.OffSet()).
		Limit(q.Limit()).
		All(context.TODO())
	return
}

func (p *PermissionManager) Delete(id int) (err error) {
	return p.db.Permission.DeleteOneID(id).Exec(context.TODO())
}

func (p *PermissionManager) Update(id int, perm ent.Permission) (permission *ent.Permission, err error) {
	return p.db.Permission.UpdateOneID(id).
		SetAction(perm.Action).
		SetControlLevel(perm.ControlLevel).
		SetFullpath(perm.Fullpath).
		SetMethod(perm.Method).
		SetSummary(perm.Summary).
		SetStatus(perm.Status).
		Save(context.TODO())
}
