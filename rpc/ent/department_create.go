// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/neumathe/simple-admin-core/rpc/ent/department"
	"github.com/neumathe/simple-admin-core/rpc/ent/user"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DepartmentCreate) SetCreatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DepartmentCreate) SetUpdatedAt(t time.Time) *DepartmentCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdatedAt(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DepartmentCreate) SetStatus(u uint8) *DepartmentCreate {
	dc.mutation.SetStatus(u)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableStatus(u *uint8) *DepartmentCreate {
	if u != nil {
		dc.SetStatus(*u)
	}
	return dc
}

// SetSort sets the "sort" field.
func (dc *DepartmentCreate) SetSort(u uint32) *DepartmentCreate {
	dc.mutation.SetSort(u)
	return dc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableSort(u *uint32) *DepartmentCreate {
	if u != nil {
		dc.SetSort(*u)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetAncestors sets the "ancestors" field.
func (dc *DepartmentCreate) SetAncestors(s string) *DepartmentCreate {
	dc.mutation.SetAncestors(s)
	return dc
}

// SetNillableAncestors sets the "ancestors" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableAncestors(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetAncestors(*s)
	}
	return dc
}

// SetLeader sets the "leader" field.
func (dc *DepartmentCreate) SetLeader(s string) *DepartmentCreate {
	dc.mutation.SetLeader(s)
	return dc
}

// SetNillableLeader sets the "leader" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableLeader(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetLeader(*s)
	}
	return dc
}

// SetPhone sets the "phone" field.
func (dc *DepartmentCreate) SetPhone(s string) *DepartmentCreate {
	dc.mutation.SetPhone(s)
	return dc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillablePhone(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetPhone(*s)
	}
	return dc
}

// SetEmail sets the "email" field.
func (dc *DepartmentCreate) SetEmail(s string) *DepartmentCreate {
	dc.mutation.SetEmail(s)
	return dc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableEmail(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetEmail(*s)
	}
	return dc
}

// SetRemark sets the "remark" field.
func (dc *DepartmentCreate) SetRemark(s string) *DepartmentCreate {
	dc.mutation.SetRemark(s)
	return dc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableRemark(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetRemark(*s)
	}
	return dc
}

// SetParentID sets the "parent_id" field.
func (dc *DepartmentCreate) SetParentID(u uint64) *DepartmentCreate {
	dc.mutation.SetParentID(u)
	return dc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableParentID(u *uint64) *DepartmentCreate {
	if u != nil {
		dc.SetParentID(*u)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DepartmentCreate) SetID(u uint64) *DepartmentCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetParent sets the "parent" edge to the Department entity.
func (dc *DepartmentCreate) SetParent(d *Department) *DepartmentCreate {
	return dc.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Department entity by IDs.
func (dc *DepartmentCreate) AddChildIDs(ids ...uint64) *DepartmentCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Department entity.
func (dc *DepartmentCreate) AddChildren(d ...*Department) *DepartmentCreate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DepartmentCreate) AddUserIDs(ids ...uuid.UUID) *DepartmentCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DepartmentCreate) AddUsers(u ...*User) *DepartmentCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := department.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := department.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := department.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Sort(); !ok {
		v := department.DefaultSort
		dc.mutation.SetSort(v)
	}
	if _, ok := dc.mutation.ParentID(); !ok {
		v := department.DefaultParentID
		dc.mutation.SetParentID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Department.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Department.updated_at"`)}
	}
	if _, ok := dc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Department.sort"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Department.name"`)}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(department.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(department.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(department.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Sort(); ok {
		_spec.SetField(department.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Ancestors(); ok {
		_spec.SetField(department.FieldAncestors, field.TypeString, value)
		_node.Ancestors = value
	}
	if value, ok := dc.mutation.Leader(); ok {
		_spec.SetField(department.FieldLeader, field.TypeString, value)
		_node.Leader = value
	}
	if value, ok := dc.mutation.Phone(); ok {
		_spec.SetField(department.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := dc.mutation.Email(); ok {
		_spec.SetField(department.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := dc.mutation.Remark(); ok {
		_spec.SetField(department.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	err      error
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
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
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
