package zombiedice

import (
	"math/rand"
	"errors"
)

type Dice struct {
	dice []Die
}

func (d *Dice) pick1() *Die {
	n:=rand.Intn(len(d.dice))
	r:=d.dice[n]
	copy(d.dice[n:], d.dice[n+1:])
	d.dice=d.dice[:len(d.dice)-1]
	return &r
}

func (di *Dice) roll3() ([3]*Die, error) {
	if len(di.dice)<3 {
		return [3]*Die{}, errors.New("too few dice")
	}
	picked:= [3]*Die{di.pick1(), di.pick1(), di.pick1()}
	for _, d := range picked {
		val:=d.roll()
		if val=="escape" {
			di.dice=append(di.dice, *d)
		}
	}
	return picked, nil
}

func NewDice() Dice {
	return Dice{ []Die{NewRedDie(),
			NewRedDie(),
			NewRedDie(),
			NewYellowDie(),
			NewYellowDie(),
			NewYellowDie(),
			NewYellowDie(),
			NewGreenDie(),
			NewGreenDie(),
			NewGreenDie(),
			NewGreenDie(),
			NewGreenDie(),
			NewGreenDie()} }
}

