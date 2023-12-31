// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbipuser"
)

// CmdbIpUserCreate is the builder for creating a CmdbIpUser entity.
type CmdbIpUserCreate struct {
	config
	mutation *CmdbIpUserMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (ciuc *CmdbIpUserCreate) SetStatus(u uint8) *CmdbIpUserCreate {
	ciuc.mutation.SetStatus(u)
	return ciuc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ciuc *CmdbIpUserCreate) SetNillableStatus(u *uint8) *CmdbIpUserCreate {
	if u != nil {
		ciuc.SetStatus(*u)
	}
	return ciuc
}

// SetCreatedAt sets the "created_at" field.
func (ciuc *CmdbIpUserCreate) SetCreatedAt(t time.Time) *CmdbIpUserCreate {
	ciuc.mutation.SetCreatedAt(t)
	return ciuc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuc *CmdbIpUserCreate) SetNillableCreatedAt(t *time.Time) *CmdbIpUserCreate {
	if t != nil {
		ciuc.SetCreatedAt(*t)
	}
	return ciuc
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuc *CmdbIpUserCreate) SetUpdatedAt(t time.Time) *CmdbIpUserCreate {
	ciuc.mutation.SetUpdatedAt(t)
	return ciuc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ciuc *CmdbIpUserCreate) SetNillableUpdatedAt(t *time.Time) *CmdbIpUserCreate {
	if t != nil {
		ciuc.SetUpdatedAt(*t)
	}
	return ciuc
}

// SetUser sets the "user" field.
func (ciuc *CmdbIpUserCreate) SetUser(s string) *CmdbIpUserCreate {
	ciuc.mutation.SetUser(s)
	return ciuc
}

// SetDepartment sets the "department" field.
func (ciuc *CmdbIpUserCreate) SetDepartment(s string) *CmdbIpUserCreate {
	ciuc.mutation.SetDepartment(s)
	return ciuc
}

// SetMobile sets the "mobile" field.
func (ciuc *CmdbIpUserCreate) SetMobile(s string) *CmdbIpUserCreate {
	ciuc.mutation.SetMobile(s)
	return ciuc
}

// SetRemark sets the "remark" field.
func (ciuc *CmdbIpUserCreate) SetRemark(s string) *CmdbIpUserCreate {
	ciuc.mutation.SetRemark(s)
	return ciuc
}

// SetID sets the "id" field.
func (ciuc *CmdbIpUserCreate) SetID(u uint64) *CmdbIpUserCreate {
	ciuc.mutation.SetID(u)
	return ciuc
}

// Mutation returns the CmdbIpUserMutation object of the builder.
func (ciuc *CmdbIpUserCreate) Mutation() *CmdbIpUserMutation {
	return ciuc.mutation
}

// Save creates the CmdbIpUser in the database.
func (ciuc *CmdbIpUserCreate) Save(ctx context.Context) (*CmdbIpUser, error) {
	ciuc.defaults()
	return withHooks(ctx, ciuc.sqlSave, ciuc.mutation, ciuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ciuc *CmdbIpUserCreate) SaveX(ctx context.Context) *CmdbIpUser {
	v, err := ciuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ciuc *CmdbIpUserCreate) Exec(ctx context.Context) error {
	_, err := ciuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuc *CmdbIpUserCreate) ExecX(ctx context.Context) {
	if err := ciuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuc *CmdbIpUserCreate) defaults() {
	if _, ok := ciuc.mutation.Status(); !ok {
		v := cmdbipuser.DefaultStatus
		ciuc.mutation.SetStatus(v)
	}
	if _, ok := ciuc.mutation.CreatedAt(); !ok {
		v := cmdbipuser.DefaultCreatedAt()
		ciuc.mutation.SetCreatedAt(v)
	}
	if _, ok := ciuc.mutation.UpdatedAt(); !ok {
		v := cmdbipuser.DefaultUpdatedAt()
		ciuc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuc *CmdbIpUserCreate) check() error {
	if _, ok := ciuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CmdbIpUser.created_at"`)}
	}
	if _, ok := ciuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CmdbIpUser.updated_at"`)}
	}
	if _, ok := ciuc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "CmdbIpUser.user"`)}
	}
	if _, ok := ciuc.mutation.Department(); !ok {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required field "CmdbIpUser.department"`)}
	}
	if _, ok := ciuc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "CmdbIpUser.mobile"`)}
	}
	if v, ok := ciuc.mutation.Mobile(); ok {
		if err := cmdbipuser.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "CmdbIpUser.mobile": %w`, err)}
		}
	}
	if _, ok := ciuc.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "CmdbIpUser.remark"`)}
	}
	return nil
}

func (ciuc *CmdbIpUserCreate) sqlSave(ctx context.Context) (*CmdbIpUser, error) {
	if err := ciuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ciuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ciuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ciuc.mutation.id = &_node.ID
	ciuc.mutation.done = true
	return _node, nil
}

func (ciuc *CmdbIpUserCreate) createSpec() (*CmdbIpUser, *sqlgraph.CreateSpec) {
	var (
		_node = &CmdbIpUser{config: ciuc.config}
		_spec = sqlgraph.NewCreateSpec(cmdbipuser.Table, sqlgraph.NewFieldSpec(cmdbipuser.FieldID, field.TypeUint64))
	)
	if id, ok := ciuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ciuc.mutation.Status(); ok {
		_spec.SetField(cmdbipuser.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := ciuc.mutation.CreatedAt(); ok {
		_spec.SetField(cmdbipuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ciuc.mutation.UpdatedAt(); ok {
		_spec.SetField(cmdbipuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ciuc.mutation.User(); ok {
		_spec.SetField(cmdbipuser.FieldUser, field.TypeString, value)
		_node.User = value
	}
	if value, ok := ciuc.mutation.Department(); ok {
		_spec.SetField(cmdbipuser.FieldDepartment, field.TypeString, value)
		_node.Department = value
	}
	if value, ok := ciuc.mutation.Mobile(); ok {
		_spec.SetField(cmdbipuser.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := ciuc.mutation.Remark(); ok {
		_spec.SetField(cmdbipuser.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	return _node, _spec
}

// CmdbIpUserCreateBulk is the builder for creating many CmdbIpUser entities in bulk.
type CmdbIpUserCreateBulk struct {
	config
	err      error
	builders []*CmdbIpUserCreate
}

// Save creates the CmdbIpUser entities in the database.
func (ciucb *CmdbIpUserCreateBulk) Save(ctx context.Context) ([]*CmdbIpUser, error) {
	if ciucb.err != nil {
		return nil, ciucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ciucb.builders))
	nodes := make([]*CmdbIpUser, len(ciucb.builders))
	mutators := make([]Mutator, len(ciucb.builders))
	for i := range ciucb.builders {
		func(i int, root context.Context) {
			builder := ciucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CmdbIpUserMutation)
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
					_, err = mutators[i+1].Mutate(root, ciucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ciucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ciucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ciucb *CmdbIpUserCreateBulk) SaveX(ctx context.Context) []*CmdbIpUser {
	v, err := ciucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ciucb *CmdbIpUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ciucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciucb *CmdbIpUserCreateBulk) ExecX(ctx context.Context) {
	if err := ciucb.Exec(ctx); err != nil {
		panic(err)
	}
}
