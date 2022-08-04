package models

import "fmt"

type HandHistory struct {
	EntryPart    string // hand id, stake size, hand time, seats info
	PreflopPart  string // preflop actions
	FlopPart     string // flop actions
	TurnPart     string // turn actions
	RiverPart    string // river actions,
	ShowdownPart string // showdown part
	SummaryPart  string // hand summary
}

func (hh HandHistory) Print() {
	fmt.Println("EntryPart:")
	fmt.Println(hh.EntryPart)
	fmt.Println("PreflopPart:")
	fmt.Println(hh.PreflopPart)
	fmt.Println("FlopPart:")
	fmt.Println(hh.FlopPart)
	fmt.Println("TurnPart:")
	fmt.Println(hh.TurnPart)
	fmt.Println("RiverPart:")
	fmt.Println(hh.RiverPart)
	fmt.Println("ShowdownPart:")
	fmt.Println(hh.ShowdownPart)
	fmt.Println("SummaryPart:")
	fmt.Println(hh.SummaryPart)
}