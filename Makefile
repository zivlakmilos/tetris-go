all: run

clean:
	@go clean
	@rm -Rf bin/*

run: build
	@./bin/tetris

build:
	@GOOS=linux go build -o bin/tetris ./cmd/tetris
