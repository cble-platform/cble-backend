// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/google/uuid"
)

// DeploymentNode is the model entity for the DeploymentNode schema.
type DeploymentNode struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The state of the deployed resource (should only by updated by the deploy engine)
	State deploymentnode.State `json:"state,omitempty"`
	// Stores metadata about the deployed resource for use with the provider
	Vars map[string]string `json:"vars,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeploymentNodeQuery when eager-loading is set.
	Edges                      DeploymentNodeEdges `json:"edges"`
	deployment_node_deployment *uuid.UUID
	deployment_node_resource   *uuid.UUID
	selectValues               sql.SelectValues
}

// DeploymentNodeEdges holds the relations/edges for other nodes in the graph.
type DeploymentNodeEdges struct {
	// The deployment for this node
	Deployment *Deployment `json:"deployment,omitempty"`
	// The resource this node represents
	Resource *Resource `json:"resource,omitempty"`
	// The previous nodes in the dependency tree
	PrevNodes []*DeploymentNode `json:"prev_nodes,omitempty"`
	// The next nodes in the dependency tree
	NextNodes []*DeploymentNode `json:"next_nodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// DeploymentOrErr returns the Deployment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentNodeEdges) DeploymentOrErr() (*Deployment, error) {
	if e.loadedTypes[0] {
		if e.Deployment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: deployment.Label}
		}
		return e.Deployment, nil
	}
	return nil, &NotLoadedError{edge: "deployment"}
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeploymentNodeEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[1] {
		if e.Resource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// PrevNodesOrErr returns the PrevNodes value or an error if the edge
// was not loaded in eager-loading.
func (e DeploymentNodeEdges) PrevNodesOrErr() ([]*DeploymentNode, error) {
	if e.loadedTypes[2] {
		return e.PrevNodes, nil
	}
	return nil, &NotLoadedError{edge: "prev_nodes"}
}

// NextNodesOrErr returns the NextNodes value or an error if the edge
// was not loaded in eager-loading.
func (e DeploymentNodeEdges) NextNodesOrErr() ([]*DeploymentNode, error) {
	if e.loadedTypes[3] {
		return e.NextNodes, nil
	}
	return nil, &NotLoadedError{edge: "next_nodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeploymentNode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deploymentnode.FieldVars:
			values[i] = new([]byte)
		case deploymentnode.FieldState:
			values[i] = new(sql.NullString)
		case deploymentnode.FieldCreatedAt, deploymentnode.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deploymentnode.FieldID:
			values[i] = new(uuid.UUID)
		case deploymentnode.ForeignKeys[0]: // deployment_node_deployment
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case deploymentnode.ForeignKeys[1]: // deployment_node_resource
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeploymentNode fields.
func (dn *DeploymentNode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deploymentnode.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dn.ID = *value
			}
		case deploymentnode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dn.CreatedAt = value.Time
			}
		case deploymentnode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dn.UpdatedAt = value.Time
			}
		case deploymentnode.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				dn.State = deploymentnode.State(value.String)
			}
		case deploymentnode.FieldVars:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vars", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dn.Vars); err != nil {
					return fmt.Errorf("unmarshal field vars: %w", err)
				}
			}
		case deploymentnode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_node_deployment", values[i])
			} else if value.Valid {
				dn.deployment_node_deployment = new(uuid.UUID)
				*dn.deployment_node_deployment = *value.S.(*uuid.UUID)
			}
		case deploymentnode.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field deployment_node_resource", values[i])
			} else if value.Valid {
				dn.deployment_node_resource = new(uuid.UUID)
				*dn.deployment_node_resource = *value.S.(*uuid.UUID)
			}
		default:
			dn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DeploymentNode.
// This includes values selected through modifiers, order, etc.
func (dn *DeploymentNode) Value(name string) (ent.Value, error) {
	return dn.selectValues.Get(name)
}

// QueryDeployment queries the "deployment" edge of the DeploymentNode entity.
func (dn *DeploymentNode) QueryDeployment() *DeploymentQuery {
	return NewDeploymentNodeClient(dn.config).QueryDeployment(dn)
}

// QueryResource queries the "resource" edge of the DeploymentNode entity.
func (dn *DeploymentNode) QueryResource() *ResourceQuery {
	return NewDeploymentNodeClient(dn.config).QueryResource(dn)
}

// QueryPrevNodes queries the "prev_nodes" edge of the DeploymentNode entity.
func (dn *DeploymentNode) QueryPrevNodes() *DeploymentNodeQuery {
	return NewDeploymentNodeClient(dn.config).QueryPrevNodes(dn)
}

// QueryNextNodes queries the "next_nodes" edge of the DeploymentNode entity.
func (dn *DeploymentNode) QueryNextNodes() *DeploymentNodeQuery {
	return NewDeploymentNodeClient(dn.config).QueryNextNodes(dn)
}

// Update returns a builder for updating this DeploymentNode.
// Note that you need to call DeploymentNode.Unwrap() before calling this method if this DeploymentNode
// was returned from a transaction, and the transaction was committed or rolled back.
func (dn *DeploymentNode) Update() *DeploymentNodeUpdateOne {
	return NewDeploymentNodeClient(dn.config).UpdateOne(dn)
}

// Unwrap unwraps the DeploymentNode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dn *DeploymentNode) Unwrap() *DeploymentNode {
	_tx, ok := dn.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeploymentNode is not a transactional entity")
	}
	dn.config.driver = _tx.drv
	return dn
}

// String implements the fmt.Stringer.
func (dn *DeploymentNode) String() string {
	var builder strings.Builder
	builder.WriteString("DeploymentNode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dn.ID))
	builder.WriteString("created_at=")
	builder.WriteString(dn.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dn.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", dn.State))
	builder.WriteString(", ")
	builder.WriteString("vars=")
	builder.WriteString(fmt.Sprintf("%v", dn.Vars))
	builder.WriteByte(')')
	return builder.String()
}

// DeploymentNodes is a parsable slice of DeploymentNode.
type DeploymentNodes []*DeploymentNode