package main

import (
	"game/pkg/sgame"
	tl "github.com/JoelOtter/termloop"
)

func main() {
	game := tl.NewGame()

	shp := sgame.NewShip(30, 15, sgame.OriR, sgame.FourMast)
	// oceMe := sgame.NewOcean(8, 14)
	// oceOp := sgame.NewOcean(50, 14)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	// level.AddEntity(oceMe)
	// level.AddEntity(oceOp)
	level.AddEntity(shp)

	game.Screen().SetLevel(level)
	game.Start()
}
