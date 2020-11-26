package model

import (
	"github.com/lingfohn/lime/ent"
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
	return
}

func (i *InstanceManage)QueryInstance(qm QueryMap)(count int,ins *ent.Instances,err error)  {
	return
}