package game

type Board struct{}

type Position struct{}

func NewBoard() *Board {
	return &Board{}
}

func FreePositions(b *Board) []Position {
	return make([]Position, 0)
}

func ClaimedPositions(b *Board) []Position {
	return make([]Position, 0)
}

func IsPositionClaimed(b *Board, p Position) bool {
	return false
}

func PlayerPositions(b *Board, pID PID) []Position {
	return make([]Position, 0)
}

func PlayerTerritory(b *Board, pID PID) []Position {
	return make([]Position, 0)
}

func PlayerArea(b *Board, pID PID) []Position {
	return make([]Position, 0)
}

func PlayerPrisoners(b *Board, pID PID) uint {
	return 0
}

func PositionLiberties(b *Board, p Position) (liberties []Liberty, count uint) {
	return make([]Liberty, 0), 0
}
