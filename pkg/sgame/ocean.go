package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

const OceanSize int = 10

var frameCol = tl.RgbTo256Color(90, 90, 90)

type Ocean struct {
	x int
	y int

	topF    *Block
	rightF  *Block
	leftF   *Block
	bottomF *Block

	ships       []*Block
	shipsToAdd  []*Block
	initialized bool
}

func (o *Ocean) Position() (int, int) {
	return o.x, o.y
}

func (o *Ocean) Size() (int, int) {
	return OceanSize + 2, OceanSize + 2
}

func NewOcean(x, y int) *Ocean {
	o := &Ocean{
		x: x,
		y: y,
	}
	return o
}

func (o *Ocean) Draw(s *tl.Screen) {
	if !o.initialized {
		o.initialize()
		s.AddEntity(o.topF)
		s.AddEntity(o.rightF)
		s.AddEntity(o.bottomF)
		s.AddEntity(o.leftF)
	}

	for _, shp := range o.shipsToAdd {
		s.Level().AddEntity(shp)
		o.ships = append(o.ships, shp)
	}
	o.shipsToAdd = o.shipsToAdd[:0]
}

func (o *Ocean) initialize() {
	o.topF = NewText(o.x, o.y, " 0123456789", OriR, tl.ColorWhite, frameCol)
	o.rightF = NewBar(o.x+OceanSize+1, o.y, OceanSize+2, OriD, tl.ColorWhite, frameCol)
	o.bottomF = NewBar(o.x, o.y+OceanSize+1, OceanSize+1, OriR, tl.ColorWhite, frameCol)
	o.leftF = NewText(o.x, o.y+1, "ABCDEFGHIJ", OriD, tl.ColorWhite, frameCol)
	o.initialized = true
}

func (o *Ocean) Tick(ev tl.Event) {
}

func (o *Ocean) Definition() (int, int, int, int) {
	return o.x, o.y, o.x + OceanSize + 1, o.y + OceanSize + 1
}

func (o *Ocean) AddShip(size ShipSize) {
	for _, shp := range o.ships {
		shp.SetFocus(false)
	}

	shp := NewShip(o.x+1, o.y+1, OriR, size)
	shp.SetFocus(true)
	o.shipsToAdd = append(o.shipsToAdd, shp)
}
