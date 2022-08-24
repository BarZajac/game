package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

// Orientation represents Ship orientation on the board.
type Orientation byte

const (
	OriL Orientation = 'L' // Left from anchor point.
	OriR Orientation = 'R' // Right from anchor point.
	OriU Orientation = 'U' // Up from anchor point.
	OriD Orientation = 'D' // Down from anchor point.
)

// ShipSize represents the length of a ship.
type ShipSize int

// Ship sizes
const (
	OneMast   = 1
	TwoMast   = 2
	ThreeMast = 3
	FourMast  = 4
)

// Ship represents single ship on the board.
type Ship struct {
	*tl.Entity
	size    ShipSize
	prevX   int
	prevY   int
	currOri Orientation
	prevOri Orientation
}

// NewShip  returns new instance of a Ship.
func NewShip(x, y int, ori Orientation, s ShipSize) *Ship {
	shp := &Ship{
		Entity:  tl.NewEntity(x, y, int(s), 1),
		size:    s,
		prevX:   x,
		prevY:   y,
		prevOri: ori,
		currOri: ori,
	}
	return shp
}

func (shp *Ship) Draw(s *tl.Screen) {
	x, y := shp.Position()

	c := &tl.Cell{
		Fg: tl.ColorWhite,
		Bg: tl.ColorBlue,
		Ch: ' ',
	}

	for i := x; i < x+int(shp.size); i++ {
		s.RenderCell(x+i, y, c)
	}
}

func (shp *Ship) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		shp.prevX, shp.prevY = shp.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyF2:
			shp.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: '*'})
		case tl.KeyPgdn:
			// tl.NewEntity(shp.prevX, shp.prevY, 1, 4)
			// canvas := tl.NewCanvas(2, 5)
			// shp.SetCanvas(&canvas)

			shp.Entity = tl.NewEntity(shp.prevX, shp.prevY, 1, 4)
			shp.Fill(&tl.Cell{Fg: tl.ColorRed, Bg: tl.RgbTo256Color(130, 130, 130), Ch: ' '})
			//shp.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: 'x'})
			shp.prevOri = 'R'
			shp.currOri = 'D'

		case tl.KeyPgup:
			shp.Entity = tl.NewEntity(shp.prevX, shp.prevY, 4, 1)
			shp.Fill(&tl.Cell{Fg: tl.ColorRed, Bg: tl.RgbTo256Color(130, 130, 130), Ch: ' '})
			shp.SetCell(0, 0, &tl.Cell{Fg: tl.RgbTo256Color(255, 255, 255), Ch: 'x'})
			shp.prevOri = 'D'
			shp.currOri = 'R'
		case tl.KeySpace:
			shp.SetPosition(shp.prevX, shp.prevY-5)
		case tl.KeyArrowRight:
			shp.SetPosition(shp.prevX+1, shp.prevY)
		case tl.KeyArrowLeft:
			shp.SetPosition(shp.prevX-1, shp.prevY)
		case tl.KeyArrowUp:
			shp.SetPosition(shp.prevX, shp.prevY-1)
		case tl.KeyArrowDown:
			shp.SetPosition(shp.prevX, shp.prevY+1)
		}
	}
}

func (shp *Ship) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	// glog("%+v %T %+v", collision, collision, shp)
	if _, ok := collision.(*tl.Rectangle); ok {
		shp.SetPosition(shp.prevX, shp.prevY)
	}

}
