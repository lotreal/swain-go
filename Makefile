PROJECT = github.com/lotreal/swain-go
VERSION = 0.0.1

OWNER := $(dir $(PROJECT))
OWNER := $(notdir $(OWNER:/=))
NAME := $(notdir $(PROJECT))


GO ?= go
GOPATH := $(CURDIR)

BUILD := $(CURDIR)/build/$(NAME)


run: build
	$(BUILD)

.PHONY: build
build:
	@echo "==> Building ${PROJECT}..."
	@PROJECT=$(PROJECT) \
	VERSION=$(VERSION) \
	NAME=$(NAME) \
	scripts/compile.sh
