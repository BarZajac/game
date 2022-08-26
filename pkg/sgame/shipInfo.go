package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

const TableSize = 9
const TableWidth = 18

var a = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

type Table struct {
	x    int
	y    int
	i    int
	dep  *tl.Rectangle
	next *tl.Rectangle
}

func (t *Table) Position() (int, int) {
	return t.x, t.y
}

func (t *Table) Size() (int, int) {
	return TableSize, TableSize
}

func NewTable(x, y int) *Table {
	o := &Table{
		x: x,
		y: y,
	}
	return o
}

func (t *Table) Draw(s *tl.Screen) {
	if t.dep == nil {
		t.dep = tl.NewRectangle(t.x+7, t.y+3, a[t.i], 1, tl.ColorBlue)
		s.AddEntity(t.dep)

	}
	if t.next == nil {
		t.next = tl.NewRectangle(t.x+7, t.y+6, a[t.i+1], 1, tl.ColorBlue)
		s.AddEntity(t.next)

	}

	top := NewBar(t.x, t.y, TableWidth, OriR, tl.ColorWhite, frameCol)
	right := NewBar(t.x+17, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	depTxt := NewText(t.x+2, t.y+2, "NOW DEPLOYING:", OriR, tl.ColorWhite, tl.ColorBlack)
	nextTxt := NewText(t.x+4, t.y+5, "NEXT SHIP:", OriR, tl.ColorWhite, tl.ColorBlack)
	left := NewBar(t.x, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	bottom := NewBar(t.x, t.y+9, TableWidth, OriR, tl.ColorWhite, frameCol)

	t.dep.SetSize(a[t.i], 1)

	s.AddEntity(top)
	s.AddEntity(right)
	s.AddEntity(depTxt)
	s.AddEntity(nextTxt)
	s.AddEntity(left)
	s.AddEntity(bottom)
}

func (t *Table) Tick(ev tl.Event) {

	if ev.Type == tl.EventKey {

		switch ev.Key {
		case tl.KeyEnter:
			t.i += 1
		}

	}

}
