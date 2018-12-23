GO        = go
GOFMT     = gofmt
GOIMPORTS = goimports
GOLINT    = golint

all: test

format:
	$(GOIMPORTS) -l -w *.go
	$(GOFMT) -l -w *.go

lint:
	$(GOLINT) *.go

test: format lint
	$(GO) test -coverprofile=cover.out

cover: test
	$(GO) tool cover -html=cover.out -o coverage.html
