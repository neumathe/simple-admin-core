package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/neumathe/simple-admin-common/orm/ent/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").
			Comment("Password | 密码"),
		field.String("salt").
			Comment("salt | 密码加盐"),
		field.String("nickname").Unique().
			Comment("Nickname | 昵称"),
		field.String("description").Optional().
			Comment("The description of user | 用户的描述信息"),
		field.String("home_path").Default("/dashboard").
			Comment("The home page that the user enters after logging in | 用户登陆后进入的首页"),
		field.String("mobile").Optional().Unique().
			Comment("Mobile number | 手机号"),
		field.String("email").Optional().Unique().
			Comment("Email | 邮箱号"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("Avatar | 头像路径"),
		field.Uint64("department_id").Optional().Default(1).
			Comment("Department ID | 部门ID"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("departments", Department.Type).Unique().Field("department_id"),
		edge.To("positions", Position.Type),
		edge.To("roles", Role.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "sys_users"},
	}
}
