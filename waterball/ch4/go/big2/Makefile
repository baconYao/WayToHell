# Test case name (no .in/.out), e.g. make run NAME=normal-no-error-play1
NAME ?= normal-no-error-play1

.PHONY: run run-all dev build list help

# Show help (default target)
help:
	@echo "Big2 - Make targets"
	@echo ""
	@echo "  make help          Show this help"
	@echo "  make list          List all test case names (use with make run NAME=<name>)"
	@echo "  make run           Run one test: stdin from tests/<NAME>.in, diff against tests/<NAME>.out (default NAME=normal-no-error-play1)"
	@echo "  make run NAME=<n>  Example: make run NAME=straight"
	@echo "  make run-all       Run all tests in order; exit non-zero if any fails"
	@echo "  make dev           Run game reading from stdin (debug)"
	@echo "  make build         Build binary to ./api"
	@echo ""

# Run one test: go run ./cmd/api < tests/<NAME>.in | diff - tests/<NAME>.out
run:
	go run ./cmd/api < tests/$(NAME).in | diff - tests/$(NAME).out

# Run all tests (NAME is ignored)
run-all:
	@for f in tests/*.in; do \
		n=$$(basename "$$f" .in); \
		echo "=== $$n ==="; \
		go run ./cmd/api < "tests/$$n.in" | diff - "tests/$$n.out" || exit 1; \
	done && echo "All passed."

# List all test case names
list:
	@for f in tests/*.in; do basename "$$f" .in; done

# Run game reading from stdin (debug)
dev:
	go run ./cmd/api

build:
	go build -o api ./cmd/api
