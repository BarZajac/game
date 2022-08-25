package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

const OceanSize int = 10

var frameCol = tl.RgbTo256Color(90, 90, 90)

type Ocean struct {
	x int
	y int
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
	txtTop := NewText(o.x, o.y, " 0123456789", OriR, tl.ColorWhite, frameCol)
	right := NewBar(o.x+OceanSize+1, o.y, OceanSize+2, OriD, tl.ColorWhite, frameCol)
	bottom := NewBar(o.x, o.y+OceanSize+1, OceanSize+1, OriR, tl.ColorWhite, frameCol)
	txtLeft := NewText(o.x, o.y+1, "ABCDEFGHIJ", OriD, tl.ColorWhite, frameCol)

	s.AddEntity(txtTop)
	s.AddEntity(right)
	s.AddEntity(bottom)
	s.AddEntity(txtLeft)
}

func (o *Ocean) Tick(_ tl.Event) {}

func (o *Ocean) Definition() (int, int, int, int) {
	return o.x, o.y, o.x + OceanSize + 1, o.y + OceanSize + 1
}
