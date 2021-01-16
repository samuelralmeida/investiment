package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/bolsa"
	"apps/investimento/pkg/ent/cblc"
	"apps/investimento/pkg/ent/despesa"
	"apps/investimento/pkg/ent/movimentacao"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
	"fmt"
)

type entNotaRepository struct {
	Client *ent.Client
}

func NewEntNotaRepository(client *ent.Client) entity.NotaRepository {
	return &entNotaRepository{Client: client}
}

func (e *entNotaRepository) Fetch(ctx context.Context) ([]*entity.Nota, error) {
	notas, err := e.Client.Nota.
		Query().
		WithMovimentacoes().
		WithCblcs().
		WithBolsas().
		WithDespesas().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap("ent:repository:fetch:all", err)
	}

	resp := make([]*entity.Nota, len(notas))

	for i, nota := range notas {
		n, err := nota.ToModel()
		if err != nil {
			return nil, errors.Wrap("ent:repository:fetch:tomodel", err)
		}
		resp[i] = n
	}

	return resp, nil

}

func (e *entNotaRepository) GetByID(ctx context.Context, id int) (*entity.Nota, error) {
	nota, err := e.Client.Nota.
		Query().
		Where(nota.ID(id)).
		WithMovimentacoes().
		WithCblcs().
		WithBolsas().
		WithDespesas().
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap("ent:repository:getbyid:only", err)
	}

	n, err := nota.ToModel()
	return n, errors.Wrap("ent:repository:getbyid:tomodel", err)
}

func (e *entNotaRepository) Save(ctx context.Context, nota *entity.Nota) error {

	err := WithTx(ctx, e.Client, func(tx *ent.Tx) error {
		newNota, err := tx.Nota.Create().
			SetDate(nota.Date).
			SetReceiptID(nota.ReceiptID).
			SetBroker(nota.Broker).
			Save(ctx)

		if err != nil {
			return errors.Wrap("ent:repository:save:createnota", err)
		}

		nota.ID = newNota.ID

		for _, ativo := range nota.Ativos {
			newAtivo, err := tx.Movimentacao.Create().
				SetMercado(ativo.Mercado).
				SetCV(ativo.CV).
				SetTipoMercado(ativo.TipoMercado).
				SetTitulo(ativo.Titulo).
				SetQtde(ativo.Qtde).
				SetValor(ativo.Valor).
				SetDC(ativo.DC).
				SetNota(newNota).
				Save(ctx)

			if err != nil {
				return errors.Wrap("ent:repository:save:createativo", err)
			}

			ativo.ID = newAtivo.ID
		}

		newCblc, err := tx.Cblc.Create().
			SetTaxaLiquidacao(nota.Cblc.TaxaLiquidacao).
			SetTaxaRegistro(nota.Cblc.TaxaRegistro).
			SetNota(newNota).
			Save(ctx)

		if err != nil {
			return errors.Wrap("ent:repository:save:createcblc", err)
		}

		nota.Cblc.ID = newCblc.ID

		newBolsa, err := tx.Bolsa.Create().
			SetTaxaTermoOpcoes(nota.Bolsa.TaxaTermoOpcoes).
			SetTaxaAna(nota.Bolsa.TaxaAna).
			SetEmolumentos(nota.Bolsa.Emolumentos).
			SetNota(newNota).
			Save(ctx)

		if err != nil {
			return errors.Wrap("ent:repository:save:createbolsa", err)
		}

		nota.Bolsa.ID = newBolsa.ID

		newDespesa, err := tx.Despesa.Create().
			SetCorretagem(nota.Despesa.Corretagem).
			SetIss(nota.Despesa.Iss).
			SetIrrf(nota.Despesa.Irrf).
			SetOutros(nota.Despesa.Outros).
			SetNota(newNota).
			Save(ctx)

		if err != nil {
			return errors.Wrap("ent:repository:save:createdespesa", err)
		}

		nota.Despesa.ID = newDespesa.ID

		return nil

	})

	fmt.Println("aaa", err) // DEBUG

	return errors.Wrap("ent:repository:save:withtx", err)
}

func (e *entNotaRepository) DeleteByID(ctx context.Context, id int) error {

	err := WithTx(ctx, e.Client, func(tx *ent.Tx) error {
		var err error

		err = tx.Nota.DeleteOneID(id).Exec(ctx)
		if err != nil {
			return errors.Wrap("ent:repository:deletebyid:deletenota", err)
		}

		_, err = tx.Movimentacao.Delete().Where(movimentacao.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return errors.Wrap("ent:repository:deletebyid:deletemovimentcao", err)
		}

		_, err = tx.Cblc.Delete().Where(cblc.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return errors.Wrap("ent:repository:deletebyid:deletecblc", err)
		}

		_, err = tx.Bolsa.Delete().Where(bolsa.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return errors.Wrap("ent:repository:deletebyid:deletebolsa", err)
		}

		_, err = tx.Despesa.Delete().Where(despesa.HasNotaWith(nota.ID(id))).Exec(ctx)
		return errors.Wrap("ent:repository:deletebyid:deletedespesa", err)

	})

	return errors.Wrap("ent:repository:deletebyid:withtx", err)
}
