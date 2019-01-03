GO        = go
GOFMT     = gofmt
GOIMPORTS = goimports
GOLINT    = golint

all: test

format:
ifeq ($(CI_TEST),true)
	@imports=`$(GOIMPORTS) -l *.go`; \
	fmts=`$(GOFMT) -l -s *.go`; \
	all=`(echo $$imports; echo $$fmts) | sort -u`; \
	if [ "$$all" != "" ]; then \
		echo "Following files need updates:"; \
		echo; \
		echo $$all; \
		exit 1;\
	fi
else
	$(GOIMPORTS) -l -w *.go
	$(GOFMT) -l -s -w *.go
endif

lint:
	$(GOLINT) -set_exit_status *.go

vet:
	$(GO) vet

test: format lint vet
	$(GO) test -coverprofile=cover.out

cover: test
	$(GO) tool cover -html=cover.out -o coverage.html
