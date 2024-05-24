
.PHONY : tools mock docs

all: example

mod:
	go mod tidy

example: mod
	mkdir -p build
	go build -o example/example ./example/*.go
