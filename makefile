test:
	go test -p 1 -race ./...  -coverprofile coverage.out

coverage:
	go tool cover -html=coverage.out