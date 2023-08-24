package main

import (
	"fmt"
	"strings"
)

var StrIn string
var Brackets = make(map[string]int)
var BracketList = []string{"{", "}", "(", ")", "[", "]"}
var Balance = true

func main() {

	culredLeft := "{"
	culredRight := "}"
	roundLeft := "("
	roundRight := ")"
	squareLeft := "["
	squareRight := "]"

	fmt.Println("Введите строку:")
	fmt.Scan(&StrIn)

	if len(StrIn) == 0 {
		fmt.Println("Пустая строка")
		return
	}

	var strSet = strings.Split(StrIn, "")

	for _, value := range strSet {
		if containsString(BracketList, value) {
			Brackets[value]++
		}
	}

	if isSet(Brackets, culredLeft) && isSet(Brackets, culredRight) &&
		Brackets[culredLeft] != Brackets[culredRight] {
		Balance = false
	} else if (!isSet(Brackets, culredLeft) && isSet(Brackets, culredRight)) ||
		(isSet(Brackets, culredLeft) && !isSet(Brackets, culredRight)) {
		Balance = false
	}

	if isSet(Brackets, roundLeft) && isSet(Brackets, roundRight) &&
		Brackets[roundLeft] != Brackets[roundRight] {
		Balance = false
	} else if (!isSet(Brackets, roundLeft) && isSet(Brackets, roundRight)) ||
		(isSet(Brackets, roundLeft) && !isSet(Brackets, roundRight)) {
		Balance = false
	}

	if isSet(Brackets, squareLeft) && isSet(Brackets, squareRight) &&
		Brackets[squareLeft] != Brackets[squareRight] {
		Balance = false
	} else if (!isSet(Brackets, squareLeft) && isSet(Brackets, squareRight)) ||
		(isSet(Brackets, squareLeft) && !isSet(Brackets, squareRight)) {
		Balance = false
	}

	if !Balance {
		fmt.Println("Скобки НЕсбалансированы!")
		return
	}

	fmt.Println("Скобки сбалансированы!")
}

func containsString(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func isSet(m map[string]int, k string) bool {
	_, ok := m[k]
	return ok
}
