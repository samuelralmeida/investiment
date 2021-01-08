package model

type Ativos struct {
	ID          int     `json:"id,omitempty"`
	Mercado     string  `json:"mercado,omitempty"`
	CV          string  `json:"c_v,omitempty"`
	TipoMercado string  `json:"tipo_mercado,omitempty"`
	Titulo      string  `json:"titulo,omitempty"`
	Qtde        int     `json:"qtde,omitempty"`
	Valor       float64 `json:"valor,omitempty"`
	DC          string  `json:"d_c,omitempty"`
}
