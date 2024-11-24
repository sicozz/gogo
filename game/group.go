package game

// group represents a group of connected positions on a board.
type group struct{}

func newGroup() *group {
	return &group{}
}

// UNIMPLEMENTED
func GroupTerritory(g *group) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func GroupLiberties(g *group) (liberties []liberty, count int) {
	return make([]liberty, 0), 0
}

// UNIMPLEMENTED
func GroupArea(g *group) []position {
	return make([]position, 0)
}
