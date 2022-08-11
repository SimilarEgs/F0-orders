.PHONY:

# ====================================================================================================================================
# DB migration

DB_NAME = orders_db
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable


migrate-up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations up 1

migrate-down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations down 1