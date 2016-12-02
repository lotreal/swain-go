PROJECT = github.com/lotreal/swain-go
VERSION = 0.0.1

OWNER := $(dir $(PROJECT))
OWNER := $(notdir $(OWNER:/=))
NAME := $(notdir $(PROJECT))


GO ?= go
GOPATH := $(CURDIR)

BUILD := $(CURDIR)/build/$(NAME)


run: build
	@echo "==> Running ${BUILD}..."
	@$(BUILD) --consul=123

.PHONY: build
build:
	@echo "==> Building ${PROJECT}..."
	@env -i \
	PROJECT=$(PROJECT) \
	VERSION=$(VERSION) \
	NAME=$(NAME) \
	scripts/compile.sh
