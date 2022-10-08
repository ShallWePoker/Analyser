package utils

import (
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"strings"
)

// full ring 6 max table
func ParsePreflop6MaxTableToRFIRange(rfiHandCnt *int, preflopPart string, playerPositions models.PlayerPosition, rfiMap *map[string]*models.PreflopRFIRange) {
	nameToPos, _ := playerPositions.ConvertToPositionMap()
	heroPos := nameToPos["Hero"]
	if heroPos == "BB" {
		return
	}
	preflopLines := strings.Split(preflopPart, "\n")
	heroHoleCards := ""
	for i := 0; i < 6; i++ {
		if strings.Contains(preflopLines[i], "Hero") {
			heroHoleCards = strings.TrimPrefix(preflopLines[i], "Dealt to Hero")
			break
		}
	}
	heroHoleCards = strings.Trim(heroHoleCards, " ")
	heroHoleCards = UniformHoleCardStr(heroHoleCards)
	utgActionLine := preflopLines[6]
	mpActionLine := preflopLines[7]
	coActionLine := preflopLines[8]
	btnActionLine := preflopLines[9]
	sbActionLine := preflopLines[10]
	switch heroPos {
	case "UTG":
		*rfiHandCnt += 1
		contributeToRFI(utgActionLine, (*rfiMap)["UTG"], heroHoleCards)
	case "MP":
		if strings.Contains(utgActionLine, "fold") {
			*rfiHandCnt += 1
			contributeToRFI(mpActionLine, (*rfiMap)["MP"], heroHoleCards)
		}
	case "CO":
		if strings.Contains(utgActionLine, "fold") &&
			strings.Contains(mpActionLine, "fold") {
			*rfiHandCnt += 1
			contributeToRFI(coActionLine, (*rfiMap)["CO"], heroHoleCards)
		}
	case "BTN":
		if strings.Contains(utgActionLine, "fold") &&
			strings.Contains(mpActionLine, "fold") &&
			strings.Contains(coActionLine, "fold") {
			*rfiHandCnt += 1
			contributeToRFI(btnActionLine, (*rfiMap)["BTN"], heroHoleCards)
		}
	case "SB":
		if strings.Contains(utgActionLine, "fold") &&
			strings.Contains(mpActionLine, "fold") &&
			strings.Contains(coActionLine, "fold") &&
			strings.Contains(btnActionLine, "fold") {
			*rfiHandCnt += 1
			contributeToRFI(sbActionLine, (*rfiMap)["SB"], heroHoleCards)
		}
	}
}

func contributeToRFI(actionLine string, rfi *models.PreflopRFIRange, holecard string) {
	if strings.Contains(actionLine, "raise") {
		if value, exists := rfi.HolecardMap[holecard]; exists {
			value.Raise += 1
			rfi.HolecardMap[holecard] = value
		} else {
			rfi.HolecardMap[holecard] = models.Actions{Raise: 1}
		}
	} else if strings.Contains(actionLine, "fold") {
		if value, exists := rfi.HolecardMap[holecard]; exists {
			value.Fold += 1
			rfi.HolecardMap[holecard] = value
		} else {
			rfi.HolecardMap[holecard] = models.Actions{Fold: 1}
		}
	} else {
		if value, exists := rfi.HolecardMap[holecard]; exists {
			value.Call += 1
			rfi.HolecardMap[holecard] = value
		} else {
			rfi.HolecardMap[holecard] = models.Actions{Call: 1}
		}
	}
}
