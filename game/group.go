package game

type Group struct{}

func NewGroup() *Group {
	return &Group{}
}

func GroupPositions(g *Group) []Position {
	return make([]Position, 0)
}

func GroupLiberties(g *Group) (liberties []Liberty, count uint) {
	return make([]Liberty, 0), 0
}

func GroupArea(g *Group) []Position {
	return make([]Position, 0)
}
