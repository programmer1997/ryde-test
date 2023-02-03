DB_NAME = rydedb
COLLECTION_NAME = users

db-setup:
	mongosh --eval 'db.getSiblingDB("$(DB_NAME)").createCollection("$(COLLECTION_NAME)")'

start-server:
	 go mod download && cd cmd && go run main.go

run-tests:
	cd internal/router && go test