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

func (pp PlayerPosition) ConvertToPositionMap() (nameToPos map[string]string, posToName map[string]string) {
	nameToPos = make(map[string]string)
	posToName = make(map[string]string)
	if pp.UTG != nil {
		posToName["UTG"] = *pp.UTG
	}
	if pp.MP != nil {
		posToName["MP"] = *pp.MP
	}
	if pp.CO != nil {
		posToName["CO"] = *pp.CO
	}
	if pp.BTN != nil {
		posToName["BTN"] = *pp.BTN
	}
	if pp.SB != nil {
		posToName["SB"] = *pp.SB
	}
	if pp.BB != nil {
		posToName["BB"] = *pp.BB
	}
	for pos, name := range posToName {
		nameToPos[name] = pos
	}
	return
}