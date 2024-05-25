package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and cases in GOlang")
	dice := rand.Intn(8) + 1
	fmt.Println(dice)
	switch dice {
	case 1:
		fmt.Println("you got 1 move 1 step")
	case 2:
		fmt.Println("you got 2 move 2 step")
		fallthrough // fall to next case also ie case 3
	case 3:
		fmt.Println("you got 3 move 3 step")
	case 4:
		fmt.Println("you got 4 move 4 step")
	case 5:
		fmt.Println("you got 5 move 5 step")
	case 6:
		fmt.Println("you got 6 move 6 step")
	default:
		fmt.Println("What was that!")
	}
}
