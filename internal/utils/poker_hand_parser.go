package utils

import (
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"io/ioutil"
	"strings"
)

func ReadGGHandsFile(handsFilePath string) (separateHands []string, err error) {
	// TODO ignore how large the file may be, assuming you have infinite memory to load the file
	fileContentBytes, err := ioutil.ReadFile(handsFilePath)
	if err != nil {
		return
	}
	fileLines := strings.Split(string(fileContentBytes), "\n")
	for i := 0; i < len(fileLines); {
		line := fileLines[i]
		line = strings.Trim(line, " ")
		if line == "" {
			continue
		}
		handContent := ""
		if strings.HasPrefix(line, "Poker Hand #") {
			handContentStartIndex := i
			handContent = fileLines[i]+"\n"
			i++
			for j := handContentStartIndex+1; j < len(fileLines); j++ {
				if strings.HasPrefix(fileLines[j], "Poker Hand #") {
					break
				}
				if strings.Trim(fileLines[j], " ") == "" {
					i++
					continue
				}
				handContent = handContent + fileLines[j]+"\n"
				i++
			}
		}
		separateHands = append(separateHands, handContent)
	}
	return separateHands, nil
}

func ParseHandStrToHandHistory(handStr string) (handHistory models.HandHistory) {
	handLines := strings.Split(handStr, "\n")
	preflopStartLineIndex := 0
	entryPartStr := ""
	for i := 0; i < len(handLines); i++ {
		if strings.HasPrefix(handLines[i], "***") {
			preflopStartLineIndex = i
			break
		}
		entryPartStr = entryPartStr + handLines[i] + "\n"
	}
	handHistory.EntryPart = entryPartStr
	for i := preflopStartLineIndex; i < len(handLines); {
		partTitle := strings.TrimPrefix(handLines[i], "***")
		partTitle = strings.TrimPrefix(partTitle, " ")
		starIndex := strings.Index(partTitle, "*")
		partTitle = strings.Trim(partTitle[:starIndex], " ")
		partStr := ""
		i++
		for j := i; j < len(handLines); j++ {
			if strings.HasPrefix(handLines[j], "***") {
				break
			}
			partStr = partStr + handLines[j] + "\n"
			i++
		}
		switch partTitle {
		case "HOLE CARDS":
			handHistory.PreflopPart = partStr
		case "FLOP":
			handHistory.FlopPart = partStr
		case "TURN":
			handHistory.TurnPart = partStr
		case "RIVER":
			handHistory.RiverPart = partStr
		case "SHOWDOWN":
			handHistory.ShowdownPart = partStr
		case "SUMMARY":
			handHistory.SummaryPart = partStr
		}
	}
	return
}

