package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/project"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
)

type ProjectHandler struct {
	m *model.ProjectManage
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{m:model.NewProjectManage()}
}

func (p *ProjectHandler) QueryProject(c *gin.Context) (err error) {
	query := []string{project.FieldProjectName,project.FieldProType}
	qm := model.ParesQuery(c,query)
	count,ps,err := p.m.Query(qm)
	if err != nil {
		return err
	}
	c.Set("data",ps)
	c.Set("total",count)
	c.Set("query",query)
	return nil
}

func (p *ProjectHandler)CreateProject(c *gin.Context)(err error)  {
	var m ent.Project
	if err := c.BindJSON(&m);err!=nil{
		return svcerr.Error(err,svcerr.JSONBindErr)
	}
	var pro *ent.Project
	pro,err = p.m.CreateProject(m)
	if err!=nil{
		return
	}
	c.Set("data",pro)
	return nil
}

func (p *ProjectHandler)DeleteProject(c *gin.Context)(err error)  {
	id := c.Param("id")
	idn,err:= strconv.Atoi(id)
	if err != nil {
		return err
	}
	return p.m.Delete(idn)
}