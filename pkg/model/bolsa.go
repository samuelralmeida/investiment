package model

type Bolsa struct {
	ID              int     `json:"id"`
	TaxaTermoOpcoes float64 `json:"taxa_termo_opcoes"`
	TaxaAna         float64 `json:"taxa_ana"`
	Emolumentos     float64 `json:"emolumentos"`
}

func (b *Bolsa) IsValid() bool {
	if b == nil {
		return false
	}
	return true
}
