package hero

type LevelSheet struct{}

// NewLevelSheet creates a new LevelSheet with default values.
func NewLevelSheet() *LevelSheet {
	return &LevelSheet{}
}

// QueryLevel calculates the level based on total experience points.
func (ls LevelSheet) QueryLevel(totalExp int) int {
	return totalExp/1000 + 1
}
