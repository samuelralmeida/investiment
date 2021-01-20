package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

func createNota(ctx context.Context, client *ent.Client, nota *entity.Nota) error {
	newNota, err := client.Nota.Create().
		SetDate(nota.Date).
		SetReceiptID(nota.ReceiptID).
		SetBroker(nota.Broker).
		Save(ctx)

	if err != nil {
		return errors.Wrap("ent:repository:save:createnota", err)
	}

	nota.ID = newNota.ID
	return nil
}

func deleteNota(ctx context.Context, client *ent.Client, notaID int) error {
	query := client.Nota.Delete()

	if notaID != 0 {
		query = query.Where(nota.ID(notaID))
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return errors.Wrap("ent:repository:deletenota", err)
	}
	return nil
}
