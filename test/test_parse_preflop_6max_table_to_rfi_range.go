package main

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
	"os"
)

func main() {
	hands, err := utils.ReadGGHandsFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	utgRfi := models.InitPreflopRFIRange("Hero", "UTG")
	mpRfi := models.InitPreflopRFIRange("Hero", "MP")
	coRfi := models.InitPreflopRFIRange("Hero", "CO")
	btnRfi := models.InitPreflopRFIRange("Hero", "BTN")
	sbRfi := models.InitPreflopRFIRange("Hero", "SB")
	rfiMap := map[string]*models.PreflopRFIRange{
		"UTG": &utgRfi,
		"MP":  &mpRfi,
		"CO":  &coRfi,
		"BTN": &btnRfi,
		"SB":  &sbRfi,
	}
	fmt.Printf("total hands: %d\n", len(hands))
	rfiHandCnt := 0
	for _, handStr := range hands {
		handHistory := utils.ParseHandStrToHandHistory(handStr)
		playerPostitions := utils.ParseHandHistoryToPlayerPositions(handHistory)
		if playerPostitions.NumOfPlayers != 6 {
			continue
		}
		utils.ParsePreflop6MaxTableToRFIRange(&rfiHandCnt, handHistory.PreflopPart, playerPostitions, &rfiMap)
	}
	fmt.Printf("rfi hand cnt: %d\n", rfiHandCnt)
	rfiMap["UTG"].PrintRFIMatrix()
	rfiMap["MP"].PrintRFIMatrix()
	rfiMap["CO"].PrintRFIMatrix()
	rfiMap["BTN"].PrintRFIMatrix()
	rfiMap["SB"].PrintRFIMatrix()

}
