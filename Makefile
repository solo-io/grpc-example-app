#----------------------------------------------------------------------------------
# Base
#----------------------------------------------------------------------------------
OUTDIR?=_output
PROJECT?=store

PROJECT_IMAGE?=soloio/$(PROJECT)

SOURCES := $(shell find . -name "*.go" | grep -v test.go)

.PHONY: generate-protos
generate-protos:
	protoc -I./api/ --include_source_info --go_out=plugins=grpc:api --include_imports --descriptor_set_out=api/store/books_descriptors.pb ./api/store/books.proto
	protoc -I./api/ --include_source_info --go_out=plugins=grpc:api --include_imports --descriptor_set_out=api/store/records_descriptors.pb ./api/store/records.proto
