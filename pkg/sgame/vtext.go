package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

// Text represents a string that can be drawn to the screen.
type Text struct {
	x      int
	y      int
	fg     tl.Attr
	bg     tl.Attr
	ori    Orientation
	text   []rune
	canvas []tl.Cell
}

// NewText creates a new Text, at position (x, y). It sets the Text's
// background and foreground colors to fg and bg respectively, and sets the
// Text's text to be text.
// Returns a pointer to the new Text.
func NewText(x, y int, text string, ori Orientation, fg, bg tl.Attr) *Text {
	str := []rune(text)
	c := make([]tl.Cell, len(str))
	for i := range c {
		c[i] = tl.Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Text{
		x:      x,
		y:      y,
		fg:     fg,
		bg:     bg,
		ori:    ori,
		text:   str,
		canvas: c,
	}
}

func (t *Text) Tick(ev tl.Event) {}

// Draw draws the Text to the Screen s.
func (t *Text) Draw(s *tl.Screen) {
	w, _ := t.Size()
	newX, newY := t.x, t.y

	for i := 0; i < w; i++ {
		switch t.ori {
		case OriR:
			newX, newY = t.x+i, t.y

		case OriD:
			newX, newY = t.x, t.y+i
		}

		s.RenderCell(newX, newY, &t.canvas[i])
	}
}

// Position returns the (x, y) coordinates of the Text.
func (t *Text) Position() (int, int) {
	return t.x, t.y
}

// Size returns the width and height of the Text.
func (t *Text) Size() (int, int) {
	return len(t.text), 1
}

// SetPosition sets the coordinates of the Text to be (x, y).
func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}

// Text returns the text of the Text.
func (t *Text) Text() string {
	return string(t.text)
}

// SetText sets the text of the Text to be text.
func (t *Text) SetText(text string) {
	t.text = []rune(text)
	c := make([]tl.Cell, len(t.text))
	for i := range c {
		c[i] = tl.Cell{Ch: t.text[i], Fg: t.fg, Bg: t.bg}
	}
	t.canvas = c
}

// Color returns the (foreground, background) colors of the Text.
func (t *Text) Color() (tl.Attr, tl.Attr) {
	return t.fg, t.bg
}

// SetColor sets the (foreground, background) colors of the Text
// to fg, bg respectively.
func (t *Text) SetColor(fg, bg tl.Attr) {
	t.fg = fg
	t.bg = bg
	for i := range t.canvas {
		t.canvas[i].Fg = fg
		t.canvas[i].Bg = bg
	}
}
