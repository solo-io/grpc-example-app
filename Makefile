.PHONY: all
all: generate-protos generated-code install/grpc-example-app-default.yaml

.PHONY: generate-protos
fmt:
	goimports -w $(shell ls -d */ | grep -v vendor)

.PHONY: generate-protos
generate-protos:
	protoc -I./api/ --include_source_info --go_out=plugins=grpc:api --include_imports --descriptor_set_out=api/store/books_descriptors.pb ./api/store/books.proto
	protoc -I./api/ --include_source_info --go_out=plugins=grpc:api --include_imports --descriptor_set_out=api/store/records_descriptors.pb ./api/store/records.proto

.PHONY: generated-code
generated-code:
	go run generate.go

.PHONY: manifest
manifest: install/grpc-example-app-default.yaml
install/grpc-example-app-default.yaml: $(shell find install/helm)
	helm template --include-crds --namespace grpc-example-app install/helm/store > $@
