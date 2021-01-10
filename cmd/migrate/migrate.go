package main

import (
	"context"
	"log"

	"apps/investimento/pkg/ent"
	"apps/investimento/pkg/ent/migrate"

	_ "github.com/lib/pq"
)

func main() {

	dbconn := "host=localhost port=5432 user=postgres dbname=investimento password=mochileiro sslmode=disable"

	client, err := ent.Open("postgres", dbconn)
	if err != nil {
		log.Fatalf("main.go - main - ent open - %s", err.Error())
	}
	defer client.Close()

	ctx := context.Background()

	// Run migration.
	err = client.Debug().Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
