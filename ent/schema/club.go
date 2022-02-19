package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", User.Type).Ref("memberOf"),
	}
}
