// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/application"
	"github.com/lingfohn/lime/ent/helmconfig"
	"github.com/lingfohn/lime/ent/instance"
)

// Instance is the model entity for the Instance schema.
type Instance struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstanceQuery when eager-loading is set.
	Edges                 InstanceEdges `json:"edges"`
	application_instances *int
	instance_config       *int
}

// InstanceEdges holds the relations/edges for other nodes in the graph.
type InstanceEdges struct {
	// Application holds the value of the application edge.
	Application *Application
	// Builds holds the value of the builds edge.
	Builds []*Build
	// Config holds the value of the config edge.
	Config *HelmConfig
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ApplicationOrErr returns the Application value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) ApplicationOrErr() (*Application, error) {
	if e.loadedTypes[0] {
		if e.Application == nil {
			// The edge application was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: application.Label}
		}
		return e.Application, nil
	}
	return nil, &NotLoadedError{edge: "application"}
}

// BuildsOrErr returns the Builds value or an error if the edge
// was not loaded in eager-loading.
func (e InstanceEdges) BuildsOrErr() ([]*Build, error) {
	if e.loadedTypes[1] {
		return e.Builds, nil
	}
	return nil, &NotLoadedError{edge: "builds"}
}

// ConfigOrErr returns the Config value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InstanceEdges) ConfigOrErr() (*HelmConfig, error) {
	if e.loadedTypes[2] {
		if e.Config == nil {
			// The edge config was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: helmconfig.Label}
		}
		return e.Config, nil
	}
	return nil, &NotLoadedError{edge: "config"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Instance) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Instance) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // application_instances
		&sql.NullInt64{}, // instance_config
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Instance fields.
func (i *Instance) assignValues(values ...interface{}) error {
	if m, n := len(values), len(instance.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		i.Name = value.String
	}
	values = values[1:]
	if len(values) == len(instance.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field application_instances", value)
		} else if value.Valid {
			i.application_instances = new(int)
			*i.application_instances = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field instance_config", value)
		} else if value.Valid {
			i.instance_config = new(int)
			*i.instance_config = int(value.Int64)
		}
	}
	return nil
}

// QueryApplication queries the application edge of the Instance.
func (i *Instance) QueryApplication() *ApplicationQuery {
	return (&InstanceClient{config: i.config}).QueryApplication(i)
}

// QueryBuilds queries the builds edge of the Instance.
func (i *Instance) QueryBuilds() *BuildQuery {
	return (&InstanceClient{config: i.config}).QueryBuilds(i)
}

// QueryConfig queries the config edge of the Instance.
func (i *Instance) QueryConfig() *HelmConfigQuery {
	return (&InstanceClient{config: i.config}).QueryConfig(i)
}

// Update returns a builder for updating this Instance.
// Note that, you need to call Instance.Unwrap() before calling this method, if this Instance
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Instance) Update() *InstanceUpdateOne {
	return (&InstanceClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Instance) Unwrap() *Instance {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Instance is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Instance) String() string {
	var builder strings.Builder
	builder.WriteString("Instance(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", name=")
	builder.WriteString(i.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Instances is a parsable slice of Instance.
type Instances []*Instance

func (i Instances) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}