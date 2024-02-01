// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/deploymentnode"
	"github.com/cble-platform/cble-backend/ent/resource"
	"github.com/google/uuid"
)

// DeploymentNodeCreate is the builder for creating a DeploymentNode entity.
type DeploymentNodeCreate struct {
	config
	mutation *DeploymentNodeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dnc *DeploymentNodeCreate) SetCreatedAt(t time.Time) *DeploymentNodeCreate {
	dnc.mutation.SetCreatedAt(t)
	return dnc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dnc *DeploymentNodeCreate) SetNillableCreatedAt(t *time.Time) *DeploymentNodeCreate {
	if t != nil {
		dnc.SetCreatedAt(*t)
	}
	return dnc
}

// SetUpdatedAt sets the "updated_at" field.
func (dnc *DeploymentNodeCreate) SetUpdatedAt(t time.Time) *DeploymentNodeCreate {
	dnc.mutation.SetUpdatedAt(t)
	return dnc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dnc *DeploymentNodeCreate) SetNillableUpdatedAt(t *time.Time) *DeploymentNodeCreate {
	if t != nil {
		dnc.SetUpdatedAt(*t)
	}
	return dnc
}

// SetState sets the "state" field.
func (dnc *DeploymentNodeCreate) SetState(d deploymentnode.State) *DeploymentNodeCreate {
	dnc.mutation.SetState(d)
	return dnc
}

// SetVars sets the "vars" field.
func (dnc *DeploymentNodeCreate) SetVars(m map[string]string) *DeploymentNodeCreate {
	dnc.mutation.SetVars(m)
	return dnc
}

// SetID sets the "id" field.
func (dnc *DeploymentNodeCreate) SetID(u uuid.UUID) *DeploymentNodeCreate {
	dnc.mutation.SetID(u)
	return dnc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dnc *DeploymentNodeCreate) SetNillableID(u *uuid.UUID) *DeploymentNodeCreate {
	if u != nil {
		dnc.SetID(*u)
	}
	return dnc
}

// SetDeploymentID sets the "deployment" edge to the Deployment entity by ID.
func (dnc *DeploymentNodeCreate) SetDeploymentID(id uuid.UUID) *DeploymentNodeCreate {
	dnc.mutation.SetDeploymentID(id)
	return dnc
}

// SetDeployment sets the "deployment" edge to the Deployment entity.
func (dnc *DeploymentNodeCreate) SetDeployment(d *Deployment) *DeploymentNodeCreate {
	return dnc.SetDeploymentID(d.ID)
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (dnc *DeploymentNodeCreate) SetResourceID(id uuid.UUID) *DeploymentNodeCreate {
	dnc.mutation.SetResourceID(id)
	return dnc
}

// SetResource sets the "resource" edge to the Resource entity.
func (dnc *DeploymentNodeCreate) SetResource(r *Resource) *DeploymentNodeCreate {
	return dnc.SetResourceID(r.ID)
}

// AddPrevNodeIDs adds the "prev_nodes" edge to the DeploymentNode entity by IDs.
func (dnc *DeploymentNodeCreate) AddPrevNodeIDs(ids ...uuid.UUID) *DeploymentNodeCreate {
	dnc.mutation.AddPrevNodeIDs(ids...)
	return dnc
}

// AddPrevNodes adds the "prev_nodes" edges to the DeploymentNode entity.
func (dnc *DeploymentNodeCreate) AddPrevNodes(d ...*DeploymentNode) *DeploymentNodeCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dnc.AddPrevNodeIDs(ids...)
}

// AddNextNodeIDs adds the "next_nodes" edge to the DeploymentNode entity by IDs.
func (dnc *DeploymentNodeCreate) AddNextNodeIDs(ids ...uuid.UUID) *DeploymentNodeCreate {
	dnc.mutation.AddNextNodeIDs(ids...)
	return dnc
}

// AddNextNodes adds the "next_nodes" edges to the DeploymentNode entity.
func (dnc *DeploymentNodeCreate) AddNextNodes(d ...*DeploymentNode) *DeploymentNodeCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dnc.AddNextNodeIDs(ids...)
}

// Mutation returns the DeploymentNodeMutation object of the builder.
func (dnc *DeploymentNodeCreate) Mutation() *DeploymentNodeMutation {
	return dnc.mutation
}

