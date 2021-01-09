package usecases

import (
	"apps/investimento/pkg/model"
	"context"
)

type notaUsecase struct {
	notaRepository model.NotaRepository
}

func NewNotaUsecase(repo model.NotaRepository) model.NotaUsecase {
	return &notaUsecase{
		notaRepository: repo,
	}
}

func (n *notaUsecase) Fetch(ctx context.Context) ([]model.Nota, error) {
	panic("ahhhh")
}

func (n *notaUsecase) GetByID(ctx context.Context, id int) (model.Nota, error) {
	panic("ahhhh")
}

func (n *notaUsecase) Save(ctx context.Context, nota *model.Nota) error {
	// if !nota.IsValid() {
	// 	return errors.New("nota invalid")
	// }
	return n.notaRepository.Save(ctx, nota)
}

func (n *notaUsecase) Delete(ctx context.Context, nota *model.Nota) error {
	panic("ahhhh")
}
