// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/degenerat3/meteor/meteor/core/ent/action"
	"github.com/degenerat3/meteor/meteor/core/ent/host"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ActionCreate is the builder for creating a Action entity.
type ActionCreate struct {
	config
	mutation *ActionMutation
	hooks    []Hook
}

// SetUUID sets the uuid field.
func (ac *ActionCreate) SetUUID(s string) *ActionCreate {
	ac.mutation.SetUUID(s)
	return ac
}

// SetMode sets the mode field.
func (ac *ActionCreate) SetMode(s string) *ActionCreate {
	ac.mutation.SetMode(s)
	return ac
}

// SetArgs sets the args field.
func (ac *ActionCreate) SetArgs(s string) *ActionCreate {
	ac.mutation.SetArgs(s)
	return ac
}

// SetQueued sets the queued field.
func (ac *ActionCreate) SetQueued(b bool) *ActionCreate {
	ac.mutation.SetQueued(b)
	return ac
}

// SetNillableQueued sets the queued field if the given value is not nil.
func (ac *ActionCreate) SetNillableQueued(b *bool) *ActionCreate {
	if b != nil {
		ac.SetQueued(*b)
	}
	return ac
}

// SetResponded sets the responded field.
func (ac *ActionCreate) SetResponded(b bool) *ActionCreate {
	ac.mutation.SetResponded(b)
	return ac
}

// SetNillableResponded sets the responded field if the given value is not nil.
func (ac *ActionCreate) SetNillableResponded(b *bool) *ActionCreate {
	if b != nil {
		ac.SetResponded(*b)
	}
	return ac
}

// SetResult sets the result field.
func (ac *ActionCreate) SetResult(s string) *ActionCreate {
	ac.mutation.SetResult(s)
	return ac
}

// SetNillableResult sets the result field if the given value is not nil.
func (ac *ActionCreate) SetNillableResult(s *string) *ActionCreate {
	if s != nil {
		ac.SetResult(*s)
	}
	return ac
}

// SetTargetingID sets the targeting edge to Host by id.
func (ac *ActionCreate) SetTargetingID(id int) *ActionCreate {
	ac.mutation.SetTargetingID(id)
	return ac
}

// SetNillableTargetingID sets the targeting edge to Host by id if the given value is not nil.
func (ac *ActionCreate) SetNillableTargetingID(id *int) *ActionCreate {
	if id != nil {
		ac = ac.SetTargetingID(*id)
	}
	return ac
}

// SetTargeting sets the targeting edge to Host.
func (ac *ActionCreate) SetTargeting(h *Host) *ActionCreate {
	return ac.SetTargetingID(h.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (ac *ActionCreate) Mutation() *ActionMutation {
	return ac.mutation
}

// Save creates the Action in the database.
func (ac *ActionCreate) Save(ctx context.Context) (*Action, error) {
	if err := ac.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Action
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActionCreate) SaveX(ctx context.Context) *Action {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *ActionCreate) preSave() error {
	if _, ok := ac.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := ac.mutation.Mode(); !ok {
		return &ValidationError{Name: "mode", err: errors.New("ent: missing required field \"mode\"")}
	}
	if _, ok := ac.mutation.Args(); !ok {
		return &ValidationError{Name: "args", err: errors.New("ent: missing required field \"args\"")}
	}
	if _, ok := ac.mutation.Queued(); !ok {
		v := action.DefaultQueued
		ac.mutation.SetQueued(v)
	}
	if _, ok := ac.mutation.Responded(); !ok {
		v := action.DefaultResponded
		ac.mutation.SetResponded(v)
	}
	if _, ok := ac.mutation.Result(); !ok {
		v := action.DefaultResult
		ac.mutation.SetResult(v)
	}
	return nil
}

func (ac *ActionCreate) sqlSave(ctx context.Context) (*Action, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *ActionCreate) createSpec() (*Action, *sqlgraph.CreateSpec) {
	var (
		a     = &Action{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: action.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: action.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldUUID,
		})
		a.UUID = value
	}
	if value, ok := ac.mutation.Mode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldMode,
		})
		a.Mode = value
	}
	if value, ok := ac.mutation.Args(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldArgs,
		})
		a.Args = value
	}
	if value, ok := ac.mutation.Queued(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldQueued,
		})
		a.Queued = value
	}
	if value, ok := ac.mutation.Responded(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldResponded,
		})
		a.Responded = value
	}
	if value, ok := ac.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldResult,
		})
		a.Result = value
	}
	if nodes := ac.mutation.TargetingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   action.TargetingTable,
			Columns: []string{action.TargetingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}

// ActionCreateBulk is the builder for creating a bulk of Action entities.
type ActionCreateBulk struct {
	config
	builders []*ActionCreate
}

// Save creates the Action entities in the database.
func (acb *ActionCreateBulk) Save(ctx context.Context) ([]*Action, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Action, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ActionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (acb *ActionCreateBulk) SaveX(ctx context.Context) []*Action {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
