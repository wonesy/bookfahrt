package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Invitation holds the schema definition for the Invitation entity.
type Invitation struct {
	ent.Schema
}

// Fields of the Invitation.
func (Invitation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Invitation.
func (Invitation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sponsor", User.Type).Unique().Required(),
		edge.To("club", Club.Type).Unique().Required(),
	}
}
