BINARY=bin/main

build:
	go build -o ${BINARY} cmd/*.go

web: build
	./${BINARY}
