package psql

import (
	"apps/investimento/pkg/model"
	"context"
	"database/sql"
)

type psqlNotaRepository struct {
	DB *sql.DB
}

func NewPsqlNotaRepository(db *sql.DB) model.NotaRepository {
	return &psqlNotaRepository{db}
}

func (p *psqlNotaRepository) Fetch(ctx context.Context) ([]model.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) GetByID(ctx context.Context, id int) (model.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) Save(ctx context.Context, nota *model.Nota) error {
	panic("ahhh")
}

func (p *psqlNotaRepository) Delete(ctx context.Context, nota *model.Nota) error {
	panic("ahhh")
}
