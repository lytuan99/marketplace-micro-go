DB_URL=postgresql://postgres:password@localhost:5432/auth-service?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

create_migrate:
	migrate create --ext sql -dir db/migration -seq ${name}

sqlc:
	sqlc generate

server-dev:
	go run cmd/main.go