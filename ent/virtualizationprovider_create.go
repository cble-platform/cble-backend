// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/backend/ent/blueprint"
	"github.com/cble-platform/backend/ent/virtualizationprovider"
	"github.com/google/uuid"
)

// VirtualizationProviderCreate is the builder for creating a VirtualizationProvider entity.
type VirtualizationProviderCreate struct {
	config
	mutation *VirtualizationProviderMutation
	hooks    []Hook
}

// SetDisplayName sets the "display_name" field.
func (vpc *VirtualizationProviderCreate) SetDisplayName(s string) *VirtualizationProviderCreate {
	vpc.mutation.SetDisplayName(s)
	return vpc
}

// SetProviderGitURL sets the "provider_git_url" field.
func (vpc *VirtualizationProviderCreate) SetProviderGitURL(s string) *VirtualizationProviderCreate {
	vpc.mutation.SetProviderGitURL(s)
	return vpc
}

// SetProviderVersion sets the "provider_version" field.
func (vpc *VirtualizationProviderCreate) SetProviderVersion(s string) *VirtualizationProviderCreate {
	vpc.mutation.SetProviderVersion(s)
	return vpc
}

// SetConfigPath sets the "config_path" field.
func (vpc *VirtualizationProviderCreate) SetConfigPath(s string) *VirtualizationProviderCreate {
	vpc.mutation.SetConfigPath(s)
	return vpc
}

// SetID sets the "id" field.
func (vpc *VirtualizationProviderCreate) SetID(u uuid.UUID) *VirtualizationProviderCreate {
	vpc.mutation.SetID(u)
	return vpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vpc *VirtualizationProviderCreate) SetNillableID(u *uuid.UUID) *VirtualizationProviderCreate {
	if u != nil {
		vpc.SetID(*u)
	}
	return vpc
}

// AddBlueprintIDs adds the "blueprints" edge to the Blueprint entity by IDs.
func (vpc *VirtualizationProviderCreate) AddBlueprintIDs(ids ...uuid.UUID) *VirtualizationProviderCreate {
	vpc.mutation.AddBlueprintIDs(ids...)
	return vpc
}

// AddBlueprints adds the "blueprints" edges to the Blueprint entity.
func (vpc *VirtualizationProviderCreate) AddBlueprints(b ...*Blueprint) *VirtualizationProviderCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return vpc.AddBlueprintIDs(ids...)
}

// Mutation returns the VirtualizationProviderMutation object of the builder.
func (vpc *VirtualizationProviderCreate) Mutation() *VirtualizationProviderMutation {
	return vpc.mutation
}

// Save creates the VirtualizationProvider in the database.
func (vpc *VirtualizationProviderCreate) Save(ctx context.Context) (*VirtualizationProvider, error) {
	vpc.defaults()
	return withHooks(ctx, vpc.sqlSave, vpc.mutation, vpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vpc *VirtualizationProviderCreate) SaveX(ctx context.Context) *VirtualizationProvider {
	v, err := vpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vpc *VirtualizationProviderCreate) Exec(ctx context.Context) error {
	_, err := vpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpc *VirtualizationProviderCreate) ExecX(ctx context.Context) {
	if err := vpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vpc *VirtualizationProviderCreate) defaults() {
	if _, ok := vpc.mutation.ID(); !ok {
		v := virtualizationprovider.DefaultID()
		vpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vpc *VirtualizationProviderCreate) check() error {
	if _, ok := vpc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "VirtualizationProvider.display_name"`)}
	}
	if _, ok := vpc.mutation.ProviderGitURL(); !ok {
		return &ValidationError{Name: "provider_git_url", err: errors.New(`ent: missing required field "VirtualizationProvider.provider_git_url"`)}
	}
	if _, ok := vpc.mutation.ProviderVersion(); !ok {
		return &ValidationError{Name: "provider_version", err: errors.New(`ent: missing required field "VirtualizationProvider.provider_version"`)}
	}
	if _, ok := vpc.mutation.ConfigPath(); !ok {
		return &ValidationError{Name: "config_path", err: errors.New(`ent: missing required field "VirtualizationProvider.config_path"`)}
	}
	return nil
}

func (vpc *VirtualizationProviderCreate) sqlSave(ctx context.Context) (*VirtualizationProvider, error) {
	if err := vpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vpc.driver, _spec); err != nil {
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
	vpc.mutation.id = &_node.ID
	vpc.mutation.done = true
	return _node, nil
}

func (vpc *VirtualizationProviderCreate) createSpec() (*VirtualizationProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &VirtualizationProvider{config: vpc.config}
		_spec = sqlgraph.NewCreateSpec(virtualizationprovider.Table, sqlgraph.NewFieldSpec(virtualizationprovider.FieldID, field.TypeUUID))
	)
	if id, ok := vpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vpc.mutation.DisplayName(); ok {
		_spec.SetField(virtualizationprovider.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := vpc.mutation.ProviderGitURL(); ok {
		_spec.SetField(virtualizationprovider.FieldProviderGitURL, field.TypeString, value)
		_node.ProviderGitURL = value
	}
	if value, ok := vpc.mutation.ProviderVersion(); ok {
		_spec.SetField(virtualizationprovider.FieldProviderVersion, field.TypeString, value)
		_node.ProviderVersion = value
	}
	if value, ok := vpc.mutation.ConfigPath(); ok {
		_spec.SetField(virtualizationprovider.FieldConfigPath, field.TypeString, value)
		_node.ConfigPath = value
	}
	if nodes := vpc.mutation.BlueprintsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   virtualizationprovider.BlueprintsTable,
			Columns: []string{virtualizationprovider.BlueprintsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blueprint.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VirtualizationProviderCreateBulk is the builder for creating many VirtualizationProvider entities in bulk.
type VirtualizationProviderCreateBulk struct {
	config
	builders []*VirtualizationProviderCreate
}

// Save creates the VirtualizationProvider entities in the database.
func (vpcb *VirtualizationProviderCreateBulk) Save(ctx context.Context) ([]*VirtualizationProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vpcb.builders))
	nodes := make([]*VirtualizationProvider, len(vpcb.builders))
	mutators := make([]Mutator, len(vpcb.builders))
	for i := range vpcb.builders {
		func(i int, root context.Context) {
			builder := vpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VirtualizationProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, vpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vpcb *VirtualizationProviderCreateBulk) SaveX(ctx context.Context) []*VirtualizationProvider {
	v, err := vpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vpcb *VirtualizationProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := vpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpcb *VirtualizationProviderCreateBulk) ExecX(ctx context.Context) {
	if err := vpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
