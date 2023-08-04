package main

import (
	"strings"
)

var dict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func fromRomanToInt(roman string) int {
	var res int
	arr := strings.Split(roman, "")
	for index, value := range arr {
		if index+1 != len(arr) && dict[value] < dict[arr[index+1]] {
			res -= dict[value]
		} else {
			res += dict[value]
		}
	}
	return res
}

func fromIntToRoman(number int) (string, error) {
	if number <= 0 {
		return "", errorHandler(7)
	}
	var str string
	for key, value := range dict {
		if value == number {
			return key, nil
		}
	}
	for number > 0 {
		for key, value := range dict {
			if dict[key] <= number {
				str += key
				number -= value
				break
			}
		}
	}
	return str, nil
}
