package utils

// full ring 6 max table
/*
func ParsePreflop6MaxTableToVsRFIRange(preflopPart string, playerPositions models.PlayerPosition, vsRFIRanges *map[string]*models.PreflopVsRFIRange) {
	nameToPos, _ := playerPositions.ConvertToPositionMap()
	heroPos := nameToPos["Hero"]
	if heroPos == "UTG" {
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

}


 */