package schema

import (
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			NotEmpty(),
		field.Int("post_id").
			Optional(),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", User.Type).
			Ref(article.Table).
			Unique().
			Field("post_id"),
	}
}
