migrate:
	migrate -database postgres://myuser:${DBPASS}@localhost:5432/db1?sslmode=disable -path migrations/ up

migrate-seq:
	migrate create -ext sql -dir migrations -seq ${SEQ}

sql-gen:
	sqlc generate

go-generate:
	go generate ./...

api-gen:
	oapi-codegen swagger.yaml > swagger/dbpractice.gen.go
