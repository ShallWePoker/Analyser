package consts

var CardCharToWeight = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

var WeightToCardChar = map[int]string{
	14: "A",
	13: "K",
	12: "Q",
	11: "J",
	10: "T",
	9:  "9",
	8:  "8",
	7:  "7",
	6:  "6",
	5:  "5",
	4:  "4",
	3:  "3",
	2:  "2",
}
