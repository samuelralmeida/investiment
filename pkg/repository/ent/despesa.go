package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/despesa"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

func createDespesa(ctx context.Context, client *ent.Client, despesa *entity.Despesa, notaID int) error {
	newDespesa, err := client.Despesa.Create().
		SetCorretagem(despesa.Corretagem).
		SetIss(despesa.Iss).
		SetIrrf(despesa.Irrf).
		SetOutros(despesa.Outros).
		SetNotaID(notaID).
		Save(ctx)

	if err != nil {
		return errors.Wrap("ent:repository:save:createdespesa", err)
	}

	despesa.ID = newDespesa.ID
	return nil
}

// deleteDespesa truncates table if despesaID and notaID are zero.
func deleteDespesa(ctx context.Context, client *ent.Client, despesaID int, notaID int) error {
	query := client.Despesa.Delete()

	if despesaID != 0 {
		query = query.Where(despesa.ID(despesaID))
	}

	if notaID != 0 {
		query = query.Where(despesa.HasNotaWith(nota.ID(notaID)))
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return errors.Wrap("ent:repository:deletedespesa", err)
	}
	return nil
}
