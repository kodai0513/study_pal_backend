package schema

import (
	"regexp"
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("name").
			MaxLen(30).
			Unique().
			Match(regexp.MustCompile("[a-zA-Z_0-9]+$")),
		field.String("nick_name"),
		field.String("password").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(article.Table, Article.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
