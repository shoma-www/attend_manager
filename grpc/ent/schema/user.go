package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/rs/xid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
// nolint
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", xid.ID{}).
			Unique().
			Immutable(),
		field.String("LoginID").
			MaxLen(80).
			NotEmpty(),
		field.String("Password").
			MaxLen(200).
			NotEmpty().
			Sensitive(),
		field.String("Name").
			MaxLen(20).
			Nillable().
			Optional(),
		field.Time("CreatedAt").
			Optional(),
		field.Time("UpdatedAt").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", AttendanceGroup.Type).
			Ref("users").
			Unique(),
	}
}

// Indexes of this User.
func (User) Indexes() []ent.Index {
	return nil
}
