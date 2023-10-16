// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/liuzw3018/simple-admin-cmdb-rpc/ent/cmdbassets"
)

// CmdbAssetsCreate is the builder for creating a CmdbAssets entity.
type CmdbAssetsCreate struct {
	config
	mutation *CmdbAssetsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cac *CmdbAssetsCreate) SetCreatedAt(t time.Time) *CmdbAssetsCreate {
	cac.mutation.SetCreatedAt(t)
	return cac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableCreatedAt(t *time.Time) *CmdbAssetsCreate {
	if t != nil {
		cac.SetCreatedAt(*t)
	}
	return cac
}

// SetUpdatedAt sets the "updated_at" field.
func (cac *CmdbAssetsCreate) SetUpdatedAt(t time.Time) *CmdbAssetsCreate {
	cac.mutation.SetUpdatedAt(t)
	return cac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableUpdatedAt(t *time.Time) *CmdbAssetsCreate {
	if t != nil {
		cac.SetUpdatedAt(*t)
	}
	return cac
}

// SetUser sets the "user" field.
func (cac *CmdbAssetsCreate) SetUser(s string) *CmdbAssetsCreate {
	cac.mutation.SetUser(s)
	return cac
}

// SetDepartment sets the "department" field.
func (cac *CmdbAssetsCreate) SetDepartment(s string) *CmdbAssetsCreate {
	cac.mutation.SetDepartment(s)
	return cac
}

// SetMobile sets the "mobile" field.
func (cac *CmdbAssetsCreate) SetMobile(s string) *CmdbAssetsCreate {
	cac.mutation.SetMobile(s)
	return cac
}

// SetRemark sets the "remark" field.
func (cac *CmdbAssetsCreate) SetRemark(s string) *CmdbAssetsCreate {
	cac.mutation.SetRemark(s)
	return cac
}

// SetIP sets the "ip" field.
func (cac *CmdbAssetsCreate) SetIP(s string) *CmdbAssetsCreate {
	cac.mutation.SetIP(s)
	return cac
}

// SetMask sets the "mask" field.
func (cac *CmdbAssetsCreate) SetMask(s string) *CmdbAssetsCreate {
	cac.mutation.SetMask(s)
	return cac
}

// SetGateway sets the "gateway" field.
func (cac *CmdbAssetsCreate) SetGateway(s string) *CmdbAssetsCreate {
	cac.mutation.SetGateway(s)
	return cac
}

// SetOnlineTime sets the "online_time" field.
func (cac *CmdbAssetsCreate) SetOnlineTime(t time.Time) *CmdbAssetsCreate {
	cac.mutation.SetOnlineTime(t)
	return cac
}

// SetOfflineTime sets the "offline_time" field.
func (cac *CmdbAssetsCreate) SetOfflineTime(t time.Time) *CmdbAssetsCreate {
	cac.mutation.SetOfflineTime(t)
	return cac
}

// SetPowerStatus sets the "power_status" field.
func (cac *CmdbAssetsCreate) SetPowerStatus(u uint) *CmdbAssetsCreate {
	cac.mutation.SetPowerStatus(u)
	return cac
}

// SetNillablePowerStatus sets the "power_status" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillablePowerStatus(u *uint) *CmdbAssetsCreate {
	if u != nil {
		cac.SetPowerStatus(*u)
	}
	return cac
}

// SetIsServer sets the "is_server" field.
func (cac *CmdbAssetsCreate) SetIsServer(u uint) *CmdbAssetsCreate {
	cac.mutation.SetIsServer(u)
	return cac
}

// SetNillableIsServer sets the "is_server" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableIsServer(u *uint) *CmdbAssetsCreate {
	if u != nil {
		cac.SetIsServer(*u)
	}
	return cac
}

// SetServerType sets the "server_type" field.
func (cac *CmdbAssetsCreate) SetServerType(u uint) *CmdbAssetsCreate {
	cac.mutation.SetServerType(u)
	return cac
}

