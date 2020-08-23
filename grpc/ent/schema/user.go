package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/rs/xid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", xid.ID{}).
			Unique().
			Immutable(),
		field.String("UserID").
			MaxLen(80).
			NotEmpty(),
		field.String("Password").
			MaxLen(200).
			NotEmpty().
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of this User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("UserID", "Password"),
	}
}
