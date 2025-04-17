package schema

import (
	"study-pal-backend/ent/article"
	"study-pal-backend/ent/articlelike"
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("page_id").Nillable().Optional().Unique(),
		field.String("description").MaxLen(400).NotEmpty(),
		field.UUID(user.Label+"_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", User.Type).Ref(article.Table).Unique().Required().Field(user.Label + "_id"),
		edge.To(articlelike.Table, ArticleLike.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
	}
}
