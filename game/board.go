package game

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNeutralClaim  = errors.New("board: neutral can not claim position")
	ErrPositionTaken = errors.New("board: position taken")
)

const (
	BoardSize9  int = 9
	BoardSize11 int = 11
	BoardSize16 int = 16
	BoardSize19 int = 19
)

// Board represents the current governance of all positions on a Go board.
type Board [][]governance

// position represents a position on a coordinate system.
type position struct {
	x int
	y int
}

// UNIMPLEMENTED
func Territory(b Board, g governance) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func Area(b Board, g governance) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func PositionLiberties(b Board, p position) ([]liberty, int) {
	/*
	   Algorithm:
	       For each direction
	       createa a liberty
	       check if it is valid (end is inbound)
	       check liberty end governance
	           - Neutral: add liberty to response
	           - Ally piece: extend liberty (DFS search for does ally have any liberty?)
	           - Enemy piece: dont add liberty
	*/
	return make([]liberty, 0), 0
}

// NewBoard creates an empty board of size by size.
func NewBoard(size int) Board {
	b := make(Board, size)
	for i0 := 0; i0 < size; i0++ {
		b[i0] = make([]governance, size)
	}
	return b
}

// Positions retrieves the positions on a board matching a certain governance.
func Positions(b Board, g governance) []position {
	size := len(b)
	positions := make([]position, 0)
	for i0 := 0; i0 < size; i0++ {
		for i1 := 0; i1 < size; i1++ {
			if b[i0][i1] == g {
				positions = append(positions, NewPosition(i0, i1))
			}
		}
	}
	return positions
}

// ClaimPosition computes the next board by claiming a position for a governance.
func ClaimPosition(b Board, p position, g governance) (Board, error) {
	// TODO: Rules
	// [x] Position not taken
	// [ ] Suicide rule
	// [ ] Ko rule
	if g == GovNeutral {
		return nil, ErrNeutralClaim
	}

	err := b.checkIndexSafe(p)
	if err != nil {
		return nil, err
	}

	if isPositionClaimed(b, p) {
		return nil, ErrPositionTaken
	}

	return b.seize(p, g), nil
}

// isPositionClaimed checks that a position has a non-neutral governance.
func isPositionClaimed(b Board, p position) bool {
	return b[p.x][p.y] != GovNeutral
}

// seize produces the resulting board of placing a governance on a position on the board.
func (b Board) seize(p position, g governance) Board {
	size := len(b)
	rB := make(Board, size)
	for i0 := 0; i0 < size; i0++ {
		rB[i0] = make([]governance, size)
		copy(rB[i0], b[i0])
	}
	rB[p.x][p.y] = g
	return rB
}

// checkIndexSafe checks that a position is inside the bounds of the board.
func (b Board) checkIndexSafe(p position) error {
	if p.x >= len(b) || p.y >= len(b) {
		return errors.New("Trying to index out of bound of a board")
	}
	return nil
}

// governance retrieves the governance for a position on the board.
func (b Board) governance(p position) governance {
	return b[p.x][p.y]
}

func NewPosition(x, y int) position {
	return position{x, y}
}

// Utils

// Display shows the governance of a board in a rich and structured format.
func (b Board) Display() {
	positionFmt := "[%v]\t"
	l := len(b)

	var header strings.Builder
	header.WriteString(fmt.Sprintf(positionFmt, "*"))
	for i0 := 0; i0 < l; i0++ {
		header.WriteString(fmt.Sprintf(positionFmt, i0))
	}
	fmt.Println(header.String())

	for i0 := 0; i0 < l; i0++ {
		var line strings.Builder
		for i1 := 0; i1 < l; i1++ {
			if i1 == 0 {
				line.WriteString(fmt.Sprintf(positionFmt, i0))
			}
			line.WriteString(fmt.Sprintf(positionFmt, b[i0][i1].string()))
		}
		fmt.Println(line.String())
	}
}
