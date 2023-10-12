MAIN_PACKAGE_PATH := ./cmd/goofy
BINARY_NAME := goofy

.PHONY: build
build:
	go build -o=${BINARY_NAME} ${MAIN_PACKAGE_PATH}
