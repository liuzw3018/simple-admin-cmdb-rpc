package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// CmdbAssets holds the schema definition for the CmdbAssets entity.
type CmdbAssets struct {
	ent.Schema
}

// Fields of the Cmdb.
func (CmdbAssets) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Comment("IP").Unique(),
		field.String("mask").Comment("掩码"),
		field.String("gateway").Comment("网关"),
		field.Time("online_time").Comment("上线时间"),
		field.Time("offline_time").Comment("下线时间"),
		field.Uint("power_status").Comment("电源状态 | 1: 开机 | 0: 关机").Default(1).Annotations(entsql.WithComments(true)),
		field.Uint("is_server").Comment("是否服务器 | 1: 是 | 0: 否").Default(1).Annotations(entsql.WithComments(true)),
		field.Uint("server_type").Comment("服务器类型 | 1: 物理机 | 2: 虚拟机 | 3: 容器").Default(1).Annotations(entsql.WithComments(true)),
		field.String("server_hostname").Comment("服务器主机名").Default("localhost").Annotations(entsql.WithComments(true)),
		field.String("server_os").Comment("服务器操作系统").Default("Linux").Annotations(entsql.WithComments(true)),
		field.String("server_os_version").Comment("服务器操作系统版本").Default("CentOS 7.6.1810").Annotations(entsql.WithComments(true)),
		field.String("server_os_arch").Comment("服务器操作系统架构").Default("AMD64").Annotations(entsql.WithComments(true)),
		field.String("cpu").Comment("cpu").Annotations(entsql.WithComments(true)),
		field.String("memory").Comment("内存").Annotations(entsql.WithComments(true)),
		field.String("disk").Comment("硬盘").Annotations(entsql.WithComments(true)),
		field.String("NetworkSpeed").Comment("网卡速率 | 100M | 1G | 10G | 100G").Default("1G").Annotations(entsql.WithComments(true)),
		field.String("device_address").Comment("设备位置").Default("机房").Annotations(entsql.WithComments(true)),
	}
}

func (CmdbAssets) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CmdbIpUser{},
	}
}

// Edges of the Cmdb.
func (CmdbAssets) Edges() []ent.Edge {
	return nil
}
