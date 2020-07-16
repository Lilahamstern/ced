
test-unit:
	cd server && go test -v -short -coverprofile cp.out ./...

test-unit-coverage: test-unit
	cd server && go tool cover -html=cp.out

generate-graphql:
	cd server/internal/resolver && go generate

stores-up: db-up

stores-down: db-down

db-up:
	docker-compose up -d db

db-down:
	docker-compose stop db

apq-down:
	docker-compose stop redis



