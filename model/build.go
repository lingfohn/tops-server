package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/storage"
)

type BuildManage struct {
	db *ent.Client
}

func NewBuildManage() *BuildManage {
	return &BuildManage{ db:storage.GetDB() }
}

func (b *BuildManage)CreateBuild(iid int,m ent.Build) (build *ent.Build,err error) {
	build,err=b.db.Build.
		Create().
		SetInstanceID(iid).
		SetTag(m.Tag).
		SetName(m.Name).
		SetCommitId(m.CommitId).
		SetCommittedData(m.CommittedData).
		SetShortId(m.ShortId).
		SetMessage(m.Message).
		SetCommitterName(m.CommitterName).
		SetJenkinsQueueId(m.JenkinsQueueId).
		SetBuildTime(m.BuildTime).
		SetJenkinsBuildId(m.JenkinsBuildId).
		Save(context.TODO())

	return
}

func (b *BuildManage)UpdateJenkinsBuildId(id int,status int,bid int)(err error)  {
	return b.db.Build.
		UpdateOneID(id).
		SetStatus(status).
		SetJenkinsBuildId(bid).
		Exec(context.TODO())
}