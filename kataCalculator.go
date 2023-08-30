package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var romanNums = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func romanConverter(res int) string {
	var result string
	roman := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabic := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(arabic); i++ {
		for res >= arabic[i] {
			res -= arabic[i]
			result += roman[i]
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите математическую операцию с римскими или с арабскими целыми числами от 1 до 10: ")
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")

	if len(input) != 3 {
		log.Fatal("Ошибка. Допустимо использование только двух операторов и одного операнда (+, -, *, /).")
		return
	}

	num1, err1 := strconv.Atoi(input[0])
	num2, err2 := strconv.Atoi(input[2])

	if err1 != nil || err2 != nil {
		if roman, ok := romanNums[input[0]]; ok {
			num1 = roman //fmt.Println(roman, num1, num2) --Проверка первого римского
		} else {
			log.Fatal("Ошибка. Математические операции допустимы только с арабскими или с рискими целыми числами.")
			return
		}
		if roman, ok := romanNums[input[2]]; ok {
			num2 = roman //fmt.Println(roman, num1, num2) --Проверка второго римского
		} else {
			log.Fatal("Ошибка. Математические операции допустимы только с арабскими или с рискими целыми числами.")
			return
		}
	}

	if num1 > 10 {
		log.Fatal("Первое вводимое число не может быть больше 10")
	} else if num2 > 10 {
		log.Fatal("Второе вводимое число не может быть больше 10")
	}

	var res int
	switch input[1] {
	case "+":
		res = num1 + num2
	case "-":
		if err1 != nil || err2 != nil { //Проверка на римское число перед вычитанием
			if num2 >= num1 {
				log.Fatal("Ошибка. Римская система счисления имеет только натуральные числа.")
				return
			}
		}
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
		if num2 == 0 {
			log.Fatal("Ошибка. Нельзя делить на ноль.")
			return
		}
	default:
		log.Fatal("Ошибка. Неизвестная операция")
		return
	}

	if err1 != nil || err2 != nil {
		fmt.Println(romanConverter(res)) //romanConverter --Конвертация res int  в Roman
	} else {
		fmt.Println(res)
	}
}
