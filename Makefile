run:
	go fmt && go test . && go mod tidy && go run .

test:
	go test ./... -coverprofile=coverage.out

test-cover:
	go tool cover -html=coverage.out