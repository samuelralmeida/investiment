package rest

import (
	"apps/investimento/pkg/entity"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type NotaHandler struct {
	NotaUsecase entity.NotaUsecase
}

func NewNotaHandler(r *chi.Mux, notaUsecase entity.NotaUsecase) {
	handler := &NotaHandler{
		NotaUsecase: notaUsecase,
	}

	r.Get("/nota/{notaID}/", handler.FindNota)
	r.Get("/notas/", handler.Notas)
	r.Post("/nota/new/", handler.NewNota)
	r.Delete("/nota/{notaID}/", handler.DeleteNota)
}

func (h *NotaHandler) NewNota(w http.ResponseWriter, r *http.Request) {
	var nota entity.Nota

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
		http.Error(w, http.StatusText(400), 400)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nota)
}

func (h *NotaHandler) DeleteNota(w http.ResponseWriter, r *http.Request) {

	notaID := chi.URLParam(r, "notaID")

	id, err := strconv.Atoi(notaID)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.NotaUsecase.DeleteByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, http.StatusText(204), 204)
}

func (h *NotaHandler) Notas(w http.ResponseWriter, r *http.Request) {

	notas, err := h.NotaUsecase.Fetch(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		w.Header().Set("Content-Type", "application/json")
		log.Print(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notas)
}
