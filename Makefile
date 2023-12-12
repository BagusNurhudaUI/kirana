build:
	@go build -o /bin/app

run: build
	@./bin/app

run-node: 
	nodemon --exec go run cmd/*.go 

test: 
	@go test -v ./...

test-profile:
	@go test -v ./... -coverprofile=profile.out
	
test-html:
	@go tool cover -html=profile.out -o coverage.html