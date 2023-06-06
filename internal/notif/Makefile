include ../../.env

.PHONY: test
test:
	@tput setaf 2;echo "generating sqlc go code"
	go test -v -cover -short ./...

.PHONY: server
server:
	@tput setaf 2;echo "running server"
	go run main.go

.PHONY: sqlc
sqlc:
	@tput setaf 2;echo "generating sqlc go code"
	@docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

.PHONY: migrateup
migrateup:
	@tput setaf 2;echo "migrating up"
	@migrate -path postgres/migration -database "$(DB_URL)" -verbose up

.PHONY: migratedown
migratedown:
	@tput setaf 2;echo "migrating down"
	@migrate -path postgres/migration -database "$(DB_URL)" -verbose up	

.PHONY: migrateforce 
migrateforce:
	@tput setaf 2;echo "migrating force to $(VERSION)"
	@migrate -path postgres/migration -database "$(DB_URL)" -verbose force $(VERSION)
