// Code generated by ent, DO NOT EDIT.

package groupmembership

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the groupmembership type in the database.
	Label = "group_membership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the groupmembership in the database.
	Table = "group_memberships"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "group_memberships"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "group_memberships"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
)

// Columns holds all SQL columns for groupmembership fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldProjectID,
	FieldGroupID,
	FieldRole,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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

// Role defines the type for the "role" enum field.
type Role string

// RoleDeployer is the default value of the Role enum.
const DefaultRole = RoleDeployer

// Role values.
const (
	RoleViewer    Role = "viewer"
	RoleDeployer  Role = "deployer"
	RoleDeveloper Role = "developer"
	RoleAdmin     Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleViewer, RoleDeployer, RoleDeveloper, RoleAdmin:
		return nil
	default:
		return fmt.Errorf("groupmembership: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the GroupMembership queries.
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

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByGroupID orders the results by the group_id field.
func ByGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
	)
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
	)
}