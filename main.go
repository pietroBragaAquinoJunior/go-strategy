package main

import (
	"strategy/action"
	player "strategy/player"
)

func main() {
	d := &action.Defense{}
	a := &action.Attack{}
	h := &action.Heal{}
	p := player.NewPlayer("pietro", a)
	p.TakeAction()
	p.SetStrategy(d)
	p.TakeAction()
	p.SetStrategy(h)
	p.TakeAction()
}
