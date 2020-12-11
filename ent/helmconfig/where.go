// Code generated by entc, DO NOT EDIT.

package helmconfig

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
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
func IDGT(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ChartVersion applies equality check predicate on the "chartVersion" field. It's identical to ChartVersionEQ.
func ChartVersion(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChartVersion), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// EnableService applies equality check predicate on the "enableService" field. It's identical to EnableServiceEQ.
func EnableService(v bool) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableService), v))
	})
}

// ServiceType applies equality check predicate on the "serviceType" field. It's identical to ServiceTypeEQ.
func ServiceType(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceType), v))
	})
}

// NodePort applies equality check predicate on the "nodePort" field. It's identical to NodePortEQ.
func NodePort(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodePort), v))
	})
}

// LimitMem applies equality check predicate on the "limitMem" field. It's identical to LimitMemEQ.
func LimitMem(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitMem), v))
	})
}

// LimitCPU applies equality check predicate on the "limitCPU" field. It's identical to LimitCPUEQ.
func LimitCPU(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitCPU), v))
	})
}

// ReqCPU applies equality check predicate on the "reqCPU" field. It's identical to ReqCPUEQ.
func ReqCPU(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReqCPU), v))
	})
}

// ReqMem applies equality check predicate on the "reqMem" field. It's identical to ReqMemEQ.
func ReqMem(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReqMem), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ChartVersionEQ applies the EQ predicate on the "chartVersion" field.
func ChartVersionEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChartVersion), v))
	})
}

// ChartVersionNEQ applies the NEQ predicate on the "chartVersion" field.
func ChartVersionNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChartVersion), v))
	})
}

// ChartVersionIn applies the In predicate on the "chartVersion" field.
func ChartVersionIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChartVersion), v...))
	})
}

// ChartVersionNotIn applies the NotIn predicate on the "chartVersion" field.
func ChartVersionNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChartVersion), v...))
	})
}

// ChartVersionGT applies the GT predicate on the "chartVersion" field.
func ChartVersionGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChartVersion), v))
	})
}

// ChartVersionGTE applies the GTE predicate on the "chartVersion" field.
func ChartVersionGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChartVersion), v))
	})
}

// ChartVersionLT applies the LT predicate on the "chartVersion" field.
func ChartVersionLT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChartVersion), v))
	})
}

// ChartVersionLTE applies the LTE predicate on the "chartVersion" field.
func ChartVersionLTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChartVersion), v))
	})
}

// ChartVersionContains applies the Contains predicate on the "chartVersion" field.
func ChartVersionContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChartVersion), v))
	})
}

// ChartVersionHasPrefix applies the HasPrefix predicate on the "chartVersion" field.
func ChartVersionHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChartVersion), v))
	})
}

// ChartVersionHasSuffix applies the HasSuffix predicate on the "chartVersion" field.
func ChartVersionHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChartVersion), v))
	})
}

// ChartVersionEqualFold applies the EqualFold predicate on the "chartVersion" field.
func ChartVersionEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChartVersion), v))
	})
}

// ChartVersionContainsFold applies the ContainsFold predicate on the "chartVersion" field.
func ChartVersionContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChartVersion), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// ActiveIn applies the In predicate on the "active" field.
func ActiveIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActive), v...))
	})
}

// ActiveNotIn applies the NotIn predicate on the "active" field.
func ActiveNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActive), v...))
	})
}

// ActiveGT applies the GT predicate on the "active" field.
func ActiveGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActive), v))
	})
}

// ActiveGTE applies the GTE predicate on the "active" field.
func ActiveGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActive), v))
	})
}

// ActiveLT applies the LT predicate on the "active" field.
func ActiveLT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActive), v))
	})
}

// ActiveLTE applies the LTE predicate on the "active" field.
func ActiveLTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActive), v))
	})
}

// ActiveContains applies the Contains predicate on the "active" field.
func ActiveContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActive), v))
	})
}

// ActiveHasPrefix applies the HasPrefix predicate on the "active" field.
func ActiveHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActive), v))
	})
}

// ActiveHasSuffix applies the HasSuffix predicate on the "active" field.
func ActiveHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActive), v))
	})
}

// ActiveEqualFold applies the EqualFold predicate on the "active" field.
func ActiveEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActive), v))
	})
}

// ActiveContainsFold applies the ContainsFold predicate on the "active" field.
func ActiveContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActive), v))
	})
}

// EnableServiceEQ applies the EQ predicate on the "enableService" field.
func EnableServiceEQ(v bool) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableService), v))
	})
}

// EnableServiceNEQ applies the NEQ predicate on the "enableService" field.
func EnableServiceNEQ(v bool) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnableService), v))
	})
}

// ServiceTypeEQ applies the EQ predicate on the "serviceType" field.
func ServiceTypeEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceType), v))
	})
}

