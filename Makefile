# start postgres container
.PHONY: START_PG

START_PG:
	docker-compose --file ./internal/repo/postgres/docker-compose.yml up -d

# psql connection string (limited user):
# 	export PGPASSWORD='appp@$$w0rd'; psql -h 127.0.0.1 -p 5432 -U appuser -d app_db