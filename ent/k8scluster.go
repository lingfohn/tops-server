// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/k8scluster"
)

// K8sCluster is the model entity for the K8sCluster schema.
type K8sCluster struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Cluster holds the value of the "cluster" field.
	Cluster string `json:"cluster,omitempty"`
	// HelmApi holds the value of the "helmApi" field.
	HelmApi string `json:"helmApi"`
	// AccessToken holds the value of the "accessToken" field.
	AccessToken string `json:"accessToken"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the K8sClusterQuery when eager-loading is set.
	Edges K8sClusterEdges `json:"edges"`
}

// K8sClusterEdges holds the relations/edges for other nodes in the graph.
type K8sClusterEdges struct {
	// Namespaces holds the value of the namespaces edge.
	Namespaces []*Namespace `json:"namespaces,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NamespacesOrErr returns the Namespaces value or an error if the edge
// was not loaded in eager-loading.
func (e K8sClusterEdges) NamespacesOrErr() ([]*Namespace, error) {
	if e.loadedTypes[0] {
		return e.Namespaces, nil
	}
	return nil, &NotLoadedError{edge: "namespaces"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*K8sCluster) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // cluster
		&sql.NullString{}, // helmApi
		&sql.NullString{}, // accessToken
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the K8sCluster fields.
func (kc *K8sCluster) assignValues(values ...interface{}) error {
	if m, n := len(values), len(k8scluster.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	kc.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cluster", values[0])
	} else if value.Valid {
		kc.Cluster = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field helmApi", values[1])
	} else if value.Valid {
		kc.HelmApi = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field accessToken", values[2])
	} else if value.Valid {
		kc.AccessToken = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[3])
	} else if value.Valid {
		kc.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[4])
	} else if value.Valid {
		kc.UpdatedAt = value.Time
	}
	return nil
}

// QueryNamespaces queries the namespaces edge of the K8sCluster.
func (kc *K8sCluster) QueryNamespaces() *NamespaceQuery {
	return (&K8sClusterClient{config: kc.config}).QueryNamespaces(kc)
}

// Update returns a builder for updating this K8sCluster.
// Note that, you need to call K8sCluster.Unwrap() before calling this method, if this K8sCluster
// was returned from a transaction, and the transaction was committed or rolled back.
func (kc *K8sCluster) Update() *K8sClusterUpdateOne {
	return (&K8sClusterClient{config: kc.config}).UpdateOne(kc)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (kc *K8sCluster) Unwrap() *K8sCluster {
	tx, ok := kc.config.driver.(*txDriver)
	if !ok {
		panic("ent: K8sCluster is not a transactional entity")
	}
	kc.config.driver = tx.drv
	return kc
}

// String implements the fmt.Stringer.
func (kc *K8sCluster) String() string {
	var builder strings.Builder
	builder.WriteString("K8sCluster(")
	builder.WriteString(fmt.Sprintf("id=%v", kc.ID))
	builder.WriteString(", cluster=")
	builder.WriteString(kc.Cluster)
	builder.WriteString(", helmApi=")
	builder.WriteString(kc.HelmApi)
	builder.WriteString(", accessToken=")
	builder.WriteString(kc.AccessToken)
	builder.WriteString(", createdAt=")
	builder.WriteString(kc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(kc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// K8sClusters is a parsable slice of K8sCluster.
type K8sClusters []*K8sCluster

func (kc K8sClusters) config(cfg config) {
	for _i := range kc {
		kc[_i].config = cfg
	}
}
