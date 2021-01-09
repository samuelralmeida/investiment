package main

import (
	"apps/investimento/pkg/deliver/rest"
	"apps/investimento/pkg/repository/psql"
	"apps/investimento/pkg/usecases"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const (
	dbconn = "host=localhost port=5432 user=postgres dbname=investimento password=mochileiro sslmode=disable"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbconn)
	if err != nil {
		panic(fmt.Sprintf("main.go - main - connect - %s", err.Error()))
	}
	defer conn.Close(ctx)

	repoNota := psql.NewPsqlNotaRepository(conn)
	usecasesNota := usecases.NewNotaUsecase(repoNota)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	rest.NewNotaHandler(r, usecasesNota)
	log.Print("ListenAndServe port :3000")
	http.ListenAndServe(":3000", r)
}
