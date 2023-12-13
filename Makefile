build:
	@go build -o ./bin/app ./cmd

run: build
	@echo "Running the application..."
	@./bin/app

run-node: 
	nodemon --exec go run cmd/main.go 

test: 
	@go test -v ./...

test-profile:
	@go test -v ./... -coverprofile=profile.out
	
test-html:
	@go tool cover -html=profile.out -o coverage.html