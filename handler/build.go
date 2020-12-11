package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
	"time"
)

type BuildHandler struct {
	m *model.BuildManage
	i *model.InstanceManage
	jenkins *model.JenkinsManage
	git *model.GitlabManage
	bs *model.BuildSession
}

func NewBuildHandler() *BuildHandler {
	return &BuildHandler{m : model.NewBuildManage(),
		i: model.NewInstanceManage(),
		jenkins: model.NewJenkinsManage(),
		git: model.NewGitlabManage(),
		bs: model.GetBuildSession(),
	}
}

func (b *BuildHandler) CreateBuild(c *gin.Context) (err error)  {
	id := c.Param("id")
	idn,_:= strconv.Atoi(id)
	var t model.Tag
	if err=c.BindJSON(&t);err!=nil{
		return svcerr.Error(err,svcerr.JSONBindErr)
	}
	proj,_:=b.i.GetInstanceProject(idn)
	namespace,_ := b.i.GetInstanceNamespace(idn)
	application,_ := b.i.GetInstanceApplication(idn)
	pid := GetNameWithNamespace(proj.Gitlab)
	tag,_:=b.git.GetProjectTag(pid,t.Name)
	b.jenkins.GetJob(application.Name)
	buildTime := time.Now()
	params := map[string]string{
		"docker_repo": namespace.DockerRepo,
		"namespace": namespace.RepoNamespace,
		"image_tag": generateImageTag(buildTime,tag.Name),
		"service_name": application.Name,
		"gitlab": proj.Gitlab,
		"tag": tag.Name,
	}
	qid,err:=b.jenkins.CreateJobBuild(application.Name,params)
	if err != nil {
		return
	}
	bm:=ent.Build{
		Tag: tag.Name,
		Name: tag.Name,
		CommitId: tag.Commit.ID,
		ShortId: tag.Commit.ShortID,
		Message: tag.Commit.Title,
		BuildTime: buildTime,
		Status: 0,
		JenkinsQueueId: qid,
		CommitterName: tag.Commit.CommitterName,
		CommittedData: *tag.Commit.CommittedDate,
	}
	build,err :=b.m.CreateBuild(idn,bm)
	if err != nil {
		return
	}
	go b.bs.CheckBuildId(application.Name,build.ID,build.JenkinsQueueId)
	c.Set("data",build)
	// 创建Build 数据
	// 创建Build Job
	// 查询Git Tag的commit历史
	// 开启Build
	// 循环查询Build的状态
	// 创建Instance
	// 启动Instance
	return nil
}

func generateImageTag(bt time.Time,t string)(imt string)  {
	bts:=bt.Format("20060102T150405")
	return fmt.Sprintf("%s-%s",bts,t)
}