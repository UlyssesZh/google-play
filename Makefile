.PHONY: build install tidy tidy-build tidy-doc clean

PREFIX ?= /usr/local

tidy-build: tidy build

tidy-doc: tidy doc

tidy:
	go mod tidy

build:
	go build -C internal/play -o $(shell pwd)/bin/play
	go build -C internal/badging -o $(shell pwd)/bin/badging

doc:
	go tool vanitydoc vanitydoc.json
	echo '{"html": {"output": "doc/google-play"}}' | go tool vanitydoc -

install: build
	mkdir -p $(PREFIX)/bin
	install -m 0755 build/play $(PREFIX)/bin/play
	install -m 0755 build/badging $(PREFIX)/bin/badging

clean:
	[[ -d bin ]] && rm -r bin || true
	[[ -d doc ]] && rm -r doc || true
