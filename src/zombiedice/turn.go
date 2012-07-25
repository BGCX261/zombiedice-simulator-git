package zombiedice

import (
	"errors"
)

type Turn struct {
	dice Dice
	shots int
	Brains int
}

func NewTurn() Turn {
	return Turn{NewDice(), 0, 0}
}

func (r *Turn) Roll3() ([3]*Die, error) {
	if r.shots>=3 {
		return [3]*Die{}, errors.New("too many shots")
	}
	res, err:= r.dice.roll3()
	if err!=nil {
		if err.Error()=="too few dice" {
			
			r.dice=NewDice()
			return r.Roll3()
		}
	}
	for _, d := range res {
		switch d.Value {
		case "shot": r.shots+=1
		case "brains": r.Brains+=1
		}
	}
	if r.shots>=3 {
		r.Brains=0
	}
	return res, err
}

