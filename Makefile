run-psql:
ifeq (, $(shell which wtc 2>/dev/null))
	@echo "\033[31mYOU NEED TO RUN: 'GO111MODULE=OFF go get -u github.com/rafaelsq/wtc'\033[m" && false
endif
	@wtc -t run-psql

run-ent:
ifeq (, $(shell which wtc 2>/dev/null))
	@echo "\033[31mYOU NEED TO RUN: 'GO111MODULE=OFF go get -u github.com/rafaelsq/wtc'\033[m" && false
endif
	@wtc -t run-ent

schema:
	go run github.com/facebook/ent/cmd/entc generate ./pkg/ent/schema

migrate:
	go run cmd/migrate/migrate.go