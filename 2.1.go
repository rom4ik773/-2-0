package main

import (
	"fmt"
)

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToArabic(roman string) (int, error) {
	arabic := 0
	lastValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < lastValue {
			arabic -= value
		} else {
			arabic += value
		}
		lastValue = value
	}
	return arabic, nil
}

func main() {
	var input string
	fmt.Println("Введите римское число для конвертации в арабские цифры:")
	fmt.Scanln(&input)
	arabic, err := romanToArabic(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("%s в арабских цифрах: %d\n", input, arabic)
	}
}
