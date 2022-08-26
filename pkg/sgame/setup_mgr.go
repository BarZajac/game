package sgame

import (
	tl "github.com/JoelOtter/termloop"
)

type SetupMgr struct {
	ocean       *Ocean
	si          *Table
	current     tl.Drawable
	initialized bool
	removeSI    bool
}

func NewSetupMgr(o *Ocean, si *Table) *SetupMgr {
	return &SetupMgr{
		ocean:   o,
		si:      si,
		current: o,
	}
}

func (mgr *SetupMgr) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey {
		switch ev.Key {
		case tl.KeyEnter:
			if mgr.si != nil {
				ss := mgr.si.GetShip()
				if ss == 0 {
					mgr.removeSI = true
				}
				mgr.ocean.AddShip(ss)
			}
			mgr.current.Tick(ev)

		default:
			mgr.current.Tick(ev)
		}
	}
}

func (mgr *SetupMgr) Draw(screen *tl.Screen) {
	mgr.current.Draw(screen)
	if !mgr.initialized {
		screen.AddEntity(mgr.si)
		ss := mgr.si.GetShip()
		mgr.ocean.AddShip(ss)
		mgr.initialized = true
	}

	if mgr.si != nil {
		if mgr.removeSI {
			screen.RemoveEntity(mgr.si)
			mgr.si = nil
		} else {
			// mgr.si.Draw(screen)
		}
	}
}
