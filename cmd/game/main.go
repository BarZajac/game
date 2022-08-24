package main

import (
	"game/pkg/sgame"
	tl "github.com/JoelOtter/termloop"
)

type Physical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	DynamicPhysical
}

// DynamicPhysical represents something that can process its own collisions.
// Implementing this is an optional addition to Drawable.
type DynamicPhysical interface {
	Position() (int, int) // Return position, x and y
	Size() (int, int)     // Return width and height
	Collide(Physical)     // Handle collisions with another Physical
}

type logFn func(log string, items ...interface{})

var glog logFn

func main() {
	game := tl.NewGame()
	game.SetDebugOn(true)
	glog = game.Log

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Fg: tl.ColorBlack,
		Ch: '~',
	})

	shp := sgame.NewShip(30, 15, sgame.OriR, sgame.FourMast)

	// Set the character at position (0, 0) on the entity.
	// shp.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Bg: tl.RgbTo256Color(100, 150, 200), Ch: '*'}) //옷
	shp.Fill(&tl.Cell{Fg: tl.ColorRed, Bg: tl.RgbTo256Color(130, 130, 130), Ch: ' '})
	level.AddEntity(shp)
	// Your sea
	level.AddEntity(tl.NewText(30, 10, "Your Fleet", tl.ColorGreen, tl.ColorBlue))
	level.AddEntity(tl.NewText(10, 18, "Ships Left: 10", tl.ColorGreen, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(30, 14, 10, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(29, 14, 1, 11, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(29, 25, 12, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(40, 14, 1, 11, tl.ColorBlack))
	// Coordinates
	level.AddEntity(tl.NewText(30, 13, "ABCDEFGHIJ", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 15, "1", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 16, "2", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 17, "3", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 18, "4", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 19, "5", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 20, "6", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 21, "7", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 22, "8", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 23, "9", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(28, 24, "0", tl.ColorWhite, tl.ColorBlue))
	// Task design
	level.AddEntity(tl.NewRectangle(44, 18, 18, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(44, 19, 1, 2, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(44, 21, 18, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(61, 19, 1, 2, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(45, 19, 3, 1, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(58, 19, 3, 1, tl.ColorBlue))
	// Task notifier
	level.AddEntity(tl.NewText(47, 13, "Choose a cell", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(45, 14, "you want to shoot", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(48, 19, "Objective:", tl.RgbTo256Color(17, 216, 220), tl.ColorBlue))
	level.AddEntity(tl.NewText(45, 20, "Sink enemy ships", tl.RgbTo256Color(17, 216, 220), tl.ColorBlue))
	// Enemy sea
	level.AddEntity(tl.NewText(66, 10, "Enemy Fleet", tl.ColorRed, tl.ColorBlue))
	level.AddEntity(tl.NewText(83, 18, "Ships Left: 10", tl.ColorRed, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(66, 14, 10, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(65, 14, 1, 11, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(65, 25, 12, 1, tl.ColorBlack))
	level.AddEntity(tl.NewRectangle(76, 14, 1, 11, tl.ColorBlack))
	// Coordinates
	level.AddEntity(tl.NewText(66, 13, "ABCDEFGHIJ", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 15, "1", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 16, "2", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 17, "3", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 18, "4", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 19, "5", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 20, "6", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 21, "7", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 22, "8", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 23, "9", tl.ColorWhite, tl.ColorBlue))
	level.AddEntity(tl.NewText(64, 24, "0", tl.ColorWhite, tl.ColorBlue))
	// Game start
	// fmt.Scan("Type X", &a)
	// fmt.Scan("Type Y", &b)
	level.Offset()
	game.Screen().SetLevel(level)
	game.Start()
}

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	pervO byte
	currO byte
	level *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	// screenWidth, screenHeight := screen.ShipSize()
	// x, y := player.Position()
	// player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	// We need to make sure and call Draw on the underlying Entity.
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyF2:
			player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: 'x'})
		case tl.KeyPgdn:
			// tl.NewEntity(player.prevX, player.prevY, 1, 4)
			// canvas := tl.NewCanvas(2, 5)
			// player.SetCanvas(&canvas)

			player.Entity = tl.NewEntity(player.prevX, player.prevY, 1, 4)
			player.Fill(&tl.Cell{Fg: tl.ColorRed, Bg: tl.RgbTo256Color(130, 130, 130), Ch: ' '})
			player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: 'x'})
			player.currO = 'D'

		case tl.KeyPgup:
			player.Entity = tl.NewEntity(player.prevX, player.prevY, 4, 1)
			player.Fill(&tl.Cell{Fg: tl.ColorRed, Bg: tl.RgbTo256Color(130, 130, 130), Ch: ' '})
			player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: 'x'})
			player.currO = 'R'
		case tl.KeyEnter:

		// case tl.KeyF2:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: '옷'})
		// case tl.KeyF3:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 0, 0), Ch: '옷'})
		// case tl.KeyF4:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(0, 255, 0), Ch: '옷'})
		// case tl.KeyF5:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(0, 0, 255), Ch: '옷'})
		// case tl.KeyF6:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 0, 255), Ch: '옷'})
		// case tl.KeyF7:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(0, 255, 255), Ch: '옷'})
		// case tl.KeyF8:
		//	player.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 0), Ch: '옷'})
		case tl.KeySpace:
			player.SetPosition(player.prevX, player.prevY-5)
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	glog("%+v %T %+v", collision, collision, player)
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}

}

// func (player *Player) Orient()  {
//	if
// }
