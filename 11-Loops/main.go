package main

import "fmt"

func main() {
	fmt.Println("Loops in GOLlang!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("%v\n", days[i])
	// }

	// for i := range days {
	// 	fmt.Printf("%v\n", days[i])
	// }

	// for idx, value := range days {
	// 	fmt.Printf("index: %v => days: %v\n", idx, value)
	// }

	for _, value := range days {
		fmt.Printf("days: %v\n", value)
	}

	// Similar to while loop
	var idx int = 0
	for idx < len(days) {
		fmt.Printf("days: %v\n", days[idx])

		if idx == 1 {
			idx++
			continue
		}
		if idx == 4 {
			goto getJump
		}
		if idx == 5 {
			break
		}
		fmt.Println("value : ", idx)
		idx++
	}

getJump:
	fmt.Println("get jumped to this from loop")

}
