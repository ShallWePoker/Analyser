package main

import (
	"fmt"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
	"os"
	"strconv"
)

func main() {
	iStr := os.Args[1]
	jStr := os.Args[2]
	i, _ := strconv.Atoi(iStr)
	j, _ := strconv.Atoi(jStr)
	fmt.Println(utils.MatrixToHolecards(i, j))
}
