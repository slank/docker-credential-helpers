.PHONY: all awsecr deps test validate

all: test

awsecr:
	mkdir -p bin
	go build -o bin/docker-credential-awsecr awsecr/cmd/main.go

deps:
	go get github.com/golang/lint/golint

test:
	# tests all packages except vendor
	go test -v `go list ./... | grep -v /vendor/`

vet_subhelper:
	go vet ./subhelper

validate: vet
	for p in `go list ./... | grep -v /vendor/`; do \
		golint $$p ; \
	done
	gofmt -s -l `ls **/*.go | grep -v vendor`
