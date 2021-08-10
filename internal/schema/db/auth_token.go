package db

import (
	"go-starter/common/util"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthToken holds the schema definition for the AuthToken entity.
type AuthToken struct {
	ent.Schema
}

// Fields of the AuthToken.
func (AuthToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(22).Immutable().Unique().DefaultFunc(func() string {
			return util.NanoID()
		}),
		field.Bool("disabled").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("user_id"),
	}
}

// Edges of the AuthToken.
func (AuthToken) Edges() []ent.Edge {
	return nil
}
