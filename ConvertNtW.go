package main

import (
	"fmt"
	"strings"
)

var lowNames = map[int]string{0: "", 1: "один", 2: "два", 3: "три", 4: "четыре", 5: "пять", 6: "шесть", 7: "семь", 8: "восемь", 9: "девять", 10: "десять", 11: "одинадцать", 12: "двенадцать", 13: "тринадцать", 14: "четырнадцать", 15: "пятнадцать", 16: "шестнадцать", 17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать"}
var tensNames = map[int]string{20: "двадцать", 30: "тридцать", 40: "сорок", 50: "пятьдесят", 60: "шестьдесят", 70: "семьдесят", 80: "восемьдесят", 90: "девяносто"}
var hundNames = map[int]string{1: "сто", 2: "двести", 3: "триста", 4: "четыреста", 5: "пятьсот", 6: "шестьсот", 7: "семьсот", 8: "восемьсот", 9: "девятьсот"}
var lowThausNames = map[int]string{0: "тысяч", 1: "одна тысяча", 2: "две тысячи", 3: "три тысячи", 4: "четыре тысячи", 5: "пять тысяч", 6: "шесть тысяч", 7: "семь тысяч", 8: "восемь тысяч", 9: "девять тысяч", 10: "десять тысяч", 11: "одинадцать тысяч", 12: "двенадцать тысяч", 13: "тринадцать тысяч", 14: "четырнадцать тысяч", 15: "пятнадцать тысяч", 16: "шестнадцать тысяч", 17: "семнадцать тысяч", 18: "восемьнадцать тысяч", 19: "девятнадцать тысяч"}
var lowMillNames = map[int]string{0: "миллион", 1: "один миллион", 2: "два миллиона", 3: "три миллиона", 4: "четыре миллиона", 5: "пять миллионов", 6: "шесть миллионов", 7: "семь миллионов", 8: "восемь миллионов", 9: "девять миллионов", 10: "десять миллионов", 11: "одинадцать миллионов", 12: "двенадцать миллионов", 13: "тринадцать миллионов", 14: "четырнадцать миллионов", 15: "пятнадцать миллионов", 16: "шестнадцать миллионов", 17: "семнадцать миллионов", 18: "восемьнадцать миллионов", 19: "девятнадцать миллионов"}

func convert100(num int) string {
	if num < 20 {
		return lowNames[num]
	}
	if num < 100 {
		if num%10 == 0 {
			return tensNames[num]
		} else {
			return tensNames[num/10*10] + " " + lowNames[num%10]
		}
	}
	return "convert100"
}

func convert100K(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return lowThausNames[num]
	}
	if num < 100 {
		if num%10 == 0 {
			return tensNames[num] + lowThausNames[0]
		} else {
			return tensNames[num/10*10] + " " + lowThausNames[num%10]
		}
	}
	return "convert100K"
}

func convert100M(num int) string {
	if num == 0 {
		return ""
	}
	if num < 20 {
		return lowMillNames[num]
	}
	if num < 100 {
		if num%10 == 0 {
			return tensNames[num] + lowMillNames[0]
		} else {
			return tensNames[num/10*10] + " " + lowMillNames[num%10]
		}
	}
	return "convert100M"
}

func convert999(num int) string {
	if num < 100 {
		return convert100(num)
	} else {
		return hundNames[num/100] + " " + convert100(num%100)
	}
}

func convert999K(num int) string {
	if num < 100 {
		return convert100K(num)
	} else {
		return hundNames[num/100] + " " + convert100K(num%100)
	}
}

func convert999M(num int) string {
	if num < 100 {
		return convert100M(num)
	} else {
		return hundNames[num/100] + " " + convert100M(num%100)
	}
}

func convertNumToWord(num int) string {
	if num == 0 {
		return "ноль"
	}
	if num < 1000 {
		return convert999(num)
	}
	if num < 1000000 {
		return convert999K(num/1000) + " " + convert999(num%1000)
	}
	if num < 1000000000 {
		return convert999M(num/1000000) + " " + convert999K((num/1000)%1000) + " " + convert999(num%1000)
	}
	return "Число больше 999 миллионов"
}

func main() {
	var i int
	fmt.Println("Введите число для конвертации")
	fmt.Scanln(&i)
	if i >= 0 {
		fmt.Println(i, " ", strings.Replace(convertNumToWord(i), "  ", " ", -1))
	} else {
		fmt.Println(i, " минус", strings.Replace(convertNumToWord(i*-1), "  ", " ", -1))
	}

}
