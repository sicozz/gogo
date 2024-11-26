package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardBuilder(t *testing.T) {
	g := GovBlack
	poss := positions{
		NewPosition(0, 0),
		NewPosition(0, 1),
		NewPosition(1, 0),
		NewPosition(1, 1),
	}
	b := NewBoardBuilder(BoardSize19).
		SetPosition(poss[0], g).
		SetPosition(poss[1], g).
		SetPosition(poss[2], g).
		SetPosition(poss[3], g).
		Build()

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
	p := NewPosition(0, 1)
	b := NewBoardBuilder(BoardSize9).SetPosition(p, GovBlack).Build()
	assert.Equal(t, GovBlack, b.governance(p), "Position should have had black-governance")

	_, err := ClaimPosition(b, p, GovWhite)
	assert.EqualError(t, err, ErrPositionTaken.Error(), "Failed prevention of claimed position claim")
}

// Test the position claiming for the neutral governance.
func TestClaimPosition_NeutralGovernance(t *testing.T) {
	b := NewBoard(BoardSize9)
	p := NewPosition(0, 1)
	_, err := ClaimPosition(b, p, GovNeutral)
	assert.EqualError(t, err, ErrNeutralClaim.Error(), "Failed prevention of Neutral claiming position")
}

// Test the computation of the life state of a group of positions in Atari.
func TestIsAlive_Atari(t *testing.T) {
	pS := []struct {
		pos position
		gov governance
	}{
		{pos: position{x: 0, y: 0}, gov: GovBlack},
		{pos: position{x: 0, y: 1}, gov: GovBlack},
		{pos: position{x: 0, y: 2}, gov: GovBlack},
		{pos: position{x: 0, y: 3}, gov: GovBlack},

		{pos: position{x: 1, y: 0}, gov: GovWhite},
		{pos: position{x: 1, y: 1}, gov: GovWhite},
		{pos: position{x: 1, y: 2}, gov: GovWhite},
		{pos: position{x: 1, y: 3}, gov: GovWhite},
	}
	builder := NewBoardBuilder(BoardSize11)
	for _, p := range pS {
		builder.SetPosition(p.pos, p.gov)
	}
	b := builder.Build()

	for _, p := range pS {
		assert.True(t, IsAlive(b, p.pos), "Failed to tell position life status")
	}
}

// Test the computation of the life state of a group of positions at capture moment.
func TestIsAlive_Capture(t *testing.T) {
	pSBlack := positions{
		{x: 0, y: 0},
		{x: 0, y: 1},
		{x: 0, y: 2},
		{x: 0, y: 3},
	}
	pSWhite := positions{
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 0, y: 4},
	}
	builder := NewBoardBuilder(BoardSize11)
	for _, p := range pSBlack {
		builder.SetPosition(p, GovBlack)
	}
	for _, p := range pSWhite {
		builder.SetPosition(p, GovWhite)
	}
	b := builder.Build()

	for _, p := range pSBlack {
		assert.False(t, IsAlive(b, p), "Failed to tell position life status")
	}
	for _, p := range pSWhite {
		assert.True(t, IsAlive(b, p), "Failed to tell position life status")
	}
}

func TestIsAlive_Case00(t *testing.T) {
	pSBlack := positions{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 1, y: 2},
		{x: 2, y: 1},
	}
	pWhite := position{x: 1, y: 1}
	builder := NewBoardBuilder(BoardSize11)
	for _, p := range pSBlack {
		builder.SetPosition(p, GovBlack)
	}
	builder.SetPosition(pWhite, GovWhite)
	b := builder.Build()

	assert.False(t, IsAlive(b, pWhite), "Failed to tell position life status")
	for _, p := range pSBlack {
		assert.True(t, IsAlive(b, p), "Failed to tell position life status")
	}
}
