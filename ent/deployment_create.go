// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/cble-backend/ent/blueprint"
	"github.com/cble-platform/cble-backend/ent/deployment"
	"github.com/cble-platform/cble-backend/ent/user"
	"github.com/google/uuid"
)

// DeploymentCreate is the builder for creating a Deployment entity.
type DeploymentCreate struct {
	config
	mutation *DeploymentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeploymentCreate) SetCreatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableCreatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeploymentCreate) SetUpdatedAt(t time.Time) *DeploymentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableUpdatedAt(t *time.Time) *DeploymentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DeploymentCreate) SetName(s string) *DeploymentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DeploymentCreate) SetDescription(s string) *DeploymentCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableDescription(s *string) *DeploymentCreate {
	if s != nil {
		dc.SetDescription(*s)
	}
	return dc
}

// SetTemplateVars sets the "template_vars" field.
func (dc *DeploymentCreate) SetTemplateVars(m map[string]interface{}) *DeploymentCreate {
	dc.mutation.SetTemplateVars(m)
	return dc
}

// SetDeploymentVars sets the "deployment_vars" field.
func (dc *DeploymentCreate) SetDeploymentVars(m map[string]interface{}) *DeploymentCreate {
	dc.mutation.SetDeploymentVars(m)
	return dc
}

// SetDeploymentState sets the "deployment_state" field.
func (dc *DeploymentCreate) SetDeploymentState(m map[string]string) *DeploymentCreate {
	dc.mutation.SetDeploymentState(m)
	return dc
}

// SetState sets the "state" field.
func (dc *DeploymentCreate) SetState(d deployment.State) *DeploymentCreate {
	dc.mutation.SetState(d)
	return dc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableState(d *deployment.State) *DeploymentCreate {
	if d != nil {
		dc.SetState(*d)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DeploymentCreate) SetID(u uuid.UUID) *DeploymentCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DeploymentCreate) SetNillableID(u *uuid.UUID) *DeploymentCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// SetBlueprintID sets the "blueprint" edge to the Blueprint entity by ID.
func (dc *DeploymentCreate) SetBlueprintID(id uuid.UUID) *DeploymentCreate {
	dc.mutation.SetBlueprintID(id)
	return dc
}

// SetBlueprint sets the "blueprint" edge to the Blueprint entity.
func (dc *DeploymentCreate) SetBlueprint(b *Blueprint) *DeploymentCreate {
	return dc.SetBlueprintID(b.ID)
}

// SetRequesterID sets the "requester" edge to the User entity by ID.
func (dc *DeploymentCreate) SetRequesterID(id uuid.UUID) *DeploymentCreate {
	dc.mutation.SetRequesterID(id)
	return dc
}

// SetRequester sets the "requester" edge to the User entity.
func (dc *DeploymentCreate) SetRequester(u *User) *DeploymentCreate {
	return dc.SetRequesterID(u.ID)
}

// Mutation returns the DeploymentMutation object of the builder.
func (dc *DeploymentCreate) Mutation() *DeploymentMutation {
	return dc.mutation
}

// Save creates the Deployment in the database.
func (dc *DeploymentCreate) Save(ctx context.Context) (*Deployment, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeploymentCreate) SaveX(ctx context.Context) *Deployment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeploymentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeploymentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeploymentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := deployment.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := deployment.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Description(); !ok {
		v := deployment.DefaultDescription
		dc.mutation.SetDescription(v)
	}
	if _, ok := dc.mutation.TemplateVars(); !ok {
		v := deployment.DefaultTemplateVars
		dc.mutation.SetTemplateVars(v)
	}
	if _, ok := dc.mutation.DeploymentVars(); !ok {
		v := deployment.DefaultDeploymentVars
		dc.mutation.SetDeploymentVars(v)
	}
	if _, ok := dc.mutation.DeploymentState(); !ok {
		v := deployment.DefaultDeploymentState
		dc.mutation.SetDeploymentState(v)
	}
	if _, ok := dc.mutation.State(); !ok {
		v := deployment.DefaultState
		dc.mutation.SetState(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := deployment.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeploymentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Deployment.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Deployment.updated_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Deployment.name"`)}
	}
	if _, ok := dc.mutation.TemplateVars(); !ok {
		return &ValidationError{Name: "template_vars", err: errors.New(`ent: missing required field "Deployment.template_vars"`)}
	}
	if _, ok := dc.mutation.DeploymentVars(); !ok {
		return &ValidationError{Name: "deployment_vars", err: errors.New(`ent: missing required field "Deployment.deployment_vars"`)}
	}
	if _, ok := dc.mutation.DeploymentState(); !ok {
		return &ValidationError{Name: "deployment_state", err: errors.New(`ent: missing required field "Deployment.deployment_state"`)}
	}
	if _, ok := dc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "Deployment.state"`)}
	}
	if v, ok := dc.mutation.State(); ok {
		if err := deployment.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Deployment.state": %w`, err)}
		}
	}
	if _, ok := dc.mutation.BlueprintID(); !ok {
		return &ValidationError{Name: "blueprint", err: errors.New(`ent: missing required edge "Deployment.blueprint"`)}
	}
	if _, ok := dc.mutation.RequesterID(); !ok {
		return &ValidationError{Name: "requester", err: errors.New(`ent: missing required edge "Deployment.requester"`)}
	}
	return nil
}

func (dc *DeploymentCreate) sqlSave(ctx context.Context) (*Deployment, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeploymentCreate) createSpec() (*Deployment, *sqlgraph.CreateSpec) {
	var (
		_node = &Deployment{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(deployment.Table, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(deployment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(deployment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(deployment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(deployment.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dc.mutation.TemplateVars(); ok {
		_spec.SetField(deployment.FieldTemplateVars, field.TypeJSON, value)
		_node.TemplateVars = value
	}
	if value, ok := dc.mutation.DeploymentVars(); ok {
		_spec.SetField(deployment.FieldDeploymentVars, field.TypeJSON, value)
		_node.DeploymentVars = value
	}
	if value, ok := dc.mutation.DeploymentState(); ok {
		_spec.SetField(deployment.FieldDeploymentState, field.TypeJSON, value)
		_node.DeploymentState = value
	}
	if value, ok := dc.mutation.State(); ok {
		_spec.SetField(deployment.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if nodes := dc.mutation.BlueprintIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.BlueprintTable,
			Columns: []string{deployment.BlueprintColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_blueprint = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.RequesterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deployment.RequesterTable,
			Columns: []string{deployment.RequesterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.deployment_requester = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeploymentCreateBulk is the builder for creating many Deployment entities in bulk.
type DeploymentCreateBulk struct {
	config
	err      error
	builders []*DeploymentCreate
}

// Save creates the Deployment entities in the database.
func (dcb *DeploymentCreateBulk) Save(ctx context.Context) ([]*Deployment, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Deployment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeploymentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) SaveX(ctx context.Context) []*Deployment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeploymentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeploymentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
