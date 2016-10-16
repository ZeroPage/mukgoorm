build:
	go build -o mukgoorm .
clean:
	go clean
install:
	go get github.com/zeropage/mukgoorm
test:
	go test ./...
server: build
	./mukgoorm
dev-server:
	sentry -c "make test && make server" -w "*.go" -w "templates"
install-tools:
	go get "github.com/bluemir/sentry"

.PONEY: clean build test server dev-server install-tools install
