package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("wlecome to GOLang!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rate GOlang!")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating!, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 is added to your rating : ", numRating+1)
	}

}
