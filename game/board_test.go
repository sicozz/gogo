package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositions(t *testing.T) {
	b := NewBoard(BoardSize19)
	g := GovBlack
	poss := positions{
		NewPosition(0, 0),
		NewPosition(0, 1),
		NewPosition(1, 0),
		NewPosition(1, 1),
	}
	b = b.seize(poss[0], g)
	b = b.seize(poss[1], g)
	b = b.seize(poss[2], g)
	b = b.seize(poss[3], g)

	resPositions := Positions(b, GovBlack)

	assert.ElementsMatch(t, poss, resPositions, "Failed to query positions")
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

func TestIsAlive_Atari(t *testing.T) {
    b := NewBoard(BoardSize11)
    pSBlack := positions{
        { x: 0, y: 0 },
        { x: 0, y: 1 },
        { x: 0, y: 2 },
        { x: 0, y: 3 },
    }
    pSWhite := positions{
        { x: 1, y: 0 },
        { x: 1, y: 1 },
        { x: 1, y: 2 },
        { x: 1, y: 3 },
    }
    for _, p := range pSBlack {
        b = b.seize(p, GovBlack)
    }
    for _, p := range pSWhite {
        b = b.seize(p, GovWhite)
    }

    for _, p := range append(pSBlack, pSWhite...) {
        assert.True(t, IsAlive(b, p), "Failed to tell position life status")
    }
}

func TestIsAlive_Capture(t *testing.T) {
    b := NewBoard(BoardSize11)
    pSBlack := positions{
        { x: 0, y: 0 },
        { x: 0, y: 1 },
        { x: 0, y: 2 },
        { x: 0, y: 3 },
    }
    pSWhite := positions{
        { x: 1, y: 0 },
        { x: 1, y: 1 },
        { x: 1, y: 2 },
        { x: 1, y: 3 },
        { x: 0, y: 4 },
    }
    for _, p := range pSBlack {
        b = b.seize(p, GovBlack)
    }
    for _, p := range pSWhite {
        b = b.seize(p, GovWhite)
    }

    for _, p := range pSBlack {
        assert.False(t, IsAlive(b, p), "Failed to tell position life status")
    }
    for _, p := range pSWhite {
        assert.True(t, IsAlive(b, p), "Failed to tell position life status")
    }
}

func TestIsAlive_Case00(t *testing.T) {
    b := NewBoard(BoardSize11)
    pSBlack := positions{
        { x: 0, y: 1 },
        { x: 1, y: 0 },
        { x: 1, y: 2 },
        { x: 2, y: 1 },
    }
    pWhite := position{ x: 1, y: 1 }
    for _, p := range pSBlack {
        b = b.seize(p, GovBlack)
    }
    b = b.seize(pWhite, GovWhite)

    assert.False(t, IsAlive(b, pWhite), "Failed to tell position life status")
    for _, p := range pSBlack {
        assert.True(t, IsAlive(b, p), "Failed to tell position life status")
    }
}
