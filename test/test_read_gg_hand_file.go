package main

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
	"os"
)

func main() {
	hands, err := utils.ReadGGHandsFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, hand := range hands {
		fmt.Printf("%s", hand)
	}
}
