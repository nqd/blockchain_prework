GO ?= go

.PHONY: test install.deps
.DEFAULT_GOAL := test

# Install the development dependencies.
install.deps:
	@$(GO) get github.com/davecgh/go-spew/spew

# Test
test:
	$(GO) test -v ./...