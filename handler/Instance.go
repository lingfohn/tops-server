package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/instance"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
	"strings"
)

type InstanceHandler struct {
	m *model.InstanceManage
	g *model.GitlabManage
}

func NewInstanceHandler() *InstanceHandler {
	return &InstanceHandler{ m: model.NewInstanceManage(),g:model.NewGitlabManage() }
}

func (i *InstanceHandler)CreateInstance(c *gin.Context)(err error)  {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var m ent.Instance
	if err= c.BindJSON(&m);err!=nil{
		return svcerr.Error(err,svcerr.JSONBindErr)
	}
	m.ApplicationId=idn
	var ins *ent.Instance
	ins,err = i.m.CreateInstance(m)
	if err != nil {
		return err
	}
	c.Set("data",ins)
	return nil
}

func (i *InstanceHandler)QueryInstance(c *gin.Context)(err error)  {
	id := c.Param("id")
	appid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	query := []string{instance.FieldApplicationId}
	qm := model.ParesQuery(c,query)
	qm.SetInt(instance.FieldApplicationId,appid)
	count,ins,err:=i.m.QueryInstance(qm)
	if err != nil {
		return err
	}
	c.Set("data",ins)
	c.Set("total",count)
	return nil
}

func (i *InstanceHandler)LoadInstanceTags(c *gin.Context)(err error) {
	id := c.Param("id")
	idn,_:= strconv.Atoi(id)
	ins,_ := i.m.GetInstance(idn)
	proj,_ := i.m.GetInstanceProject(idn)
	n:=GetNameWithNamespace(proj.Gitlab)
	ts,err:=i.g.ListProjectTags(n,ins.Name)
	ret:=gin.H{
		"project" : proj.ProjectName,
		"tags": ts,
	}
	c.Set("data",ret)
	c.Set("total",len(ts))
	return nil
}

func GetNameWithNamespace(gitlab string) string{
	return strings.TrimRight(strings.TrimPrefix(gitlab,model.GitLab()),".git")
}