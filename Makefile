run-psql:
	CompileDaemon -build="go build -o app-psql cmd/psql/main.go" -command="./app-psql" -exclude-dir=.git -color=true -log-prefix=false

run-ent:
	CompileDaemon -build="go build -o app-ent cmd/ent/main.go" -command="./app-ent" -exclude-dir=.git -color=true -log-prefix=false

env:
	asdf local golang 1.15.6 && \
	go get github.com/facebook/ent/cmd/ent && \
	GO111MODULE=off go get -u github.com/githubnemo/CompileDaemon && \
	asdf reshim golang 1.15.6

schema:
	go run github.com/facebook/ent/cmd/entc generate ./pkg/ent/schema

migrate:
	go run cmd/migrate/migrate.go

run-postgres-container:
	docker run --name apps-postgres -v pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=1593574562585 postgres

start-container:
	docker start apps-postgres

stop-container:
	docker stop apps-postgres