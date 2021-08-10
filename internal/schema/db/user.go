package db

import (
	"go-starter/common/util"
	"time"

	"entgo.io/ent"
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
		field.String("id").MaxLen(22).Immutable().Unique().DefaultFunc(func() string {
			return util.NanoID()
		}),
		field.String("username").Unique().MaxLen(255),
		field.String("email").Unique().Nillable().Optional().MaxLen(255),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
