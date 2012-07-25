package zombiedice

import (
	"fmt"
)

type NaiveStrategy struct {
	totalbrains int
	numrolls int
	numturns int
	totalbrainsforround int
	shots int
	brains int
}

func (n *NaiveStrategy) Newround() {
	n.totalbrains+=n.totalbrainsforround
	n.totalbrainsforround=0
}

func NewNaiveStrategy() Strategy {
	return &NaiveStrategy{}
}

func (n *NaiveStrategy) String() string {
	return fmt.Sprintf("%s. totalbrains: %d, rolls: %d, turns: %d (%f brains/turn)", n.Getname(), n.totalbrains, n.numrolls, n.numturns, float64(n.totalbrains)/float64(n.numturns))
}

func (n *NaiveStrategy) Getname() string {
	return "naive"
}

func (n *NaiveStrategy) Adddice(di [3]*Die) {
	for _, d := range di {
		switch d.Value {
		case "shot": n.shots+=1
		case "brains": n.brains+=1
		}
	}
}

func (n *NaiveStrategy) Hitme() bool {
	if n.shots<2 {
		n.numrolls+=1
		return true
	}
	n.numturns+=1
	if n.shots==2 {
		n.totalbrainsforround+=n.brains
		n.brains=0
		n.shots=0
		return false
	}
	n.brains=0
	n.shots=0
	return false
}
