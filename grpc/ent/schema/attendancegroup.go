package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/rs/xid"
)

// AttendanceGroup holds the schema definition for the AttendanceGroup entity.
type AttendanceGroup struct {
	ent.Schema
}

// Fields of the AttendanceGroup.
// nolint
func (AttendanceGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", xid.ID{}).
			Unique().
			Immutable(),
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

// Edges of the AttendanceGroup.
func (AttendanceGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
