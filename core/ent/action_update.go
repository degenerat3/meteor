// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/degenerat3/meteor/core/ent/action"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/degenerat3/meteor/core/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ActionUpdate is the builder for updating Action entities.
type ActionUpdate struct {
	config
	hooks    []Hook
	mutation *ActionMutation
}

// Where adds a new predicate for the ActionUpdate builder.
func (au *ActionUpdate) Where(ps ...predicate.Action) *ActionUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetUUID sets the "uuid" field.
func (au *ActionUpdate) SetUUID(s string) *ActionUpdate {
	au.mutation.SetUUID(s)
	return au
}

// SetMode sets the "mode" field.
func (au *ActionUpdate) SetMode(s string) *ActionUpdate {
	au.mutation.SetMode(s)
	return au
}

// SetArgs sets the "args" field.
func (au *ActionUpdate) SetArgs(s string) *ActionUpdate {
	au.mutation.SetArgs(s)
	return au
}

// SetQueued sets the "queued" field.
func (au *ActionUpdate) SetQueued(b bool) *ActionUpdate {
	au.mutation.SetQueued(b)
	return au
}

// SetNillableQueued sets the "queued" field if the given value is not nil.
func (au *ActionUpdate) SetNillableQueued(b *bool) *ActionUpdate {
	if b != nil {
		au.SetQueued(*b)
	}
	return au
}

// SetResponded sets the "responded" field.
func (au *ActionUpdate) SetResponded(b bool) *ActionUpdate {
	au.mutation.SetResponded(b)
	return au
}

// SetNillableResponded sets the "responded" field if the given value is not nil.
func (au *ActionUpdate) SetNillableResponded(b *bool) *ActionUpdate {
	if b != nil {
		au.SetResponded(*b)
	}
	return au
}

// SetResult sets the "result" field.
func (au *ActionUpdate) SetResult(s string) *ActionUpdate {
	au.mutation.SetResult(s)
	return au
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (au *ActionUpdate) SetNillableResult(s *string) *ActionUpdate {
	if s != nil {
		au.SetResult(*s)
	}
	return au
}

// SetTargetingID sets the "targeting" edge to the Host entity by ID.
func (au *ActionUpdate) SetTargetingID(id int) *ActionUpdate {
	au.mutation.SetTargetingID(id)
	return au
}

// SetNillableTargetingID sets the "targeting" edge to the Host entity by ID if the given value is not nil.
func (au *ActionUpdate) SetNillableTargetingID(id *int) *ActionUpdate {
	if id != nil {
		au = au.SetTargetingID(*id)
	}
	return au
}

// SetTargeting sets the "targeting" edge to the Host entity.
func (au *ActionUpdate) SetTargeting(h *Host) *ActionUpdate {
	return au.SetTargetingID(h.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (au *ActionUpdate) Mutation() *ActionMutation {
	return au.mutation
}

// ClearTargeting clears the "targeting" edge to the Host entity.
func (au *ActionUpdate) ClearTargeting() *ActionUpdate {
	au.mutation.ClearTargeting()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActionUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActionUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActionUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: action.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldUUID,
		})
	}
	if value, ok := au.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldMode,
		})
	}
	if value, ok := au.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldArgs,
		})
	}
	if value, ok := au.mutation.Queued(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldQueued,
		})
	}
	if value, ok := au.mutation.Responded(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldResponded,
		})
	}
	if value, ok := au.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldResult,
		})
	}
	if au.mutation.TargetingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TargetingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ActionUpdateOne is the builder for updating a single Action entity.
type ActionUpdateOne struct {
	config
	hooks    []Hook
	mutation *ActionMutation
}

// SetUUID sets the "uuid" field.
func (auo *ActionUpdateOne) SetUUID(s string) *ActionUpdateOne {
	auo.mutation.SetUUID(s)
	return auo
}

// SetMode sets the "mode" field.
func (auo *ActionUpdateOne) SetMode(s string) *ActionUpdateOne {
	auo.mutation.SetMode(s)
	return auo
}

// SetArgs sets the "args" field.
func (auo *ActionUpdateOne) SetArgs(s string) *ActionUpdateOne {
	auo.mutation.SetArgs(s)
	return auo
}

// SetQueued sets the "queued" field.
func (auo *ActionUpdateOne) SetQueued(b bool) *ActionUpdateOne {
	auo.mutation.SetQueued(b)
	return auo
}

// SetNillableQueued sets the "queued" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableQueued(b *bool) *ActionUpdateOne {
	if b != nil {
		auo.SetQueued(*b)
	}
	return auo
}

// SetResponded sets the "responded" field.
func (auo *ActionUpdateOne) SetResponded(b bool) *ActionUpdateOne {
	auo.mutation.SetResponded(b)
	return auo
}

// SetNillableResponded sets the "responded" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableResponded(b *bool) *ActionUpdateOne {
	if b != nil {
		auo.SetResponded(*b)
	}
	return auo
}

// SetResult sets the "result" field.
func (auo *ActionUpdateOne) SetResult(s string) *ActionUpdateOne {
	auo.mutation.SetResult(s)
	return auo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableResult(s *string) *ActionUpdateOne {
	if s != nil {
		auo.SetResult(*s)
	}
	return auo
}

// SetTargetingID sets the "targeting" edge to the Host entity by ID.
func (auo *ActionUpdateOne) SetTargetingID(id int) *ActionUpdateOne {
	auo.mutation.SetTargetingID(id)
	return auo
}

// SetNillableTargetingID sets the "targeting" edge to the Host entity by ID if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableTargetingID(id *int) *ActionUpdateOne {
	if id != nil {
		auo = auo.SetTargetingID(*id)
	}
	return auo
}

// SetTargeting sets the "targeting" edge to the Host entity.
func (auo *ActionUpdateOne) SetTargeting(h *Host) *ActionUpdateOne {
	return auo.SetTargetingID(h.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (auo *ActionUpdateOne) Mutation() *ActionMutation {
	return auo.mutation
}

// ClearTargeting clears the "targeting" edge to the Host entity.
func (auo *ActionUpdateOne) ClearTargeting() *ActionUpdateOne {
	auo.mutation.ClearTargeting()
	return auo
}

// Save executes the query and returns the updated Action entity.
func (auo *ActionUpdateOne) Save(ctx context.Context) (*Action, error) {
	var (
		err  error
		node *Action
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActionUpdateOne) SaveX(ctx context.Context) *Action {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActionUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActionUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ActionUpdateOne) sqlSave(ctx context.Context) (_node *Action, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: action.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Action.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldUUID,
		})
	}
	if value, ok := auo.mutation.Mode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldMode,
		})
	}
	if value, ok := auo.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldArgs,
		})
	}
	if value, ok := auo.mutation.Queued(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldQueued,
		})
	}
	if value, ok := auo.mutation.Responded(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: action.FieldResponded,
		})
	}
	if value, ok := auo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldResult,
		})
	}
	if auo.mutation.TargetingCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TargetingIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Action{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
