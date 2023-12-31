VERSION=$(shell git describe --tags --always)

build:
	go build -o filesystem -ldflags "-X main.Version=$(VERSION)" main.go 

.PHONY: docker
docker: build
	docker image build -t filesystem:$(VERSION) .

