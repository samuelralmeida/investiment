package usecases

import (
	"apps/investimento/pkg/entity"
	"apps/investimento/pkg/support/errors"
	"context"
)

type notaUsecase struct {
	notaRepository entity.NotaRepository
}

func NewNotaUsecase(repo entity.NotaRepository) entity.NotaUsecase {
	return &notaUsecase{
		notaRepository: repo,
	}
}

func (n *notaUsecase) Fetch(ctx context.Context) ([]*entity.Nota, error) {
	return n.notaRepository.Fetch(ctx)
}

func (n *notaUsecase) GetByID(ctx context.Context, id int) (*entity.Nota, error) {
	return n.notaRepository.GetByID(ctx, id)
}

func (n *notaUsecase) Save(ctx context.Context, nota *entity.Nota) error {
	if !nota.IsValid() {
		return errors.Wrap("usecases:save:nota:valid", errors.ErrInvalidNota)
	}
	return n.notaRepository.Save(ctx, nota)
}

func (n *notaUsecase) DeleteByID(ctx context.Context, id int) error {
	return n.notaRepository.DeleteByID(ctx, id)
}
