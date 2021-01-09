package psql

import (
	"apps/investimento/pkg/model"
	"apps/investimento/pkg/repository/query"
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type psqlNotaRepository struct {
	Conn *pgx.Conn
}

func NewPsqlNotaRepository(conn *pgx.Conn) model.NotaRepository {
	return &psqlNotaRepository{Conn: conn}
}

func (p *psqlNotaRepository) Fetch(ctx context.Context) ([]model.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) GetByID(ctx context.Context, id int) (model.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) Save(ctx context.Context, nota *model.Nota) error {
	tx, err := p.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var id int
	err = tx.QueryRow(ctx, query.InsertNota, nota.Date, nota.ReceiptID, nota.Broker).Scan(&id)
	if err != nil {
		log.Printf("psql/nota.go - save - quert row - %s", err.Error())
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("psql/nota.go - save - commit - %s", err.Error())
		return err
	}

	nota.ID = id

	return nil

}

func (p *psqlNotaRepository) Delete(ctx context.Context, nota *model.Nota) error {
	panic("ahhh")
}
