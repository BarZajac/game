package sgame

import (
	"bytes"

	tl "github.com/JoelOtter/termloop"
)

// Orientation represents Ship orientation on the board.
type Orientation byte

// orientations
const (
	OriU Orientation = 'U' // Up from anchor point.
	OriR Orientation = 'R' // Right from anchor point.
	OriD Orientation = 'D' // Down from anchor point.
	OriL Orientation = 'L' // Left from anchor point.
)

// orientations represents all valid orientations of a ship.
var orientations = [4]byte{
	byte(OriU),
	byte(OriR),
	byte(OriD),
	byte(OriL),
}

// RotDir represents rotation direction.
type RotDir int

// Rotation directions.
const (
	RotL RotDir = -1
	RotR RotDir = 1
)

// ShipSize represents the length of a ship.
type ShipSize int

// Ship sizes.
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
func NewShip(x, y int, ori Orientation, size ShipSize) *Ship {
	shp := &Ship{
		size:    size,
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

	newX, newY := shp.currX, shp.currY
	for i := 0; i < int(shp.size); i++ {
		switch shp.currOri {
		case OriR:
			newX, newY = shp.currX+i, shp.currY
		case OriL:
			newX, newY = shp.currX-i, shp.currY
		case OriU:
			newX, newY = shp.currX, shp.currY-i
		case OriD:
			newX, newY = shp.currX, shp.currY+i
		}

		s.RenderCell(newX, newY, c)
	}
}

func (shp *Ship) Tick(event tl.Event) {
	// Is it a keyboard event?
	if event.Type == tl.EventKey {
		shp.prevX = shp.currX
		shp.prevY = shp.currY
		shp.prevOri = shp.currOri

		switch event.Key {
		case tl.KeyF2:

		case tl.KeyPgdn:
			shp.currOri = nextOri(shp.currOri, RotR)

		case tl.KeyPgup:
			shp.currOri = nextOri(shp.currOri, RotL)

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
	if _, ok := collision.(*tl.Rectangle); ok {
		shp.currX = shp.prevX
		shp.currY = shp.prevY
	}
}

// nextOri returns orientation after rotating left or right from current one.
func nextOri(cur Orientation, dir RotDir) Orientation {
	i := bytes.IndexByte(orientations[:], byte(cur))
	if i == -1 {
		return cur
	}

	next := i + int(dir)
	if next == len(orientations) {
		next = 0
	}

	if next == -1 {
		next = len(orientations) - 1
	}

	return Orientation(orientations[next])
}
