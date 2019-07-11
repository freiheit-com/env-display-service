#!make
include Makefile.variables
export

getver:
	$(eval GIT_COMMIT=$(shell git log -1 --pretty=format:"%H"))

docker-tag: getver
	$(eval DOCKER_TAG="$(IMAGE_NAME):$(GIT_COMMIT)")

run:
	go run cmd/main.go

build-for-docker:
	mkdir -p deployments/tmp
	GOOS=linux GOARCH=amd64 go build -o deployments/tmp/env-display cmd/main.go

docker: build-for-docker docker-tag
	docker build --tag $(DOCKER_TAG) deployments
