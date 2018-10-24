BINARY_NAME = data-service

build:
	go build -o $(BINARY_NAME) -v

run:
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)