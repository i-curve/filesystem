VERSION=$(shell git describe --tags --always | sed 's/^v//')

build:
	go build -o filesystem -ldflags "-X main.Version=$(VERSION)" main.go 

.PHONY: docker
docker: build
	docker image build -t filesystem:$(VERSION) .

.PHONY: docker-push
docker-push:
	docker tag filesystem:$(VERSION) wjuncurve/filesystem
	docker tag filesystem:$(VERSION) wjuncurve/filesystem:$(VERSION)
	docker push wjuncurve/filesystem:$(VERSION)
	docker push wjuncurve/filesystem

.PHONY: run-docker
run-docker: docker
	docker run -d -p 8000:8000 -p 8001:8001 -e "MODE=DEBUG" -e "MYSQL_HOST=host.docker.internal" filesystem:$(VERSION)