// ServiceTypeNEQ applies the NEQ predicate on the "serviceType" field.
func ServiceTypeNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceType), v))
	})
}

// ServiceTypeIn applies the In predicate on the "serviceType" field.
func ServiceTypeIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldServiceType), v...))
	})
}

// ServiceTypeNotIn applies the NotIn predicate on the "serviceType" field.
func ServiceTypeNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldServiceType), v...))
	})
}

// ServiceTypeGT applies the GT predicate on the "serviceType" field.
func ServiceTypeGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceType), v))
	})
}

// ServiceTypeGTE applies the GTE predicate on the "serviceType" field.
func ServiceTypeGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceType), v))
	})
}

// ServiceTypeLT applies the LT predicate on the "serviceType" field.
func ServiceTypeLT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceType), v))
	})
}

// ServiceTypeLTE applies the LTE predicate on the "serviceType" field.
func ServiceTypeLTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceType), v))
	})
}

// ServiceTypeContains applies the Contains predicate on the "serviceType" field.
func ServiceTypeContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldServiceType), v))
	})
}

// ServiceTypeHasPrefix applies the HasPrefix predicate on the "serviceType" field.
func ServiceTypeHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldServiceType), v))
	})
}

// ServiceTypeHasSuffix applies the HasSuffix predicate on the "serviceType" field.
func ServiceTypeHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldServiceType), v))
	})
}

// ServiceTypeEqualFold applies the EqualFold predicate on the "serviceType" field.
func ServiceTypeEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldServiceType), v))
	})
}

// ServiceTypeContainsFold applies the ContainsFold predicate on the "serviceType" field.
func ServiceTypeContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldServiceType), v))
	})
}

// NodePortEQ applies the EQ predicate on the "nodePort" field.
func NodePortEQ(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNodePort), v))
	})
}

// NodePortNEQ applies the NEQ predicate on the "nodePort" field.
func NodePortNEQ(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNodePort), v))
	})
}

// NodePortIn applies the In predicate on the "nodePort" field.
func NodePortIn(vs ...int) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNodePort), v...))
	})
}

// NodePortNotIn applies the NotIn predicate on the "nodePort" field.
func NodePortNotIn(vs ...int) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNodePort), v...))
	})
}

// NodePortGT applies the GT predicate on the "nodePort" field.
func NodePortGT(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNodePort), v))
	})
}

// NodePortGTE applies the GTE predicate on the "nodePort" field.
func NodePortGTE(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNodePort), v))
	})
}

// NodePortLT applies the LT predicate on the "nodePort" field.
func NodePortLT(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNodePort), v))
	})
}

// NodePortLTE applies the LTE predicate on the "nodePort" field.
func NodePortLTE(v int) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNodePort), v))
	})
}

// LimitMemEQ applies the EQ predicate on the "limitMem" field.
func LimitMemEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitMem), v))
	})
}

// LimitMemNEQ applies the NEQ predicate on the "limitMem" field.
func LimitMemNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLimitMem), v))
	})
}

// LimitMemIn applies the In predicate on the "limitMem" field.
func LimitMemIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLimitMem), v...))
	})
}

// LimitMemNotIn applies the NotIn predicate on the "limitMem" field.
func LimitMemNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLimitMem), v...))
	})
}

// LimitMemGT applies the GT predicate on the "limitMem" field.
func LimitMemGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLimitMem), v))
	})
}

// LimitMemGTE applies the GTE predicate on the "limitMem" field.
func LimitMemGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLimitMem), v))
	})
}

// LimitMemLT applies the LT predicate on the "limitMem" field.
func LimitMemLT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLimitMem), v))
	})
}

// LimitMemLTE applies the LTE predicate on the "limitMem" field.
func LimitMemLTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLimitMem), v))
	})
}

// LimitMemContains applies the Contains predicate on the "limitMem" field.
func LimitMemContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLimitMem), v))
	})
}

// LimitMemHasPrefix applies the HasPrefix predicate on the "limitMem" field.
func LimitMemHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLimitMem), v))
	})
}

// LimitMemHasSuffix applies the HasSuffix predicate on the "limitMem" field.
func LimitMemHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLimitMem), v))
	})
}

// LimitMemEqualFold applies the EqualFold predicate on the "limitMem" field.
func LimitMemEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLimitMem), v))
	})
}

// LimitMemContainsFold applies the ContainsFold predicate on the "limitMem" field.
func LimitMemContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLimitMem), v))
	})
}

// LimitCPUEQ applies the EQ predicate on the "limitCPU" field.
func LimitCPUEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUNEQ applies the NEQ predicate on the "limitCPU" field.
func LimitCPUNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUIn applies the In predicate on the "limitCPU" field.
func LimitCPUIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLimitCPU), v...))
	})
}

// LimitCPUNotIn applies the NotIn predicate on the "limitCPU" field.
func LimitCPUNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLimitCPU), v...))
	})
}

