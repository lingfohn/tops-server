// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/build"
	"github.com/lingfohn/lime/ent/instance"
)

// Build is the model entity for the Build schema.
type Build struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Status holds the value of the "status" field.
	Status int `json:"status"`
	// CommitId holds the value of the "commitId" field.
	CommitId string `json:"commitId"`
	// ShortId holds the value of the "shortId" field.
	ShortId string `json:"shortId"`
	// Message holds the value of the "message" field.
	Message string `json:"message"`
	// CommitterName holds the value of the "committerName" field.
	CommitterName string `json:"committerName"`
	// CommittedData holds the value of the "committedData" field.
	CommittedData time.Time `json:"committedData"`
	// BuildTime holds the value of the "buildTime" field.
	BuildTime time.Time `json:"buildTime"`
	// JenkinsQueueId holds the value of the "jenkinsQueueId" field.
	JenkinsQueueId int `json:"jenkinsQueueId"'`
	// JenkinsBuildId holds the value of the "jenkinsBuildId" field.
	JenkinsBuildId int `json:"jenkinsBuildId"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BuildQuery when eager-loading is set.
	Edges           BuildEdges `json:"edges"`
	instance_builds *int
}

// BuildEdges holds the relations/edges for other nodes in the graph.
type BuildEdges struct {
	// Instance holds the value of the instance edge.
	Instance *Instance
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) InstanceOrErr() (*Instance, error) {
	if e.loadedTypes[0] {
		if e.Instance == nil {
			// The edge instance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Build) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // tag
		&sql.NullString{}, // name
		&sql.NullInt64{},  // status
		&sql.NullString{}, // commitId
		&sql.NullString{}, // shortId
		&sql.NullString{}, // message
		&sql.NullString{}, // committerName
		&sql.NullTime{},   // committedData
		&sql.NullTime{},   // buildTime
		&sql.NullInt64{},  // jenkinsQueueId
		&sql.NullInt64{},  // jenkinsBuildId
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Build) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // instance_builds
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Build fields.
func (b *Build) assignValues(values ...interface{}) error {
	if m, n := len(values), len(build.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tag", values[0])
	} else if value.Valid {
		b.Tag = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		b.Name = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[2])
	} else if value.Valid {
		b.Status = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field commitId", values[3])
	} else if value.Valid {
		b.CommitId = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field shortId", values[4])
	} else if value.Valid {
		b.ShortId = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field message", values[5])
	} else if value.Valid {
		b.Message = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field committerName", values[6])
	} else if value.Valid {
		b.CommitterName = value.String
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field committedData", values[7])
	} else if value.Valid {
		b.CommittedData = value.Time
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field buildTime", values[8])
	} else if value.Valid {
		b.BuildTime = value.Time
	}
	if value, ok := values[9].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field jenkinsQueueId", values[9])
	} else if value.Valid {
		b.JenkinsQueueId = int(value.Int64)
	}
	if value, ok := values[10].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field jenkinsBuildId", values[10])
	} else if value.Valid {
		b.JenkinsBuildId = int(value.Int64)
	}
	if value, ok := values[11].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[11])
	} else if value.Valid {
		b.CreatedAt = value.Time
	}
	if value, ok := values[12].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[12])
	} else if value.Valid {
		b.UpdatedAt = value.Time
	}
	values = values[13:]
	if len(values) == len(build.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field instance_builds", value)
		} else if value.Valid {
			b.instance_builds = new(int)
			*b.instance_builds = int(value.Int64)
		}
	}
	return nil
}

// QueryInstance queries the instance edge of the Build.
func (b *Build) QueryInstance() *InstanceQuery {
	return (&BuildClient{config: b.config}).QueryInstance(b)
}

// Update returns a builder for updating this Build.
// Note that, you need to call Build.Unwrap() before calling this method, if this Build
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Build) Update() *BuildUpdateOne {
	return (&BuildClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Build) Unwrap() *Build {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Build is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Build) String() string {
	var builder strings.Builder
	builder.WriteString("Build(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", tag=")
	builder.WriteString(b.Tag)
	builder.WriteString(", name=")
	builder.WriteString(b.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", commitId=")
	builder.WriteString(b.CommitId)
	builder.WriteString(", shortId=")
	builder.WriteString(b.ShortId)
	builder.WriteString(", message=")
	builder.WriteString(b.Message)
	builder.WriteString(", committerName=")
	builder.WriteString(b.CommitterName)
	builder.WriteString(", committedData=")
	builder.WriteString(b.CommittedData.Format(time.ANSIC))
	builder.WriteString(", buildTime=")
	builder.WriteString(b.BuildTime.Format(time.ANSIC))
	builder.WriteString(", jenkinsQueueId=")
	builder.WriteString(fmt.Sprintf("%v", b.JenkinsQueueId))
	builder.WriteString(", jenkinsBuildId=")
	builder.WriteString(fmt.Sprintf("%v", b.JenkinsBuildId))
	builder.WriteString(", createdAt=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Builds is a parsable slice of Build.
type Builds []*Build

func (b Builds) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
