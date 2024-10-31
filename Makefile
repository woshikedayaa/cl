# article name
NAME = cl
OUTPUT = $(WORKDIR)/bin/$(NAME)
WORKDIR = .
# injures
VERSION = $(shell cat $(WORKDIR)/VERSION)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_HASH = $(shell git rev-parse --short HEAD)
# build args
VERSION_FLAG= -X github.com/woshikedayaa/$(NAME)/cmd/$(NAME)/version.Version=$(VERSION)
BRANCH_FLAG= -X github.com/woshikedayaa/$(NAME)/cmd/$(NAME)/version.Branch=$(BRANCH)
COMMIT_HASH_FLAG= -X github.com/woshikedayaa/$(NAME)/cmd/$(NAME)/version.CommitHash=$(COMMIT_HASH)

LDFLAGS = $(VERSION_FLAG) $(BRANCH_FLAG) $(COMMIT_HASH_FLAG) -s -w

GOOS = $(shell go env GOOS)
ifeq ($(GOOS),windows)
OUTPUT:=$(OUTPUT).exe
endif
PARAMS = -trimpath -ldflags "$(LDFLAGS)" -v -o $(OUTPUT)
# Cgo
CGO=1
build: clean
	CGO_ENABLED=$(CGO) go build $(PARAMS)  $(WORKDIR)
.PHONY: build

clean:
	@rm -rf $(OUTPUT)
.PHONY: clean
