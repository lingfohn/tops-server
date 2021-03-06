// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/lingfohn/lime/ent/menu"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Component holds the value of the "component" field.
	Component string `json:"component,omitempty"`
	// ParentId holds the value of the "parentId" field.
	ParentId int `json:"parentId"`
	// Redirect holds the value of the "redirect" field.
	Redirect string `json:"redirect,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Icon holds the value of the "icon" field.
	Icon string `json:"icon,omitempty"`
	// Target holds the value of the "target" field.
	Target string `json:"target,omitempty"`
	// KeepAlive holds the value of the "keepAlive" field.
	KeepAlive bool `json:"keepAlive"`
	// Show holds the value of the "show" field.
	Show *bool `json:"show,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges         MenuEdges `json:"edges"`
	menu_submenus *int
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Submenus holds the value of the submenus edge.
	Submenus []*Menu `json:"submenus,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: menu.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// SubmenusOrErr returns the Submenus value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) SubmenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[1] {
		return e.Submenus, nil
	}
	return nil, &NotLoadedError{edge: "submenus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // path
		&sql.NullString{}, // name
		&sql.NullString{}, // component
		&sql.NullInt64{},  // parentId
		&sql.NullString{}, // redirect
		&sql.NullInt64{},  // weight
		&sql.NullInt64{},  // level
		&sql.NullString{}, // title
		&sql.NullString{}, // icon
		&sql.NullString{}, // target
		&sql.NullBool{},   // keepAlive
		&sql.NullBool{},   // show
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Menu) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // menu_submenus
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(values ...interface{}) error {
	if m, n := len(values), len(menu.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field path", values[0])
	} else if value.Valid {
		m.Path = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		m.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field component", values[2])
	} else if value.Valid {
		m.Component = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field parentId", values[3])
	} else if value.Valid {
		m.ParentId = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field redirect", values[4])
	} else if value.Valid {
		m.Redirect = value.String
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field weight", values[5])
	} else if value.Valid {
		m.Weight = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field level", values[6])
	} else if value.Valid {
		m.Level = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[7])
	} else if value.Valid {
		m.Title = value.String
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field icon", values[8])
	} else if value.Valid {
		m.Icon = value.String
	}
	if value, ok := values[9].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field target", values[9])
	} else if value.Valid {
		m.Target = value.String
	}
	if value, ok := values[10].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field keepAlive", values[10])
	} else if value.Valid {
		m.KeepAlive = value.Bool
	}
	if value, ok := values[11].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field show", values[11])
	} else if value.Valid {
		m.Show = new(bool)
		*m.Show = value.Bool
	}
	if value, ok := values[12].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[12])
	} else if value.Valid {
		m.CreatedAt = value.Time
	}
	if value, ok := values[13].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[13])
	} else if value.Valid {
		m.UpdatedAt = value.Time
	}
	values = values[14:]
	if len(values) == len(menu.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field menu_submenus", value)
		} else if value.Valid {
			m.menu_submenus = new(int)
			*m.menu_submenus = int(value.Int64)
		}
	}
	return nil
}

// QueryParent queries the parent edge of the Menu.
func (m *Menu) QueryParent() *MenuQuery {
	return (&MenuClient{config: m.config}).QueryParent(m)
}

// QuerySubmenus queries the submenus edge of the Menu.
func (m *Menu) QuerySubmenus() *MenuQuery {
	return (&MenuClient{config: m.config}).QuerySubmenus(m)
}

// Update returns a builder for updating this Menu.
// Note that, you need to call Menu.Unwrap() before calling this method, if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return (&MenuClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", path=")
	builder.WriteString(m.Path)
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", component=")
	builder.WriteString(m.Component)
	builder.WriteString(", parentId=")
	builder.WriteString(fmt.Sprintf("%v", m.ParentId))
	builder.WriteString(", redirect=")
	builder.WriteString(m.Redirect)
	builder.WriteString(", weight=")
	builder.WriteString(fmt.Sprintf("%v", m.Weight))
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", m.Level))
	builder.WriteString(", title=")
	builder.WriteString(m.Title)
	builder.WriteString(", icon=")
	builder.WriteString(m.Icon)
	builder.WriteString(", target=")
	builder.WriteString(m.Target)
	builder.WriteString(", keepAlive=")
	builder.WriteString(fmt.Sprintf("%v", m.KeepAlive))
	if v := m.Show; v != nil {
		builder.WriteString(", show=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", createdAt=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu

func (m Menus) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
