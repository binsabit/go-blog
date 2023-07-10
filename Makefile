run:
	go run main.go -config=config.yaml

new-psql-migration:
	@read -p "Enter migration name: " name;\
		migrate create -ext sql -dir migrations/postgres -seq $$name

migrate-psql-up:
	migrate -path migrations/postgres \
		-database "postgresql://yerda:postgres@localhost:5432/dbblog?sslmode=disable" -verbose up
migrate-psql-down:
	migrate -path migrations/postgres \
		-database "postgresql://yerda:postgres@localhost:5432/dbblog?sslmode=disable" -verbose down
migrate-psql-force:
	@read -p "Enter version to force :" version;\
		migrate -path migrations/postgres \
			-database "postgresql://yerda:postgres@localhost:5432/dbblog?sslmode=disable" -verbose force $$version

gensql:
	sqlboiler psql
