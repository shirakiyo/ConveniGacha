api/run:
	go run cmd/api/main.go

task/run:
	go run cmd/task/main.go

test: 
	@go test -cover ./...
