package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

type Game struct {
	current *SetupMgr
}

func NewGame(mgr *SetupMgr) *Game {
	return &Game{
		current: mgr,
	}
}

func (g *Game) Tick(ev tl.Event) {
	g.current.Tick(ev)
}

func (g *Game) Draw(screen *tl.Screen) {
	g.current.Draw(screen)
}
