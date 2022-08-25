package sgame

import (
	"bytes"

	tl "github.com/JoelOtter/termloop"
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

type Block struct {
	x       int
	y       int
	prevX   int
	prevY   int
	fg      tl.Attr
	bg      tl.Attr
	ori     Orientation
	prevOri Orientation
	text    []rune
	canvas  []tl.Cell
	locked  bool
}

// NewShip  returns new instance of a Ship.
func NewShip(x, y int, ori Orientation, size ShipSize) *Block {
	bg := tl.ColorBlue
	fg := tl.ColorWhite

	str := make([]rune, int(size))
	c := make([]tl.Cell, int(size))
	for i := range c {
		str[i] = ' '
		c[i] = tl.Cell{Ch: ' ', Fg: fg, Bg: bg}
	}
	return &Block{
		x:       x,
		prevX:   x,
		y:       y,
		prevY:   y,
		fg:      fg,
		bg:      bg,
		ori:     ori,
		prevOri: ori,
		text:    str,
		canvas:  c,
	}
}

func NewText(x, y int, text string, ori Orientation, fg, bg tl.Attr) *Block {
	str := []rune(text)
	c := make([]tl.Cell, len(str))
	for i := range c {
		c[i] = tl.Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Block{
		x:       x,
		prevX:   x,
		y:       y,
		prevY:   y,
		fg:      fg,
		bg:      bg,
		ori:     ori,
		prevOri: ori,
		text:    str,
		canvas:  c,
		locked:  true,
	}
}

func NewBar(x, y, width int, ori Orientation, fg, bg tl.Attr) *Block {
	str := make([]rune, width)
	c := make([]tl.Cell, width)
	for i := range c {
		str[i] = ' '
		c[i] = tl.Cell{Ch: ' ', Fg: fg, Bg: bg}
	}
	return &Block{
		x:       x,
		prevX:   x,
		y:       y,
		prevY:   y,
		fg:      fg,
		bg:      bg,
		ori:     ori,
		prevOri: ori,
		text:    str,
		canvas:  c,
		locked:  true,
	}
}

func (blk *Block) Tick(ev tl.Event) {
	if blk.locked {
		return
	}

	if ev.Type == tl.EventKey {
		blk.prevX = blk.x
		blk.prevY = blk.y
		blk.prevOri = blk.ori

		switch ev.Key {
		case tl.KeyF2:

		case tl.KeyPgdn:
			blk.ori = nextOri(blk.ori, RotR)

		case tl.KeyPgup:
			blk.ori = nextOri(blk.ori, RotL)

		case tl.KeySpace:

		case tl.KeyArrowRight:
			blk.x += 1

		case tl.KeyArrowLeft:
			blk.x -= 1

		case tl.KeyArrowUp:
			blk.y -= 1

		case tl.KeyArrowDown:
			blk.y += 1
		}
	}
}

func (blk *Block) Draw(s *tl.Screen) {
	w, _ := blk.Size()
	newX, newY := blk.x, blk.y

	for i := 0; i < w; i++ {
		switch blk.ori {
		case OriU:
			newX, newY = blk.x, blk.y-i

		case OriR:
			newX, newY = blk.x+i, blk.y

		case OriD:
			newX, newY = blk.x, blk.y+i

		case OriL:
			newX, newY = blk.x-i, blk.y
		}

		s.RenderCell(newX, newY, &blk.canvas[i])
	}
}

func (blk *Block) Position() (int, int) {
	return blk.x, blk.y
}

func (blk *Block) Size() (int, int) {
	return len(blk.text), 1
}

func (blk *Block) SetPosition(x, y int) {
	blk.x = x
	blk.y = y
}

func (blk *Block) Text() string {
	return string(blk.text)
}

func (blk *Block) SetText(text string) {
	blk.text = []rune(text)
	c := make([]tl.Cell, len(blk.text))
	for i := range c {
		c[i] = tl.Cell{Ch: blk.text[i], Fg: blk.fg, Bg: blk.bg}
	}
	blk.canvas = c
}

func (blk *Block) Color() (tl.Attr, tl.Attr) {
	return blk.fg, blk.bg
}

func (blk *Block) SetColor(fg, bg tl.Attr) {
	blk.fg = fg
	blk.bg = bg
	for i := range blk.canvas {
		blk.canvas[i].Fg = fg
		blk.canvas[i].Bg = bg
	}
}

// func (shp *Ship) Collide(collision tl.Physical) {
// 	// Check if it's a Rectangle we're colliding with
// 	if _, ok := collision.(*tl.Rectangle); ok {
// 		shp.currX = shp.prevX
// 		shp.currY = shp.prevY
// 	}
// }

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
