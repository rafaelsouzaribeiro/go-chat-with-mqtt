create-keyspace:
	go run  internal/infra/database/cassandra/migrations/keyspace/main.go

migrateup: 
	go run internal/infra/database/cassandra/migrations/up/main.go 

migratedown:
	go run  internal/infra/database/cassandra/migrations/down/main.go

.PHONY: migrateup migratedown create-keyspace
