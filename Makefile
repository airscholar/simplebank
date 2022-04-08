# You can set these variables from the command line, and also
# from the environment.
DATABASE_M    	=
DB_USERNAME     = 
DB_URL			= 


createdb: 
	psql -U $(DB_USERNAME) -d $(DATABASE_M) -c "create database simplebank"

dropdb:
	psql -U $(DB_USERNAME) -d $(DATABASE_M) -c "drop database simplebank"

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test