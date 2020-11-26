// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/project"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProjectName holds the value of the "projectName" field.
	ProjectName string `json:"projectName"`
	// ProType holds the value of the "proType" field.
	ProType string `json:"proType,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Gitlab holds the value of the "gitlab" field.
	Gitlab string `json:"gitlab,omitempty" gitlab`
	// Port holds the value of the "port" field.
	Port int `json:"port,omitempty" port`
	// DebugPort holds the value of the "debugPort" field.
	DebugPort int `json:"debugPort"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges ProjectEdges `json:"edges"`
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Applications holds the value of the applications edge.
	Applications []*Application
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplicationsOrErr returns the Applications value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ApplicationsOrErr() ([]*Application, error) {
	if e.loadedTypes[0] {
		return e.Applications, nil
	}
	return nil, &NotLoadedError{edge: "applications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // projectName
		&sql.NullString{}, // proType
		&sql.NullString{}, // description
		&sql.NullString{}, // gitlab
		&sql.NullInt64{},  // port
		&sql.NullInt64{},  // debugPort
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(values ...interface{}) error {
	if m, n := len(values), len(project.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field projectName", values[0])
	} else if value.Valid {
		pr.ProjectName = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field proType", values[1])
	} else if value.Valid {
		pr.ProType = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[2])
	} else if value.Valid {
		pr.Description = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field gitlab", values[3])
	} else if value.Valid {
		pr.Gitlab = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field port", values[4])
	} else if value.Valid {
		pr.Port = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field debugPort", values[5])
	} else if value.Valid {
		pr.DebugPort = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[6])
	} else if value.Valid {
		pr.CreatedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[7])
	} else if value.Valid {
		pr.UpdatedAt = value.Time
	}
	return nil
}

// QueryApplications queries the applications edge of the Project.
func (pr *Project) QueryApplications() *ApplicationQuery {
	return (&ProjectClient{config: pr.config}).QueryApplications(pr)
}

// Update returns a builder for updating this Project.
// Note that, you need to call Project.Unwrap() before calling this method, if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return (&ProjectClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", projectName=")
	builder.WriteString(pr.ProjectName)
	builder.WriteString(", proType=")
	builder.WriteString(pr.ProType)
	builder.WriteString(", description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", gitlab=")
	builder.WriteString(pr.Gitlab)
	builder.WriteString(", port=")
	builder.WriteString(fmt.Sprintf("%v", pr.Port))
	builder.WriteString(", debugPort=")
	builder.WriteString(fmt.Sprintf("%v", pr.DebugPort))
	builder.WriteString(", createdAt=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project

func (pr Projects) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}