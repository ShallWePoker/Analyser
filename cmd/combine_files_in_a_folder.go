package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	handsTxtFolder := flag.String("hands-txt-folder", "", "hands txt folder")
	combinedTxtPath := flag.String("combined-txt-path", "", "combined txt path")
	flag.Parse()
	allTxtFiles, err := ListAllTxtFiles1(*handsTxtFolder)
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
	contents := []string{}
	for _, txtFile := range allTxtFiles {
		content, err := ioutil.ReadFile(txtFile)
		if err != nil {
			panic(err)
		}
		contents = append(contents, string(content))
	}
	err = ioutil.WriteFile(*combinedTxtPath, []byte(strings.Join(contents, "\n\n")), 0644)

}

func ListAllTxtFiles1(dir string) (filePaths []string, err error) {
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