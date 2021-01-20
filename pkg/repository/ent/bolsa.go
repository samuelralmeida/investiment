package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/bolsa"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

func createBolsa(ctx context.Context, client *ent.Client, bolsa *entity.Bolsa, notaID int) error {
	newBolsa, err := client.Bolsa.Create().
		SetTaxaTermoOpcoes(bolsa.TaxaTermoOpcoes).
		SetTaxaAna(bolsa.TaxaAna).
		SetEmolumentos(bolsa.Emolumentos).
		SetNotaID(notaID).
		Save(ctx)

	if err != nil {
		return errors.Wrap("ent:repository:save:createbolsa", err)
	}

	bolsa.ID = newBolsa.ID
	return nil
}

func deleteBolsa(ctx context.Context, client *ent.Client, bolsaID int, notaID int) error {
	query := client.Bolsa.Delete()

	if bolsaID != 0 {
		query = query.Where(bolsa.ID(bolsaID))
	}

	if notaID != 0 {
		query = query.Where(bolsa.HasNotaWith(nota.ID(notaID)))
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return errors.Wrap("ent:repository:deletebolsa", err)
	}
	return nil
}
