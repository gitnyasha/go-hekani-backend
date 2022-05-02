DB_URL=postgresql://hekanidb:topgear12@localhost:5432/hekani?sslmode=disable

# network:
# 	docker network create bank-network

postgres:
	docker run --name hekani14 -p 127.0.0.1:5432:5432 --env POSTGRES_USER=hekanidb --env POSTGRES_PASSWORD=topgear12 --detach --restart unless-stopped postgres:14-alpine

# mysql:
# 	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it hekani14 createdb --username=hekanidb --owner=hekanidb hekani

dropdb:
	docker exec -it hekani14 dropdb --username=hekanidb hekani

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# migrateup1:
# 	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# migratedown1:
# 	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# db_docs:
# 	dbdocs build doc/db.dbml

# db_schema:
# 	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

# server:
# 	go run main.go

# mock:
# 	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown
