MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make
ROOT := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export PATH := $(ROOT)/scripts:$(PATH)

.PHONY: build

gen-proto:
	mkdir -p proto/pb
	rm -rf proto/pb/*
	protoc --proto_path=proto/. --go-grpc_opt require_unimplemented_servers=false,paths=source_relative --go-grpc_out proto/pb/ --go_opt paths=source_relative --go_out proto/pb/ proto/*.proto

gen-injector:
	cd di && wire

gen: gen-proto gen-injector

run-local-api:
	go run main.go

run-local-grpcui:
	grpcui -plaintext -port 4000 localhost:3000
