BINARY_NAME = data_service_runner

build:
	go get -t -v
	go build -o $(BINARY_NAME) -v

run:
	go get -t -v
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build-linux:
	go get -t -v
	shell export GOOS=linux
	shell export GOARCH=amd64
	go build -o $(BINARY_NAME) -v