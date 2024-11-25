package main

import (
	"fmt"

	"github.com/sicozz/gogo/game"
)

func main() {
	b := game.
		NewBoardBuilder(game.BoardSize11).
		SetPosition(game.NewPosition(8, 2), game.GovBlack).
		SetPosition(game.NewPosition(8, 3), game.GovBlack).
		SetPosition(game.NewPosition(9, 3), game.GovBlack).
		SetPosition(game.NewPosition(9, 1), game.GovBlack).
		SetPosition(game.NewPosition(10, 3), game.GovBlack).
		SetPosition(game.NewPosition(10, 0), game.GovBlack).
		SetPosition(game.NewPosition(9, 2), game.GovWhite).
		SetPosition(game.NewPosition(10, 2), game.GovWhite).
		SetPosition(game.NewPosition(10, 1), game.GovWhite).
		Build()
	b.Display()
	liberties := game.PositionLiberties(b, game.NewPosition(9, 2))
	fmt.Println("----- Liberties -----")
	for _, l := range liberties {
		fmt.Printf("[ liberty: %v ]\n", l)
	}
	fmt.Println("----- ========= -----")
}
