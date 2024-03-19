initmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

updb:
	migrate -path db/migration -database "postgres://postgres:postgres123@localhost:5432/fp_gdsc?sslmode=disable" up

downdb:
	migrate -path db/migration -database "postgres://postgres:postgres123@localhost:5432/fp_gdsc?sslmode=disable" down

forcedb:
	migrate -path db/migration -database "postgres://postgres:postgres123@localhost:5432/fp_gdsc?sslmode=disable" force 000001