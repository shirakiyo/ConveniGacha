api/run:
	docker-compose up -d

api/stop:
	docker-compose down

api/build:
	docker build -t conveni-api -f './cmd/api/Dockerfile' .

famima/scrape:
	@go run cmd/task/main.go --conveni famima

lawson/scrape:
	@go run cmd/task/main.go --conveni lawson

test:
	@go test -cover ./...
