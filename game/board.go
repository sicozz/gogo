package game

import (
	"container/list"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNeutralClaim  = errors.New("board: neutral can not claim position")
	ErrPositionTaken = errors.New("board: position taken")
	ErrPositionOOB   = errors.New("Trying to index out of bound of a board")
	ErrPositionNeg   = errors.New("Trying to index at a negative coordinate")
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

// positions represents a collection of position
type positions []position

// UNIMPLEMENTED
func Territory(b Board, g governance) positions {
	return make(positions, 0)
}

// UNIMPLEMENTED
func Area(b Board, g governance) positions {
	return make(positions, 0)
}

// NewBoard creates an empty board of size by size.
func NewBoard(size int) Board {
	b := make(Board, size)
	for i0 := 0; i0 < size; i0++ {
		b[i0] = make([]governance, size)
	}
	return b
}

// PositionLiberties computes the liberties of a position
func PositionLiberties(b Board, p position) []liberty {
	g := b.governance(p)
    if g == GovNeutral {
        return nil
    }
	libertyCandidateCriteria := func(p0 position) bool {
		g0 := b.governance(p0)
		return g0 == g || g0 == GovNeutral
	}
	libertyCandidates := b.neighbors(p).filter(libertyCandidateCriteria)
	liberties := make([]liberty, 0)
	for _, lC := range libertyCandidates {
		if b.governance(lC) == GovNeutral {
			liberties = append(liberties, newLiberty(lC))
		}
		if b.governance(lC) == g && IsAlive(b, lC) {
			liberties = append(liberties, newLiberty(lC))
		}
	}
	return liberties
}

// IsAlive computes the life status for a position
func IsAlive(b Board, p position) bool {
    g := b.governance(p)
    if g == GovNeutral {
        return false
    }

	neutralCriteria := governanceCriteria(b, GovNeutral)
	allyCriteria := governanceCriteria(b, g)

    visited := make(map[position]bool)
    queue := list.New()
    queue.PushBack(p)
    visited[p] = true
    for queue.Len() > 0 {
        current := queue.Remove(queue.Front()).(position)
        visited[current] = true

        currentNeighbors := b.neighbors(current)
        if len(currentNeighbors.filter(neutralCriteria)) > 0 {
        	return true
        }

	    for _, currentAlly := range currentNeighbors.filter(allyCriteria) {
            if !visited[currentAlly] {
                visited[currentAlly] = true
                queue.PushBack(currentAlly)
            }
        }
    }

    // No adjacent neutral positions or alive allies
    return false
}

// Positions retrieves the positions on a board matching a certain governance.
func Positions(b Board, g governance) positions {
	size := len(b)
	positions := make(positions, 0)
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

// governanceCriteria creates a function to check if a position governance matches with another one on a board
func governanceCriteria(b Board, g governance) func(position) bool {
	return func(p0 position) bool {
		return b.governance(p0) == g
	}
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

// neighbors computes the nearest 4 adjacent positions for the position
func (b Board) neighbors(p position) positions {
	size := len(b)
	nC := make(positions, 0)
	if p.x+1 < size {
		nC = append(nC, position{x: p.x + 1, y: p.y})
	}
	if p.x > 0 {
		nC = append(nC, position{x: p.x - 1, y: p.y})
	}
	if p.y+1 < size {
		nC = append(nC, position{x: p.x, y: p.y + 1})
	}
	if p.y > 0 {
		nC = append(nC, position{x: p.x, y: p.y - 1})
	}
	return nC
}

// checkIndexSafe checks that a position is inside the bounds of the board.
func (b Board) checkIndexSafe(p position) error {
	if p.x < 0 || p.y < 0 {
		return ErrPositionNeg
	}
	if p.x >= len(b) || p.y >= len(b) {
		return ErrPositionOOB
	}
	return nil
}

// filter filters positions by a matching criteria
func (pS positions) filter(criteria func(position) bool) positions {
	filtered := make(positions, 0)
	for _, p := range pS {
		if criteria(p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
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

	for row := 0; row < l; row++ {
		var line strings.Builder
		for col := 0; col < l; col++ {
			if col == 0 {
				line.WriteString(fmt.Sprintf(positionFmt, row))
			}
			line.WriteString(fmt.Sprintf(positionFmt, b[col][row].string()))
		}
		fmt.Println(line.String())
	}
}
