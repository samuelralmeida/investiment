package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Despesa holds the schema definition for the Despesa entity.
type Despesa struct {
	ent.Schema
}

// Fields of the Despesa.
func (Despesa) Fields() []ent.Field {
	return []ent.Field{
		field.Float("corretagem"),
		field.Float("iss"),
		field.Float("irrf"),
		field.Float("outros"),
	}
}

// Edges of the Despesa.
func (Despesa) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nota", Nota.Type).Ref("despesas").Unique().Required(),
	}
}
