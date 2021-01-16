package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/bolsa"
	"apps/investimento/pkg/ent/cblc"
	"apps/investimento/pkg/ent/despesa"
	"apps/investimento/pkg/ent/movimentacao"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"context"
	"fmt"
	"log"
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
		return nil, err
	}

	resp := make([]*entity.Nota, len(notas))

	for i, nota := range notas {
		n, err := nota.ToModel()
		if err != nil {
			return nil, err
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
		return nil, err
	}

	return nota.ToModel()
}

func (e *entNotaRepository) Save(ctx context.Context, nota *entity.Nota) error {

	if err := WithTx(ctx, e.Client, func(tx *ent.Tx) error {
		newNota, err := tx.Nota.Create().
			SetDate(nota.Date).
			SetReceiptID(nota.ReceiptID).
			SetBroker(nota.Broker).
			Save(ctx)

		if err != nil {
			log.Printf("ent/ent_nota.go - save - create nota - %s", err.Error())
			return err
		}

		nota.ID = newNota.ID

		fmt.Println("ativos", nota.Ativos) // DEBUG

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
				log.Printf("ent/ent_nota.go - save - create ativo - %s", err.Error())
				return err
			}

			ativo.ID = newAtivo.ID
		}

		newCblc, err := tx.Cblc.Create().
			SetTaxaLiquidacao(nota.Cblc.TaxaLiquidacao).
			SetTaxaRegistro(nota.Cblc.TaxaRegistro).
			SetNota(newNota).
			Save(ctx)

		if err != nil {
			log.Printf("ent/ent_nota.go - save - create cblc - %s", err.Error())
			return err
		}

		nota.Cblc.ID = newCblc.ID

		newBolsa, err := tx.Bolsa.Create().
			SetTaxaTermoOpcoes(nota.Bolsa.TaxaTermoOpcoes).
			SetTaxaAna(nota.Bolsa.TaxaAna).
			SetEmolumentos(nota.Bolsa.Emolumentos).
			SetNota(newNota).
			Save(ctx)

		if err != nil {
			log.Printf("ent/ent_nota.go - save - create bolsa - %s", err.Error())
			return err
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
			log.Printf("ent/ent_nota.go - save - create despesa - %s", err.Error())
			return err
		}

		nota.Despesa.ID = newDespesa.ID

		return nil

	}); err != nil {
		log.Printf("ent_nota.go - save - withtx - %s", err.Error())
		return err
	}

	return nil

}

func (e *entNotaRepository) DeleteByID(ctx context.Context, id int) error {

	if err := WithTx(ctx, e.Client, func(tx *ent.Tx) error {
		var err error

		err = tx.Nota.DeleteOneID(id).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.Movimentacao.Delete().Where(movimentacao.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.Cblc.Delete().Where(cblc.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.Bolsa.Delete().Where(bolsa.HasNotaWith(nota.ID(id))).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.Despesa.Delete().Where(despesa.HasNotaWith(nota.ID(id))).Exec(ctx)
		return err

	}); err != nil {
		log.Printf("ent_nota.go - delete - withtx - %s", err.Error())
		return err
	}

	return nil
}
