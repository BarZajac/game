package sgame

import tl "github.com/JoelOtter/termloop"

const TableSize = 9
const TableWidth = 18

type Table struct {
	x int
	y int
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
	a := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
	top := NewBar(t.x, t.y, TableWidth, OriR, tl.ColorWhite, frameCol)
	right := NewBar(t.x+17, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	depTxt := NewText(t.x+2, t.y+2, "NOW DEPLOYING:", OriR, tl.ColorWhite, tl.ColorBlack)
	dep := tl.NewRectangle(t.x+7, t.y+3, a[0], 1, tl.ColorBlue)
	nextTxt := NewText(t.x+4, t.y+5, "NEXT SHIP:", OriR, tl.ColorWhite, tl.ColorBlack)
	next := tl.NewRectangle(t.x+7, t.y+6, a[1], 1, tl.ColorBlue)
	left := NewBar(t.x, t.y, TableSize, OriD, tl.ColorWhite, frameCol)
	bottom := NewBar(t.x, t.y+9, TableWidth, OriR, tl.ColorWhite, frameCol)

	s.AddEntity(top)
	s.AddEntity(right)
	s.AddEntity(depTxt)
	s.AddEntity(dep)
	s.AddEntity(nextTxt)
	s.AddEntity(next)
	s.AddEntity(left)
	s.AddEntity(bottom)
}

func (t *Table) Tick(_ tl.Event) {}
