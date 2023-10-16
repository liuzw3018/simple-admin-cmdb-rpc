package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// CmdbIp holds the schema definition for the CmdbIp entity.
type CmdbIp struct {
	ent.Schema
}

// Fields of the Cmdb.
func (CmdbIp) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Comment("IP").Unique().Annotations(entsql.WithComments(true)),
		field.String("mask").Comment("掩码").Annotations(entsql.WithComments(true)),
		field.String("gateway").Comment("网关").Annotations(entsql.WithComments(true)),
		field.Time("online_time").Comment("上线时间").Annotations(entsql.WithComments(true)),
		field.Time("offline_time").Comment("下线时间").Annotations(entsql.WithComments(true)),
		field.Uint("is_leisure").Comment("是否空闲 | 1: 空闲 | 0: 使用中").Default(1).Annotations(entsql.WithComments(true)),
		field.String("device_type").Comment("设备类型 | 服务器 | 网络设备 | 存储设备 | PC | 其他").Default("").Annotations(entsql.WithComments(true)),
	}
}

func (CmdbIp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CmdbIpUser{},
	}
}

// Edges of the Cmdb.
func (CmdbIp) Edges() []ent.Edge {
	return nil
}
