include .env
GO_BIN := $(GOPATH)/bin
#GOIMPORTS := go run golang.org/x/tools/cmd/goimports@v0.1.11
#GOFUMPT := go run mvdan.cc/gofumpt@v0.3.1
#GOLINES := go run github.com/segmentio/golines@v0.11.0
OAPI_CODEGEN := $(GO_BIN)/oapi-codegen
MERGED_OAPI_V1=$(PWD)/api/openapi/v1/merged.json
OAPI_MERGER := $(GO_BIN)/oapi-merger
GOLANGCI := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5

init: \
	docker-clean \
	postgres-init \
	docs \
	up

init-prod: \
	docker-clean \
	postgres-init \
	docs \
	lint \
	up


up: \
	stop \
	docker-up

stop:
	docker-compose stop

docker-clean:
	docker-compose down --volumes --remove-orphans --rmi local

docker-up:
	docker-compose up --build --remove-orphans

build:
	docker-compose build

.PHONY: postgres-init
postgres-init:
	docker-compose up -d postgres

.PHONY: docs
docs: openapi_merge openapi_http

.PHONY: openapi_merge
openapi_merge: $(OAPI_MERGER)
	$(OAPI_MERGER) -wdir api/openapi/v1 -spec openapi.yaml -o $(MERGED_OAPI_V1)

.PHONY: openapi_http
openapi_http: $(OAPI_CODEGEN)
	$(OAPI_CODEGEN) --old-config-style  -generate types,skip-prune -o ./internal/controllers/v1/view/types.gen.go -package view $(MERGED_OAPI_V1)
	$(OAPI_CODEGEN) --old-config-style -generate spec -o ./internal/controllers/v1/spec.gen.go -package v1 $(MERGED_OAPI_V1)
	rm -f $(MERGED_OAPI_V1)

$(OAPI_MERGER):
	go install github.com/felicson/oapi-merger/cmd/oapi-merger@v0.0.2
$(OAPI_CODEGEN):
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0

.PHONY: lint
lint:
	$(GOLANGCI) run

generate_models:
	tables-to-go -fn-format s -u ${POSTGRES_USER} -p ${POSTGRES_PASSWORD} -d ${POSTGRES_DB} -v -of ./models