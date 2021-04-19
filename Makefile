run:
	go build .
	./go-fiber-auth-test

install:
	go mod tidy

test:
	go test --failfast