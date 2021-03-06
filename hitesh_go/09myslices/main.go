package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Mango", "Tomato", "Peach"}
	fmt.Println("FruitList is:", fruitList)
	fmt.Println("FruitList length is:", len(fruitList))
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Apple", "Banana")
	fmt.Println("FruitList is:", fruitList)
	fmt.Println("FruitList length is:", len(fruitList))
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList[1:3])
	fruitList = fruitList[1:]
	fmt.Println("FruitList is:", fruitList)
	fmt.Println("FruitList length is:", len(fruitList))
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println(highScores)
	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
