.PHONY: init build tidy test bash
GO_VERSION := 1.17
PLUGIN_NAME := cached-router

build:
	docker run -v "${PWD}:/app" golang:$(GO_VERSION) bash -c "cd /app && go build -o ./build/go-swaggit ./" 

tidy:
	docker run -v "${PWD}:/app" golang:$(GO_VERSION) bash -c "cd /app && go mod tidy" 

test:
	docker run -v "${PWD}:/app" golang:$(GO_VERSION) bash -c "cd /app && go test -v ./..." 

bash:
	docker run -v "${PWD}:/app" -it golang:$(GO_VERSION) bash