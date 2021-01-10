package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Bolsa holds the schema definition for the Bolsa entity.
type Bolsa struct {
	ent.Schema
}

// Fields of the Bolsa.
func (Bolsa) Fields() []ent.Field {
	return []ent.Field{
		field.Float("taxa_termo_opcoes"),
		field.Float("taxa_ana"),
		field.Float("emolumentos"),
	}
}

// Edges of the Bolsa.
func (Bolsa) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nota", Nota.Type).Ref("bolsas").Unique().Required(),
	}
}
