package rest

import (
	"apps/investimento/pkg/model"
	"encoding/json"
	"fmt"
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

	nota := &model.Nota{
		ReceiptID: 100,
		Date:      "2020-01-08",
		Broker:    "Easynvest",
	}

	err := h.NotaUsecase.Save(r.Context(), nota)
	fmt.Println("err", err) // DEBUG

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nota)
}
