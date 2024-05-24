package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's study Array!")

	var movieList [5]string

	movieList[0] = "Inception"
	movieList[1] = "JohnWick"
	movieList[2] = "Avatar"
	fmt.Println("movieList : ", movieList)
	fmt.Println("movieList length : ", len(movieList))

	var seriesList = [2]string{"Ozark", "Lucifer"}
	fmt.Println("seriesList : ", seriesList)
	fmt.Println("seriesList length : ", len(seriesList))
	var fruitList = []string{"Apple", "Mango", "Watermelon", "Banana"}
	fmt.Println("fruitList : ", fruitList)
	fmt.Println("modified fruitList : ", fruitList[1:])
	fmt.Println("modified fruitList : ", fruitList[:3])
	fmt.Println("modified fruitList : ", fruitList[1:3])
	fmt.Println("modified fruitList : ", append(fruitList, "StrawBerry"))

	highScore := make([]int, 4) //Allocate default memory according to size
	highScore[0] = 984
	highScore[1] = 543
	highScore[3] = 6239
	fmt.Println("highScore : ", highScore)
	highScore = append(highScore, 1010, 61) //Relocate memory along with changing the size
	fmt.Println("highScore : ", sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println("highScore : ", highScore)
	fmt.Println("highScore : ", sort.IntsAreSorted(highScore))

	index := 2
	fmt.Println("highScore : ", append(highScore[index:len(highScore)-index], highScore[index+1:]...))

}
