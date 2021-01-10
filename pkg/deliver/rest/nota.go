package rest

import (
	"apps/investimento/pkg/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type NotaHandler struct {
	NotaUsecase model.NotaUsecase
}

func NewNotaHandler(r *chi.Mux, notaUsecase model.NotaUsecase) {
	handler := &NotaHandler{
		NotaUsecase: notaUsecase,
	}

	r.Get("/nota/{notaID}/", handler.FindNota)
	r.Post("/nota/new/", handler.NewNota)
}

func (h *NotaHandler) NewNota(w http.ResponseWriter, r *http.Request) {

	var nota model.Nota

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&nota)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	err = h.NotaUsecase.Save(r.Context(), &nota)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nota)
}

func (h *NotaHandler) FindNota(w http.ResponseWriter, r *http.Request) {

	notaID := chi.URLParam(r, "notaID")

	id, err := strconv.Atoi(notaID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	nota, err := h.NotaUsecase.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nota)
}
