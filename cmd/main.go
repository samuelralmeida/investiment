package main

import (
	"apps/investimento/pkg/deliver/rest"
	"apps/investimento/pkg/repository/psql"
	"apps/investimento/pkg/usecases"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"
)

const (
	dbconn = "host=localhost port=5432 user=postgres dbname=investimento password=mochileiro sslmode=disable"
)

func main() {

	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		panic(fmt.Sprintf("main.go - main - open - %s", err.Error()))
	}
	defer db.Close()

	repoNota := psql.NewPsqlNotaRepository(db)
	usecasesNota := usecases.NewNotaUsecase(repoNota)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	rest.NewNotaHandler(r, usecasesNota)
	http.ListenAndServe(":3000", r)
}
