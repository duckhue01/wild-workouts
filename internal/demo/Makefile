DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

.PHONY: migrate_up
migrate_up:
 	docker run --rm -v $(shell pwd)/postgres/migration:/migrations --network host migrate/migrate -path=/migrations/ -database "$(DB_URL)" database up

.PHONY: migrate_down
migrate_down:
 	docker run --rm -v $(shell pwd)/postgres/migration:/migrations --network host migrate/migrate -path=/migrations/ -database "$(DB_URL)" database up

.PHONY: test
test:
	go test -v -cover -short ./...

.PHONY: server
server:
	go run main.go

.PHONY: sqlc
sqlc:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

