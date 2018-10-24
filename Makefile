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
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o $(BINARY_NAME) -v