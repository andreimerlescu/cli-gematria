BINARY_NAME=gematria-cli
BIN_DIR=bin
VERSION_FILE=VERSION
SOURCES=main.go
VERSION_FILE=VERSION
VERSION=$(shell cat $(VERSION_FILE) | tr -d '[:space:]')
PLATFORMS=linux windows darwin
ARCHS=amd64 arm64

.PHONY: all
all: clean build

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

.PHONY: build
build: $(PLATFORMS)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@echo "Building for $@ with version $(VERSION)..."
	@mkdir -p $(BIN_DIR)
	@for arch in $(ARCHS); do \
		GOOS=$@ GOARCH=$$arch go build -o $(BIN_DIR)/$(BINARY_NAME)-$(VERSION)-$@-$$arch $(SOURCES); \
		if [ $$? -eq 0 ]; then \
			echo "  Built $(BIN_DIR)/$(BINARY_NAME)-$(VERSION)-$@-$$arch"; \
		else \
			echo "  Failed to build for $@-$$arch"; \
		fi; \
	done
	@if [ "$@" = "windows" ]; then \
		for arch in $(ARCHS); do \
			mv $(BIN_DIR)/$(BINARY_NAME)-$(VERSION)-$@-$$arch $(BIN_DIR)/$(BINARY_NAME)-$(VERSION)-$@-$$arch.exe; \
			echo "  Renamed to $(BIN_DIR)/$(BINARY_NAME)-$(VERSION)-$@-$$arch.exe"; \
		done; \
	fi

	
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go get github.com/andreimerlescu/gematria
	@go get github.com/fatih/color
	@go get github.com/olekukonko/tablewriter

.PHONY: run
run:
	@go run $(SOURCES)