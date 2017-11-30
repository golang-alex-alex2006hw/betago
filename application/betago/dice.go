package betago

import "fmt"

const (
	// ErrorDiceIsNotValid ...
	ErrorDiceIsNotValid = "Dice is not valid"
)

var diceTable = []string{
	"3,1", "3,2",
	"4,1", "4,2", "4,3",
	"5,1", "5,2", "5,3", "5,4",
	"6,1", "6,2", "6,3", "6,4", "6,5",
	"1,1", "2,2", "3,3", "4,4", "5,5", "6,6",
	"2,1",
}

// ParseDice ...
func ParseDice(dice string) (int, error) {
	for val, str := range diceTable {
		if str == dice {
			return val, nil
		}
	}
	return -1, fmt.Errorf(ErrorDiceIsNotValid)
}

// DiceToString ...
func DiceToString(value int) string {
	if value > 20 {
		return diceTable[20]
	} else if value < 0 {
		return diceTable[0]
	}
	return diceTable[value]
}
