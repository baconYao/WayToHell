package matchmaking

// MatchType defines the interface for matching logic
type MatchType interface {
	Match(matcher *Individual, matchees []*Individual) (*Individual, error) // Returns the best matchee
}
