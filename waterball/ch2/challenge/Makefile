# Default match strategy
MATCH_STRATEGY ?= distance

# Valid match strategies
VALID_STRATEGIES := distance habit

# Check if MATCH_STRATEGY is valid
ifneq ($(filter $(MATCH_STRATEGY),$(VALID_STRATEGIES)),$(MATCH_STRATEGY))
$(error Invalid MATCH_STRATEGY: $(MATCH_STRATEGY). Must be one of: $(VALID_STRATEGIES))
endif

build:
	go build -o bin/matchmaking ./cmd/api
test:
	go test ./...
run:
	go run ./cmd/api --match_strategy=$(MATCH_STRATEGY)
dev:
	go run ./cmd/api --log_level DEBUG --match_strategy=$(MATCH_STRATEGY)
	