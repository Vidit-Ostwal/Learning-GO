package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Vidit is here")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:], "a")
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 211
	highScores[1] = 242
	highScores[2] = 241
	highScores[3] = 240

	highScores = append(highScores, 239, 238)
	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
