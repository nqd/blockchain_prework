TEST ?= go test
GO ?= go

.PHONY: test

test:
	$(TEST) -v ./...