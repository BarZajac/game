package sgame

import (
	"bytes"

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
	size    ShipSize
	currX   int
	currY   int
	prevX   int
	prevY   int
	currOri Orientation
	prevOri Orientation
}

// NewShip  returns new instance of a Ship.
func NewShip(x, y int, ori Orientation, s ShipSize) *Ship {
	shp := &Ship{
		size:    s,
		currX:   x,
		currY:   y,
		prevX:   x,
		prevY:   y,
		prevOri: ori,
		currOri: ori,
	}
	return shp
}

func (shp *Ship) Draw(s *tl.Screen) {
	c := &tl.Cell{
		Fg: tl.ColorWhite,
		Bg: tl.ColorBlue,
		Ch: ' ',
	}

	for i := 0; i < int(shp.size); i++ {
		switch shp.currOri {
		case OriR:
			s.RenderCell(shp.currX+i, shp.currY, c)

		case OriL:
			s.RenderCell(shp.currX-i, shp.currY, c)

		case OriU:
			s.RenderCell(shp.currX, shp.currY-i, c)

		case OriD:
			s.RenderCell(shp.currX, shp.currY+i, c)
		}
	}
}

func (shp *Ship) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		shp.prevX = shp.currX
		shp.prevY = shp.currY
		shp.prevOri = shp.currOri

		switch event.Key {
		case tl.KeyF2:

		case tl.KeyPgdn:

			shp.currOri = nextOri(shp.currOri, true)

		case tl.KeyPgup:
			shp.currOri = nextOri(shp.currOri, false)

		case tl.KeySpace:

		case tl.KeyArrowRight:
			shp.currX += 1

		case tl.KeyArrowLeft:
			shp.currX -= 1

		case tl.KeyArrowUp:
			shp.currY -= 1

		case tl.KeyArrowDown:
			shp.currY += 1
		}
	}
}

func (shp *Ship) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	// glog("%+v %T %+v", collision, collision, shp)
	if _, ok := collision.(*tl.Rectangle); ok {
		shp.currX = shp.prevX
		shp.currY = shp.prevY
	}

}

var s = []byte{'U', 'R', 'D', 'L'}

func nextOri(cur Orientation, dir bool) Orientation {

	i := bytes.IndexByte(s, byte(cur))
	if i == -1 {
		return cur
	}

	if dir {
		// Clockwise.
		switch {
		case i == 0:
			return OriR

		case i == 3:
			return OriU

		default:
			return Orientation(s[i+1])
		}
	} else {
		// Counterclockwise.
		switch {
		case i == 0:
			return OriL

		case i == 3:
			return OriD

		default:
			return Orientation(s[i-1])
		}
	}
}