// LimitCPUGT applies the GT predicate on the "limitCPU" field.
func LimitCPUGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUGTE applies the GTE predicate on the "limitCPU" field.
func LimitCPUGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLimitCPU), v))
	})
}

// LimitCPULT applies the LT predicate on the "limitCPU" field.
func LimitCPULT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLimitCPU), v))
	})
}

// LimitCPULTE applies the LTE predicate on the "limitCPU" field.
func LimitCPULTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUContains applies the Contains predicate on the "limitCPU" field.
func LimitCPUContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUHasPrefix applies the HasPrefix predicate on the "limitCPU" field.
func LimitCPUHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUHasSuffix applies the HasSuffix predicate on the "limitCPU" field.
func LimitCPUHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUEqualFold applies the EqualFold predicate on the "limitCPU" field.
func LimitCPUEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLimitCPU), v))
	})
}

// LimitCPUContainsFold applies the ContainsFold predicate on the "limitCPU" field.
func LimitCPUContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLimitCPU), v))
	})
}

// ReqCPUEQ applies the EQ predicate on the "reqCPU" field.
func ReqCPUEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReqCPU), v))
	})
}

// ReqCPUNEQ applies the NEQ predicate on the "reqCPU" field.
func ReqCPUNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReqCPU), v))
	})
}

// ReqCPUIn applies the In predicate on the "reqCPU" field.
func ReqCPUIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReqCPU), v...))
	})
}

// ReqCPUNotIn applies the NotIn predicate on the "reqCPU" field.
func ReqCPUNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReqCPU), v...))
	})
}

// ReqCPUGT applies the GT predicate on the "reqCPU" field.
func ReqCPUGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReqCPU), v))
	})
}

// ReqCPUGTE applies the GTE predicate on the "reqCPU" field.
func ReqCPUGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReqCPU), v))
	})
}

// ReqCPULT applies the LT predicate on the "reqCPU" field.
func ReqCPULT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReqCPU), v))
	})
}

// ReqCPULTE applies the LTE predicate on the "reqCPU" field.
func ReqCPULTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReqCPU), v))
	})
}

// ReqCPUContains applies the Contains predicate on the "reqCPU" field.
func ReqCPUContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReqCPU), v))
	})
}

// ReqCPUHasPrefix applies the HasPrefix predicate on the "reqCPU" field.
func ReqCPUHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReqCPU), v))
	})
}

// ReqCPUHasSuffix applies the HasSuffix predicate on the "reqCPU" field.
func ReqCPUHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReqCPU), v))
	})
}

// ReqCPUEqualFold applies the EqualFold predicate on the "reqCPU" field.
func ReqCPUEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReqCPU), v))
	})
}

// ReqCPUContainsFold applies the ContainsFold predicate on the "reqCPU" field.
func ReqCPUContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReqCPU), v))
	})
}

// ReqMemEQ applies the EQ predicate on the "reqMem" field.
func ReqMemEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReqMem), v))
	})
}

// ReqMemNEQ applies the NEQ predicate on the "reqMem" field.
func ReqMemNEQ(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReqMem), v))
	})
}

// ReqMemIn applies the In predicate on the "reqMem" field.
func ReqMemIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReqMem), v...))
	})
}

// ReqMemNotIn applies the NotIn predicate on the "reqMem" field.
func ReqMemNotIn(vs ...string) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReqMem), v...))
	})
}

// ReqMemGT applies the GT predicate on the "reqMem" field.
func ReqMemGT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReqMem), v))
	})
}

// ReqMemGTE applies the GTE predicate on the "reqMem" field.
func ReqMemGTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReqMem), v))
	})
}

// ReqMemLT applies the LT predicate on the "reqMem" field.
func ReqMemLT(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReqMem), v))
	})
}

// ReqMemLTE applies the LTE predicate on the "reqMem" field.
func ReqMemLTE(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReqMem), v))
	})
}

// ReqMemContains applies the Contains predicate on the "reqMem" field.
func ReqMemContains(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldReqMem), v))
	})
}

// ReqMemHasPrefix applies the HasPrefix predicate on the "reqMem" field.
func ReqMemHasPrefix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldReqMem), v))
	})
}

// ReqMemHasSuffix applies the HasSuffix predicate on the "reqMem" field.
func ReqMemHasSuffix(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldReqMem), v))
	})
}

// ReqMemEqualFold applies the EqualFold predicate on the "reqMem" field.
func ReqMemEqualFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldReqMem), v))
	})
}

// ReqMemContainsFold applies the ContainsFold predicate on the "reqMem" field.
func ReqMemContainsFold(v string) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldReqMem), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.HelmConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.HelmConfig(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.HelmConfig) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.HelmConfig) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
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
func Not(p predicate.HelmConfig) predicate.HelmConfig {
	return predicate.HelmConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
