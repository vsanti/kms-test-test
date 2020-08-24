#!/usr/bin/env bash
#
# makefile for Kyani Micro Service (kms) project
#

BUILD := $(shell date -u +%y%m%d%H%M)
KMS_NAME := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))
TEST_PORT := 5000

GOPRIVATE := github.com/kyani-inc
export GOPRIVATE

.PHONY: proto test
default: help

help:
	@echo "Kyani Inc. Microservice $(KMS_NAME)"
	@echo ""
	@echo "make [help]     This help screen."
	@echo "make migrate    Update DB_MIGRATE_DSN with latest changes."
	@echo "make api        Build and run the api service."
	@echo "make build-api  Build api service."
	@echo "make run-api    Run the api service."
	@echo "make tests      Run all tests."
	@echo ""
	@echo "See readme for more information: https://github.com/kyani-inc/$(KMS_NAME)"
	@echo ""
	@echo "Local development: copy env-sample as env and adjust variables as appropriate."
	@echo "Run 'make migrate' (if this service has a db). Run 'make api'."

migrate:
	@go run src/migrate/main.go

api: build-api run-api

# Build proto services
proto:
	cd $(shell go list -f '{{ .Dir }}' -m github.com/kyani-inc/proto) && \
	make proto

build-api:
	@cd src && go build -i -o ../bin/$(KMS_NAME) -ldflags "-X main.BUILD=$(BUILD) -X main.AppName=$(KMS_NAME) -X main.ENV=local"

run-api:
	@./bin/$(KMS_NAME)

test:
	@make build-api
	@KMS_NAME=$(KMS_NAME) TEST_PORT=$(TEST_PORT) sh .scripts/test.sh
