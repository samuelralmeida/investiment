package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/cblc"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

func createCblc(ctx context.Context, client *ent.Client, cblc *entity.Cblc, notaID int) error {
	newCblc, err := client.Cblc.Create().
		SetTaxaLiquidacao(cblc.TaxaLiquidacao).
		SetTaxaRegistro(cblc.TaxaRegistro).
		SetNotaID(notaID).
		Save(ctx)

	if err != nil {
		return errors.Wrap("ent:repository:save:createcblc", err)
	}

	cblc.ID = newCblc.ID
	return nil
}

func deleteCblc(ctx context.Context, client *ent.Client, cblcID int, notaID int) error {
	query := client.Cblc.Delete()

	if cblcID != 0 {
		query = query.Where(cblc.ID(cblcID))
	}

	if notaID != 0 {
		query = query.Where(cblc.HasNotaWith(nota.ID(notaID)))
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return errors.Wrap("ent:repository:deletecblc", err)
	}
	return nil
}
