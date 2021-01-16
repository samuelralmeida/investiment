run-psql:
	CompileDaemon -build="go build -o app-psql cmd/psql/main.go" -command="./app-psql" -exclude-dir=.git -color=true -log-prefix=false

run-ent:
	CompileDaemon -build="go build -o app-ent cmd/ent/main.go" -command="./app-ent" -exclude-dir=.git -color=true -log-prefix=false

env:
	GO111MODULE=off go get -u github.com/githubnemo/CompileDaemon && asdf reshim golang 1.15.6

schema:
	go run github.com/facebook/ent/cmd/entc generate ./pkg/ent/schema

migrate:
	go run cmd/migrate/migrate.go

