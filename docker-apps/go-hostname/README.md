## Simple Go Webserver

This container returns the hostname of the container with a timestamp
Corresponds to the docker image sarun87/go-hostname

Entrypoint: /server

### To compile & build docker image

Run `make` from this directory

### Running the container

When starting a container specify port using environment variable `PORT`

```bash
docker run -d --name=goserver -e "PORT=80" sarun87/go-hostname:v1 /server
```
