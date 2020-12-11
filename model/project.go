package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/ent/project"
	"github.com/lingfohn/lime/storage"
)

type ProjectManage struct {
	db *ent.Client
}

func NewProjectManage() *ProjectManage{
	return &ProjectManage{ db: storage.GetDB()}
}

func (p *ProjectManage) predicate(qm QueryMap) []predicate.Project {
	var query []predicate.Project
	if val,ok :=qm.GetString(project.FieldProjectName);ok {
		query = append(query,project.ProjectNameContains(val))
	}
	if val,ok :=qm.GetString(project.FieldProType);ok {
		query = append(query,project.ProTypeEQ(val))
	}
	return query
}

func (p *ProjectManage) Query(q QueryMap)(count int,ps ent.Projects,err error)  {
	proQuery := p.db.Project.Query().Where(p.predicate(q)...)
	if count,err = proQuery.Count(context.TODO());err!=nil{
		return
	}
	ps,err = proQuery.
		Order(q.Order()...).
		Offset(q.OffSet()).
		Limit(q.Limit()).
		All(context.TODO())
	return
}

func (p *ProjectManage)Delete(id int)(err error)  {
	return p.db.Project.DeleteOneID(id).Exec(context.TODO())
}

func (p *ProjectManage) CreateProject(pt ent.Project) (pro *ent.Project,err error) {
	pc := p.db.Project.Create()
	pc.SetProjectName(pt.ProjectName).
		SetGitlab(pt.Gitlab).
		SetProType(pt.ProType).
		SetPort(pt.Port).
		SetDebugPort(pt.DebugPort)
	return  pc.Save(context.TODO())
}