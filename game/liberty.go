package game

import (
	"errors"
	"log"
)

type liberty struct {
	origin position
	dir    direction
}

type direction int

var ErrDirectionNotFound = errors.New("liberty: liberty with unknown direction: %v")

const (
	up direction = iota
	down
	left
	right
)

func newLiberty(p position, d direction) liberty {
	return liberty{origin: p, dir: d}
}

// UNIMPLEMENTED
func Origin(l liberty) position {
	return l.origin
}

// UNIMPLEMENTED
func End(l liberty) position {
	switch l.dir {
	case up:
		return NewPosition(l.origin.x, l.origin.y+1)
	case down:
		return NewPosition(l.origin.x, l.origin.y-1)
	case left:
		return NewPosition(l.origin.x-1, l.origin.y)
	case right:
		return NewPosition(l.origin.x+1, l.origin.y)
	default:
		log.Fatalf(ErrDirectionNotFound.Error(), l)
		return NewPosition(-1, -1)
	}
}

// UNIMPLEMENTED
func GetDirection(l liberty) direction {
	return up
}

func (d direction) string() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	default:
		log.Fatalf(ErrDirectionNotFound.Error(), d)
		return "unknown direction"
	}
}
