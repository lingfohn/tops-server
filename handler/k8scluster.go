package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/k8scluster"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
)

type K8sClusterHandler struct {
	m *model.K8sClusterManage
}

func NewK8sClusterHandler()  *K8sClusterHandler{
	return &K8sClusterHandler{m:model.NewK8sClusterManage()}
}

func (k *K8sClusterHandler)CreateK8sCluster(c *gin.Context) (err error)  {
	var m ent.K8sCluster
	if err:=c.BindJSON(&m);err!=nil{
		return svcerr.Error(err, svcerr.JSONBindErr)
	}
	clu,err:=k.m.CreateK8sCluser(m)
	if err != nil {
		return err
	}
	c.Set("data",clu)
	return
}

func (k *K8sClusterHandler)QueryK8sCluster(c *gin.Context)(err error)  {
	query := []string{k8scluster.FieldCluster}
	qm := model.ParesQuery(c,query)
	count,cluster,err:=k.m.QueryAll(qm)
	if err != nil {
		return err
	}
	c.Set("data",cluster)
	c.Set("total",count)
	return
}