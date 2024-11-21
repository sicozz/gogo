package game

import (
	"errors"
	"fmt"
	"strings"
)

var (
    ErrNeutralClaim = errors.New("board: neutral can not claim position")
    ErrPositionTaken = errors.New("board: position taken")
)

const (
    BoardSize9 int = 9
    BoardSize11 int = 11
    BoardSize16 int = 16
    BoardSize19 int = 19
)

type Board [][]governance

type position struct {
	x int
	y int
}

// UNIMPLEMENTED
func NeutralArea(b Board) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func ClaimedArea(b Board) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func PlayerTerritory(b Board, g governance) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func PlayerArea(b Board, g governance) []position {
	return make([]position, 0)
}

// UNIMPLEMENTED
func PlayerPrisoners(b Board, g governance) int {
	return 0
}

// UNIMPLEMENTED
func PositionLiberties(b Board, p position) (liberties []liberty, count int) {
	return make([]liberty, 0), 0
}

func NewBoard(size int) Board {
	b := make(Board, size)
	for i0 := 0; i0 < size; i0++ {
		b[i0] = make([]governance, size)
	}
	return b
}

func ClaimPosition(b Board, p position, g governance) (Board, error) {
	// TODO: Rules
	// [ ] Position not taken
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

func isPositionClaimed(b Board, p position) bool {
	return b[p.x][p.y] != GovNeutral
}

func (b Board) seize(p position, g governance) Board {
	rB := NewBoard(len(b))
	rB[p.x][p.y] = g
	return rB
}

func (b Board) checkIndexSafe(p position) error {
	if p.x >= len(b) || p.y >= len(b) {
		return errors.New("Trying to index out of bound of a board")
	}
	return nil
}

func (b Board) governance(p position) governance {
    return b[p.x][p.y]
}

func NewPosition(x, y int) position {
	return position{x, y}
}

// Utils

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
