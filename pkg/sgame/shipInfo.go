package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

const TableSize = 9
const TableWidth = 18

var ships = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

type Table struct {
	x           int
	y           int
	i           int
	dep         *tl.Rectangle
	next        *tl.Rectangle
	ocean       *Ocean
	initialized bool
}

func (t *Table) Position() (int, int) {
	return t.x, t.y
}

func (t *Table) Size() (int, int) {
	return TableSize, TableSize
}

func NewTable(x, y int, o *Ocean) *Table {
	tbl := &Table{
		x:     x,
		y:     y,
		ocean: o,
	}
	return tbl
}

func (t *Table) Draw(s *tl.Screen) {
	if !t.initialized {
		t.initialize(s)
	}

	if t.i == len(ships) {
		return
	}

	t.dep.SetSize(ships[t.i], 1)

	if t.i+1 < len(ships) {
		t.next.SetSize(ships[t.i+1], 1)
	}
}

func (t *Table) Tick(ev tl.Event) {
	// 	if ev.Type == tl.EventKey {
	// 		switch ev.Key {
	// 		case tl.KeyEnter:
	// 			t.ocean.AddShip(ShipSize(ships[t.i]))
	// 			t.i += 1
	// 		}
	// 	}
}

func (t *Table) GetShip() ShipSize {
	i := t.i
	if len(ships)-1 == i {
		return ShipSize(0)
	}

	if i < len(ships)-1 {
		t.i++
	}

	return ShipSize(ships[i])
}

func (t *Table) initialize(s *tl.Screen) {
	top := NewBar(t.x, t.y, TableWidth, OriR, tl.ColorWhite, frameCol)
	s.AddEntity(top)

	right := NewBar(t.x+17, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	s.AddEntity(right)

	left := NewBar(t.x, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	s.AddEntity(left)

	bottom := NewBar(t.x, t.y+9, TableWidth, OriR, tl.ColorWhite, frameCol)
	s.AddEntity(bottom)

	depTxt := NewText(t.x+2, t.y+2, "NOW DEPLOYING:", OriR, tl.ColorWhite, tl.ColorBlack)
	s.AddEntity(depTxt)

	t.dep = tl.NewRectangle(t.x+7, t.y+3, ships[t.i], 1, tl.ColorBlue)
	s.AddEntity(t.dep)

	nextTxt := NewText(t.x+4, t.y+5, "NEXT SHIP:", OriR, tl.ColorWhite, tl.ColorBlack)
	s.AddEntity(nextTxt)

	t.next = tl.NewRectangle(t.x+7, t.y+6, ships[t.i+1], 1, tl.ColorBlue)
	s.AddEntity(t.next)

	t.initialized = true
}
