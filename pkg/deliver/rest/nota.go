package rest

import (
	"apps/investimento/pkg/model"
	"net/http"

	"github.com/go-chi/chi"
)

type NotaHandler struct {
	NotaUsecase model.NotaUsecase
}

func NewNotaHandler(r *chi.Mux, notaUsecase model.NotaUsecase) {
	handler := &NotaHandler{
		NotaUsecase: notaUsecase,
	}

	r.Get("/notas", handler.FindNotas)
}

func (h *NotaHandler) FindNotas(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("find notas"))
}
