package entrepository

import (
	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/nota"
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
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

	createNota := func(tx *ent.Tx) error {
		return createNota(ctx, tx.Client(), nota)
	}

	createAtivos := func(tx *ent.Tx) error {
		return createAtivos(ctx, tx.Client(), nota.Ativos, nota.ID)
	}

	createCblc := func(tx *ent.Tx) error {
		return createCblc(ctx, tx.Client(), nota.Cblc, nota.ID)
	}

	createBolsa := func(tx *ent.Tx) error {
		return createBolsa(ctx, tx.Client(), nota.Bolsa, nota.ID)
	}

	createDespesa := func(tx *ent.Tx) error {
		return createDespesa(ctx, tx.Client(), nota.Despesa, nota.ID)
	}

	create := []func(tx *ent.Tx) error{
		createNota,
		createAtivos,
		createCblc,
		createBolsa,
		createDespesa,
	}

	err := WithTxFns(ctx, e.Client, create)
	return errors.Wrap("ent:repository:save:withtxfns", err)
}

func (e *entNotaRepository) DeleteByID(ctx context.Context, id int) error {

	deleteDespesa := func(tx *ent.Tx) error {
		return deleteDespesa(ctx, tx.Client(), 0, id)
	}

	deleteBolsa := func(tx *ent.Tx) error {
		return deleteBolsa(ctx, tx.Client(), 0, id)
	}

	deleteCblc := func(tx *ent.Tx) error {
		return deleteCblc(ctx, tx.Client(), 0, id)
	}

	deleteAtivos := func(tx *ent.Tx) error {
		return deleteAtivos(ctx, tx.Client(), []int{}, id)
	}

	deleteNota := func(tx *ent.Tx) error {
		return deleteNota(ctx, tx.Client(), id)
	}

	delete := []func(tx *ent.Tx) error{
		deleteDespesa,
		deleteBolsa,
		deleteCblc,
		deleteAtivos,
		deleteNota,
	}

	err := WithTxFns(ctx, e.Client, delete)
	return errors.Wrap("ent:repository:deletebyid:withtxfns", err)
}
