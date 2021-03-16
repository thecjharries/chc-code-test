# Run provided tests
test:
	go test -v ./... -cover -race -coverprofile=.coverage.out

# Get test coverage
coverage: test
	go tool cover -func=.coverage.out

# Execute the program
run:
	go run main.go
