// 1
package main

import (
	"fmt"
)

var move int

func main() {

	disks := 0
	move = 0
	fmt.Println("Введите кол-во дисков")
	fmt.Scan(&disks)
	fmt.Println()
	hanoi(disks, "A", "B", "C")
}

func hanoi(disks int, source, helper, destination string) {
	if disks == 1 {
		fmt.Println("move ", move, " Disk ", disks, " moves from tower ", source, " to tower ", destination)
		return
	}

	hanoi(disks-1, source, destination, helper)
	move = move + 1
	fmt.Println("move ", move, " Disk ", disks, " moves from tower ", source, " to tower ", destination)
	move = move + 1
	hanoi(disks-1, helper, source, destination)
}
