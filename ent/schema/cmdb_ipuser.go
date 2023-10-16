package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	"time"
)

// CmdbIpUser holds the schema definition for the CmdbIpUser entity.
type CmdbIpUser struct {
	ent.Schema
}

// Fields of the Cmdb.
func (CmdbIpUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Comment("Create Time | 创建日期").
			Annotations(entsql.WithComments(true)),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Update Time | 修改日期").
			Annotations(entsql.WithComments(true)),
		field.String("user").Comment("用户").Unique().Annotations(entsql.WithComments(true)),
		field.String("department").Comment("部门").Annotations(entsql.WithComments(true)),
		field.String("mobile").Comment("联系方式").MinLen(11).MaxLen(11).Annotations(entsql.WithComments(true)),
		field.String("remark").Comment("备注").Annotations(entsql.WithComments(true)),
	}
}

func (CmdbIpUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.StatusMixin{},
	}
}

// Edges of the Cmdb.
func (CmdbIpUser) Edges() []ent.Edge {
	return nil
}
