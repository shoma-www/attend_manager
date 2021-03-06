// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/rs/xid"
	"github.com/shoma-www/attend_manager/grpc/ent/attendancegroup"
	"github.com/shoma-www/attend_manager/grpc/ent/predicate"
	"github.com/shoma-www/attend_manager/grpc/ent/user"
)

// AttendanceGroupUpdate is the builder for updating AttendanceGroup entities.
type AttendanceGroupUpdate struct {
	config
	hooks      []Hook
	mutation   *AttendanceGroupMutation
	predicates []predicate.AttendanceGroup
}

// Where adds a new predicate for the builder.
func (agu *AttendanceGroupUpdate) Where(ps ...predicate.AttendanceGroup) *AttendanceGroupUpdate {
	agu.predicates = append(agu.predicates, ps...)
	return agu
}

// SetName sets the Name field.
func (agu *AttendanceGroupUpdate) SetName(s string) *AttendanceGroupUpdate {
	agu.mutation.SetName(s)
	return agu
}

// SetNillableName sets the Name field if the given value is not nil.
func (agu *AttendanceGroupUpdate) SetNillableName(s *string) *AttendanceGroupUpdate {
	if s != nil {
		agu.SetName(*s)
	}
	return agu
}

// ClearName clears the value of Name.
func (agu *AttendanceGroupUpdate) ClearName() *AttendanceGroupUpdate {
	agu.mutation.ClearName()
	return agu
}

// SetCreatedAt sets the CreatedAt field.
func (agu *AttendanceGroupUpdate) SetCreatedAt(t time.Time) *AttendanceGroupUpdate {
	agu.mutation.SetCreatedAt(t)
	return agu
}

// SetNillableCreatedAt sets the CreatedAt field if the given value is not nil.
func (agu *AttendanceGroupUpdate) SetNillableCreatedAt(t *time.Time) *AttendanceGroupUpdate {
	if t != nil {
		agu.SetCreatedAt(*t)
	}
	return agu
}

// ClearCreatedAt clears the value of CreatedAt.
func (agu *AttendanceGroupUpdate) ClearCreatedAt() *AttendanceGroupUpdate {
	agu.mutation.ClearCreatedAt()
	return agu
}

// SetUpdatedAt sets the UpdatedAt field.
func (agu *AttendanceGroupUpdate) SetUpdatedAt(t time.Time) *AttendanceGroupUpdate {
	agu.mutation.SetUpdatedAt(t)
	return agu
}

// SetNillableUpdatedAt sets the UpdatedAt field if the given value is not nil.
func (agu *AttendanceGroupUpdate) SetNillableUpdatedAt(t *time.Time) *AttendanceGroupUpdate {
	if t != nil {
		agu.SetUpdatedAt(*t)
	}
	return agu
}

// ClearUpdatedAt clears the value of UpdatedAt.
func (agu *AttendanceGroupUpdate) ClearUpdatedAt() *AttendanceGroupUpdate {
	agu.mutation.ClearUpdatedAt()
	return agu
}

// AddUserIDs adds the users edge to User by ids.
func (agu *AttendanceGroupUpdate) AddUserIDs(ids ...xid.ID) *AttendanceGroupUpdate {
	agu.mutation.AddUserIDs(ids...)
	return agu
}

// AddUsers adds the users edges to User.
func (agu *AttendanceGroupUpdate) AddUsers(u ...*User) *AttendanceGroupUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return agu.AddUserIDs(ids...)
}

