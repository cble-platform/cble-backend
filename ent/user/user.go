// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeDeployments holds the string denoting the deployments edge name in mutations.
	EdgeDeployments = "deployments"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "user_groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// DeploymentsTable is the table that holds the deployments relation/edge.
	DeploymentsTable = "deployments"
	// DeploymentsInverseTable is the table name for the Deployment entity.
	// It exists in this package in order to avoid circular dependency with the "deployment" package.
	DeploymentsInverseTable = "deployments"
	// DeploymentsColumn is the table column denoting the deployments relation/edge.
	DeploymentsColumn = "deployment_requester"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldEmail,
	FieldPassword,
	FieldFirstName,
	FieldLastName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
)

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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeploymentsCount orders the results by deployments count.
func ByDeploymentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeploymentsStep(), opts...)
	}
}

// ByDeployments orders the results by deployments terms.
func ByDeployments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeploymentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, GroupsTable, GroupsPrimaryKey...),
	)
}
func newDeploymentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeploymentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DeploymentsTable, DeploymentsColumn),
	)
}
