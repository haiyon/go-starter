package schema

import "entgo.io/ent"

// Sample holds the schema definition for the Sample entity.
type Sample struct {
	ent.Schema
}

// Fields of the Sample.
func (Sample) Fields() []ent.Field {
	return nil
}

// Edges of the Sample.
func (Sample) Edges() []ent.Edge {
	return nil
}
