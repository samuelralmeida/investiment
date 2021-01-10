package main

import (
	"apps/investimento/pkg/deliver/rest"
	entr "apps/investimento/pkg/repository/ent"
	"apps/investimento/pkg/usecases"

	"log"
	"net/http"
	"time"

	"apps/investimento/pkg/ent"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"
)

const (
	dbconn = "host=localhost port=5432 user=postgres dbname=investimento password=mochileiro sslmode=disable"
)

func main() {

	client, err := ent.Open("postgres", dbconn)
	if err != nil {
		log.Fatalf("main.go - main - ent open - %s", err.Error())
	}
	defer client.Close()
	repoNota := entr.NewEntNotaRepository(client)

	// common
	usecasesNota := usecases.NewNotaUsecase(repoNota)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	rest.NewNotaHandler(r, usecasesNota)
	log.Print("Serving --repo=ent --port=3000")
	http.ListenAndServe(":3000", r)
}
