# Default card game strategy
CARD_GAME_STRATEGY ?= showdown

# Valid card game strategies
VALID_STRATEGIES := showdown uno

# Check if CARD_GAME_STRATEGY is valid
ifneq ($(filter $(CARD_GAME_STRATEGY),$(VALID_STRATEGIES)),$(CARD_GAME_STRATEGY))
$(error Invalid CARD_GAME_STRATEGY: $(CARD_GAME_STRATEGY). Must be one of: $(VALID_STRATEGIES))
endif

build:
	go build -o bin/cardkit ./cmd/api
run:
	go run ./cmd/api --card_game_type=$(CARD_GAME_STRATEGY)
dev:
	go run ./cmd/api --log_level DEBUG --card_game_type=$(CARD_GAME_STRATEGY)