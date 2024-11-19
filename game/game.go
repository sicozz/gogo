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

type GameEngine struct{}

type GameState struct{}

type action int

type governance int

func NewGameEngine() *GameEngine {
	return &GameEngine{}
}

// UNIMPLEMENTED
func Players(g *GameEngine) *GameEngine {
	return nil
}

// UNIMPLEMENTED
func InitHandicap(g *GameEngine) int {
	return 0
}

// UNIMPLEMENTED
func ContinuousPasses(g *GameEngine) int {
	return 0
}

// UNIMPLEMENTED
func History(g *GameEngine) []GameState {
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
