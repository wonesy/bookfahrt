package schema

import "entgo.io/ent"

// Completion holds the schema definition for the Completion entity.
type Completion struct {
	ent.Schema
}

// Fields of the Completion.
func (Completion) Fields() []ent.Field {
	return nil
}

// Edges of the Completion.
func (Completion) Edges() []ent.Edge {
	return nil
}
