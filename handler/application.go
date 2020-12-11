package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
)

type ApplicationHandler struct {
	m *model.ApplicationManage
	jenkins *model.JenkinsManage
}

func NewApplicationHandler() *ApplicationHandler {
	return &ApplicationHandler{
		m:model.NewApplicationManage(),
		jenkins: model.NewJenkinsManage(),
	}
}

func (a *ApplicationHandler)CreateApplication(c *gin.Context)(err error)  {
	var m ent.Application
	if err=c.BindJSON(&m);err!=nil{
		return svcerr.Error(err,svcerr.JSONBindErr)
	}
	app,err:=a.m.CreateApplication(m)
	if err != nil {
		return
	}
	if err =a.jenkins.GetJob(app.Name);err!=nil{
		// 测试用
		err =a.jenkins.CreateDevServiceJob(app.Name)
		if err != nil {
			return
		}
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

func (a *ApplicationHandler)DeleteApplication(c *gin.Context)(err error)  {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return a.m.Delete(idn)
}

func (a *ApplicationHandler)GetApplication(c *gin.Context)(err error)  {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	app,err:=a.m.GetApplication(idn)
	if err != nil {
		return err
	}
	c.Set("data",app)
	return
}