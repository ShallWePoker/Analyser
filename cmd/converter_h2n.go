package main

import (
	"bufio"
	"os"
	"regexp"
)

func main() {
	singleFile, err := os.Open("/path/to/combined.txt")
	if err != nil {
		panic(err)
	}
	writeFile, err := os.Create("/path/to/converted.txt")
	if err != nil {
		panic(err)
	}
	datawriter := bufio.NewWriter(writeFile)
	defer singleFile.Close()
	defer writeFile.Close()
	scanner := bufio.NewScanner(singleFile)
	reg1, err := regexp.Compile("Dealt to Hero")
	reg2, err := regexp.Compile("Dealt to [0-9A-Za-z_]{0,}\\s+")
	reg3, err := regexp.Compile("XXXXXXXXXXXXXXX")
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		replaced := reg1.ReplaceAllString(scanner.Text(), "XXXXXXXXXXXXXXX")
		replaced = reg2.ReplaceAllString(replaced, "")
		replaced = reg3.ReplaceAllString(replaced, "Dealt to Hero")
		_, _ = datawriter.WriteString(replaced + "\n")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
