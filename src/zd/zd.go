package main

import (
	"fmt"
	"math/rand"
	"time"
	"zombiedice"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	player1:=zombiedice.NewPlayer(zombiedice.NewNaiveStrategy())
	player2:=zombiedice.NewPlayer(zombiedice.NewImprovedStrategy())
	table:=zombiedice.NewTable()
	table.Players=append(table.Players, player1, player2)

	wonbystarter:=0
	starter:=0
	numRounds:=100000

	for ;table.RoundsPlayed<numRounds;table.RoundsPlayed+=1 {
		starter=table.Currentplayer
		for p:=table.GetNextPlayer();;p=table.GetNextPlayer() {
			table.Turn=zombiedice.NewTurn()
			for ; p.Strategy.Hitme(); {
				res, err:=table.Turn.Roll3()
				if err!=nil {
					fmt.Print(err)
				}
				p.Strategy.Adddice(res)
			}
			p.Brains+=table.Turn.Brains
			if table.Iswon() {
				break
			}
		}
		for i, p:= range table.Players {
			if p.Haswon() {
				if i==starter {
					wonbystarter+=1
				}
				p.Wins+=1
			}
			p.Newround()
		}
	}

	for _, p:= range table.Players {
		fmt.Println(p)
	}
	fmt.Printf("won by starter: %d\n", wonbystarter)
}