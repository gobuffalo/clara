TAGS ?= ""
GO_BIN ?= "go"

install:
	$(GO_BIN) install -tags ${TAGS} -v ./.
	make tidy

tidy:
	$(GO_BIN) mod tidy
	echo skipping go mod tidy

deps:
	$(GO_BIN) get -tags ${TAGS} -t ./...
	make tidy

build:
	$(GO_BIN) build -v .
	make tidy

test:
	$(GO_BIN) test -cover -tags ${TAGS} ./...
	make tidy

packr:
	$(GO_BIN) get github.com/gobuffalo/packr/v2/packr2
	packr2
	make tidy

