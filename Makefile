ENTRYPOINT=./cmd/gowc
BINARY_PATH=./bin/gowc

run:
	go run $(ENTRYPOINT)

build:
	go build -o $(BINARY_PATH) $(ENTRYPOINT)


