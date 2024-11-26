package game

// liberty represents a liberty in the Go game.
type liberty struct {
	position
}

func newLiberty(p position) liberty {
	return liberty{p}
}
