package game

type Liberty struct{}

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewLiberty() Liberty {
	return Liberty{}
}

func Origin(l Liberty) Position {
	return Position{}
}

func End(l Liberty) Position {
	return Position{}
}

func GetDirection(l Liberty) Direction {
	return Up
}
