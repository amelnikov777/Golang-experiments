//To Do update 1.enter maximum value, 2.sorting before writing to file, 3. Maiking +,*,!,stepen 4.Use method
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	intCh := make(chan int, 20)
	intChnum := make(chan int, 20)
	var val, valnum int
	file, err := os.Create("учеба.txt")
	if err != nil {
		fmt.Println("Error to create file", err)
		os.Exit(1)
	}
	defer file.Close()
	for i := 1; i < 20; i++ {
		go factorial(i, intCh, intChnum)
	}
	for i := 1; i < 20; i++ {
		val = <-intCh
		valnum = <-intChnum
		fmt.Println(valnum, " - ", val)
		file.WriteString(strconv.Itoa(valnum) + " - " + strconv.Itoa(val) + "\n")
	}
	fmt.Println("The End")
}

func factorial(n int, ch, chnum chan int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
	chnum <- n
}
