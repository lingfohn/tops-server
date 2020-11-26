// Code generated by entc, DO NOT EDIT.

package k8scluster

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/lingfohn/lime/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Cluster applies equality check predicate on the "cluster" field. It's identical to ClusterEQ.
func Cluster(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCluster), v))
	})
}

// HelmApi applies equality check predicate on the "helmApi" field. It's identical to HelmApiEQ.
func HelmApi(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHelmApi), v))
	})
}

// AccessToken applies equality check predicate on the "accessToken" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ClusterEQ applies the EQ predicate on the "cluster" field.
func ClusterEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCluster), v))
	})
}

// ClusterNEQ applies the NEQ predicate on the "cluster" field.
func ClusterNEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCluster), v))
	})
}

// ClusterIn applies the In predicate on the "cluster" field.
func ClusterIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCluster), v...))
	})
}

// ClusterNotIn applies the NotIn predicate on the "cluster" field.
func ClusterNotIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCluster), v...))
	})
}

// ClusterGT applies the GT predicate on the "cluster" field.
func ClusterGT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCluster), v))
	})
}

// ClusterGTE applies the GTE predicate on the "cluster" field.
func ClusterGTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCluster), v))
	})
}

// ClusterLT applies the LT predicate on the "cluster" field.
func ClusterLT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCluster), v))
	})
}

// ClusterLTE applies the LTE predicate on the "cluster" field.
func ClusterLTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCluster), v))
	})
}

// ClusterContains applies the Contains predicate on the "cluster" field.
func ClusterContains(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCluster), v))
	})
}

// ClusterHasPrefix applies the HasPrefix predicate on the "cluster" field.
func ClusterHasPrefix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCluster), v))
	})
}

// ClusterHasSuffix applies the HasSuffix predicate on the "cluster" field.
func ClusterHasSuffix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCluster), v))
	})
}

// ClusterEqualFold applies the EqualFold predicate on the "cluster" field.
func ClusterEqualFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCluster), v))
	})
}

// ClusterContainsFold applies the ContainsFold predicate on the "cluster" field.
func ClusterContainsFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCluster), v))
	})
}

// HelmApiEQ applies the EQ predicate on the "helmApi" field.
func HelmApiEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHelmApi), v))
	})
}

// HelmApiNEQ applies the NEQ predicate on the "helmApi" field.
func HelmApiNEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHelmApi), v))
	})
}

// HelmApiIn applies the In predicate on the "helmApi" field.
func HelmApiIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHelmApi), v...))
	})
}

// HelmApiNotIn applies the NotIn predicate on the "helmApi" field.
func HelmApiNotIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHelmApi), v...))
	})
}

// HelmApiGT applies the GT predicate on the "helmApi" field.
func HelmApiGT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHelmApi), v))
	})
}

// HelmApiGTE applies the GTE predicate on the "helmApi" field.
func HelmApiGTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHelmApi), v))
	})
}

// HelmApiLT applies the LT predicate on the "helmApi" field.
func HelmApiLT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHelmApi), v))
	})
}

// HelmApiLTE applies the LTE predicate on the "helmApi" field.
func HelmApiLTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHelmApi), v))
	})
}

// HelmApiContains applies the Contains predicate on the "helmApi" field.
func HelmApiContains(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHelmApi), v))
	})
}

// HelmApiHasPrefix applies the HasPrefix predicate on the "helmApi" field.
func HelmApiHasPrefix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHelmApi), v))
	})
}

// HelmApiHasSuffix applies the HasSuffix predicate on the "helmApi" field.
func HelmApiHasSuffix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHelmApi), v))
	})
}

// HelmApiEqualFold applies the EqualFold predicate on the "helmApi" field.
func HelmApiEqualFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHelmApi), v))
	})
}

// HelmApiContainsFold applies the ContainsFold predicate on the "helmApi" field.
func HelmApiContainsFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHelmApi), v))
	})
}

// AccessTokenEQ applies the EQ predicate on the "accessToken" field.
func AccessTokenEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenNEQ applies the NEQ predicate on the "accessToken" field.
func AccessTokenNEQ(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIn applies the In predicate on the "accessToken" field.
func AccessTokenIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenNotIn applies the NotIn predicate on the "accessToken" field.
func AccessTokenNotIn(vs ...string) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccessToken), v...))
	})
}

// AccessTokenGT applies the GT predicate on the "accessToken" field.
func AccessTokenGT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenGTE applies the GTE predicate on the "accessToken" field.
func AccessTokenGTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLT applies the LT predicate on the "accessToken" field.
func AccessTokenLT(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccessToken), v))
	})
}

// AccessTokenLTE applies the LTE predicate on the "accessToken" field.
func AccessTokenLTE(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContains applies the Contains predicate on the "accessToken" field.
func AccessTokenContains(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "accessToken" field.
func AccessTokenHasPrefix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "accessToken" field.
func AccessTokenHasSuffix(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccessToken), v))
	})
}

// AccessTokenIsNil applies the IsNil predicate on the "accessToken" field.
func AccessTokenIsNil() predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAccessToken)))
	})
}

// AccessTokenNotNil applies the NotNil predicate on the "accessToken" field.
func AccessTokenNotNil() predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAccessToken)))
	})
}

// AccessTokenEqualFold applies the EqualFold predicate on the "accessToken" field.
func AccessTokenEqualFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccessToken), v))
	})
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "accessToken" field.
func AccessTokenContainsFold(v string) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccessToken), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.K8sCluster {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.K8sCluster(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasNamespaces applies the HasEdge predicate on the "namespaces" edge.
func HasNamespaces() predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespacesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NamespacesTable, NamespacesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespacesWith applies the HasEdge predicate on the "namespaces" edge with a given conditions (other predicates).
func HasNamespacesWith(preds ...predicate.Namespace) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespacesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NamespacesTable, NamespacesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.K8sCluster) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.K8sCluster) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.K8sCluster) predicate.K8sCluster {
	return predicate.K8sCluster(func(s *sql.Selector) {
		p(s.Not())
	})
}
