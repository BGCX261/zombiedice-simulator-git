package zombiedice

import (
	"fmt"
)

type ImprovedStrategy struct {
	Name string
	totalbrains int
	numrolls int
	numturns int
	totalbrainsforround int
	shots int
	brains int
	redDice int
	yellowDice int
	greenDice int
}


func NewImprovedStrategy() Strategy {
	return &ImprovedStrategy{}
}

func (n *ImprovedStrategy) String() string {
	return fmt.Sprintf("%s. totalbrains: %d, rolls: %d, turns: %d (%f brains/turn)", n.Getname(), n.totalbrains, n.numrolls, n.numturns, float64(n.totalbrains)/float64(n.numturns))
}

func (n *ImprovedStrategy) Getname() string {
	return "improved"
}

func (n *ImprovedStrategy) Newround() {
	n.reset()
	n.totalbrains+=n.totalbrainsforround
	n.totalbrainsforround=0
}

func (n *ImprovedStrategy) reset() {
	n.shots=0
	n.brains=0
	n.redDice=0
	n.yellowDice=0
	n.greenDice=0
}

func (n *ImprovedStrategy) getWeightedShots() float64 {
	return (0.8*float64(n.redDice) + float64(n.yellowDice) + 1.2*float64(n.greenDice))*float64(n.shots)
}

func (n *ImprovedStrategy) bank() {
	n.totalbrainsforround+=n.brains
	n.reset()
}

func (n *ImprovedStrategy) Hitme() bool {
	if n.brains+n.totalbrainsforround>=13 {
		n.bank()
		n.numturns+=1
		return false
	}
	if n.shots>=3 {
		n.reset()
		n.numturns+=1
		return false
	}
	if n.shots>=2 && n.brains>=1 {
		n.bank()
		n.numturns+=1
		return false
	}
	if n.shots>=1 && n.brains>=5 {
		n.bank()
		n.numturns+=1
		return false
	}
	if n.shots>=0 && n.brains>=9 {
		n.bank()
		n.numturns+=1
		return false
	}
	n.numrolls+=1
	return true
}

func (n *ImprovedStrategy) Adddice(di [3]*Die) {
	for _,d := range di {
		switch d.Value {
		case "brains":
			n.brains+=1
		case "shot":
			n.shots+=1
		}
		if d.Value!="escape" {
			switch d.Name {
			case "Red":
				n.redDice+=1
			case "Yellow":
				n.yellowDice+=1
			case "Green":
				n.greenDice+=1
			}
		}
	}
}

