package services

import (
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
)

func GenerateRFIRanges(txtFilePath string) (map[string]*models.PreflopRFIRange, error) {
	hands, err := utils.ReadGGHandsFile(txtFilePath)
	if err != nil {
		return nil, err
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
	rfiHandCnt := 0
	for _, handStr := range hands {
		handHistory := utils.ParseHandStrToHandHistory(handStr)
		playerPostitions := utils.ParseHandHistoryToPlayerPositions(handHistory)
		if playerPostitions.NumOfPlayers != 6 {
			continue
		}
		utils.ParsePreflop6MaxTableToRFIRange(&rfiHandCnt, handHistory.PreflopPart, playerPostitions, &rfiMap)
	}
	return rfiMap, nil
}
