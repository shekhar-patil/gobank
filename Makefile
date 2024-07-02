build:
	go build -o bin/gobank

run: 
	go run main.go

test:
	go test -v ./..
