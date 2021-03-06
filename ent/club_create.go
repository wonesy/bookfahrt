// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wonesy/bookfahrt/ent/club"
	"github.com/wonesy/bookfahrt/ent/invitation"
	"github.com/wonesy/bookfahrt/ent/user"
)

// ClubCreate is the builder for creating a Club entity.
type ClubCreate struct {
	config
	mutation *ClubMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (cc *ClubCreate) SetName(s string) *ClubCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetID sets the "id" field.
func (cc *ClubCreate) SetID(u uuid.UUID) *ClubCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ClubCreate) SetNillableID(u *uuid.UUID) *ClubCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cc *ClubCreate) AddMemberIDs(ids ...int) *ClubCreate {
	cc.mutation.AddMemberIDs(ids...)
	return cc
}

// AddMembers adds the "members" edges to the User entity.
func (cc *ClubCreate) AddMembers(u ...*User) *ClubCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddMemberIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (cc *ClubCreate) AddInvitationIDs(ids ...uuid.UUID) *ClubCreate {
	cc.mutation.AddInvitationIDs(ids...)
	return cc
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (cc *ClubCreate) AddInvitations(i ...*Invitation) *ClubCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cc.AddInvitationIDs(ids...)
}

// Mutation returns the ClubMutation object of the builder.
func (cc *ClubCreate) Mutation() *ClubMutation {
	return cc.mutation
}

// Save creates the Club in the database.
func (cc *ClubCreate) Save(ctx context.Context) (*Club, error) {
	var (
		err  error
		node *Club
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClubMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClubCreate) SaveX(ctx context.Context) *Club {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClubCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClubCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ClubCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := club.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClubCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Club.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := club.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Club.name": %w`, err)}
		}
	}
	return nil
}

func (cc *ClubCreate) sqlSave(ctx context.Context) (*Club, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (cc *ClubCreate) createSpec() (*Club, *sqlgraph.CreateSpec) {
	var (
		_node = &Club{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: club.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: club.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: club.FieldName,
		})
		_node.Name = value
	}
	if nodes := cc.mutation.MembersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.InvitationsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Club.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ClubUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (cc *ClubCreate) OnConflict(opts ...sql.ConflictOption) *ClubUpsertOne {
	cc.conflict = opts
	return &ClubUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Club.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *ClubCreate) OnConflictColumns(columns ...string) *ClubUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ClubUpsertOne{
		create: cc,
	}
}

type (
	// ClubUpsertOne is the builder for "upsert"-ing
	//  one Club node.
	ClubUpsertOne struct {
		create *ClubCreate
	}

	// ClubUpsert is the "OnConflict" setter.
	ClubUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ClubUpsert) SetName(v string) *ClubUpsert {
	u.Set(club.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClubUpsert) UpdateName() *ClubUpsert {
	u.SetExcluded(club.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Club.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(club.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ClubUpsertOne) UpdateNewValues() *ClubUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(club.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Club.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ClubUpsertOne) Ignore() *ClubUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ClubUpsertOne) DoNothing() *ClubUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ClubCreate.OnConflict
// documentation for more info.
func (u *ClubUpsertOne) Update(set func(*ClubUpsert)) *ClubUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ClubUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ClubUpsertOne) SetName(v string) *ClubUpsertOne {
	return u.Update(func(s *ClubUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClubUpsertOne) UpdateName() *ClubUpsertOne {
	return u.Update(func(s *ClubUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ClubUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ClubCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ClubUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ClubUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ClubUpsertOne.ID is not supported by MySQL driver. Use ClubUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ClubUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ClubCreateBulk is the builder for creating many Club entities in bulk.
type ClubCreateBulk struct {
	config
	builders []*ClubCreate
	conflict []sql.ConflictOption
}

// Save creates the Club entities in the database.
func (ccb *ClubCreateBulk) Save(ctx context.Context) ([]*Club, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Club, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClubMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ClubCreateBulk) SaveX(ctx context.Context) []*Club {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClubCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClubCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Club.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ClubUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *ClubCreateBulk) OnConflict(opts ...sql.ConflictOption) *ClubUpsertBulk {
	ccb.conflict = opts
	return &ClubUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Club.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *ClubCreateBulk) OnConflictColumns(columns ...string) *ClubUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ClubUpsertBulk{
		create: ccb,
	}
}

// ClubUpsertBulk is the builder for "upsert"-ing
// a bulk of Club nodes.
type ClubUpsertBulk struct {
	create *ClubCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Club.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(club.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ClubUpsertBulk) UpdateNewValues() *ClubUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(club.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Club.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ClubUpsertBulk) Ignore() *ClubUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ClubUpsertBulk) DoNothing() *ClubUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ClubCreateBulk.OnConflict
// documentation for more info.
func (u *ClubUpsertBulk) Update(set func(*ClubUpsert)) *ClubUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ClubUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ClubUpsertBulk) SetName(v string) *ClubUpsertBulk {
	return u.Update(func(s *ClubUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ClubUpsertBulk) UpdateName() *ClubUpsertBulk {
	return u.Update(func(s *ClubUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ClubUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ClubCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ClubCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ClubUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
