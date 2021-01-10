package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Nota holds the schema definition for the NotaCorretagem entity.
type Nota struct {
	ent.Schema
}

// Fields of the Nota.
func (Nota) Fields() []ent.Field {
	return []ent.Field{
		field.String("date"),
		field.Int("receiptID"),
		field.String("broker"),
	}
}

// Edges of the Nota.
func (Nota) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cblcs", Cblc.Type).Unique(),
		edge.To("bolsas", Bolsa.Type).Unique(),
		edge.To("despesas", Despesa.Type).Unique(),
		edge.To("movimentacoes", Movimentacao.Type),
	}
}
