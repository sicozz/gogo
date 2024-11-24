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

// Test the position claiming for a not-claimed position.
func TestClaimPosition_NotClaimedPosition(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
    assert.Equal(t, GovNeutral, b.governance(p), "Position should have been neutral")

	rB, err := ClaimPosition(b, p, GovBlack)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, GovBlack, rB.governance(p), "Failed to claim position")
}

// Test the position claiming for a claimed position.
func TestClaimPosition_ClaimedPosition(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	bB, err := ClaimPosition(b, p, GovBlack)
	if err != nil {
		t.Fatal(err)
	}
    assert.Equal(t, GovBlack, bB.governance(p), "Position should have had black-governance")

	_, err = ClaimPosition(bB, p, GovWhite)
	assert.EqualError(t, err, ErrPositionTaken.Error(), "Failed prevention of claimed position claim")
}

// Test the position claiming for the neutral governance.
func TestClaimPosition_NeutralGovernance(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	_, err := ClaimPosition(b, p, GovNeutral)
	assert.EqualError(t, err, ErrNeutralClaim.Error(), "Failed prevention of Neutral claiming position")
}
