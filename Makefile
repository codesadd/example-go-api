test:
	@echo "Start testing API..."
	go test -v -cover ./...
	@echo "Done testing API..."

build:
	@echo "Building app binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o app ./cmd/api
	@echo "Done!"

.PHONY: test build