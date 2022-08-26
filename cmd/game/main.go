package main

import (
	"game/pkg/sgame"
	tl "github.com/JoelOtter/termloop"
)

// func main() {
// 	game := tl.NewGame()
// 	// game.DebugOn()
// 	game.SetDebugOn(true)
// 	sgame.Glog = game.Log
//
// 	shpL := sgame.NewShip(60, 15, sgame.OriL, sgame.FourMast)
//
// 	oceMe := sgame.NewOcean(8, 14)
// 	oceOp := sgame.NewOcean(50, 14)
//
// 	level := tl.NewBaseLevel(tl.Cell{
// 		Bg: tl.ColorBlack,
// 		Fg: tl.ColorWhite,
// 		Ch: ' ',
// 	})
//
// 	level.AddEntity(oceMe)
// 	level.AddEntity(oceOp)
// 	level.AddEntity(shpL)
//
// 	game.Screen().SetLevel(level)
// 	game.Start()
// }

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	sgame.Glog = game.Log

	oceMe := sgame.NewOcean(10, 14)
	oceOp := sgame.NewOcean(50, 14)
	shpInf := sgame.NewTable(70, 15, oceMe)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	level.AddEntity(oceMe)
	level.AddEntity(oceOp)
	level.AddEntity(shpInf)

	// shpL := sgame.NewShip(11, 16, sgame.OriR, sgame.FourMast)
	// shpL.SetFocus(true)
	// oceMe.AddShip(sgame.FourMast)
	// level.AddEntity(shpL)

	game.Screen().SetLevel(level)
	game.Start()
}
