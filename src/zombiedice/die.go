package zombiedice

import (
	"fmt"
	"math/rand"
)

type Die struct {
	Name string
	Values [6]Result
	Value Result
}

func (d *Die) printMe() string {
	return fmt.Sprintf("%s: %s\n", d.Name, d.Value)
}

func (d *Die) roll() Result {
	d.Value=d.Values[rand.Int31n(6)]
	return d.Value
}

func NewRedDie() Die {
	return Die{"Red", [6]Result{"shot",
			"shot",
			"shot",
			"escape",
			"escape",
			"brains"}, ""}
}

func NewYellowDie() Die {
	return Die{"Yellow", [6]Result{"shot", "shot", "escape", "escape",
			"brains", "brains"}, ""}
}

func NewGreenDie() Die {
	return Die{"Green", [6]Result{"shot", "escape", "escape",
			"brains", "brains", "brains"}, ""}
}

