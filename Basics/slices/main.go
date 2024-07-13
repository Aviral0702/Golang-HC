package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	// when we create a slice we don't define how many elements will be there
	var fruitList = []string{"Apple", "Banana", "Strawberry", "Mango"}

	fmt.Println("The fruitlist is : ", fruitList)

	fruitList = append(fruitList, "Peach")

	fmt.Println("The fruitlist is : ", fruitList)

	fruitList = append(fruitList[1:])

	fmt.Println("The fruitlist is : ", fruitList)

	//we are going to apply slices with make function

	highScores := make([]int, 5) // by default this is array but a layer of abstraction is there so this is slices
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	highScores[4] = 234
	fmt.Printf("The type of highScores is %T\n", highScores)
	fmt.Println("The high scores are : ", highScores)

	// When appending to a full slice:
	// 1. Go allocates a new array with doubled capacity
	// 2. Copies existing elements to the new array
	// 3. Adds the new element to the end
	// 4. Updates the slice header to point to the new array
	// 5. Old array is left for garbage collection if no array is referencing it

	//This process is called "growing the slice."
	//It's an expensive operation because it involves allocating new memory and copying data.
	//However, by doubling the capacity each time, Go ensures that this
	//expensive operation happens less frequently as the slice grows larger.

	highScores = append(highScores, 553, 222, 121)

	fmt.Println("The high scores are : ", highScores)
	//The main differences from Go are that C++ vectors handle their own memory directly,
	//use move semantics when possible, and may have different growth factors depending on the implementation.

	sort.Ints(highScores)
	fmt.Println("The sorted high scores are : ", highScores)
	fmt.Println("Are the high scores sorted? : ", sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //V.V.Imp
	fmt.Println(courses)
}
