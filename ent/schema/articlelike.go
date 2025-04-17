package schema

import (
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/articlelike"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ArticleLike holds the schema definition for the ArticleLike entity.
type ArticleLike struct {
	ent.Schema
}

func (ArticleLike) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the ArticleLike.
func (ArticleLike) Fields() []ent.Field {
	return []ent.Field{
		field.UUID(article.Label+"_id", uuid.UUID{}).Unique(),
		field.UUID(user.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the ArticleLike.
func (ArticleLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(article.Label, Article.Type).Ref(articlelike.Table).Unique().Required().Field(article.Label + "_id"),
	}
}