// SetNillableServerType sets the "server_type" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableServerType(u *uint) *CmdbAssetsCreate {
	if u != nil {
		cac.SetServerType(*u)
	}
	return cac
}

// SetServerHostname sets the "server_hostname" field.
func (cac *CmdbAssetsCreate) SetServerHostname(s string) *CmdbAssetsCreate {
	cac.mutation.SetServerHostname(s)
	return cac
}

// SetNillableServerHostname sets the "server_hostname" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableServerHostname(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetServerHostname(*s)
	}
	return cac
}

// SetServerOs sets the "server_os" field.
func (cac *CmdbAssetsCreate) SetServerOs(s string) *CmdbAssetsCreate {
	cac.mutation.SetServerOs(s)
	return cac
}

// SetNillableServerOs sets the "server_os" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableServerOs(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetServerOs(*s)
	}
	return cac
}

// SetServerOsVersion sets the "server_os_version" field.
func (cac *CmdbAssetsCreate) SetServerOsVersion(s string) *CmdbAssetsCreate {
	cac.mutation.SetServerOsVersion(s)
	return cac
}

// SetNillableServerOsVersion sets the "server_os_version" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableServerOsVersion(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetServerOsVersion(*s)
	}
	return cac
}

// SetServerOsArch sets the "server_os_arch" field.
func (cac *CmdbAssetsCreate) SetServerOsArch(s string) *CmdbAssetsCreate {
	cac.mutation.SetServerOsArch(s)
	return cac
}

// SetNillableServerOsArch sets the "server_os_arch" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableServerOsArch(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetServerOsArch(*s)
	}
	return cac
}

// SetCPU sets the "cpu" field.
func (cac *CmdbAssetsCreate) SetCPU(s string) *CmdbAssetsCreate {
	cac.mutation.SetCPU(s)
	return cac
}

// SetMemory sets the "memory" field.
func (cac *CmdbAssetsCreate) SetMemory(s string) *CmdbAssetsCreate {
	cac.mutation.SetMemory(s)
	return cac
}

// SetDisk sets the "disk" field.
func (cac *CmdbAssetsCreate) SetDisk(s string) *CmdbAssetsCreate {
	cac.mutation.SetDisk(s)
	return cac
}

// SetNetworkSpeed sets the "NetworkSpeed" field.
func (cac *CmdbAssetsCreate) SetNetworkSpeed(s string) *CmdbAssetsCreate {
	cac.mutation.SetNetworkSpeed(s)
	return cac
}

// SetNillableNetworkSpeed sets the "NetworkSpeed" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableNetworkSpeed(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetNetworkSpeed(*s)
	}
	return cac
}

// SetDeviceAddress sets the "device_address" field.
func (cac *CmdbAssetsCreate) SetDeviceAddress(s string) *CmdbAssetsCreate {
	cac.mutation.SetDeviceAddress(s)
	return cac
}

// SetNillableDeviceAddress sets the "device_address" field if the given value is not nil.
func (cac *CmdbAssetsCreate) SetNillableDeviceAddress(s *string) *CmdbAssetsCreate {
	if s != nil {
		cac.SetDeviceAddress(*s)
	}
	return cac
}

// SetID sets the "id" field.
func (cac *CmdbAssetsCreate) SetID(u uint64) *CmdbAssetsCreate {
	cac.mutation.SetID(u)
	return cac
}

// Mutation returns the CmdbAssetsMutation object of the builder.
func (cac *CmdbAssetsCreate) Mutation() *CmdbAssetsMutation {
	return cac.mutation
}

