package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
