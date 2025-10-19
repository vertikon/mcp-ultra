SHELL := /bin/bash

PKG=./...
COVER=coverage.out

.PHONY: all fmt lint test cover coverhtml vet tidy ci mocks

all: fmt tidy lint test

fmt:
	go fmt ./...
	goimports -w .

tidy:
	go mod tidy

lint:
	golangci-lint run

test:
	go test $(PKG) -race -count=1 -coverprofile=$(COVER)

cover:
	@go tool cover -func=$(COVER) | tail -n1

coverhtml:
	go tool cover -html=$(COVER) -o coverage.html

vet:
	go vet $(PKG)

mocks:
	bash scripts/regenerate_mocks.sh

ci: fmt tidy vet lint test cover
	@TOTAL=$$(go tool cover -func=$(COVER) | grep total | awk '{print $$3}' | sed 's/%//'); \
	MIN=80; \
	echo "Total coverage: $$TOTAL% (min $$MIN%)"; \
	A=$${TOTAL%.*}; if [ $$A -lt $$MIN ]; then echo "Coverage gate failed"; exit 1; fi
