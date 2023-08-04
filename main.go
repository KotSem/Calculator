package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Добро пожаловать в бета-версию Calculator v.0.1\n" +
		"Правила использования:\n" +
		"1. Введите в одну строку два числа и математический знак между ними.\n" +
		"2. Можно использовать как арабские так и римские числа, главное, чтобы оба числа были в одной системе исчисления\n" +
		"3. На данный момент калькулятор поддерживает только 4 математических операции: '+', '-', '*', '/'\n" +
		"P.S.Enjoy and give feedback to https://github.com/psatin42/\n" +
		"\n\n" +
		"Введите математическое выражение:")
	intType, first, second, sign, err := readLine()
	if err != nil {
		fmt.Println("Возникла ошибка при вводе данных:\n", err)
		return
	}
	if intType == "arab" {
		firstNum, err1 := strconv.Atoi(first)
		if err1 != nil {
			fmt.Println("Возникла ошибка при переводе строки в число:\n", err1)
			return
		}
		secondNum, err2 := strconv.Atoi(second)
		if err2 != nil {
			fmt.Println("Возникла ошибка при переводе строки в число:\n", err2)
			return
		}
		res, err3 := calculator(firstNum, secondNum, sign)
		if err3 != nil {
			fmt.Println("Возникла ошибка при работе калькулятора:\n", err3)
			return
		} else {
			fmt.Println("Ответ: ", res)
		}
	} else {
		firstNum := fromRomanToInt(first)
		secondNum := fromRomanToInt(second)
		res, err1 := calculator(firstNum, secondNum, sign)
		if err1 != nil {
			fmt.Println("Возникла ошибка при работе калькулятора:\n", err1)
			return
		} else {
			final, err2 := fromIntToRoman(res)
			if err2 != nil {
				fmt.Println("Возникла ошибка при работе калькулятора:\n", err2)
				return
			}
			fmt.Println("Ответ: ", final)
		}
	}
}
