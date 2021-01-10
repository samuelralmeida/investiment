package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/model"
	"context"
	"log"
)

type entNotaRepository struct {
	Client *ent.Client
}

func NewEntNotaRepository(client *ent.Client) model.NotaRepository {
	return &entNotaRepository{Client: client}
}

func (e *entNotaRepository) Fetch(ctx context.Context) ([]model.Nota, error) {
	panic("ahhh")
}

func (e *entNotaRepository) GetByID(ctx context.Context, id int) (model.Nota, error) {
	panic("ahhh")
}

func (e *entNotaRepository) Save(ctx context.Context, nota *model.Nota) error {

	if err := WithTx(ctx, e.Client, func(tx *ent.Tx) error {
		n, err := tx.Nota.Create().
			SetDate(nota.Date).
			SetReceiptID(nota.ReceiptID).
			SetBroker(nota.Broker).
			Save(ctx)

		if err != nil {
			log.Printf("ent/ent_nota.go - save - create nota - %s", err.Error())
			return err
		}

		nota.ID = n.ID
		return nil

	}); err != nil {
		log.Printf("ent_nota.go - save - withtx - %s", err.Error())
		return err
	}

	return nil

}

func (e *entNotaRepository) Delete(ctx context.Context, nota *model.Nota) error {
	panic("ahhh")
}
