package model

type Cblc struct {
	ID             int     `json:"id,omitempty"`
	TaxaLiquidacao float64 `json:"taxa_liquidacao,omitempty"`
	TaxaRegistro   float64 `json:"taxa_registro,omitempty"`
}
