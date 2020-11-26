package model

import (
	"context"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/k8scluster"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/storage"
)

type K8sClusterManage struct {
	 db *ent.Client
}

func NewK8sClusterManage() *K8sClusterManage {
	return &K8sClusterManage{
		db: storage.GetDB(),
	}
}

func (k *K8sClusterManage) CreateK8sCluser(m ent.K8sCluster) (cluster *ent.K8sCluster,err error){
	mc := k.db.K8sCluster.Create()
	mc.SetCluster(m.Cluster).
		SetHelmApi(m.HelmApi).
		SetAccessToken(m.AccessToken)
	return mc.Save(context.TODO())
}

func (p *K8sClusterManage)predicate(q QueryMap)[]predicate.K8sCluster  {
	var query []predicate.K8sCluster
	if val,ok := q.GetString(k8scluster.FieldCluster);ok{
		query = append(query,k8scluster.ClusterContains(val))
	}
	return query
}

func (k *K8sClusterManage) QueryAll(q QueryMap)(count int,clusters ent.K8sClusters,err error)  {
	cq := k.db.K8sCluster.Query()
	cq = cq.Where(k.predicate(q)...)
	count,err = cq.Count(context.TODO())
	clusters,err = cq.Offset(q.OffSet()).Limit(q.Limit()).All(context.TODO())
	return
}