// game provides a simulation engine for a Go match.
package game

const (
	GovNeutral governance = iota
	GovBlack
	GovWhite
)

const (
	ClaimPos action = iota
	Pass
	Resign
)

// GoMatch manages the lifecycle of a Go game.
type GoMatch struct{}

type GameState struct{}

// action represents the options a player can take on his turn.
type action int

// governance represents the ownership state for a position.
type governance int

func NewGoMatch() *GoMatch {
	return &GoMatch{}
}

// UNIMPLEMENTED
func Players(g *GoMatch) *GoMatch {
	return nil
}

// UNIMPLEMENTED
func InitHandicap(g *GoMatch) int {
	return 0
}

// UNIMPLEMENTED
func ContinuousPasses(g *GoMatch) int {
	return 0
}

// UNIMPLEMENTED
func History(g *GoMatch) []GameState {
	return make([]GameState, 0)
}

// UNIMPLEMENTED
func NewGameState() *GameState {
	return &GameState{}
}

// UNIMPLEMENTED
func NextState(gst *GameState) *GameState {
	return nil
}

// UNIMPLEMENTED
func TurnEvents(gst *GameState) map[string]string {
	return make(map[string]string)
}

// UNIMPLEMENTED
func PlayerScore(gst *GameState, pID governance) int {
	return 0
}

// UNIMPLEMENTED
func ActivePlayer(gst *GameState) governance {
	return governance(0)
}

// UNIMPLEMENTED
func Handicap(gst *GameState) int {
	return 0
}

func (g governance) string() string {
	switch g {
	case GovNeutral:
		return "+"
	case GovBlack:
		return "●"
	case GovWhite:
		return "○"
	default:
		return "Unknown"
	}
}
