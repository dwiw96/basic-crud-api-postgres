runBasic:
	go run cmd/classic/main.go

runMux:
	go run cmd/mux_pgxpool/main.go

compose:
	docker compose -f docker-compose.yml up

database:
	docker run --rm --name postgres_api -p 5432:5432 -e POSTGRES_USER=db -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=crudAPI postgres

postgres:
	docker exec -t postgres_api psql -U db crudAPI