package main

import (
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
	"os"
)

func main() {
	hands, err := utils.ReadGGHandsFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	hand := hands[0]
	handHistory := utils.ParseHandStrToHandHistory(hand)
	handHistory.Print()

	playerPositions := utils.ParseHandHistoryToPlayerPositions(handHistory)
	playerPositions.Print()
}

