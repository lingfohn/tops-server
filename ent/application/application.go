// Code generated by entc, DO NOT EDIT.

package application

import (
	"time"
)

const (
	// Label holds the string label denoting the application type in the database.
	Label = "application"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMulti holds the string denoting the multi field in the database.
	FieldMulti = "multi"
	// FieldProjectId holds the string denoting the projectid field in the database.
	FieldProjectId = "projectId"
	// FieldNamespaceId holds the string denoting the namespaceid field in the database.
	FieldNamespaceId = "namespaceId"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "createdAt"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updatedAt"

	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeInstances holds the string denoting the instances edge name in mutations.
	EdgeInstances = "instances"
	// EdgeConfig holds the string denoting the config edge name in mutations.
	EdgeConfig = "config"

	// Table holds the table name of the application in the database.
	Table = "applications"
	// NamespaceTable is the table the holds the namespace relation/edge.
	NamespaceTable = "applications"
	// NamespaceInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespaceInverseTable = "namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "namespace_applications"
	// ProjectTable is the table the holds the project relation/edge.
	ProjectTable = "applications"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_applications"
	// InstancesTable is the table the holds the instances relation/edge.
	InstancesTable = "instances"
	// InstancesInverseTable is the table name for the Instance entity.
	// It exists in this package in order to avoid circular dependency with the "instance" package.
	InstancesInverseTable = "instances"
	// InstancesColumn is the table column denoting the instances relation/edge.
	InstancesColumn = "application_instances"
	// ConfigTable is the table the holds the config relation/edge.
	ConfigTable = "applications"
	// ConfigInverseTable is the table name for the HelmConfig entity.
	// It exists in this package in order to avoid circular dependency with the "helmconfig" package.
	ConfigInverseTable = "helm_configs"
	// ConfigColumn is the table column denoting the config relation/edge.
	ConfigColumn = "application_config"
)

// Columns holds all SQL columns for application fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMulti,
	FieldProjectId,
	FieldNamespaceId,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Application type.
var ForeignKeys = []string{
	"application_config",
	"namespace_applications",
	"project_applications",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMulti holds the default value on creation for the multi field.
	DefaultMulti bool
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
