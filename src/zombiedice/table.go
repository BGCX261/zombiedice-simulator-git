package zombiedice


type Table struct {
	Turn Turn
	Players []*Player
	RoundsPlayed int
	Currentplayer int
}

func (t *Table) GetNextPlayer() *Player {
	p:=t.Players[t.Currentplayer]
	
	t.Currentplayer=(t.Currentplayer+1)%len(t.Players)
	return p
}

func (t *Table) Iswon() bool {
	for _, p := range t.Players {
		if p.Haswon() {
			return true
		}
	}
	return false
}

func NewTable() Table {
	return Table{NewTurn(), []*Player{}, 0, 0}
}
	
