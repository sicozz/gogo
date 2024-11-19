package main

import (
	"log"

	"github.com/sicozz/gogo/game"
)

func main() {
	b := game.NewBoard(4)
	rB, err := game.ClaimPosition(b, game.NewPosition(0, 1), game.GovBlack)
	if err != nil {
		log.Fatal(err)
	}
	rB.Display()
}
