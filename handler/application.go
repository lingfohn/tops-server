package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
)

type ApplicationHandler struct {
	m *model.ApplicationManage
}

func NewApplicationHandler() *ApplicationHandler {
	return &ApplicationHandler{ m:model.NewApplicationManage()}
}

func (a *ApplicationHandler)CreateApplication(c *gin.Context)(err error)  {
	var m ent.Application
	if err=c.BindJSON(&m);err!=nil{
		return svcerr.Error(err,svcerr.JSONBindErr)
	}
	app,err:=a.m.CreateApplication(m)
	if err != nil {
		return err
	}
	c.Set("data",app)
	return
}

func (a *ApplicationHandler)QueryApplication(c *gin.Context)(err error)  {
	query := []string{application.FieldName,application.FieldMulti,application.FieldNamespaceId}
	qm := model.ParesQuery(c,query)
	count,apps,err:=a.m.QueryApplication(qm)
	if err != nil {
		return err
	}
	c.Set("data",apps)
	c.Set("total",count)
	return
}