package game

type Game struct{}

type GameState struct{}

type PID uint

type Action uint

const (
	ClaimPosition Action = iota
	Pass
	Resign
)

func NewGame() *Game {
	return &Game{}
}

func Players(g *Game) []PID {
	return make([]PID, 0)
}

func InitHandicap(g *Game) uint {
	return 0
}

func ContinuousPasses(g *Game) uint {
	return 0
}

func History(g *Game) []GameState {
	return make([]GameState, 0)
}

func NewGameState() *GameState {
	return &GameState{}
}

func GetBoard(gst *GameState) *Board {
	return &Board{}
}

func Handicap(gst *GameState) uint {
	return 0
}

func PlayerScore(gst *GameState, pID PID) uint {
	return 0
}

func TurnPlayer(gst *GameState) PID {
	return PID(0)
}

func TurnEvents(gst *GameState) map[string]string {
	return make(map[string]string)
}

func GetAction(gst *GameState) *Action {
	return nil
}

func NextState(gst *GameState) *GameState {
	return nil
}
