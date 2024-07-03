SCRIPT_PATH := ./scripts/run_dev.sh

.PHONY: dev
dev:
	@$(SCRIPT_PATH)

BINARY_NAME=pdf_manager

clean:
	rm -f ./tmp/bin/$(BINARY_NAME)*

build-debug: clean
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o ./tmp/bin/$(BINARY_NAME)-debug ./cmd/pdf_manager/main.go