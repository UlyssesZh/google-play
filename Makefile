.PHONY: build install tidy

PREFIX ?= /usr/local

tidy-build: tidy build

tidy:
	go mod tidy

build:
	go build -C play
	go build -C internal/play -o $(shell pwd)/bin/play
	go build -C internal/badging -o $(shell pwd)/bin/badging

install: build
	mkdir -p $(PREFIX)/bin
	install -m 0755 build/play $(PREFIX)/bin/play
	install -m 0755 build/badging $(PREFIX)/bin/badging
