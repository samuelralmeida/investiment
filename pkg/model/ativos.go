package model

type Ativo struct {
	ID          int     `json:"id,omitempty"`
	Mercado     string  `json:"mercado,omitempty"`
	CV          string  `json:"c_v,omitempty"`
	TipoMercado string  `json:"tipo_mercado,omitempty"`
	Titulo      string  `json:"titulo,omitempty"`
	Qtde        int     `json:"qtde,omitempty"`
	Valor       float64 `json:"valor,omitempty"`
	DC          string  `json:"d_c,omitempty"`
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
