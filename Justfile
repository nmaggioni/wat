#!/usr/bin/env just --justfile

GO_BUILD_CMD := "go build -v"
GOLANG_IMAGE_TAG := "alpine"

# Interactively choose the recipe to run
_default:
    @just --choose

# Build (natively)
go:
    {{ GO_BUILD_CMD }}
    @file ./wat

# Build (through a Docker container)
docker:
    docker run --rm -it -v "${PWD}:/src" "golang:{{ GOLANG_IMAGE_TAG }}" sh -c "cd /src && {{ GO_BUILD_CMD }}"
    @file ./wat
