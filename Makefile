tidy: go.mod
	@go mod tidy

verify: go.mod
	@go mod verify

build: verify
	go build -o build/qr-generator cmd/main.go

build-linux: verify
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(MAKE) build

docker:
	docker build -t qr-generator .
	docker-compose up

mongo:
	docker run -p 27017:27017 --name mongo mongo