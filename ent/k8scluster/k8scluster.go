// Code generated by entc, DO NOT EDIT.

package k8scluster

import (
	"time"
)

const (
	// Label holds the string label denoting the k8scluster type in the database.
	Label = "k8s_cluster"
	// FieldID holds the string denoting the id field in the database.
	FieldID          = "id"          // FieldCluster holds the string denoting the cluster vertex property in the database.
	FieldCluster     = "cluster"     // FieldHelmApi holds the string denoting the helmapi vertex property in the database.
	FieldHelmApi     = "helmApi"     // FieldAccessToken holds the string denoting the accesstoken vertex property in the database.
	FieldAccessToken = "accessToken" // FieldCreatedAt holds the string denoting the createdat vertex property in the database.
	FieldCreatedAt   = "createdAt"   // FieldUpdatedAt holds the string denoting the updatedat vertex property in the database.
	FieldUpdatedAt   = "updatedAt"

	// EdgeNamespaces holds the string denoting the namespaces edge name in mutations.
	EdgeNamespaces = "namespaces"

	// Table holds the table name of the k8scluster in the database.
	Table = "k8s_clusters"
	// NamespacesTable is the table the holds the namespaces relation/edge.
	NamespacesTable = "namespaces"
	// NamespacesInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespacesInverseTable = "namespaces"
	// NamespacesColumn is the table column denoting the namespaces relation/edge.
	NamespacesColumn = "k8s_cluster_namespaces"
)

// Columns holds all SQL columns for k8scluster fields.
var Columns = []string{
	FieldID,
	FieldCluster,
	FieldHelmApi,
	FieldAccessToken,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)
