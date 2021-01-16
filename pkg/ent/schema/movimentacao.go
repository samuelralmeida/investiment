package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Movimentacao holds the schema definition for the Movimentacao entity.
type Movimentacao struct {
	ent.Schema
}

// Fields of the Movimentacao.
func (Movimentacao) Fields() []ent.Field {
	return []ent.Field{
		field.String("mercado"),
		field.String("c_v"),
		field.String("tipo_mercado"),
		field.String("titulo"),
		field.Int("qtde"),
		field.Float("valor"),
		field.String("d_c"),
	}
}

// Edges of the Movimentacao.
func (Movimentacao) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nota", Nota.Type).Ref("movimentacaos").Unique().Required(),
	}
}
