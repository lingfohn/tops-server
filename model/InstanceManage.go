package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/instance"
	"github.com/lingfohn/lime/ent/namespace"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/ent/project"
	"github.com/lingfohn/lime/storage"
)

type InstanceManage struct {
	db *ent.Client
}

func NewInstanceManage() *InstanceManage  {
	return &InstanceManage{
		db: storage.GetDB(),
	}
}

func (i *InstanceManage)CreateInstance(m ent.Instance)(ins *ent.Instance,err error)  {
	return i.db.Instance.
		Create().
		SetName(m.Name).
		SetApplicationId(m.ApplicationId).
		SetApplicationID(m.ApplicationId).
		Save(context.TODO())
}

func (i *InstanceManage) predicate(qm QueryMap) []predicate.Instance  {
	var query []predicate.Instance
	if val,ok := qm.GetInt(instance.FieldApplicationId);ok {
		query = append(query,instance.ApplicationIdEQ(val))
	}
	return query
}

func (i *InstanceManage)QueryInstance(qm QueryMap)(count int,ins ent.Instances,err error)  {
	insQuery := i.db.Instance.Query().Where(i.predicate(qm)...)

	if count, err = insQuery.Count(context.TODO()); err != nil {
		return
	}
	ins,err = insQuery.
		Order(qm.Order()...).
		Offset(qm.OffSet()).
		Limit(qm.Limit()).
		All(context.TODO())
	return
}

func (i *InstanceManage)GetInstance(id int)(ins *ent.Instance,err error)  {
	return i.db.Instance.Get(context.TODO(),id)
}

func (i *InstanceManage)GetInstanceProject(id int)(ins *ent.Project,err error)  {
	return i.db.Project.
		Query().
		Where(project.HasApplicationsWith(application.HasInstancesWith(instance.ID(id)))).
		Only(context.TODO())
}

func (i *InstanceManage)GetInstanceNamespace(id int)(ns *ent.Namespace,err error)  {
	return i.db.Namespace.
		Query().
		Where(namespace.HasApplicationsWith(application.HasInstancesWith(instance.ID(id)))).
		Only(context.TODO())
}

func (i *InstanceManage)GetInstanceApplication(id int)(app *ent.Application,err error)  {
	return i.db.Application.
		Query().
		Where(application.HasInstancesWith(instance.ID(id))).
		Only(context.TODO())
}