package zombiedice

import (
	"fmt"
)

type Player struct {
	Strategy Strategy
	Brains int
	Wins int
}

func (p *Player) String() string {
	return fmt.Sprintf("%s: wins: %d", p.Strategy, p.Wins)
}

func (p *Player) Haswon() bool {
	return p.Brains>=13
}

func (p *Player) Newround() {
	p.Brains=0
	p.Strategy.Newround()
}

func NewPlayer(s Strategy) *Player {
	return &Player{s, 0, 0}
}
