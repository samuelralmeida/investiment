package model

type Despesa struct {
	ID         int     `json:"id,omitempty"`
	Corretagem float64 `json:"corretagem,omitempty"`
	Iss        float64 `json:"iss,omitempty"`
	Irrf       float64 `json:"irrf,omitempty"`
	Outros     float64 `json:"outros,omitempty"`
}

func (d *Despesa) IsValid() bool {
	if d == nil {
		return false
	}
	return true
}
