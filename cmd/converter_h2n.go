package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	handsTxtFolder := flag.String("hands-txt-folder", "", "hands txt folder")
	combinedTxtPath := flag.String("combined-txt-path", "", "combined txt path")
	playerName := flag.String("player-name", "", "player name")
	flag.Parse()
	fmt.Printf("hands txt folder: %s\n", *handsTxtFolder)
	fmt.Printf("combined txt file path: %s\n", *combinedTxtPath)
	fmt.Printf("player name: %s\n", *playerName)

	allTxtFiles, err := ListAllTxtFiles(*handsTxtFolder)
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(*combinedTxtPath)
	if err == nil {
		fmt.Printf("combined txt file already exists. will delete it and create a new one.")
		err = os.Remove(*combinedTxtPath)
		if err != nil {
			fmt.Printf("cannot delete %s.", *combinedTxtPath)
			panic(err)
		}
	}
	writeFile, err := os.Create(*combinedTxtPath)
	if err != nil {
		panic(err)
	}
	datawriter := bufio.NewWriter(writeFile)
	defer writeFile.Close()

	pokerHandRegexp, err := regexp.Compile("(Poker Hand #)(RC|HD)")
	if err != nil {
		panic(err)
	}
	dealtToHeroRegexp, err := regexp.Compile("Dealt to Hero")
	if err != nil {
		panic(err)
	}
	dealtToRandomRegexp, err := regexp.Compile("Dealt to [0-9A-Za-z_]{0,}\\s+")
	if err != nil {
		panic(err)
	}
	xxxxxRegexp, err := regexp.Compile("XXXXXXXXXXXXXXX")
	if err != nil {
		panic(err)
	}
	heroRegexp, err := regexp.Compile("Hero")
	if err != nil {
		panic(err)
	}
	for _, singleTxtFile := range allTxtFiles {
		err = readFilesAppendToCombinedFile(singleTxtFile, datawriter, *playerName, pokerHandRegexp, dealtToHeroRegexp, dealtToRandomRegexp, xxxxxRegexp, heroRegexp)
		if err != nil {
			fmt.Printf("err dealing with txt file %s: %+v", singleTxtFile, err)
			panic(err)
		}
	}
}

func ListAllTxtFiles(dir string) (filePaths []string, err error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		return
	}
	if !dirInfo.IsDir() {
		return filePaths, errors.New(dir + "is not a directory")
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && info.Mode()&os.ModeSymlink == 0 {
			pathParts := strings.Split(path, "/")
			if strings.HasSuffix(path, ".txt") && strings.HasPrefix(pathParts[len(pathParts)-1], "GG") {
				filePaths = append(filePaths, path)
			}
		}
		return nil
	})
	if err != nil {
		return filePaths, err
	}
	return
}

func readFilesAppendToCombinedFile(handsTxtFile string, datawriter *bufio.Writer, playerName string, pokerHandRegexp, dealtHeroRegexp, dealtToRandomRegexp, xxxxxRegexp, heroRegexp *regexp.Regexp) error {
	fileContent, err := os.Open(handsTxtFile)
	if err != nil {
		fmt.Printf("cannot open txt file %s: %+v; skipping.", handsTxtFile, err)
		return nil
	}
	defer fileContent.Close()
	scanner := bufio.NewScanner(fileContent)
	for scanner.Scan() {
		// TODO PokerStars Hand #20 meaning ?
		replaced := pokerHandRegexp.ReplaceAllString(scanner.Text(), "PokerStars Hand #20")
		replaced = dealtHeroRegexp.ReplaceAllString(replaced, "XXXXXXXXXXXXXXX")
		replaced = dealtToRandomRegexp.ReplaceAllString(replaced, "")
		replaced = xxxxxRegexp.ReplaceAllString(replaced, fmt.Sprintf("Dealt to %s", playerName))
		replaced = heroRegexp.ReplaceAllString(replaced, playerName)
		_, err = datawriter.WriteString(replaced + "\n")
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
