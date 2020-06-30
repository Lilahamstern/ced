
test-unit:
	cd server && go test -v -short -coverprofile cp.out ./...

test-unit-coverage: test-unit
	cd server && go tool cover -html=cp.out

generate-graphql:
	cd server/internal/resolver && go generate

stores-up: db-up apq-up

stores-down: db-down apq-down

db-up:
	docker-compose up -d db

apq-up:
	docker-compose up -d redis

db-down:
	docker-compose stop db

apq-down:
	docker-compose stop redis



