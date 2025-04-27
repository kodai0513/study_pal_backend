package schema

import (
	"study-pal-backend/ent/mixin"
	"study-pal-backend/ent/role"
	"study-pal-backend/ent/user"
	"study-pal-backend/ent/workbook"
	"study-pal-backend/ent/workbookinvitationmember"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WorkbookInvitationMember holds the schema definition for the WorkbookInvitationMember entity.
type WorkbookInvitationMember struct {
	ent.Schema
}

func (WorkbookInvitationMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IdMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the WorkbookInvitationMember.
func (WorkbookInvitationMember) Fields() []ent.Field {
	return []ent.Field{
		field.Time("effective_at"),
		field.Bool("is_invited"),
		field.UUID("role_id", uuid.UUID{}).Unique(),
		field.UUID("user_id", uuid.UUID{}).Unique(),
		field.UUID("workbook_id", uuid.UUID{}).Unique(),
	}
}

// Edges of the WorkbookInvitationMember.
func (WorkbookInvitationMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(role.Label, Role.Type).Ref(workbookinvitationmember.Table).Unique().Required().Field("role_id"),
		edge.From(user.Label, User.Type).Ref(workbookinvitationmember.Table).Unique().Required().Field("user_id"),
		edge.From(workbook.Label, Workbook.Type).Ref(workbookinvitationmember.Table).Unique().Required().Field("workbook_id"),
	}
}
