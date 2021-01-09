package model

type Bolsa struct {
	ID              int     `json:"id,omitempty"`
	TaxaTermoOpcoes float64 `json:"taxa_termo_opcoes,omitempty"`
	TaxaAna         float64 `json:"taxa_ana,omitempty"`
	Emolumentos     float64 `json:"emolumentos,omitempty"`
}

func (b *Bolsa) IsValid() bool {
	if b == nil {
		return false
	}
	return true
}
