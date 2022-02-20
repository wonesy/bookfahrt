package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("password").NotEmpty().Sensitive(),
		field.String("email").Optional(),
		field.Time("created_at").Default(time.Now().UTC),
		field.Time("last_login_at").Default(time.Now().UTC),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("clubs", Club.Type),
		edge.From("invitations", Invitation.Type).Ref("sponsor"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
	}
}
