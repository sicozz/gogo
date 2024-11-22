package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositions(t *testing.T) {
	b := NewBoard(BoardSize19)
	g := GovBlack
	positions := []position{
		NewPosition(0, 0),
		NewPosition(0, 1),
		NewPosition(1, 0),
		NewPosition(1, 1),
	}
	b = b.seize(positions[0], g)
	b = b.seize(positions[1], g)
	b = b.seize(positions[2], g)
	b = b.seize(positions[3], g)

	resPositions := Positions(b, GovBlack)

	assert.ElementsMatch(t, positions, resPositions, "Failed to query positions")
}

func TestClaimPosition(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	rB, err := ClaimPosition(b, p, GovBlack)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, GovBlack, rB.governance(p), "Failed to claim position")
}

func TestClaimPositionClaimed(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	bB, err := ClaimPosition(b, p, GovBlack)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ClaimPosition(bB, p, GovWhite)
	assert.EqualError(t, err, ErrPositionTaken.Error(), "Failed prevention of claimed position claim")
}

func TestNeutralClaimsPosition(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	_, err := ClaimPosition(b, p, GovNeutral)
	assert.EqualError(t, err, ErrNeutralClaim.Error(), "Failed prevention of Neutral claiming position")
}
