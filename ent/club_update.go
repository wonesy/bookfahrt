// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wonesy/bookfahrt/ent/club"
	"github.com/wonesy/bookfahrt/ent/invitation"
	"github.com/wonesy/bookfahrt/ent/predicate"
	"github.com/wonesy/bookfahrt/ent/user"
)

// ClubUpdate is the builder for updating Club entities.
type ClubUpdate struct {
	config
	hooks    []Hook
	mutation *ClubMutation
}

// Where appends a list predicates to the ClubUpdate builder.
func (cu *ClubUpdate) Where(ps ...predicate.Club) *ClubUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ClubUpdate) SetName(s string) *ClubUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cu *ClubUpdate) AddMemberIDs(ids ...int) *ClubUpdate {
	cu.mutation.AddMemberIDs(ids...)
	return cu
}

// AddMembers adds the "members" edges to the User entity.
func (cu *ClubUpdate) AddMembers(u ...*User) *ClubUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddMemberIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (cu *ClubUpdate) AddInvitationIDs(ids ...uuid.UUID) *ClubUpdate {
	cu.mutation.AddInvitationIDs(ids...)
	return cu
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (cu *ClubUpdate) AddInvitations(i ...*Invitation) *ClubUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddInvitationIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cu *ClubUpdate) Mutation() *ClubMutation {
	return cu.mutation
}

// ClearMembers clears all "members" edges to the User entity.
func (cu *ClubUpdate) ClearMembers() *ClubUpdate {
	cu.mutation.ClearMembers()
	return cu
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cu *ClubUpdate) RemoveMemberIDs(ids ...int) *ClubUpdate {
	cu.mutation.RemoveMemberIDs(ids...)
	return cu
}

// RemoveMembers removes "members" edges to User entities.
func (cu *ClubUpdate) RemoveMembers(u ...*User) *ClubUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveMemberIDs(ids...)
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (cu *ClubUpdate) ClearInvitations() *ClubUpdate {
	cu.mutation.ClearInvitations()
	return cu
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (cu *ClubUpdate) RemoveInvitationIDs(ids ...uuid.UUID) *ClubUpdate {
	cu.mutation.RemoveInvitationIDs(ids...)
	return cu
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (cu *ClubUpdate) RemoveInvitations(i ...*Invitation) *ClubUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveInvitationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClubUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClubUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClubUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClubUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClubUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := club.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Club.name": %w`, err)}
		}
	}
	return nil
}

func (cu *ClubUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   club.Table,
			Columns: club.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: club.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
	}
	if cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !cu.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{club.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ClubUpdateOne is the builder for updating a single Club entity.
type ClubUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClubMutation
}

// SetName sets the "name" field.
func (cuo *ClubUpdateOne) SetName(s string) *ClubUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cuo *ClubUpdateOne) AddMemberIDs(ids ...int) *ClubUpdateOne {
	cuo.mutation.AddMemberIDs(ids...)
	return cuo
}

// AddMembers adds the "members" edges to the User entity.
func (cuo *ClubUpdateOne) AddMembers(u ...*User) *ClubUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddMemberIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (cuo *ClubUpdateOne) AddInvitationIDs(ids ...uuid.UUID) *ClubUpdateOne {
	cuo.mutation.AddInvitationIDs(ids...)
	return cuo
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (cuo *ClubUpdateOne) AddInvitations(i ...*Invitation) *ClubUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddInvitationIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cuo *ClubUpdateOne) Mutation() *ClubMutation {
	return cuo.mutation
}

// ClearMembers clears all "members" edges to the User entity.
func (cuo *ClubUpdateOne) ClearMembers() *ClubUpdateOne {
	cuo.mutation.ClearMembers()
	return cuo
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cuo *ClubUpdateOne) RemoveMemberIDs(ids ...int) *ClubUpdateOne {
	cuo.mutation.RemoveMemberIDs(ids...)
	return cuo
}

// RemoveMembers removes "members" edges to User entities.
func (cuo *ClubUpdateOne) RemoveMembers(u ...*User) *ClubUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveMemberIDs(ids...)
}

// ClearInvitations clears all "invitations" edges to the Invitation entity.
func (cuo *ClubUpdateOne) ClearInvitations() *ClubUpdateOne {
	cuo.mutation.ClearInvitations()
	return cuo
}

// RemoveInvitationIDs removes the "invitations" edge to Invitation entities by IDs.
func (cuo *ClubUpdateOne) RemoveInvitationIDs(ids ...uuid.UUID) *ClubUpdateOne {
	cuo.mutation.RemoveInvitationIDs(ids...)
	return cuo
}

// RemoveInvitations removes "invitations" edges to Invitation entities.
func (cuo *ClubUpdateOne) RemoveInvitations(i ...*Invitation) *ClubUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveInvitationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClubUpdateOne) Select(field string, fields ...string) *ClubUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Club entity.
func (cuo *ClubUpdateOne) Save(ctx context.Context) (*Club, error) {
	var (
		err  error
		node *Club
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClubUpdateOne) SaveX(ctx context.Context) *Club {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClubUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClubUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClubUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := club.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Club.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *ClubUpdateOne) sqlSave(ctx context.Context) (_node *Club, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   club.Table,
			Columns: club.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: club.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Club.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, club.FieldID)
		for _, f := range fields {
			if !club.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != club.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
	}
	if cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   club.MembersTable,
			Columns: club.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedInvitationsIDs(); len(nodes) > 0 && !cuo.mutation.InvitationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   club.InvitationsTable,
			Columns: []string{club.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: invitation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Club{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{club.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
