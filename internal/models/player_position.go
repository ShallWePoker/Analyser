package models

import "fmt"

type PlayerPosition struct {
	NumOfPlayers int
	UTG          *string
	MP           *string
	CO           *string
	BTN          *string
	SB           *string
	BB           *string
}

func (pp PlayerPosition) Print() {
	fmt.Println(fmt.Sprintf("num of players: %d", pp.NumOfPlayers))
	if pp.UTG != nil {
		fmt.Println(fmt.Sprintf("UTG: %s", *pp.UTG))
	}
	if pp.MP != nil {
		fmt.Println(fmt.Sprintf("MP: %s", *pp.MP))
	}
	if pp.CO != nil {
		fmt.Println(fmt.Sprintf("CO: %s", *pp.CO))
	}
	if pp.BTN != nil {
		fmt.Println(fmt.Sprintf("BTN: %s", *pp.BTN))
	}
	if pp.SB != nil {
		fmt.Println(fmt.Sprintf("SB: %s", *pp.SB))
	}
	if pp.BB != nil {
		fmt.Println(fmt.Sprintf("BB: %s", *pp.BB))
	}
}