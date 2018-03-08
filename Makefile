TEST ?= go test
GO ?= go
GET ?= go get

.PHONY: test

# Install the development dependencies.
install.deps:
	@$(GO) get github.com/davecgh/go-spew/spew

# Test
test:
	$(TEST) -v ./...