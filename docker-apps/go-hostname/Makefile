
VERSION ?= v1

default: all

all: server docker-image

clean:
	rm -f server
	sudo docker rmi sarun87/go-hostname:$(VERSION) || true

.PHONY: server

server: 
	GOPATH=$(shell pwd) \
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=0 \
	go build \
		-v \
		--buildmode=exe \
		-o server \
		server
docker-image: server
	sudo docker build -t sarun87/go-hostname:$(VERSION) .

