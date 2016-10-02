build:
	go build .
clean:
	go clean
test:
	go test

.PONEY: clean build test
