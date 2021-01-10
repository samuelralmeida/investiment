package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/nota"
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

func (e *entNotaRepository) Fetch(ctx context.Context) ([]*model.Nota, error) {
	panic("ahhh")
}

func (e *entNotaRepository) GetByID(ctx context.Context, id int) (*model.Nota, error) {
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

func (e *entNotaRepository) Save(ctx context.Context, nota *model.Nota) error {

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

func (e *entNotaRepository) Delete(ctx context.Context, nota *model.Nota) error {
	panic("ahhh")
}
