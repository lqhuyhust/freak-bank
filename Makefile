migrateup:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/fbank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123456@localhost:5432/fbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go  freak-bank/db/sqlc Store

.PHONY: migrateup migratedown sqlc test server mock