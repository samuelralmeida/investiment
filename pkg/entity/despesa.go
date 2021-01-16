package entity

type Despesa struct {
	ID         int     `json:"id"`
	Corretagem float64 `json:"corretagem"`
	Iss        float64 `json:"iss"`
	Irrf       float64 `json:"irrf"`
	Outros     float64 `json:"outros"`
}

func (d *Despesa) IsValid() bool {
	if d == nil {
		return false
	}
	return true
}