// Mutation returns the AttendanceGroupMutation object of the builder.
func (agu *AttendanceGroupUpdate) Mutation() *AttendanceGroupMutation {
	return agu.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (agu *AttendanceGroupUpdate) RemoveUserIDs(ids ...xid.ID) *AttendanceGroupUpdate {
	agu.mutation.RemoveUserIDs(ids...)
	return agu
}

// RemoveUsers removes users edges to User.
func (agu *AttendanceGroupUpdate) RemoveUsers(u ...*User) *AttendanceGroupUpdate {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return agu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (agu *AttendanceGroupUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := agu.mutation.Name(); ok {
		if err := attendancegroup.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(agu.hooks) == 0 {
		affected, err = agu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agu.mutation = mutation
			affected, err = agu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agu.hooks) - 1; i >= 0; i-- {
			mut = agu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (agu *AttendanceGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := agu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (agu *AttendanceGroupUpdate) Exec(ctx context.Context) error {
	_, err := agu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agu *AttendanceGroupUpdate) ExecX(ctx context.Context) {
	if err := agu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (agu *AttendanceGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendancegroup.Table,
			Columns: attendancegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attendancegroup.FieldID,
			},
		},
	}
	if ps := agu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attendancegroup.FieldName,
		})
	}
	if agu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: attendancegroup.FieldName,
		})
	}
	if value, ok := agu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendancegroup.FieldCreatedAt,
		})
	}
	if agu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: attendancegroup.FieldCreatedAt,
		})
	}
	if value, ok := agu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendancegroup.FieldUpdatedAt,
		})
	}
	if agu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: attendancegroup.FieldUpdatedAt,
		})
	}
	if nodes := agu.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendancegroup.UsersTable,
			Columns: []string{attendancegroup.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := agu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendancegroup.UsersTable,
			Columns: []string{attendancegroup.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, agu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendancegroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AttendanceGroupUpdateOne is the builder for updating a single AttendanceGroup entity.
type AttendanceGroupUpdateOne struct {
	config
	hooks    []Hook
	mutation *AttendanceGroupMutation
}

// SetName sets the Name field.
func (aguo *AttendanceGroupUpdateOne) SetName(s string) *AttendanceGroupUpdateOne {
	aguo.mutation.SetName(s)
	return aguo
}

// SetNillableName sets the Name field if the given value is not nil.
func (aguo *AttendanceGroupUpdateOne) SetNillableName(s *string) *AttendanceGroupUpdateOne {
	if s != nil {
		aguo.SetName(*s)
	}
	return aguo
}

// ClearName clears the value of Name.
func (aguo *AttendanceGroupUpdateOne) ClearName() *AttendanceGroupUpdateOne {
	aguo.mutation.ClearName()
	return aguo
}

// SetCreatedAt sets the CreatedAt field.
func (aguo *AttendanceGroupUpdateOne) SetCreatedAt(t time.Time) *AttendanceGroupUpdateOne {
	aguo.mutation.SetCreatedAt(t)
	return aguo
}

// SetNillableCreatedAt sets the CreatedAt field if the given value is not nil.
func (aguo *AttendanceGroupUpdateOne) SetNillableCreatedAt(t *time.Time) *AttendanceGroupUpdateOne {
	if t != nil {
		aguo.SetCreatedAt(*t)
	}
	return aguo
}

// ClearCreatedAt clears the value of CreatedAt.
func (aguo *AttendanceGroupUpdateOne) ClearCreatedAt() *AttendanceGroupUpdateOne {
	aguo.mutation.ClearCreatedAt()
	return aguo
}

// SetUpdatedAt sets the UpdatedAt field.
func (aguo *AttendanceGroupUpdateOne) SetUpdatedAt(t time.Time) *AttendanceGroupUpdateOne {
	aguo.mutation.SetUpdatedAt(t)
	return aguo
}

// SetNillableUpdatedAt sets the UpdatedAt field if the given value is not nil.
func (aguo *AttendanceGroupUpdateOne) SetNillableUpdatedAt(t *time.Time) *AttendanceGroupUpdateOne {
	if t != nil {
		aguo.SetUpdatedAt(*t)
	}
	return aguo
}

// ClearUpdatedAt clears the value of UpdatedAt.
func (aguo *AttendanceGroupUpdateOne) ClearUpdatedAt() *AttendanceGroupUpdateOne {
	aguo.mutation.ClearUpdatedAt()
	return aguo
}

// AddUserIDs adds the users edge to User by ids.
func (aguo *AttendanceGroupUpdateOne) AddUserIDs(ids ...xid.ID) *AttendanceGroupUpdateOne {
	aguo.mutation.AddUserIDs(ids...)
	return aguo
}

// AddUsers adds the users edges to User.
func (aguo *AttendanceGroupUpdateOne) AddUsers(u ...*User) *AttendanceGroupUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return aguo.AddUserIDs(ids...)
}

// Mutation returns the AttendanceGroupMutation object of the builder.
func (aguo *AttendanceGroupUpdateOne) Mutation() *AttendanceGroupMutation {
	return aguo.mutation
}

// RemoveUserIDs removes the users edge to User by ids.
func (aguo *AttendanceGroupUpdateOne) RemoveUserIDs(ids ...xid.ID) *AttendanceGroupUpdateOne {
	aguo.mutation.RemoveUserIDs(ids...)
	return aguo
}

// RemoveUsers removes users edges to User.
func (aguo *AttendanceGroupUpdateOne) RemoveUsers(u ...*User) *AttendanceGroupUpdateOne {
	ids := make([]xid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return aguo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (aguo *AttendanceGroupUpdateOne) Save(ctx context.Context) (*AttendanceGroup, error) {
	if v, ok := aguo.mutation.Name(); ok {
		if err := attendancegroup.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err  error
		node *AttendanceGroup
	)
	if len(aguo.hooks) == 0 {
		node, err = aguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttendanceGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aguo.mutation = mutation
			node, err = aguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aguo.hooks) - 1; i >= 0; i-- {
			mut = aguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aguo *AttendanceGroupUpdateOne) SaveX(ctx context.Context) *AttendanceGroup {
	ag, err := aguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ag
}

// Exec executes the query on the entity.
func (aguo *AttendanceGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := aguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aguo *AttendanceGroupUpdateOne) ExecX(ctx context.Context) {
	if err := aguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aguo *AttendanceGroupUpdateOne) sqlSave(ctx context.Context) (ag *AttendanceGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attendancegroup.Table,
			Columns: attendancegroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attendancegroup.FieldID,
			},
		},
	}
	id, ok := aguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AttendanceGroup.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := aguo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: attendancegroup.FieldName,
		})
	}
	if aguo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: attendancegroup.FieldName,
		})
	}
	if value, ok := aguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendancegroup.FieldCreatedAt,
		})
	}
	if aguo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: attendancegroup.FieldCreatedAt,
		})
	}
	if value, ok := aguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: attendancegroup.FieldUpdatedAt,
		})
	}
	if aguo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: attendancegroup.FieldUpdatedAt,
		})
	}
	if nodes := aguo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendancegroup.UsersTable,
			Columns: []string{attendancegroup.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aguo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attendancegroup.UsersTable,
			Columns: []string{attendancegroup.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ag = &AttendanceGroup{config: aguo.config}
	_spec.Assign = ag.assignValues
	_spec.ScanValues = ag.scanValues()
	if err = sqlgraph.UpdateNode(ctx, aguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendancegroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ag, nil
}
