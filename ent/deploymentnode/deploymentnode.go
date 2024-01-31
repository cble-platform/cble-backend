// Code generated by ent, DO NOT EDIT.

package deploymentnode

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deploymentnode type in the database.
	Label = "deployment_node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldVars holds the string denoting the vars field in the database.
	FieldVars = "vars"
	// EdgeDeployment holds the string denoting the deployment edge name in mutations.
	EdgeDeployment = "deployment"
	// EdgeResource holds the string denoting the resource edge name in mutations.
	EdgeResource = "resource"
	// EdgePrevNodes holds the string denoting the prev_nodes edge name in mutations.
	EdgePrevNodes = "prev_nodes"
	// EdgeNextNodes holds the string denoting the next_nodes edge name in mutations.
	EdgeNextNodes = "next_nodes"
	// Table holds the table name of the deploymentnode in the database.
	Table = "deployment_nodes"
	// DeploymentTable is the table that holds the deployment relation/edge.
	DeploymentTable = "deployment_nodes"
	// DeploymentInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentInverseTable = "deployments"
	// DeploymentColumn is the table column denoting the deployment relation/edge.
	DeploymentColumn = "deployment_node_deployment"
	// ResourceTable is the table that holds the resource relation/edge.
	ResourceTable = "deployment_nodes"
	// ResourceInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourceInverseTable = "resources"
	// ResourceColumn is the table column denoting the resource relation/edge.
	ResourceColumn = "deployment_node_resource"
	// PrevNodesTable is the table that holds the prev_nodes relation/edge. The primary key declared below.
	PrevNodesTable = "deployment_node_next_nodes"
	// NextNodesTable is the table that holds the next_nodes relation/edge. The primary key declared below.
	NextNodesTable = "deployment_node_next_nodes"
)

// Columns holds all SQL columns for deploymentnode fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldState,
	FieldVars,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deployment_nodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_node_deployment",
	"deployment_node_resource",
}

var (
	// PrevNodesPrimaryKey and PrevNodesColumn2 are the table columns denoting the
	// primary key for the prev_nodes relation (M2M).
	PrevNodesPrimaryKey = []string{"deployment_node_id", "prev_node_id"}
	// NextNodesPrimaryKey and NextNodesColumn2 are the table columns denoting the
	// primary key for the next_nodes relation (M2M).
	NextNodesPrimaryKey = []string{"deployment_node_id", "prev_node_id"}
)

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateAwaiting       State = "awaiting"
	StateParentAwaiting State = "parent_awaiting"
	StateInProgress     State = "in_progress"
	StateComplete       State = "complete"
	StateTainted        State = "tainted"
	StateFailed         State = "failed"
	StateToDelete       State = "to_delete"
	StateDeleted        State = "deleted"
	StateToRebuild      State = "to_rebuild"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateAwaiting, StateParentAwaiting, StateInProgress, StateComplete, StateTainted, StateFailed, StateToDelete, StateDeleted, StateToRebuild:
		return nil
	default:
		return fmt.Errorf("deploymentnode: invalid enum value for state field: %q", s)
	}
}

// OrderOption defines the ordering options for the DeploymentNode queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByDeploymentField orders the results by deployment field.
func ByDeploymentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeploymentStep(), sql.OrderByField(field, opts...))
	}
}

// ByResourceField orders the results by resource field.
func ByResourceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourceStep(), sql.OrderByField(field, opts...))
	}
}

// ByPrevNodesCount orders the results by prev_nodes count.
func ByPrevNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrevNodesStep(), opts...)
	}
}

// ByPrevNodes orders the results by prev_nodes terms.
func ByPrevNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrevNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNextNodesCount orders the results by next_nodes count.
func ByNextNodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNextNodesStep(), opts...)
	}
}

// ByNextNodes orders the results by next_nodes terms.
func ByNextNodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNextNodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDeploymentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeploymentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DeploymentTable, DeploymentColumn),
	)
}
func newResourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ResourceTable, ResourceColumn),
	)
}
func newPrevNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, PrevNodesTable, PrevNodesPrimaryKey...),
	)
}
func newNextNodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NextNodesTable, NextNodesPrimaryKey...),
	)
}
