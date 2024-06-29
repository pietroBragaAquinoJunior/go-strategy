package player

import "strategy/strategy"

type Player struct {
	hp       int
	name     string
	strategy strategy.Strategy
}

func NewPlayer(name string, strategy strategy.Strategy) *Player {
	return &Player{hp: 10, name: name, strategy: strategy}
}

func (p *Player) SetStrategy(strategy strategy.Strategy) {
	p.strategy = strategy
}

func (p *Player) TakeAction() {
	p.strategy.TakeAction()
}
