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

func NewOcean(x, y int) *Ocean {
	o := &Ocean{
		x: x,
		y: y,
	}
	return o
}

func (o *Ocean) Draw(s *tl.Screen) {
	txtTop := NewText(o.x, o.y, " 0123456789", OriR, tl.ColorWhite, frameCol)
	right := tl.NewRectangle(o.x+OceanSize+1, o.y, 1, OceanSize+2, tl.ColorBlue)
	bottom := tl.NewRectangle(o.x, o.y+OceanSize+1, OceanSize+1, 1, frameCol)
	txtLeft := NewText(o.x, o.y+1, "ABCDEFGHIJ", OriD, tl.ColorWhite, frameCol)

	s.AddEntity(txtTop)
	s.AddEntity(right)
	s.AddEntity(bottom)
	s.AddEntity(txtLeft)
}

// func (o *Ocean) Draw(s *tl.Screen) {
// 	wave := &tl.Cell{
// 		Fg: tl.ColorWhite,
// 		Bg: tl.ColorBlue,
// 		Ch: '~',
// 	}
//
// 	alp := 'A'
// 	dig := '0'
//
// 	for x := o.x; x <= o.x+OceanSize; x++ {
// 		for y := o.y; y <= o.y+OceanSize; y++ {
// 			// Letters.
// 			if y == o.y {
// 				label := &tl.Cell{
// 					Fg: tl.ColorWhite,
// 					Bg: frameCol,
// 					Ch: rune(alp),
// 				}
//
// 				if x == o.x {
// 					label.Ch = ' '
// 					alp -= 1
// 				}
//
// 				s.RenderCell(x, y, label)
// 				alp += 1
// 				continue
// 			}
//
// 			// Digits.
// 			if x == o.x {
// 				label := &tl.Cell{
// 					Fg: tl.ColorWhite,
// 					Bg: frameCol,
// 					Ch: rune(dig),
// 				}
// 				s.RenderCell(x, y, label)
// 				dig += 1
// 				continue
// 			}
//
// 			if x >= o.x+1 && y >= o.y {
// 				s.RenderCell(x, y, wave)
// 				continue
// 			}
// 		}
// 	}
//
// 	frame := &tl.Cell{
// 		Fg: tl.ColorWhite,
// 		Bg: frameCol,
// 		Ch: ' ',
// 	}
//
// 	for x := o.x; x <= o.x+OceanSize+1; x++ {
// 		s.RenderCell(x, o.y+OceanSize+1, frame)
// 	}
//
// 	for y := o.y; y <= o.y+OceanSize+1; y++ {
// 		s.RenderCell(o.x+OceanSize+1, y, frame)
// 	}
// }

func (o *Ocean) Tick(_ tl.Event) {}
