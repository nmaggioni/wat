GO_BUILD_CMD = go build -v
GOLANG_IMAGE_TAG = alpine

.PHONY: all go docker

all: go

go:
	${GO_BUILD_CMD}
	@file ./wat

docker:
	docker run --rm -it -v "${PWD}:/src" "golang:${GOLANG_IMAGE_TAG}" sh -c "cd /src && ${GO_BUILD_CMD}"
	@file ./wat