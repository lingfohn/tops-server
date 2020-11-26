package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/storage"
)

type ApplicationManage struct {
	db *ent.Client
}

func NewApplicationManage() *ApplicationManage{
	return &ApplicationManage{
		db: storage.GetDB(),
	}
}

func (a *ApplicationManage)CreateApplication(m ent.Application)(app *ent.Application,err error)  {
	ns,err:=a.db.Namespace.Get(context.TODO(),m.NamespaceId)
	if err != nil {
		return
	}
	pro,err:=a.db.Project.Get(context.TODO(),m.ProjectId)
	if err != nil {
		return
	}
	return a.db.Application.
		Create().
		SetName(pro.ProjectName).
		SetNamespaceId(m.NamespaceId).
		SetNamespace(ns).
		SetProjectId(m.ProjectId).
		SetProject(pro).
		SetNillableMulti(&m.Multi).
		Save(context.TODO())
}

func (a *ApplicationManage)predicate(q QueryMap) []predicate.Application  {
	var query []predicate.Application
	return query
}

func (a *ApplicationManage)QueryApplication(qm QueryMap)(ct int,apps ent.Applications,err error){
	aq :=a.db.Application.Query()
	aq = aq.Where(a.predicate(qm)...)
	if val,ok := qm.GetBool("withInstance");ok && val{
		aq = aq.WithInstances()
	}
	if val,ok := qm.GetBool("withProject");ok && val {
		aq = aq.WithProject()
	}
	if val,ok := qm.GetBool("withNamespace");ok && val {
		aq = aq.WithNamespace()
	}
	ct,err = aq.Count(context.TODO())
	apps,err = aq.Offset(qm.OffSet()).
		Limit(qm.Limit()).
		Order(qm.Order()...).
		All(context.TODO())
	return
}