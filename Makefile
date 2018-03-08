GO ?= go

.PHONY: test

# Install the development dependencies.
install.deps:
	@$(GO) get github.com/davecgh/go-spew/spew

# Test
test:
	$(GO) test -v ./...