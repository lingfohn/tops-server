// Code generated by entc, DO NOT EDIT.

package namespace

import (
	"time"
)

const (
	// Label holds the string label denoting the namespace type in the database.
	Label = "namespace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDockerRepo holds the string denoting the dockerrepo field in the database.
	FieldDockerRepo = "dockerRepo"
	// FieldRepoNamespace holds the string denoting the reponamespace field in the database.
	FieldRepoNamespace = "repoNamespace"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "createdAt"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updatedAt"

	// EdgeCluster holds the string denoting the cluster edge name in mutations.
	EdgeCluster = "cluster"
	// EdgeApplications holds the string denoting the applications edge name in mutations.
	EdgeApplications = "applications"

	// Table holds the table name of the namespace in the database.
	Table = "namespaces"
	// ClusterTable is the table the holds the cluster relation/edge.
	ClusterTable = "namespaces"
	// ClusterInverseTable is the table name for the K8sCluster entity.
	// It exists in this package in order to avoid circular dependency with the "k8scluster" package.
	ClusterInverseTable = "k8s_clusters"
	// ClusterColumn is the table column denoting the cluster relation/edge.
	ClusterColumn = "k8s_cluster_namespaces"
	// ApplicationsTable is the table the holds the applications relation/edge.
	ApplicationsTable = "applications"
	// ApplicationsInverseTable is the table name for the Application entity.
	// It exists in this package in order to avoid circular dependency with the "application" package.
	ApplicationsInverseTable = "applications"
	// ApplicationsColumn is the table column denoting the applications relation/edge.
	ApplicationsColumn = "namespace_applications"
)

// Columns holds all SQL columns for namespace fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDockerRepo,
	FieldRepoNamespace,
	FieldActive,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Namespace type.
var ForeignKeys = []string{
	"k8s_cluster_namespaces",
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
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
