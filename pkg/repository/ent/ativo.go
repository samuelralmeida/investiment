package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/movimentacao"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

func createAtivos(ctx context.Context, client *ent.Client, ativos []*entity.Ativo, notaID int) error {
	for _, ativo := range ativos {
		newAtivo, err := client.Movimentacao.Create().
			SetMercado(ativo.Mercado).
			SetCV(ativo.CV).
			SetTipoMercado(ativo.TipoMercado).
			SetTitulo(ativo.Titulo).
			SetQtde(ativo.Qtde).
			SetValor(ativo.Valor).
			SetDC(ativo.DC).
			SetNotaID(notaID).
			Save(ctx)

		if err != nil {
			return errors.Wrap("ent:repository:save:createativo", err)
		}

		ativo.ID = newAtivo.ID
	}
	return nil
}

func deleteAtivos(ctx context.Context, client *ent.Client, ativosIDs []int, notaID int) error {
	query := client.Movimentacao.Delete()

	if len(ativosIDs) > 0 {
		query = query.Where(movimentacao.IDIn(ativosIDs...))
	}

	if notaID != 0 {
		query = query.Where(movimentacao.HasNotaWith(nota.ID(notaID)))
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return errors.Wrap("ent:repository:deleteativos", err)
	}
	return nil
}
