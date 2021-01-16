package psql

import (
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/repository/query"
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type psqlNotaRepository struct {
	Conn *pgx.Conn
}

func NewPsqlNotaRepository(conn *pgx.Conn) entity.NotaRepository {
	return &psqlNotaRepository{Conn: conn}
}

func (p *psqlNotaRepository) Fetch(ctx context.Context) ([]*entity.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) GetByID(ctx context.Context, id int) (*entity.Nota, error) {
	panic("ahhh")
}

func (p *psqlNotaRepository) Save(ctx context.Context, nota *entity.Nota) error {
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

func (p *psqlNotaRepository) DeleteByID(ctx context.Context, id int) error {
	panic("ahhh")
}
