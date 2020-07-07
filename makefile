api/run:
	docker-compose up -d

api/stop:
	docker-compose down

api/build:
	docker build -t conveni-api -f './cmd/api/Dockerfile' .

task/run:
	go run cmd/task/main.go

test: 
	@go test -cover ./...