// Save creates the CmdbAssets in the database.
func (cac *CmdbAssetsCreate) Save(ctx context.Context) (*CmdbAssets, error) {
	cac.defaults()
	return withHooks(ctx, cac.sqlSave, cac.mutation, cac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cac *CmdbAssetsCreate) SaveX(ctx context.Context) *CmdbAssets {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cac *CmdbAssetsCreate) Exec(ctx context.Context) error {
	_, err := cac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cac *CmdbAssetsCreate) ExecX(ctx context.Context) {
	if err := cac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cac *CmdbAssetsCreate) defaults() {
	if _, ok := cac.mutation.CreatedAt(); !ok {
		v := cmdbassets.DefaultCreatedAt()
		cac.mutation.SetCreatedAt(v)
	}
	if _, ok := cac.mutation.UpdatedAt(); !ok {
		v := cmdbassets.DefaultUpdatedAt()
		cac.mutation.SetUpdatedAt(v)
	}
	if _, ok := cac.mutation.PowerStatus(); !ok {
		v := cmdbassets.DefaultPowerStatus
		cac.mutation.SetPowerStatus(v)
	}
	if _, ok := cac.mutation.IsServer(); !ok {
		v := cmdbassets.DefaultIsServer
		cac.mutation.SetIsServer(v)
	}
	if _, ok := cac.mutation.ServerType(); !ok {
		v := cmdbassets.DefaultServerType
		cac.mutation.SetServerType(v)
	}
	if _, ok := cac.mutation.ServerHostname(); !ok {
		v := cmdbassets.DefaultServerHostname
		cac.mutation.SetServerHostname(v)
	}
	if _, ok := cac.mutation.ServerOs(); !ok {
		v := cmdbassets.DefaultServerOs
		cac.mutation.SetServerOs(v)
	}
	if _, ok := cac.mutation.ServerOsVersion(); !ok {
		v := cmdbassets.DefaultServerOsVersion
		cac.mutation.SetServerOsVersion(v)
	}
	if _, ok := cac.mutation.ServerOsArch(); !ok {
		v := cmdbassets.DefaultServerOsArch
		cac.mutation.SetServerOsArch(v)
	}
	if _, ok := cac.mutation.NetworkSpeed(); !ok {
		v := cmdbassets.DefaultNetworkSpeed
		cac.mutation.SetNetworkSpeed(v)
	}
	if _, ok := cac.mutation.DeviceAddress(); !ok {
		v := cmdbassets.DefaultDeviceAddress
		cac.mutation.SetDeviceAddress(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *CmdbAssetsCreate) check() error {
	if _, ok := cac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CmdbAssets.created_at"`)}
	}
	if _, ok := cac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CmdbAssets.updated_at"`)}
	}
	if _, ok := cac.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "CmdbAssets.user"`)}
	}
	if _, ok := cac.mutation.Department(); !ok {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required field "CmdbAssets.department"`)}
	}
	if _, ok := cac.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "CmdbAssets.mobile"`)}
	}
	if v, ok := cac.mutation.Mobile(); ok {
		if err := cmdbassets.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "CmdbAssets.mobile": %w`, err)}
		}
	}
	if _, ok := cac.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "CmdbAssets.remark"`)}
	}
	if _, ok := cac.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "CmdbAssets.ip"`)}
	}
	if _, ok := cac.mutation.Mask(); !ok {
		return &ValidationError{Name: "mask", err: errors.New(`ent: missing required field "CmdbAssets.mask"`)}
	}
	if _, ok := cac.mutation.Gateway(); !ok {
		return &ValidationError{Name: "gateway", err: errors.New(`ent: missing required field "CmdbAssets.gateway"`)}
	}
	if _, ok := cac.mutation.OnlineTime(); !ok {
		return &ValidationError{Name: "online_time", err: errors.New(`ent: missing required field "CmdbAssets.online_time"`)}
	}
	if _, ok := cac.mutation.OfflineTime(); !ok {
		return &ValidationError{Name: "offline_time", err: errors.New(`ent: missing required field "CmdbAssets.offline_time"`)}
	}
	if _, ok := cac.mutation.PowerStatus(); !ok {
		return &ValidationError{Name: "power_status", err: errors.New(`ent: missing required field "CmdbAssets.power_status"`)}
	}
	if _, ok := cac.mutation.IsServer(); !ok {
		return &ValidationError{Name: "is_server", err: errors.New(`ent: missing required field "CmdbAssets.is_server"`)}
	}
	if _, ok := cac.mutation.ServerType(); !ok {
		return &ValidationError{Name: "server_type", err: errors.New(`ent: missing required field "CmdbAssets.server_type"`)}
	}
	if _, ok := cac.mutation.ServerHostname(); !ok {
		return &ValidationError{Name: "server_hostname", err: errors.New(`ent: missing required field "CmdbAssets.server_hostname"`)}
	}
	if _, ok := cac.mutation.ServerOs(); !ok {
		return &ValidationError{Name: "server_os", err: errors.New(`ent: missing required field "CmdbAssets.server_os"`)}
	}
	if _, ok := cac.mutation.ServerOsVersion(); !ok {
		return &ValidationError{Name: "server_os_version", err: errors.New(`ent: missing required field "CmdbAssets.server_os_version"`)}
	}
	if _, ok := cac.mutation.ServerOsArch(); !ok {
		return &ValidationError{Name: "server_os_arch", err: errors.New(`ent: missing required field "CmdbAssets.server_os_arch"`)}
	}
	if _, ok := cac.mutation.CPU(); !ok {
		return &ValidationError{Name: "cpu", err: errors.New(`ent: missing required field "CmdbAssets.cpu"`)}
	}
	if _, ok := cac.mutation.Memory(); !ok {
		return &ValidationError{Name: "memory", err: errors.New(`ent: missing required field "CmdbAssets.memory"`)}
	}
	if _, ok := cac.mutation.Disk(); !ok {
		return &ValidationError{Name: "disk", err: errors.New(`ent: missing required field "CmdbAssets.disk"`)}
	}
	if _, ok := cac.mutation.NetworkSpeed(); !ok {
		return &ValidationError{Name: "NetworkSpeed", err: errors.New(`ent: missing required field "CmdbAssets.NetworkSpeed"`)}
	}
	if _, ok := cac.mutation.DeviceAddress(); !ok {
		return &ValidationError{Name: "device_address", err: errors.New(`ent: missing required field "CmdbAssets.device_address"`)}
	}
	return nil
}

