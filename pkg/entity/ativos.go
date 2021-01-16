package entity

type Ativo struct {
	ID          int     `json:"id"`
	Mercado     string  `json:"mercado"`
	CV          string  `json:"c_v"`
	TipoMercado string  `json:"tipo_mercado"`
	Titulo      string  `json:"titulo"`
	Qtde        int     `json:"qtde"`
	Valor       float64 `json:"valor"`
	DC          string  `json:"d_c"`
}

type Ativos []*Ativo

func (as *Ativos) IsValid() bool {
	if as == nil {
		return false
	}

	if len(*as) == 0 {
		return false
	}

	for _, a := range *as {
		if !a.IsValid() {
			return false
		}
	}

	return true
}

func (a *Ativo) IsValid() bool {
	if a == nil {
		return false
	}

	if a.Mercado == "" {
		return false
	}

	if a.CV == "" {
		return false
	}

	if a.TipoMercado == "" {
		return false
	}

	if a.Titulo == "" {
		return false
	}

	if a.Qtde == 0 {
		return false
	}

	if a.Valor == 0 {
		return false
	}

	if a.DC == "" {
		return false
	}

	return true
}
