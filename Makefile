IMAGE_REPOSITORY ?= "otelcol-custom"
IMAGE_TAG ?= "latest"
COMPONENT_NAME ?= "mytestconnector"

.PHONY: build
build:
	builder --config=testfiles/ocb-manifest.yaml

# This is required to build the a proper executable that can be included in a docker image for an arm64 architecture K8s cluster
.PHONY: build-arm
build-arm:
	env GOOS=linux GOARCH=arm64 builder --config=testfiles/ocb-manifest.yaml

.PHONY: generate-metadata
generate-metadata:
	cd $(COMPONENT_NAME) && mdatagen metadata.yaml

.PHONY: run
run:
	./dist/otelcol-custom --config=testfiles/collector-config.yaml

.PHONY: debug
debug:
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient --log exec ./dist/otelcol-custom -- --config=collector-config.yaml

.PHONY: test-span
test-span:
	otel-cli span --endpoint localhost:4317 --service "test" --name "Test Span" --attrs "app.component_name=$(COMPONENT_NAME)" --tp-print