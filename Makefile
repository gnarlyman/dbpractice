migrate:
	migrate -database postgres://myuser:${DBPASS}@localhost:5432/db1?sslmode=disable -path migrations/ up

migrate-seq:
	migrate create -ext sql -dir migrations -seq ${SEQ}