func ParseHandHistoryToPlayerPositions(hh models.HandHistory) (pp models.PlayerPosition) {
	entryPartStr := hh.EntryPart
	entryParts := strings.Split(entryPartStr, "\n")
	btnSeatNo := "Seat " + string(entryParts[1][strings.Index(entryParts[1], "#")+1])
	seats := []string{}
	for i := 2; i < len(entryParts); i++ {
		if !strings.HasPrefix(entryParts[i], "Seat") {
			break
		}
		if strings.HasPrefix(entryParts[i], "Seat") {
			seats = append(seats, entryParts[i])
		}
	}
	pp.NumOfPlayers = len(seats)
	if len(seats) == 2 {
		btnIndex := findIndexOf(seats, btnSeatNo)
		btnPlayerName := GetPlayerNameFromLine(seats[btnIndex])
		pp.BTN = &btnPlayerName
		BBPlayerIndex := (btnIndex+1) % 2
		BBPlayerName := GetPlayerNameFromLine(seats[BBPlayerIndex])
		pp.BB = &BBPlayerName
		return
	}
	if len(seats) == 3 {
		btnIndex := findIndexOf(seats, btnSeatNo)
		btnPlayerName := GetPlayerNameFromLine(seats[btnIndex])
		pp.BTN = &btnPlayerName
		SBPlayerIndex := (btnIndex+1) % 3
		SBPlayerName := GetPlayerNameFromLine(seats[SBPlayerIndex])
		pp.SB = &SBPlayerName
		BBPlayerIndex := (SBPlayerIndex+1) % 3
		BBPlayerName := GetPlayerNameFromLine(seats[BBPlayerIndex])
		pp.BB = &BBPlayerName
		return
	}
	if len(seats) == 4 {
		btnIndex := findIndexOf(seats, btnSeatNo)
		btnPlayerName := GetPlayerNameFromLine(seats[btnIndex])
		pp.BTN = &btnPlayerName
		SBPlayerIndex := (btnIndex+1) % 4
		SBPlayerName := GetPlayerNameFromLine(seats[SBPlayerIndex])
		pp.SB = &SBPlayerName
		BBPlayerIndex := (SBPlayerIndex+1) % 4
		BBPlayerName := GetPlayerNameFromLine(seats[BBPlayerIndex])
		pp.BB = &BBPlayerName
		UTGPlayerIndex := (BBPlayerIndex+1) % 4
		UTGPlayerName := GetPlayerNameFromLine(seats[UTGPlayerIndex])
		pp.UTG = &UTGPlayerName
		return
	}
	if len(seats) == 5 {
		btnIndex := findIndexOf(seats, btnSeatNo)
		btnPlayerName := GetPlayerNameFromLine(seats[btnIndex])
		pp.BTN = &btnPlayerName
		SBPlayerIndex := (btnIndex+1) % 5
		SBPlayerName := GetPlayerNameFromLine(seats[SBPlayerIndex])
		pp.SB = &SBPlayerName
		BBPlayerIndex := (SBPlayerIndex+1) % 5
		BBPlayerName := GetPlayerNameFromLine(seats[BBPlayerIndex])
		pp.BB = &BBPlayerName
		UTGPlayerIndex := (BBPlayerIndex+1) % 5
		UTGPlayerName := GetPlayerNameFromLine(seats[UTGPlayerIndex])
		pp.UTG = &UTGPlayerName
		MPPlayerIndex := (UTGPlayerIndex+1) % 5
		MPPlayerName := GetPlayerNameFromLine(seats[MPPlayerIndex])
		pp.MP = &MPPlayerName
		return
	}
	if len(seats) == 6 {
		btnIndex := findIndexOf(seats, btnSeatNo)
		btnPlayerName := GetPlayerNameFromLine(seats[btnIndex])
		pp.BTN = &btnPlayerName
		SBPlayerIndex := (btnIndex+1) % 6
		SBPlayerName := GetPlayerNameFromLine(seats[SBPlayerIndex])
		pp.SB = &SBPlayerName
		BBPlayerIndex := (SBPlayerIndex+1) % 6
		BBPlayerName := GetPlayerNameFromLine(seats[BBPlayerIndex])
		pp.BB = &BBPlayerName
		UTGPlayerIndex := (BBPlayerIndex+1) % 6
		UTGPlayerName := GetPlayerNameFromLine(seats[UTGPlayerIndex])
		pp.UTG = &UTGPlayerName
		MPPlayerIndex := (UTGPlayerIndex+1) % 6
		MPPlayerName := GetPlayerNameFromLine(seats[MPPlayerIndex])
		pp.MP = &MPPlayerName
		COPlayerIndex := (MPPlayerIndex+1) % 6
		COPlayerName := GetPlayerNameFromLine(seats[COPlayerIndex])
		pp.CO = &COPlayerName
		return
	}
	return
}

func findIndexOf(strSlice []string, targetStrPrefix string) int {
	for i, str := range strSlice {
		if strings.HasPrefix(str, targetStrPrefix) {
			return i
		}
	}
	return -1
}

func GetPlayerNameFromLine(line string) string {
	index1 := strings.Index(line, ":")
	index2 := strings.Index(line, "(")
	return line[index1+2:index2-1]
}