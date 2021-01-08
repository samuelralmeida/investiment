package model

import "context"

type Nota struct {
	ID        int       `json:"id,omitempty"`
	Date      string    `json:"date,omitempty"`
	ReceiptID int       `json:"receiptID,omitempty"`
	Broker    string    `json:"broker,omitempty"`
	Ativos    []*Ativos `json:"ativos,omitempty"`
	Cblc      *Cblc     `json:"cblc,omitempty"`
	Bolsa     *Bolsa    `json:"bolsa,omitempty"`
	Despesa   *Despesa  `json:"despesa,omitempty"`
}

type NotaUsecase interface {
	Fetch(ctx context.Context) ([]Nota, error)
	GetByID(ctx context.Context, id int) (Nota, error)
	Save(ctx context.Context, nota *Nota) error
	Delete(ctx context.Context, nota *Nota) error
}

type NotaRepository interface {
	Fetch(ctx context.Context) ([]Nota, error)
	GetByID(ctx context.Context, id int) (Nota, error)
	Save(ctx context.Context, nota *Nota) error
	Delete(ctx context.Context, nota *Nota) error
}
