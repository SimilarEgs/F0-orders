.PHONY:

# ====================================================================================================================================
# DB migration
# migrate create -ext sql -dir migrations -seq init

DB_NAME = orders_db
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable


migrate-force:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations force 1

migrate-up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations up 

migrate-down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations down 1