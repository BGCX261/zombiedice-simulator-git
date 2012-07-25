package zombiedice

type Strategy interface {
	Getname() string
	Adddice([3]*Die) 
	Hitme() bool
	Newround()
}
