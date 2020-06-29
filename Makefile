
test:
	cd server && go test -coverprofile cp.out ./...
	cd server && go tool cover -html=cp.out

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



