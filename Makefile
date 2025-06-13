run:
	go run main.go

build:
	go build -o app main.go

test:
	go test ./...

lint:
	go vet ./...

dev:
	APP_ENV=development go run main.go
