package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Cblc holds the schema definition for the Cblc entity.
type Cblc struct {
	ent.Schema
}

// Fields of the Cblc.
func (Cblc) Fields() []ent.Field {
	return []ent.Field{
		field.Float("taxa_liquidacao"),
		field.Float("taxa_registro"),
	}
}

// Edges of the Cblc.
func (Cblc) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nota", Nota.Type).Ref("cblcs").Unique().Required(),
	}
}
