package model

import (
	"time"
)

var session *BuildSession


const (
	QUEUE=1
	RUNNING=2
	BUILD_SUCCESS=3
	BUILD_FAILD=4
)
type BuildSession struct {
	b *BuildManage
	jenkins *JenkinsManage
}

func InitBuildSession()  {
	session = &BuildSession{
		b: NewBuildManage(),
		jenkins: NewJenkinsManage(),
	}
}

func GetBuildSession()  *BuildSession{
	return session
}

func (bs *BuildSession)CheckBuildId(jobName string,id int,qid int)  {
	bs.b.UpdateJenkinsBuildId(id,QUEUE,0)
	bid:=bs.jenkins.GetJobBuildId(qid)
	// 查询状态
	for i:=0;i<1000;i++ {
		running,_:=bs.jenkins.GetJobBuildStatus(jobName,bid)
		if i==0 && running {
			bs.b.UpdateJenkinsBuildId(id,RUNNING,bid)
		}
		if !running {
			break
		}
		time.Sleep(2*time.Second)
	}
	_,good:= bs.jenkins.GetJobBuildStatus(jobName,bid)
	if good {
		bs.b.UpdateJenkinsBuildId(id,BUILD_SUCCESS,bid)
	} else {
		bs.b.UpdateJenkinsBuildId(id,BUILD_FAILD,bid)
		return
	}
}