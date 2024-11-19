package game

type liberty struct{}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func newLiberty() liberty {
	return liberty{}
}

// UNIMPLEMENTED
func Origin(l liberty) position {
	return position{}
}

// UNIMPLEMENTED
func End(l liberty) position {
	return position{}
}

// UNIMPLEMENTED
func GetDirection(l liberty) direction {
	return up
}
