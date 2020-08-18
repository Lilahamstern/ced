init:
	docker network create -d bridge ced_server


test-unit:
	cd server && go test -v -short -coverprofile cp.out ./...

test-unit-coverage: test-unit
	cd server && go tool cover -html=cp.out

stores-up: db-up

stores-down: db-down

db-up:
	docker-compose up -d db

db-down:
	docker-compose stop db

swag-gen:
	swag i -g ./cmd/ced/main.go -d ./server -o ./server/docs

