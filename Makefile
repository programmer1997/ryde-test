DB_NAME = rydedb
COLLECTION_NAME = users

db-setup:
	mongosh --eval 'db.getSiblingDB("$(DB_NAME)").createCollection("$(COLLECTION_NAME)")'

start-server:
	cd cmd && go run main.go