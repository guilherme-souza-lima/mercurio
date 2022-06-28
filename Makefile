up:
	docker-compose up -d
down:
	docker-compose down
env:
	cp ./.env.example ./.env
run:
	go run main.go httpserver
