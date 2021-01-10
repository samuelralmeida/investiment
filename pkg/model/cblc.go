package model

type Cblc struct {
	ID             int     `json:"id"`
	TaxaLiquidacao float64 `json:"taxa_liquidacao"`
	TaxaRegistro   float64 `json:"taxa_registro"`
}

func (c *Cblc) IsValid() bool {
	if c == nil {
		return false
	}
	return true
}
