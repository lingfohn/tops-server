package model

import (
	"github.com/lingfohn/lime/ent"
	"github.com/xanzy/go-gitlab"
	"strings"
)

var (
	token = "RA4oiamGgh1xz-v5pt9T"
	url = "https://dev-gitlab.iuctrip.com/"
)

type GitlabManage struct {
	client *gitlab.Client
}

func NewGitlabManage() *GitlabManage {
	client, _ := gitlab.NewClient(token,gitlab.WithBaseURL(url))
	gm := &GitlabManage{}
	gm.client = client
	return gm
}

type Tag struct {
	Name string `json:"tag"`
}


func (g *GitlabManage)ListProjectTags(pid string,search string)(ts []Tag,err error)  {
	tags,_,err:=g.client.Tags.ListTags(pid,&gitlab.ListTagsOptions{
		Search: &search,
	})
	if err != nil {
		return
	}
	for _, i2 := range tags {
		ts= append(ts,Tag{
			Name: i2.Name,
		})
	}
	return
}

func (g *GitlabManage)GetProjectTag(pid string,tag string) (t *gitlab.Tag,err error) {
	t,_,err =g.client.Tags.GetTag(pid,tag)
	return
}

func GitLab()  string{
	return url
}

func (g *GitlabManage) GetProjectNameWithNamespace(p ent.Project) string  {
	return strings.TrimRight(strings.TrimPrefix(p.Gitlab,GitLab()),".git")
}