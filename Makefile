.EXPORT_ALL_VARIABLES:

BIN_DIR := ./bin
OUT_DIR := ./output
$(shell mkdir -p $(BIN_DIR) $(OUT_DIR))

IMAGE_REGISTRY=zufardhiyaulhaq
IMAGE_NAME=$(IMAGE_REGISTRY)/goweekly
IMAGE_TAG=$(shell git rev-parse --short HEAD)

CURRENT_DIR=$(shell pwd)
VERSION=$(shell cat ${CURRENT_DIR}/VERSION)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_TREE_STATE=$(shell if [ -z "`git status --porcelain`" ]; then echo "clean" ; else echo "dirty"; fi)

STATIC_BUILD?=true

override LDFLAGS += \
  -X ${PACKAGE}.version=${VERSION} \
  -X ${PACKAGE}.buildDate=${BUILD_DATE} \
  -X ${PACKAGE}.gitCommit=${GIT_COMMIT} \
  -X ${PACKAGE}.gitTreeState=${GIT_TREE_STATE}

ifeq (${STATIC_BUILD}, true)
override LDFLAGS += -extldflags "-static"
endif

ifneq (${GIT_TAG},)
IMAGE_TAG=${GIT_TAG}
IMAGE_TRACK=stable
LDFLAGS += -X ${PACKAGE}.gitTag=${GIT_TAG}
else
IMAGE_TAG?=$(GIT_COMMIT)
IMAGE_TRACK=latest
endif

.PHONY: test
test:
	go get golang.org/x/tools/cmd/cover
	go test -coverprofile=./output/coverage.out -race ./...
	go tool cover -html=./output/coverage.out -o ./output/coverage.html

.PHONY: lint
lint: 
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
	golangci-lint run --verbose --timeout 300s

.PHONY: build
build:
	CGO_ENABLED=0 GO111MODULE=on go build -a -ldflags '${LDFLAGS}' -o ${BIN_DIR}/goweekly ./cmd/goweekly/main.go

.PHONY: run
run:
	go run cmd/goweekly/main.go

.PHONY: readme
readme:
	GO111MODULE=on go get github.com/norwoodj/helm-docs/cmd/helm-docs
	helm-docs -c ./charts/goweekly -d > README.md
	helm-docs -c ./charts/goweekly

.PHONY: helm.create.releases
helm.create.releases:
	helm package charts/goweekly --destination charts/releases
	helm repo index charts/releases