// Save creates the DeploymentNode in the database.
func (dnc *DeploymentNodeCreate) Save(ctx context.Context) (*DeploymentNode, error) {
	dnc.defaults()
	return withHooks(ctx, dnc.sqlSave, dnc.mutation, dnc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dnc *DeploymentNodeCreate) SaveX(ctx context.Context) *DeploymentNode {
	v, err := dnc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dnc *DeploymentNodeCreate) Exec(ctx context.Context) error {
	_, err := dnc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dnc *DeploymentNodeCreate) ExecX(ctx context.Context) {
	if err := dnc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dnc *DeploymentNodeCreate) defaults() {
	if _, ok := dnc.mutation.CreatedAt(); !ok {
		v := deploymentnode.DefaultCreatedAt()
		dnc.mutation.SetCreatedAt(v)
	}
	if _, ok := dnc.mutation.UpdatedAt(); !ok {
		v := deploymentnode.DefaultUpdatedAt()
		dnc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dnc.mutation.ID(); !ok {
		v := deploymentnode.DefaultID()
		dnc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dnc *DeploymentNodeCreate) check() error {
	if _, ok := dnc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeploymentNode.created_at"`)}
	}
	if _, ok := dnc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeploymentNode.updated_at"`)}
	}
	if _, ok := dnc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "DeploymentNode.state"`)}
	}
	if v, ok := dnc.mutation.State(); ok {
		if err := deploymentnode.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "DeploymentNode.state": %w`, err)}
		}
	}
	if _, ok := dnc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New(`ent: missing required field "DeploymentNode.vars"`)}
	}
	if _, ok := dnc.mutation.DeploymentID(); !ok {
		return &ValidationError{Name: "deployment", err: errors.New(`ent: missing required edge "DeploymentNode.deployment"`)}
	}
	if _, ok := dnc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource", err: errors.New(`ent: missing required edge "DeploymentNode.resource"`)}
	}
	return nil
}

func (dnc *DeploymentNodeCreate) sqlSave(ctx context.Context) (*DeploymentNode, error) {
	if err := dnc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dnc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dnc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dnc.mutation.id = &_node.ID
	dnc.mutation.done = true
	return _node, nil
}

func (dnc *DeploymentNodeCreate) createSpec() (*DeploymentNode, *sqlgraph.CreateSpec) {
	var (
		_node = &DeploymentNode{config: dnc.config}
		_spec = sqlgraph.NewCreateSpec(deploymentnode.Table, sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID))
	)
	if id, ok := dnc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dnc.mutation.CreatedAt(); ok {
		_spec.SetField(deploymentnode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dnc.mutation.UpdatedAt(); ok {
		_spec.SetField(deploymentnode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dnc.mutation.State(); ok {
		_spec.SetField(deploymentnode.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := dnc.mutation.Vars(); ok {
		_spec.SetField(deploymentnode.FieldVars, field.TypeJSON, value)
		_node.Vars = value
	}
	if nodes := dnc.mutation.DeploymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deploymentnode.DeploymentTable,
			Columns: []string{deploymentnode.DeploymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_node_deployment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dnc.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deploymentnode.ResourceTable,
			Columns: []string{deploymentnode.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_node_resource = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dnc.mutation.PrevNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   deploymentnode.PrevNodesTable,
			Columns: deploymentnode.PrevNodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dnc.mutation.NextNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deploymentnode.NextNodesTable,
			Columns: deploymentnode.NextNodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deploymentnode.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentNodeCreateBulk is the builder for creating many DeploymentNode entities in bulk.
type DeploymentNodeCreateBulk struct {
	config
	err      error
	builders []*DeploymentNodeCreate
}

// Save creates the DeploymentNode entities in the database.
func (dncb *DeploymentNodeCreateBulk) Save(ctx context.Context) ([]*DeploymentNode, error) {
	if dncb.err != nil {
		return nil, dncb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dncb.builders))
	nodes := make([]*DeploymentNode, len(dncb.builders))
	mutators := make([]Mutator, len(dncb.builders))
	for i := range dncb.builders {
		func(i int, root context.Context) {
			builder := dncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentNodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dncb *DeploymentNodeCreateBulk) SaveX(ctx context.Context) []*DeploymentNode {
	v, err := dncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dncb *DeploymentNodeCreateBulk) Exec(ctx context.Context) error {
	_, err := dncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dncb *DeploymentNodeCreateBulk) ExecX(ctx context.Context) {
	if err := dncb.Exec(ctx); err != nil {
		panic(err)
	}
}
