# Project settings
APP_NAME := 3ml
BIN_DIR := /bin
BUILD_DIR := build
GO_FILES := $(shell find . -type f -name '*.go')

.PHONY: all build install clean

all: build

build: $(GO_FILES)
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) .

install: build
	@echo "==> Installing $(APP_NAME) to $(BIN_DIR)..."
	@sudo cp $(BUILD_DIR)/$(APP_NAME) $(BIN_DIR)/$(APP_NAME)
	@echo "==> Installed successfully."

clean:
	@echo "==> Cleaning up..."
	@rm -rf $(BUILD_DIR)