func (cac *CmdbAssetsCreate) sqlSave(ctx context.Context) (*CmdbAssets, error) {
	if err := cac.check(); err != nil {
		return nil, err
	}
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	cac.mutation.id = &_node.ID
	cac.mutation.done = true
	return _node, nil
}

func (cac *CmdbAssetsCreate) createSpec() (*CmdbAssets, *sqlgraph.CreateSpec) {
	var (
		_node = &CmdbAssets{config: cac.config}
		_spec = sqlgraph.NewCreateSpec(cmdbassets.Table, sqlgraph.NewFieldSpec(cmdbassets.FieldID, field.TypeUint64))
	)
	if id, ok := cac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cac.mutation.CreatedAt(); ok {
		_spec.SetField(cmdbassets.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cac.mutation.UpdatedAt(); ok {
		_spec.SetField(cmdbassets.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cac.mutation.User(); ok {
		_spec.SetField(cmdbassets.FieldUser, field.TypeString, value)
		_node.User = value
	}
	if value, ok := cac.mutation.Department(); ok {
		_spec.SetField(cmdbassets.FieldDepartment, field.TypeString, value)
		_node.Department = value
	}
	if value, ok := cac.mutation.Mobile(); ok {
		_spec.SetField(cmdbassets.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := cac.mutation.Remark(); ok {
		_spec.SetField(cmdbassets.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := cac.mutation.IP(); ok {
		_spec.SetField(cmdbassets.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := cac.mutation.Mask(); ok {
		_spec.SetField(cmdbassets.FieldMask, field.TypeString, value)
		_node.Mask = value
	}
	if value, ok := cac.mutation.Gateway(); ok {
		_spec.SetField(cmdbassets.FieldGateway, field.TypeString, value)
		_node.Gateway = value
	}
	if value, ok := cac.mutation.OnlineTime(); ok {
		_spec.SetField(cmdbassets.FieldOnlineTime, field.TypeTime, value)
		_node.OnlineTime = value
	}
	if value, ok := cac.mutation.OfflineTime(); ok {
		_spec.SetField(cmdbassets.FieldOfflineTime, field.TypeTime, value)
		_node.OfflineTime = value
	}
	if value, ok := cac.mutation.PowerStatus(); ok {
		_spec.SetField(cmdbassets.FieldPowerStatus, field.TypeUint, value)
		_node.PowerStatus = value
	}
	if value, ok := cac.mutation.IsServer(); ok {
		_spec.SetField(cmdbassets.FieldIsServer, field.TypeUint, value)
		_node.IsServer = value
	}
	if value, ok := cac.mutation.ServerType(); ok {
		_spec.SetField(cmdbassets.FieldServerType, field.TypeUint, value)
		_node.ServerType = value
	}
	if value, ok := cac.mutation.ServerHostname(); ok {
		_spec.SetField(cmdbassets.FieldServerHostname, field.TypeString, value)
		_node.ServerHostname = value
	}
	if value, ok := cac.mutation.ServerOs(); ok {
		_spec.SetField(cmdbassets.FieldServerOs, field.TypeString, value)
		_node.ServerOs = value
	}
	if value, ok := cac.mutation.ServerOsVersion(); ok {
		_spec.SetField(cmdbassets.FieldServerOsVersion, field.TypeString, value)
		_node.ServerOsVersion = value
	}
	if value, ok := cac.mutation.ServerOsArch(); ok {
		_spec.SetField(cmdbassets.FieldServerOsArch, field.TypeString, value)
		_node.ServerOsArch = value
	}
	if value, ok := cac.mutation.CPU(); ok {
		_spec.SetField(cmdbassets.FieldCPU, field.TypeString, value)
		_node.CPU = value
	}
	if value, ok := cac.mutation.Memory(); ok {
		_spec.SetField(cmdbassets.FieldMemory, field.TypeString, value)
		_node.Memory = value
	}
	if value, ok := cac.mutation.Disk(); ok {
		_spec.SetField(cmdbassets.FieldDisk, field.TypeString, value)
		_node.Disk = value
	}
	if value, ok := cac.mutation.NetworkSpeed(); ok {
		_spec.SetField(cmdbassets.FieldNetworkSpeed, field.TypeString, value)
		_node.NetworkSpeed = value
	}
	if value, ok := cac.mutation.DeviceAddress(); ok {
		_spec.SetField(cmdbassets.FieldDeviceAddress, field.TypeString, value)
		_node.DeviceAddress = value
	}
	return _node, _spec
}

// CmdbAssetsCreateBulk is the builder for creating many CmdbAssets entities in bulk.
type CmdbAssetsCreateBulk struct {
	config
	err      error
	builders []*CmdbAssetsCreate
}

// Save creates the CmdbAssets entities in the database.
func (cacb *CmdbAssetsCreateBulk) Save(ctx context.Context) ([]*CmdbAssets, error) {
	if cacb.err != nil {
		return nil, cacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*CmdbAssets, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CmdbAssetsMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *CmdbAssetsCreateBulk) SaveX(ctx context.Context) []*CmdbAssets {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cacb *CmdbAssetsCreateBulk) Exec(ctx context.Context) error {
	_, err := cacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cacb *CmdbAssetsCreateBulk) ExecX(ctx context.Context) {
	if err := cacb.Exec(ctx); err != nil {
		panic(err)
	}
}
