package main

import "fmt"

var lowNames = map[int]string{0: "", 1: "один", 2: "два", 3: "три", 4: "четыре", 5: "пять", 6: "шесть", 7: "семь", 8: "восемь", 9: "девять", 10: "десять", 11: "одинадцать", 12: "двенадцать", 13: "тринадцать", 14: "четырнадцать", 15: "пятнадцать", 16: "шестнадцать", 17: "семнадцать", 18: "восемнадцать", 19: "девятнадцать"}
var tensNames = map[int]string{20: "двадцать", 30: "тридцать", 40: "сорок", 50: "пятьдесят", 60: "шестьдесят", 70: "семьдесят", 80: "восемьдесят", 90: "девяносто"}
var hundNames = map[int]string{1: "сто", 2: "двести", 3: "триста", 4: "четыреста", 5: "пятьсот", 6: "шестьсот", 7: "семьсот", 8: "восемьсот", 9: "девятьсот"}
var lowThausNames = map[int]string{1: "одна тысяча", 2: "две тысячи", 3: "три тысячи", 4: "четыре тысячи", 5: "пять тысяч", 6: "шесть тысяч", 7: "семь тысяч", 8: "восемь тысяч", 9: "девять тысяч", 10: "десять тысяч", 11: "одинадцать тысяч", 12: "двенадцать тысяч", 13: "тринадцать тысяч", 14: "четырнадцать тысяч", 15: "пятнадцать тысяч", 16: "шестнадцать тысяч", 17: "семнадцать тысяч", 18: "восемьнадцать тысяч", 19: "девятнадцать тысяч"}

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

func convert1000(num int) string {
	return hundNames[num/100] + " " + convert100(num%100)
}

func convert1000_(num int) string {
	return " " + hundNames[num/100] + " " + convert100(num%100)
}

func convert100000(num int) string {
	if num%1000 < 20 {
		return lowThausNames[num/1000] + convert1000(num%1000)
	} //else {
	//return lowThausNames[num/1000] + convert1000_(num%1000)
	//}
	//	if num < 100000 {
	//		return lowThausNames[num/1000] + " " + convert1000(num%1000)
	//	}
	//	if num/1000 < 100 {
	//		if num%10 == 0 {
	//			return tensNames[num]
	//		} else {
	//			return tensNames[num/10*10] + " " + lowNames[num%10]
	//		}
	//	}
	return "convert100000"
}

func convertNumToWord(num int) string {
	if num == 0 {
		return "ноль"
	}
	if num < 100 {
		return convert100(num)
	}
	if num < 1000 {
		return convert1000(num)
	}
	if num < 1000000 {
		return convert100000(num)
	}
	return "convertNumToWord"
}

func main() {
	for i := 0; i <= 1101; i++ {
		fmt.Println(i, " "+convertNumToWord(i))
	}

}
